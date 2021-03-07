package server

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/plaid/plaid-go/plaid"
	"github.com/rs/xid"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/khoerling/flux/api/lib/auth"
	"github.com/khoerling/flux/api/lib/db/models/onetimepasscode"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/plaid/item"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/address"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/dateofbirth"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/legalname"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/ssn"
	proto "github.com/khoerling/flux/api/lib/protocol"
)

// https://api.sendwyre.com/v3/rates?as=priced

// ViewerData is an rpc handler
func (s *Server) ViewerData(ctx context.Context, _ *emptypb.Empty) (*proto.ViewerDataResponse, error) {
	userID, err := GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, err
	}

	u, err := s.Db.GetUserByID(ctx, user.ID(userID))
	if err != nil || u == nil {
		return nil, status.Errorf(codes.Unauthenticated, codes.Unauthenticated.String())
	}

	user := proto.User{
		Id:        string(u.ID),
		CreatedAt: u.CreatedAt.Unix(),
	}
	if u.Email != nil {
		user.Email = string(*u.Email)
	}
	if u.Phone != nil {
		user.Phone = string(*u.Phone)
	}

	// todo: factor this out into separate module (db)
	plaidItems := s.Db.Firestore.
		Collection("users").
		Doc(user.Id).Collection("plaidItems").
		Limit(1).Documents(ctx)

	var hasPlaidItems bool

	_, err = plaidItems.Next()
	if err == iterator.Done {
		hasPlaidItems = false
	} else if err != nil {
		return nil, err
	} else {
		hasPlaidItems = true
	}

	flags := proto.UserFlags{
		HasPlaidItems: hasPlaidItems,
		// todo: implement first
		//HasWyreAccount:        false,
		//HasWyrePaymentMethods: false,
	}

	return &proto.ViewerDataResponse{
		User:  &user,
		Flags: &flags,
	}, nil
}

// PricingData is an rpc handler
func (s *Server) PricingData(ctx context.Context, in *proto.PricingDataRequest) (*proto.PricingDataResponse, error) {
	wyreRates, err := s.Wyre.PricedExchangeRates()
	if err != nil {
		return nil, err
	}

	rates := map[string]*proto.PricingRate{}
	resp := proto.PricingDataResponse{
		Rates: rates,
	}

	for rateMapName, rateMap := range *wyreRates {
		ratePairAry := []string{}
		for currencySymbol := range rateMap {
			if currencySymbol == rateMapName[:len(currencySymbol)] {
				ratePairAry = append(ratePairAry, currencySymbol)
				break
			}
		}
		ratePairAry = append(ratePairAry, rateMapName[len(ratePairAry[0]):])

		newRatePairName := strings.Join(ratePairAry, "_")
		rates[newRatePairName] = &proto.PricingRate{
			Rate: rateMap,
		}
	}

	return &resp, nil
}

// OneTimePasscode is an rpc handler
func (s *Server) OneTimePasscode(ctx context.Context, req *proto.OneTimePasscodeRequest) (*proto.OneTimePasscodeResponse, error) {
	loginKind, loginValue, err := ValidateAndNormalizeLogin(req.EmailOrPhone)
	if err != nil {
		return nil, err
	}

	otp, err := s.Db.CreateOneTimePasscode(ctx, loginValue, loginKind)
	if err != nil {
		return nil, err
	}

	if loginKind == onetimepasscode.LoginKindPhone {
		from := s.TwilioPhoneNumber
		to := loginValue
		message := fmt.Sprintf("Your one time passcode for flux is: %s", otp.Code)

		_, _, err := s.Twilio.SendSMS(from, to, message, "", "")
		if err != nil {
			return nil, err
		}

		return &proto.OneTimePasscodeResponse{}, nil
	}

	msg := generateOtpMessage(mail.NewEmail("Customer", loginValue), otp.Code)

	_, err = s.Sendgrid.Send(msg)
	if err != nil {
		return nil, err
	}

	return &proto.OneTimePasscodeResponse{}, nil
}

// OneTimePasscodeVerify is an rpc handler
func (s *Server) OneTimePasscodeVerify(ctx context.Context, req *proto.OneTimePasscodeVerifyRequest) (*proto.OneTimePasscodeVerifyResponse, error) {
	loginKind, loginValue, err := ValidateAndNormalizeLogin(req.EmailOrPhone)
	if err != nil {
		return nil, err
	}

	passcodes := s.Firestore.Collection("one-time-passcodes").
		Where("emailOrPhone", "==", loginValue).
		Where("code", "==", req.Code).
		Where("createdAt", ">", time.Now().Add(-10*time.Minute)).
		Documents(ctx)

	unknownMsg := "An unknown error occurred. Please try again later."

	passcode, err := passcodes.Next()
	if err == iterator.Done {
		return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedOTP(loginKind))
	}
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unauthenticated, unknownMsg)
	}

	_, err = passcode.Ref.Delete(ctx)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unauthenticated, unknownMsg)
	}

	u, err := s.Db.GetOrCreateUser(ctx, loginKind, loginValue)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unauthenticated, unknownMsg)
	}

	jwt, err := s.JwtSigner.Sign(auth.NewClaims(u.ID))
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unauthenticated, unknownMsg)
	}

	respUser := &proto.User{
		Id:        string(u.ID),
		CreatedAt: u.CreatedAt.Unix(),
	}

	if u.Email != nil {
		respUser.Email = string(*u.Email)
	}

	if u.Phone != nil {
		respUser.Phone = string(*u.Phone)
	}

	return &proto.OneTimePasscodeVerifyResponse{
		Jwt:  jwt,
		User: respUser,
	}, nil
}

// PlaidConnectBankAccounts is an rpc handler
func (s *Server) PlaidConnectBankAccounts(ctx context.Context, req *proto.PlaidConnectBankAccountsRequest) (*proto.PlaidConnectBankAccountsResponse, error) {
	userID, err := GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, err
	}

	err = req.Validate()
	if err != nil {
		return nil, err
	}

	resp, err := s.Plaid.ExchangePublicToken(req.PlaidPublicToken)
	if err != nil {
		return nil, err
	}
	log.Printf("Plaid Public Token successfuly exchanged")

	_, err = s.Db.SavePlaidItem(ctx, userID, item.ID(resp.ItemID), resp.AccessToken)
	if err != nil {
		return nil, err
	}
	log.Printf("Plaid ItemID %s saved", resp.ItemID)

	for _, plaidAccountID := range req.PlaidAccountIds {
		log.Printf("STUB > process PlaidAccountID: %s", plaidAccountID)
	}

	processorTokenResp, err := s.Plaid.CreateProcessorToken(resp.AccessToken, req.PlaidAccountIds[0], "wyre")
	if err != nil {
		return nil, err
	}
	log.Printf("processor token: %s", processorTokenResp.ProcessorToken)

	return &proto.PlaidConnectBankAccountsResponse{}, nil
}

func generateOtpMessage(to *mail.Email, code string) *mail.SGMailV3 {
	from := mail.NewEmail("Ctulhu", "ctulhu@dreamcodez.cc")
	subject := "Your one time passcode for flux"
	plainTextContent := fmt.Sprintf("Your one time passcode is: %s", code)
	htmlContent := fmt.Sprintf("Your one time passcode is: <strong>%s</strong>", code)
	return mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
}

// PlaidCreateLinkToken is an rpc handler
func (s *Server) PlaidCreateLinkToken(ctx context.Context, req *proto.PlaidCreateLinkTokenRequest) (*proto.PlaidCreateLinkTokenResponse, error) {
	userID, err := GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("Generating Plaid Link Token for User ID: %s", userID)

	u, err := s.Db.GetUserByID(ctx, user.ID(userID))
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unknown, "An unknown error ocurred; please try again.")
	}
	if u == nil {
		return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
	}

	// TODO: remove me, for testing only
	//pdata, err := s.Db.GetAllProfileData(ctx, userID)
	//if err != nil {
	//	return nil, err
	//}
	//log.Printf("%#v", pdata)

	plaidUserDetails := plaid.LinkTokenUser{
		ClientUserID: string(userID),
	}

	linkTokenResp, err := s.Plaid.CreateLinkToken(plaid.LinkTokenConfigs{
		User:         &plaidUserDetails,
		ClientName:   "Flux",
		Products:     []string{"auth"},
		CountryCodes: []string{"US"},
		Language:     "en",
		/* NOTE: may need this?
		Products:     []string{"auth", "transactions"},
			Webhook:               "https://webhook-uri.com",
			LinkCustomizationName: "default",
			AccountFilters: &map[string]map[string][]string{
				"depository": map[string][]string{
					"account_subtypes": []string{"checking", "savings"},
				},
			},
		*/
	})
	if err != nil {
		return nil, err

	}

	return &proto.PlaidCreateLinkTokenResponse{
		LinkToken: linkTokenResp.LinkToken,
	}, nil
}

// SaveProfileData is an rpc handler
func (s *Server) SaveProfileData(ctx context.Context, req *proto.SaveProfileDataRequest) (*proto.ProfileDataInfo, error) {
	userID, err := GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, err
	}

	u, err := s.Db.GetUserByID(ctx, user.ID(userID))
	if err != nil || u == nil {
		return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
	}

	err = req.Validate()
	if err != nil {
		return nil, err
	}

	var legalNameData *legalname.ProfileDataLegalName
	var dobData *dateofbirth.ProfileDataDateOfBirth
	var ssnData *ssn.ProfileDataSSN
	var addressData *address.ProfileDataAddress

	err = s.Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		profile, err := s.Db.GetAllProfileData(ctx, tx, u.ID)
		if err != nil {
			return err
		}

		{
			existingProfileData := profile.FilterKind(common.KindLegalName).First()
			if existingProfileData != nil {
				legalNameData = (*existingProfileData).(*legalname.ProfileDataLegalName)
			}
		}
		{
			existingProfileData := profile.FilterKind(common.KindDateOfBirth).First()
			if existingProfileData != nil {
				dobData = (*existingProfileData).(*dateofbirth.ProfileDataDateOfBirth)
			}
		}
		{
			existingProfileData := profile.FilterKind(common.KindSSN).First()
			if existingProfileData != nil {
				ssnData = (*existingProfileData).(*ssn.ProfileDataSSN)
			}
		}
		{
			existingProfileData := profile.FilterKind(common.KindAddress).First()
			if existingProfileData != nil {
				addressData = (*existingProfileData).(*address.ProfileDataAddress)
			}
		}

		if req.LegalName != "" {
			if legalNameData == nil {
				legalNameData = &legalname.ProfileDataLegalName{
					ID:        common.ProfileDataID(xid.New().String()),
					Status:    common.StatusReceived,
					LegalName: req.LegalName,
					CreatedAt: time.Now(),
				}
			} else {
				legalNameData.LegalName = req.LegalName
				now := time.Now()
				legalNameData.UpdatedAt = &now
			}

			_, err := s.Db.SaveProfileData(ctx, tx, userID, *legalNameData)
			if err != nil {
				return err
			}
		}

		if req.DateOfBirth != "" {
			if dobData == nil {
				dobData = &dateofbirth.ProfileDataDateOfBirth{
					ID:          common.ProfileDataID(xid.New().String()),
					Status:      common.StatusReceived,
					DateOfBirth: req.DateOfBirth,
					CreatedAt:   time.Now(),
				}
			} else {
				dobData.DateOfBirth = req.DateOfBirth
				now := time.Now()
				dobData.UpdatedAt = &now
			}

			_, err := s.Db.SaveProfileData(ctx, tx, userID, *dobData)
			if err != nil {
				return err
			}
		}

		if req.Ssn != "" {
			if ssnData == nil {
				ssnData = &ssn.ProfileDataSSN{
					ID:        common.ProfileDataID(xid.New().String()),
					Status:    common.StatusReceived,
					SSN:       req.Ssn,
					CreatedAt: time.Now(),
				}
			} else {
				ssnData.SSN = req.Ssn
				now := time.Now()
				ssnData.UpdatedAt = &now
			}

			_, err := s.Db.SaveProfileData(ctx, tx, userID, *dobData)
			if err != nil {
				return err
			}
		}

		if req.Address != nil {
			addressData := &address.ProfileDataAddress{
				ID:         common.ProfileDataID(xid.New().String()),
				Status:     common.StatusReceived,
				Street1:    req.Address.Street_1,
				Street2:    req.Address.Street_2,
				City:       req.Address.City,
				State:      req.Address.State,
				PostalCode: req.Address.PostalCode,
				Country:    req.Address.Country,
				CreatedAt:  time.Now(),
			}
			_, err := s.Db.SaveProfileData(ctx, tx, userID, addressData)
			if err != nil {
				return err
			}
		}

		return nil
	})

	var legalNameInfo *proto.ProfileDataItemInfo
	if legalNameData != nil {
		legalNameInfo = legalNameData.GetProfileDataItemInfo()
	}

	var dobInfo *proto.ProfileDataItemInfo
	if dobData != nil {
		dobInfo = dobData.GetProfileDataItemInfo()
	}

	var ssnInfo *proto.ProfileDataItemInfo
	if ssnData != nil {
		ssnInfo = ssnData.GetProfileDataItemInfo()
	}

	var addressInfo *proto.ProfileDataItemInfo
	if addressData != nil {
		addressInfo = addressData.GetProfileDataItemInfo()
	}

	var email *proto.ProfileDataItemInfo
	if (u.Email != nil && *u.Email != "" && u.EmailVerifiedAt != nil && *u.EmailVerifiedAt != time.Time{}) {
		email = &proto.ProfileDataItemInfo{
			Id:        string(u.ID),
			Kind:      proto.ProfileDataItemKind_K_EMAIL,
			Status:    proto.ProfileDataItemStatus_S_RECEIVED,
			Length:    int32(len(*u.Email)),
			CreatedAt: u.EmailVerifiedAt.Format(time.RFC3339),
		}
	}

	var phone *proto.ProfileDataItemInfo
	if (u.Phone != nil && *u.Phone != "" && u.PhoneVerifiedAt != nil && *u.PhoneVerifiedAt != time.Time{}) {
		phone = &proto.ProfileDataItemInfo{
			Id:        string(u.ID),
			Kind:      proto.ProfileDataItemKind_K_PHONE,
			Status:    proto.ProfileDataItemStatus_S_RECEIVED,
			Length:    int32(len(*u.Phone)),
			CreatedAt: u.PhoneVerifiedAt.Format(time.RFC3339),
		}
	}

	return &proto.ProfileDataInfo{
		LegalName:   legalNameInfo,
		DateOfBirth: dobInfo,
		Ssn:         ssnInfo,
		Address:     addressInfo,
		Email:       email,
		Phone:       phone,
	}, nil
}

// WyreCreateAccount is an rpc handler
func (s *Server) WyreCreateAccount(ctx context.Context, req *proto.WyreCreateAccountRequest) (*proto.WyreCreateAccountResponse, error) {
	userID, err := GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, err
	}

	u, err := s.Db.GetUserByID(ctx, user.ID(userID))
	if err != nil || u == nil {
		return nil, status.Errorf(codes.Unauthenticated, codes.Unauthenticated.String())
	}

	if u.Email == nil || *u.Email == "" {
		return nil, status.Errorf(codes.FailedPrecondition, "Sorry! We cannot create an account with Wyre; an email must be set for your user")
	}
	//email := *u.Email

	return nil, nil
	// TODO: serverside validation on this request
	// TODO: store address info and name info for future kyc use
	/*
		name := "bob"
		address := "fixme"
		truth := true

		profileFields := []wyre.ProfileField{
			{
				FieldID: wyre.ProfileFieldIDIndividualLegalName,
				Value:   name,
			},
			{
				FieldID: wyre.ProfileFieldIDIndividualEmail,
				Value:   email,
			},
			{
				FieldID: wyre.ProfileFieldIDIndividualResidenceAddress,
				Value: wyre.ProfileFieldAddress{
					Street1:    address.Street_1,
					Street2:    address.Street_2,
					City:       address.City,
					State:      address.State,
					PostalCode: address.PostalCode,
					Country:    address.Country,
				},
			},
		}

		wyreReq := wyre.CreateAccountRequest{
			SubAccount:    &truth,
			DisableEmail:  &truth,
			ProfileFields: profileFields,
		}.WithDefaults()

		wyreAccount, err := s.Wyre.CreateAccount(wyreReq)
		if err != nil {
			return nil, err
		}

		log.Printf("wyre account created: %#v", wyreAccount)

		return &proto.WyreCreateAccountResponse{}, nil
	*/
}

// UploadFile ..
/*
func (s *Server) UploadFile(ctx context.Context, req *httpbody.HttpBody) (*httpbody.HttpBody, error) {
	log.Println("UploadFile")
	log.Printf("received %d bytes of data from the upload", len(req.Data))
	return &httpbody.HttpBody{
		ContentType: "application/json",
		Data:        []byte("OK"),
	}, nil
}

*/
