package server

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/plaid/plaid-go/plaid"
	"github.com/pusher/pusher-http-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/disintegration/imaging"
	"github.com/khoerling/flux/api/lib/auth"
	"github.com/khoerling/flux/api/lib/db/models/onetimepasscode"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/file"
	"github.com/khoerling/flux/api/lib/db/models/user/plaid/item"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/address"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/dateofbirth"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/legalname"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/ssn"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/usgovernmentid"
	proto "github.com/khoerling/flux/api/lib/protocol"

	"github.com/lithammer/shortuuid/v3"
)

// https://api.sendwyre.com/v3/rates?as=priced

// ViewerData is an rpc handler
func (s *Server) ViewerData(ctx context.Context, _ *emptypb.Empty) (*proto.ViewerDataResponse, error) {
	userID := GetUserIDFromIncomingContext(ctx)
	if userID == "" {
		return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
	}

	u, err := s.Db.GetUserByID(ctx, user.ID(userID))
	if err != nil || u == nil {
		return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
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

	hasPlaidItems, err := s.Db.HasPlaidItems(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	profile, err := s.Db.GetAllProfileData(ctx, nil, u.ID)
	if err != nil {
		return nil, err
	}

	var hasWyreAccount bool

	{
		accounts, err := s.Db.GetWyreAccounts(ctx, nil, userID)
		if err != nil {
			return nil, err
		}
		if len(accounts) > 0 {
			hasWyreAccount = true
		}
	}

	flags := proto.UserFlags{
		HasPlaidItems:                  hasPlaidItems,
		HasWyreAccountPrerequisitesMet: profile.HasWyreAccountPreconditionsMet(),
		// todo: implement first
		HasWyreAccount: hasWyreAccount,
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

// OneTimePasscodeVerify is an rpc handler -- should maybe be renamed to login?
func (s *Server) OneTimePasscodeVerify(ctx context.Context, req *proto.OneTimePasscodeVerifyRequest) (*proto.OneTimePasscodeVerifyResponse, error) {
	const unknownMsg = "An unknown error occurred. Please try again later."

	loginKind, loginValue, err := ValidateAndNormalizeLogin(req.EmailOrPhone)
	if err != nil {
		return nil, err
	}

	passcode, err := s.Db.AckOneTimePasscode(ctx, loginValue, req.Code)
	if err != nil {
		return nil, err
	}
	if passcode == nil {
		return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedOTP(loginKind))
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
	userID := GetUserIDFromIncomingContext(ctx)
	if userID == "" {
		return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
	}

	err := req.Validate()
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

	userID := GetUserIDFromIncomingContext(ctx)
	if userID == "" {
		return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
	}

	/*** TEST ***/
	go func() {
		time.Sleep(5 * time.Second)

		pusherClient := pusher.Client{
			AppID:   "1171786",
			Key:     "dd280d42ccafc24e19ff",
			Secret:  "d8cfa16565ede2ae414d",
			Cluster: "us3",
			Secure:  true,
		}

		data := map[string]string{"message": "hello world"}
		err := pusherClient.Trigger(string(userID), "my-event", data)
		if err != nil {
			log.Println(err)
		}
	}()

	{
		accounts, err := s.Db.GetWyreAccounts(ctx, nil, userID)
		if err != nil {
			return nil, err
		}
		if len(accounts) > 0 {
			account := accounts[0]
			wyreAcct, err := s.Wyre.GetAccount(account.SecretKey, string(account.ID))
			if err != nil {
				return nil, err
			}
			log.Printf("WYRE_ACCOUNT %#v\n", wyreAcct)
		}
	}
	/*** TEST ***/

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
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
	}

	err = req.Validate()
	if err != nil {
		return nil, err
	}

	err = s.Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		profile, err := s.Db.GetAllProfileData(ctx, tx, u.ID)
		if err != nil {
			return err
		}

		for _, kind := range common.ProfileDataKinds {
			existingProfileData := profile.FilterKind(kind).First()

			switch kind {
			case common.KindLegalName:
				var legalNameData *legalname.ProfileDataLegalName

				if req.LegalName != "" {
					if existingProfileData == nil {
						legalNameData = &legalname.ProfileDataLegalName{
							ID:        common.ProfileDataID(shortuuid.New()),
							Status:    common.StatusReceived,
							LegalName: req.LegalName,
							CreatedAt: time.Now(),
						}
					} else {
						legalNameData = (*existingProfileData).(*legalname.ProfileDataLegalName)

						now := time.Now()

						legalNameData.LegalName = req.LegalName
						legalNameData.UpdatedAt = &now
					}

					_, err := s.Db.SaveProfileData(ctx, tx, u.ID, *legalNameData)
					if err != nil {
						return err
					}
				}
			case common.KindDateOfBirth:
				var dobData *dateofbirth.ProfileDataDateOfBirth

				if req.DateOfBirth != "" {
					if existingProfileData == nil {
						dobData = &dateofbirth.ProfileDataDateOfBirth{
							ID:          common.ProfileDataID(shortuuid.New()),
							Status:      common.StatusReceived,
							DateOfBirth: req.DateOfBirth,
							CreatedAt:   time.Now(),
						}
					} else {
						dobData = (*existingProfileData).(*dateofbirth.ProfileDataDateOfBirth)

						now := time.Now()

						dobData.DateOfBirth = req.DateOfBirth
						dobData.UpdatedAt = &now
					}

					_, err := s.Db.SaveProfileData(ctx, tx, u.ID, *dobData)
					if err != nil {
						return err
					}
				}
			case common.KindUSSSN:
				var ssnData *ssn.ProfileDataSSN

				if req.Ssn != "" {
					if existingProfileData == nil {
						ssnData = &ssn.ProfileDataSSN{
							ID:        common.ProfileDataID(shortuuid.New()),
							Status:    common.StatusReceived,
							SSN:       req.Ssn,
							CreatedAt: time.Now(),
						}
					} else {
						ssnData = (*existingProfileData).(*ssn.ProfileDataSSN)

						now := time.Now()

						ssnData.SSN = req.Ssn
						ssnData.UpdatedAt = &now
					}

					_, err := s.Db.SaveProfileData(ctx, tx, u.ID, *ssnData)
					if err != nil {
						return err
					}
				}
			case common.KindAddress:
				var addressData *address.ProfileDataAddress

				if req.Address != nil {
					if existingProfileData == nil {
						addressData = &address.ProfileDataAddress{
							ID:         common.ProfileDataID(shortuuid.New()),
							Status:     common.StatusReceived,
							Street1:    req.Address.Street_1,
							Street2:    req.Address.Street_2,
							City:       req.Address.City,
							State:      req.Address.State,
							PostalCode: req.Address.PostalCode,
							Country:    req.Address.Country,
							CreatedAt:  time.Now(),
						}
					} else {
						addressData = (*existingProfileData).(*address.ProfileDataAddress)

						now := time.Now()

						addressData.Street1 = req.Address.Street_1
						addressData.Street2 = req.Address.Street_2
						addressData.City = req.Address.City
						addressData.State = req.Address.State
						addressData.PostalCode = req.Address.PostalCode
						addressData.Country = req.Address.Country

						addressData.UpdatedAt = &now
					}

					_, err := s.Db.SaveProfileData(ctx, tx, u.ID, *addressData)
					if err != nil {
						return err
					}
				}
			case common.KindUSGovernmentID:
				var governmentIDData *usgovernmentid.ProfileDataUSGovernmentID

				if req.UsGovernmentIdDoc != nil {
					if req.UsGovernmentIdDoc.Kind == proto.UsGovernmentIdDocumentInputKind_GI_UNKNOWN {
						return status.Errorf(codes.InvalidArgument, "government id document kind needs to be specified ")
					}
					kind := usgovernmentid.KindFromGovernmentIdDocKind(req.UsGovernmentIdDoc.Kind)

					if len(req.UsGovernmentIdDoc.FileIds) != kind.FilesRequired() {
						return status.Errorf(codes.InvalidArgument, fmt.Sprintf("%s requires %d files to be attached to its input", kind, kind.FilesRequired()))
					}

					for _, fileID := range req.UsGovernmentIdDoc.FileIds {
						meta, err := s.Db.GetFileMetadata(ctx, u.ID, file.ID(fileID))
						if err != nil {
							return err
						}
						if meta == nil {
							return status.Errorf(codes.InvalidArgument, "one or more file ids is invalid")
						}
					}

					if existingProfileData == nil {
						governmentIDData = &usgovernmentid.ProfileDataUSGovernmentID{
							ID:               common.ProfileDataID(shortuuid.New()),
							Status:           common.StatusReceived,
							GovernmentIDKind: kind,
							FileIDs:          []file.ID{},
							CreatedAt:        time.Now(),
						}
					} else {
						governmentIDData = (*existingProfileData).(*usgovernmentid.ProfileDataUSGovernmentID)

						now := time.Now()

						fileIDs := []file.ID{}
						for _, id := range req.UsGovernmentIdDoc.FileIds {
							fileIDs = append(fileIDs, file.ID(id))
						}

						governmentIDData.GovernmentIDKind = usgovernmentid.Kind(req.UsGovernmentIdDoc.Kind)
						governmentIDData.FileIDs = fileIDs
						governmentIDData.UpdatedAt = &now
					}
					_, err := s.Db.SaveProfileData(ctx, tx, u.ID, *governmentIDData)
					if err != nil {
						return err
					}
				}
			case common.KindPhone:
				// do nothing we don't accept input from here (we get it from our user record, and stamp it out from there because its verified)
			case common.KindEmail:
				// do nothing we don't accept input from here (we get it from our user record, and stamp it out from there because its verified)
			default:
				panic(fmt.Sprintf("handlers.SaveProfileData: unhandled profile data kind: %s", kind))
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	profile, err := s.Db.GetAllProfileData(ctx, nil, u.ID)
	if err != nil {
		return nil, err
	}

	existingWyreAccounts, err := s.Db.GetWyreAccounts(ctx, nil, u.ID)
	if err != nil {
		return nil, err
	}

	if len(existingWyreAccounts) == 0 && profile.HasWyreAccountPreconditionsMet() {
		log.Printf("Creating new wyre account for user id: %s", u.ID)

		account, err := s.WyreManager.CreateAccount(ctx, u.ID, profile)
		if err != nil {
			return nil, err
		}
		existingWyreAccounts = append(existingWyreAccounts, account)

		// todo add payment methods
	}

	if !profile.HasWyreAccountPreconditionsMet() {
		log.Printf("Preconditions for wyre are unmet for user id: %s", u.ID)
	}

	if len(existingWyreAccounts) == 0 {
		return &proto.ProfileDataInfo{
			Profile: profile.GetProfileDataItemInfo(),
		}, nil
	} else {
		// todo: add remediations to this resp somehow (gotta think it thru)

		log.Printf("Wyre account found for user id: %s, %#v", u.ID, existingWyreAccounts[0])

		return &proto.ProfileDataInfo{
			Profile: profile.GetProfileDataItemInfo(),
			Wyre: &proto.ThirdPartyUserAccount{
				Status: existingWyreAccounts[0].Status,
			},
		}, nil
	}
}

// ViewerProfileData is an rpc handler
func (s *Server) ViewerProfileData(ctx context.Context, _ *emptypb.Empty) (*proto.ProfileDataInfo, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
	}

	profile, err := s.Db.GetAllProfileData(ctx, nil, u.ID)
	if err != nil {
		return nil, err
	}

	return &proto.ProfileDataInfo{
		Profile: profile.GetProfileDataItemInfo(),
	}, nil
}

// WyreCreateAccount is an rpc handler
func (s *Server) WyreCreateAccount(ctx context.Context, req *proto.WyreCreateAccountRequest) (*proto.WyreCreateAccountResponse, error) {
	userID := GetUserIDFromIncomingContext(ctx)
	if userID == "" {
		return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
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

// ChangeViewerEmail is an rpc handler
func (s *Server) ChangeViewerEmail(ctx context.Context, req *proto.ChangeViewerEmailRequest) (*emptypb.Empty, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
	}

	valueKind, newEmailValue, err := ValidateAndNormalizeLogin(req.Email)
	if err != nil {
		return nil, err
	}
	if valueKind != onetimepasscode.LoginKindEmail {
		return nil, status.Errorf(codes.InvalidArgument, "a valid email address must be provided")
	}

	passcode, err := s.Db.AckOneTimePasscode(ctx, newEmailValue, req.Code)
	if err != nil {
		return nil, err
	}
	if passcode == nil {
		return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedOTP(valueKind))
	}
	if passcode.EmailOrPhone != newEmailValue {
		return nil, status.Errorf(codes.InvalidArgument, "The code provided does not correlate with the desired email")
	}

	// everything checks out; modify the user and save with the new email address value
	now := time.Now()
	u.Email = &newEmailValue
	u.EmailVerifiedAt = &now

	err = s.Db.SaveUser(ctx, nil, u)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unknown, "An unknown error ocurred; please try again.")
	}

	return &emptypb.Empty{}, nil
}

// ChangeViewerPhone is an rpc handler
func (s *Server) ChangeViewerPhone(ctx context.Context, req *proto.ChangeViewerPhoneRequest) (*emptypb.Empty, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
	}

	valueKind, newPhoneValue, err := ValidateAndNormalizeLogin(req.Phone)
	if err != nil {
		return nil, err
	}
	if valueKind != onetimepasscode.LoginKindPhone {
		return nil, status.Errorf(codes.InvalidArgument, "a valid phone must be provided")
	}

	passcode, err := s.Db.AckOneTimePasscode(ctx, newPhoneValue, req.Code)
	if err != nil {
		return nil, err
	}
	if passcode == nil {
		return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedOTP(valueKind))
	}
	if passcode.EmailOrPhone != newPhoneValue {
		return nil, status.Errorf(codes.InvalidArgument, "The code provided does not correlate with the desired phone")
	}

	// everything checks out; modify the user and save with the new phone value
	now := time.Now()
	u.Phone = &newPhoneValue
	u.PhoneVerifiedAt = &now

	err = s.Db.SaveUser(ctx, nil, u)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unknown, "An unknown error ocurred; please try again.")
	}

	return &emptypb.Empty{}, nil
}

// UploadFile is an rpc handler
func (s *Server) UploadFile(ctx context.Context, req *proto.UploadFileRequest) (*proto.UploadFileResponse, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
	}

	log.Println(req.Filename)
	log.Println("       mime type:", req.MimeType)
	log.Println(" size (reported):", req.Size)
	log.Println(" size (verified):", len(req.Body))

	fileID, err := s.FileManager.UploadEncryptedFile(ctx, u.ID, req)
	if err != nil {
		return nil, err
	}

	return &proto.UploadFileResponse{
		FileId: string(fileID),
	}, nil
}

// GetImage is an rpc handler
func (s *Server) GetImage(ctx context.Context, req *proto.GetImageRequest) (*proto.GetImageResponse, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
	}

	f, err := s.FileManager.GetFile(ctx, u.ID, file.ID(req.FileId))
	if err != nil {
		return nil, err
	}

	var out []byte
	var width, height int32

	if req.ProcessingMode == proto.ImageProcessingMode_IP_FIT {
		if req.Width == 0 || req.Height == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "in fit mode you must provide width and height")
		}

		img, err := imaging.Decode(bytes.NewReader(*f.Body))
		if err != nil {
			return nil, err
		}

		fitted := imaging.Fit(img, int(req.Width), int(req.Height), imaging.MitchellNetravali)
		width = int32(fitted.Rect.Dx())
		height = int32(fitted.Rect.Dy())

		var buf bytes.Buffer
		err = imaging.Encode(&buf, fitted, imaging.JPEG)
		if err != nil {
			return nil, err
		}

		out = buf.Bytes()
	}

	if req.ProcessingMode == proto.ImageProcessingMode_IP_RESIZE {
		return nil, status.Errorf(codes.Unimplemented, "resize mode is not implemented yet")
	}

	return &proto.GetImageResponse{
		MimeType: "image/jpeg",
		Size:     int32(len(out)),
		Body:     out,
		Width:    width,
		Height:   height,
	}, nil
}

// GetImage is an rpc handler
func (s *Server) WyreWebhook(ctx context.Context, req *proto.WyreWebhookRequest) (*emptypb.Empty, error) {
	// todo: auth the webhook?!?
	log.Printf("WyreWebhook %#v", req)
	return &emptypb.Empty{}, nil
}
