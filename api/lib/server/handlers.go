package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/plaid/plaid-go/plaid"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/disintegration/imaging"
	"github.com/khoerling/flux/api/lib/db/models/gotoconfig"
	"github.com/khoerling/flux/api/lib/db/models/job"
	"github.com/khoerling/flux/api/lib/db/models/onetimepasscode"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/file"
	"github.com/khoerling/flux/api/lib/db/models/user/plaid/item"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/address"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/dateofbirth"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/legalname"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/proofofaddress"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/ssn"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/usgovernmentid"
	"github.com/khoerling/flux/api/lib/db/models/user/transaction"
	"github.com/khoerling/flux/api/lib/db/models/user/wyre/account"
	"github.com/khoerling/flux/api/lib/db/models/user/wyre/paymentmethod"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	proto "github.com/khoerling/flux/api/lib/protocol"

	"github.com/lithammer/shortuuid/v3"
	"github.com/teris-io/shortid"
)

// https://api.sendwyre.com/v3/rates?as=priced

// ViewerData is an rpc handler
func (s *Server) ViewerData(ctx context.Context, _ *emptypb.Empty) (*proto.ViewerDataResponse, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
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
	var wyreAccountID account.ID

	{
		accounts, err := s.Db.GetWyreAccounts(ctx, nil, u.ID)
		if err != nil {
			return nil, err
		}
		if len(accounts) > 0 {
			hasWyreAccount = true
			wyreAccountID = accounts[0].ID
		}
	}

	var hasWyrePaymentMethods bool

	if hasWyreAccount {
		pms, err := s.Db.GetWyrePaymentMethods(ctx, nil, u.ID, wyreAccountID)
		if err != nil {
			return nil, err
		}
		if len(pms) > 0 {
			hasWyrePaymentMethods = true
		}
	}

	flags := proto.UserFlags{
		HasPlaidItems:                  hasPlaidItems,
		HasWyreAccountPrerequisitesMet: profile.HasWyreAccountPreconditionsMet(),
		HasWyreAccount:                 hasWyreAccount,
		HasWyrePaymentMethods:          hasWyrePaymentMethods,
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
		from := s.TwilioConfig.PhoneNumber
		to := loginValue
		message := fmt.Sprintf("Your one time passcode for flux is: %s", otp.Code)

		_, _, err := s.Twilio.SendSMS(from, to, message, "", "")
		if err != nil {
			return nil, err
		}

		return &proto.OneTimePasscodeResponse{}, nil
	}

	msg, err := generateOtpMessage(mail.NewEmail("Customer", loginValue), otp.Code)

	if err != nil {
		return nil, err
	}

	_, err = s.SendEmail.Send(msg)
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

		// @chris; i figured out how to add structured details to grpc errors; details can be any proto message:
		//
		// status, err := status.New(codes.Unauthenticated, genMsgUnauthenticatedOTP(loginKind)).WithDetails(&proto.Address{})
		// if err != nil {
		// 	return nil, err
		// }

		// return nil, status.Err()
	}

	u, err := s.Db.GetOrCreateUser(ctx, loginKind, loginValue)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unauthenticated, unknownMsg)
	}

	tokenMaterial, err := s.AuthManager.NewTokenMaterial(time.Now(), u.ID, string(u.ID))
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
		Tokens: tokenMaterial,
		User:   respUser,
	}, nil
}

func (s *Server) TokenExchange(ctx context.Context, req *proto.TokenExchangeRequest) (*proto.TokenExchangeResponse, error) {
	material, err := s.AuthManager.TokenExchange(ctx, time.Now(), req.RefreshToken)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
	}

	return &proto.TokenExchangeResponse{
		Tokens: material,
	}, nil
}

// RPC handler for connecting a Wyre ACH account using their Plaid integration
func (s *Server) WyreConnectBankAccount(ctx context.Context, req *proto.WyreConnectBankAccountRequest) (*proto.WyrePaymentMethod, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
	}

	if req.PlaidPublicToken == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a valid Plaid token.")
	}

	if req.PlaidAccountId == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a valid bank account ID.")
	}

	accounts, err := s.Db.GetWyreAccounts(ctx, nil, u.ID)
	if err != nil || len(accounts) <= 0 {
		log.Printf("Could not find account for %s", u.ID)
		return nil, status.Error(codes.NotFound, "Please verify your identity before connecting a bank account.")
	}

	userAccount := accounts[0]

	wyrePublicToken := req.PlaidPublicToken + "|" + req.PlaidAccountId
	wyreReqParams := wyre.CreateWyrePaymentMethodRequest{
		// NOTE: both attrs are required. See struct definition for further explanation.
		PublicToken:       wyrePublicToken,
		PlaidPublicToken:  wyrePublicToken,
		PaymentMethodType: "LOCAL_TRANSFER",
		Country:           "US",
	}
	res, err := s.Wyre.CreateWyrePaymentMethod(userAccount.SecretKey, wyreReqParams)

	if err != nil {
		log.Printf("Error creating Wyre payment method")
		return nil, err
	}

	hookResponse, err := s.Wyre.SubscribeWebhook(userAccount.SecretKey, "paymentmethod:"+string(res.ID), string(s.APIHost)+"/wyre/hooks/"+string(u.ID))
	if err != nil {
		return nil, err
	}
	log.Printf("hook response from wyre: %#v", hookResponse)

	pm := paymentmethod.PaymentMethod{
		ID:                    paymentmethod.ID(res.ID),
		PlaidItemID:           "",
		PlaidAccountID:        item.AccountID(req.PlaidAccountId),
		Status:                res.Status,
		Name:                  res.Name,
		Last4:                 res.Last4Digits,
		ChargeableCurrencies:  res.ChargeableCurrencies,
		DepositableCurrencies: res.DepositableCurrencies,
		UpdatedAt:             time.Now(),
		CreatedAt:             time.Now(),
	}
	err = s.Db.SaveWyrePaymentMethod(ctx, nil, u.ID, userAccount.ID, &pm)

	if err != nil {
		log.Printf("Error saving Wyre payment method")
		return nil, status.Error(codes.Internal, "An error occurred while connecting your account. Please try again.")
	}

	var lifecycleStatus proto.LifecycleStatus
	pmStatus := strings.ToLower(res.Status)
	if pmStatus == "active" {
		lifecycleStatus = proto.LifecycleStatus_L_CREATED
	} else {
		lifecycleStatus = proto.LifecycleStatus_L_PENDING
	}

	return &proto.WyrePaymentMethod{
		Id:                    res.ID,
		Status:                res.Status,
		Name:                  res.Name,
		Last4:                 res.Last4Digits,
		ChargeableCurrencies:  res.ChargeableCurrencies,
		DepositableCurrencies: res.DepositableCurrencies,
		LifecycleStatus:       lifecycleStatus,
	}, nil
}

// PlaidConnectBankAccounts is an rpc handler
func (s *Server) PlaidConnectBankAccounts(ctx context.Context, req *proto.PlaidConnectBankAccountsRequest) (*proto.PlaidConnectBankAccountsResponse, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
	}

	err = req.Validate()
	if err != nil {
		return nil, err
	}

	existingPlaidItems, err := s.Db.GetAllPlaidItems(ctx, nil, u.ID)
	if err != nil {
		return nil, err
	}

	resp, err := s.Plaid.ExchangePublicToken(req.PlaidPublicToken)
	if err != nil {
		return nil, err
	}
	log.Printf("Plaid Public Token successfuly exchanged")

	var accounts []item.Account
	for _, reqAccount := range req.Accounts {
		var alreadyExists bool

		for _, epi := range existingPlaidItems {
			// https://plaid.com/docs/link/duplicate-items/#preventing-duplicate-item-adds-with-onsuccess
			// You can compare a combination of the accounts’ institution_id, account name, and account mask
			// to determine whether your user has previously linked their account to your application.
			if item.InstitutionID(req.Institution.Id) == epi.Institution.ID {
				for _, epiAccount := range epi.Accounts {
					if reqAccount.Name == epiAccount.Name && reqAccount.Mask == epiAccount.Mask {
						alreadyExists = true
						break
					}
				}
			}
		}

		if !alreadyExists {
			accounts = append(accounts, item.Account{
				ID:      item.AccountID(reqAccount.Id),
				Name:    reqAccount.Name,
				Mask:    reqAccount.Mask,
				Type:    reqAccount.Type,
				SubType: reqAccount.SubType,
			})
		}
	}

	if len(accounts) == 0 {
		// no new plaid accounts were provided (duplicates are not allowed)
		return &proto.PlaidConnectBankAccountsResponse{}, nil
	}

	item := item.Item{
		ID:          item.ID(resp.ItemID),
		AccessToken: resp.AccessToken,
		Institution: item.Institution{
			ID:   item.InstitutionID(req.Institution.Id),
			Name: req.Institution.Name,
		},
		Accounts:  accounts,
		CreatedAt: time.Now(),
	}
	err = s.Db.SavePlaidItem(ctx, u.ID, &item)
	if err != nil {
		return nil, err
	}
	log.Printf("Plaid ItemID %s saved", resp.ItemID)

	// submit wyre create payment methods job; will still need wyre account prereq but this is the other update path when they add new accounts too
	{
		now := time.Now()

		err = s.JobPublisher.PublishJob(ctx, &job.Job{
			ID:         shortuuid.New(),
			Kind:       job.KindCreateWyrePaymentMethodsForUser,
			Status:     job.StatusQueued,
			RelatedIDs: []string{string(u.ID)},
			CreatedAt:  now.Unix(),
			UpdatedAt:  now.Unix(),
		})
		if err != nil {
			log.Println(err)
		}
	}

	return &proto.PlaidConnectBankAccountsResponse{}, nil
}

func generateOtpMessage(to *mail.Email, code string) (*mail.SGMailV3, error) {
	htmlContent, err := genEmailTemplate("otpHTML", EmailTemplateVars{OTPCode: code})

	if err != nil {
		return nil, err
	}

	from := mail.NewEmail("Snap Wallet", "support@snapwallet.io")
	subject := "Security Code"
	plainTextContent := fmt.Sprintf("Your security code is: %s", code)
	return mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent), nil
}

func newMailWithAttachments(to *mail.Email, attachments []*mail.Attachment, plain string, html string) *mail.SGMailV3 {
	from := mail.NewEmail("Snap Wallet", "support@snapwallet.io")

	m := mail.NewV3Mail()

	p := mail.NewPersonalization()
	p.AddTos(to)

	cplain := mail.NewContent("text/plain", plain)
	chtml := mail.NewContent("text/html", html)

	m.SetFrom(from)
	m.AddPersonalizations(p)

	for _, attachment := range attachments {
		m.AddAttachment(attachment)
	}

	m.AddContent(cplain)
	m.AddContent(chtml)

	return m
}

func generateTransferMessage(to *mail.Email, t *wyre.TransferDetail) (*mail.SGMailV3, error) {
	htmlContent, err := genEmailTemplate("newTransactionHTML", EmailTemplateVars{TransactionID: string(t.ID), BusinessDays: 5})

	if err != nil {
		return nil, err
	}

	from := mail.NewEmail("Snap Wallet", "support@snapwallet.io")
	subject := fmt.Sprintf("Transaction Created")
	plainTextContent := fmt.Sprintf("Transaction %s created successfully", t.ID)
	return mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent), nil
}

func generateTransactionStatusMessage(to *mail.Email, t *wyre.TransferDetail) (*mail.SGMailV3, error) {
	txnStatus := strings.ToLower(t.Status)
	var status string

	if txnStatus == "failed" {
		status = "Failed"
	}
	if txnStatus == "completed" {
		status = "Completed"
	}

	htmlContent, err := genEmailTemplate("transactionStatusHTML", EmailTemplateVars{
		TransactionID: string(t.ID),
		Status:        status,
	})

	if err != nil {
		return nil, err
	}

	from := mail.NewEmail("Snap Wallet", "support@snapwallet.io")
	subject := fmt.Sprintf("Transaction %s", status)
	plainTextContent := fmt.Sprintf("Your transaction status has been changed to %s", status)
	return mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent), nil
}

func generateDebitCardTransactionMessage(to *mail.Email, order *wyre.WalletOrder) (*mail.SGMailV3, error) {
	htmlContent, err := genEmailTemplate("newTransactionHTML", EmailTemplateVars{TransactionID: order.ID, BusinessDays: 1})

	if err != nil {
		return nil, err
	}

	from := mail.NewEmail("Snap Wallet", "support@snapwallet.io")
	subject := fmt.Sprintf("Transaction Created")
	plainTextContent := fmt.Sprintf("Transaction %s created successfully", order.ID)
	return mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent), nil
}

// PlaidCreateLinkToken is an rpc handler
func (s *Server) PlaidCreateLinkToken(ctx context.Context, req *proto.PlaidCreateLinkTokenRequest) (*proto.PlaidCreateLinkTokenResponse, error) {

	userID := GetUserIDFromIncomingContext(ctx)
	if userID == "" {
		return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
	}

	/*** TEST ***/
	{
		accounts, err := s.Db.GetWyreAccounts(ctx, nil, userID)
		if err != nil {
			return nil, err
		}
		if len(accounts) > 0 {
			a := accounts[0]
			wyreAcct, err := s.Wyre.GetAccount(a.SecretKey, wyre.AccountID(a.ID))
			if err != nil {
				return nil, err
			}
			log.Printf("WYRE_ACCOUNT %#v\n", wyreAcct)

			pms, err := s.Db.GetWyrePaymentMethods(ctx, nil, userID, account.ID(wyreAcct.ID))
			if err != nil {
				return nil, err
			}
			for _, pm := range pms {
				theirPm, err := s.Wyre.GetPaymentMethod(a.SecretKey, wyre.PaymentMethodID(pm.ID))
				if err != nil {
					return nil, err
				}
				log.Printf("WYRE_PAYMENT_METHOD %#v\n", theirPm)

			}
		}
	}
	/*** TEST ***/

	log.Printf("Generating Plaid Link Token for User ID: %s", userID)

	u, err := s.Db.GetUserByID(ctx, nil, user.ID(userID))
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

	profile, err := s.Db.GetAllProfileData(ctx, nil, u.ID)
	if err != nil {
		return nil, err
	}

	for _, kind := range common.ProfileDataKinds {
		existingProfileData := profile.FilterKind(kind).First()

		switch kind {
		case common.KindLegalName:
			err = s.Db.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
				var legalNameData *legalname.ProfileDataLegalName

				if req.LegalName != "" {
					if existingProfileData == nil {
						legalNameData = &legalname.ProfileDataLegalName{
							CommonProfileData: common.CommonProfileData{
								ID:        common.ProfileDataID(shortuuid.New()),
								Status:    common.StatusReceived,
								CreatedAt: time.Now(),
							},
							LegalName: req.LegalName,
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
				return nil
			})
		case common.KindDateOfBirth:
			err = s.Db.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
				var dobData *dateofbirth.ProfileDataDateOfBirth

				if req.DateOfBirth != "" {
					if existingProfileData == nil {
						dobData = &dateofbirth.ProfileDataDateOfBirth{
							CommonProfileData: common.CommonProfileData{
								ID:        common.ProfileDataID(shortuuid.New()),
								Status:    common.StatusReceived,
								CreatedAt: time.Now(),
							},
							DateOfBirth: req.DateOfBirth,
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
				return nil
			})
		case common.KindUSSSN:
			err = s.Db.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
				var ssnData *ssn.ProfileDataSSN

				if req.Ssn != "" {
					if existingProfileData == nil {
						ssnData = &ssn.ProfileDataSSN{
							CommonProfileData: common.CommonProfileData{
								ID:        common.ProfileDataID(shortuuid.New()),
								Status:    common.StatusReceived,
								CreatedAt: time.Now(),
							},
							SSN: req.Ssn,
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
				return nil
			})
		case common.KindAddress:
			err = s.Db.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
				var addressData *address.ProfileDataAddress

				if req.Address != nil {
					if existingProfileData == nil {
						addressData = &address.ProfileDataAddress{
							CommonProfileData: common.CommonProfileData{
								ID:        common.ProfileDataID(shortuuid.New()),
								Status:    common.StatusReceived,
								CreatedAt: time.Now(),
							},
							Street1:    req.Address.Street_1,
							Street2:    req.Address.Street_2,
							City:       req.Address.City,
							State:      req.Address.State,
							PostalCode: req.Address.PostalCode,
							Country:    req.Address.Country,
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
				return nil
			})
		case common.KindProofOfAddressDoc:
			err = s.Db.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
				var proofOfAddressData *proofofaddress.ProfileDataProofOfAddressDoc

				if req.ProofOfAddressDoc != nil {
					if len(req.ProofOfAddressDoc.FileIds) == 0 {
						return status.Errorf(codes.InvalidArgument, "at least one file id must be attached")
					}

					fileIDs := []file.ID{}
					for _, fileID := range req.ProofOfAddressDoc.FileIds {
						fileIDs = append(fileIDs, file.ID(fileID))
						meta, err := s.Db.GetFileMetadata(ctx, u.ID, file.ID(fileID))
						if err != nil {
							return err
						}
						if meta == nil {
							return status.Errorf(codes.InvalidArgument, "one or more file ids is invalid")
						}
					}

					if existingProfileData == nil {
						proofOfAddressData = &proofofaddress.ProfileDataProofOfAddressDoc{
							CommonProfileData: common.CommonProfileData{
								ID:        common.ProfileDataID(shortuuid.New()),
								Status:    common.StatusReceived,
								CreatedAt: time.Now(),
							},
							FileIDs: fileIDs,
						}
					} else {
						proofOfAddressData = (*existingProfileData).(*proofofaddress.ProfileDataProofOfAddressDoc)

						now := time.Now()

						proofOfAddressData.FileIDs = fileIDs
						proofOfAddressData.UpdatedAt = &now
					}
					_, err := s.Db.SaveProfileData(ctx, tx, u.ID, *proofOfAddressData)
					if err != nil {
						return err
					}
				}
				return nil
			})
		case common.KindUSGovernmentIDDoc:
			err = s.Db.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
				var governmentIDData *usgovernmentid.ProfileDataUSGovernmentIDDoc

				if req.UsGovernmentIdDoc != nil {
					if req.UsGovernmentIdDoc.Kind == proto.UsGovernmentIdDocumentInputKind_GI_UNKNOWN {
						return status.Errorf(codes.InvalidArgument, "government id document kind needs to be specified ")
					}
					kind := usgovernmentid.KindFromUsGovernmentIdDocumentInputKind(req.UsGovernmentIdDoc.Kind)

					if len(req.UsGovernmentIdDoc.FileIds) != kind.FilesRequired() {
						return status.Errorf(codes.InvalidArgument, fmt.Sprintf("%s requires %d files to be attached to its input", kind, kind.FilesRequired()))
					}

					fileIDs := []file.ID{}
					for _, fileID := range req.UsGovernmentIdDoc.FileIds {
						fileIDs = append(fileIDs, file.ID(fileID))
						meta, err := s.Db.GetFileMetadata(ctx, u.ID, file.ID(fileID))
						if err != nil {
							return err
						}
						if meta == nil {
							return status.Errorf(codes.InvalidArgument, "one or more file ids is invalid")
						}
					}

					if existingProfileData == nil {
						governmentIDData = &usgovernmentid.ProfileDataUSGovernmentIDDoc{
							CommonProfileData: common.CommonProfileData{
								ID:        common.ProfileDataID(shortuuid.New()),
								Status:    common.StatusReceived,
								CreatedAt: time.Now(),
							},
							GovernmentIDKind: kind,
							FileIDs:          fileIDs,
						}
					} else {
						governmentIDData = (*existingProfileData).(*usgovernmentid.ProfileDataUSGovernmentIDDoc)

						now := time.Now()

						governmentIDData.GovernmentIDKind = kind
						governmentIDData.FileIDs = fileIDs
						governmentIDData.UpdatedAt = &now
					}
					_, err := s.Db.SaveProfileData(ctx, tx, u.ID, *governmentIDData)
					if err != nil {
						return err
					}
				}
				return nil
			})
		case common.KindPhone:
			// do nothing we don't accept input from here (we get it from our user record, and stamp it out from there because its verified)
		case common.KindEmail:
			// do nothing we don't accept input from here (we get it from our user record, and stamp it out from there because its verified)
		default:
			panic(fmt.Sprintf("handlers.SaveProfileData: unhandled profile data kind: %s", kind))
		}

		if err != nil {
			return nil, err
		}
	}

	profile, err = s.Db.GetAllProfileData(ctx, nil, u.ID)
	if err != nil {
		return nil, err
	}

	existingWyreAccounts, err := s.Db.GetWyreAccounts(ctx, nil, u.ID)
	if err != nil {
		return nil, err
	}

	if !profile.HasWyreAccountPreconditionsMet() {
		log.Printf("Preconditions for wyre are unmet for user id: %s", u.ID)
	}

	remediations, err := s.RemedyManager.GetRemediationsProto(u.ID, profile)
	if err != nil {
		return nil, err
	}

	resp := &proto.ProfileDataInfo{
		Profile:      profile.GetProfileDataItemInfo(),
		Remediations: remediations,
	}

	if profile.HasWyreAccountPreconditionsMet() {
		// todo, create job in db
		// todo make sure theres not a job already running
		log.Printf("Creating new wyre account for user id: %s", u.ID)

		now := time.Now()

		err = s.JobPublisher.PublishJob(ctx, &job.Job{
			ID:         shortuuid.New(),
			Kind:       job.KindUpdateWyreAccountForUser,
			Status:     job.StatusQueued,
			RelatedIDs: []string{string(u.ID)},
			CreatedAt:  now.Unix(),
			UpdatedAt:  now.Unix(),
		})
		if err != nil {
			return nil, err
		}

		// todo: store pending lifecycle status? or can use job submitted information
		resp.Wyre = &proto.ThirdPartyUserAccount{
			LifecycleStatus: proto.LifecycleStatus_L_PENDING,
		}
	} else if len(existingWyreAccounts) == 0 {
		job, err := s.Db.GetJobByKindAndStatusAndRelatedId(ctx, job.KindUpdateWyreAccountForUser, job.StatusQueued, string(u.ID))
		if err != nil {
			return nil, err
		}
		if job != nil {
			resp.Wyre = &proto.ThirdPartyUserAccount{
				LifecycleStatus: proto.LifecycleStatus_L_PENDING,
			}
		}
	}

	if len(existingWyreAccounts) > 0 {
		resp.Wyre = &proto.ThirdPartyUserAccount{
			// todo: store created lifecycle status?
			LifecycleStatus: proto.LifecycleStatus_L_CREATED,
			Status:          existingWyreAccounts[0].Status,
			// todo: remediations
		}
	}

	return resp, nil
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

	var wyre *proto.ThirdPartyUserAccount
	{
		existingWyreAccounts, err := s.Db.GetWyreAccounts(ctx, nil, u.ID)
		if err != nil {
			return nil, err
		}
		if len(existingWyreAccounts) > 0 {
			wa := existingWyreAccounts[0]

			wyre = &proto.ThirdPartyUserAccount{
				LifecycleStatus: proto.LifecycleStatus_L_CREATED,
				Status:          wa.Status,
				// todo: remediations
			}
		} else {
			job, err := s.Db.GetJobByKindAndStatusAndRelatedId(ctx, job.KindUpdateWyreAccountForUser, job.StatusQueued, string(u.ID))
			if err != nil {
				return nil, err
			}
			if job != nil {
				wyre = &proto.ThirdPartyUserAccount{
					LifecycleStatus: proto.LifecycleStatus_L_PENDING,
				}
			}
		}
	}

	remediations, err := s.RemedyManager.GetRemediationsProto(u.ID, profile)
	if err != nil {
		return nil, err
	}

	return &proto.ProfileDataInfo{
		Profile:      profile.GetProfileDataItemInfo(),
		Wyre:         wyre,
		Remediations: remediations,
	}, nil
}

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
	err = s.Db.UpdateEmail(ctx, u.ID, newEmailValue)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unknown, "An unknown error ocurred; please try again.")
	}

	now := time.Now()
	err = s.JobPublisher.PublishJob(ctx, &job.Job{
		ID:         shortuuid.New(),
		Kind:       job.KindUpdateWyreAccountForUser,
		Status:     job.StatusQueued,
		RelatedIDs: []string{string(u.ID)},
		CreatedAt:  now.Unix(),
		UpdatedAt:  now.Unix(),
	})
	if err != nil {
		return nil, err
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
	err = s.Db.UpdatePhone(ctx, u.ID, newPhoneValue)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unknown, "An unknown error ocurred; please try again.")
	}

	now := time.Now()
	err = s.JobPublisher.PublishJob(ctx, &job.Job{
		ID:         shortuuid.New(),
		Kind:       job.KindUpdateWyreAccountForUser,
		Status:     job.StatusQueued,
		RelatedIDs: []string{string(u.ID)},
		CreatedAt:  now.Unix(),
		UpdatedAt:  now.Unix(),
	})
	if err != nil {
		return nil, err
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
		Size:   req.Size,
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
	// todo, check api and store status update/updated at in our db
	log.Printf("WyreWebhook %#v", req)

	now := time.Now()
	userID := user.ID(req.HookId)

	parts := strings.Split(req.Trigger, ":")
	objectKind := parts[0]
	objectID := parts[1]

	var msg *pusher.Message
	switch objectKind {
	case "account":
		msg = &pusher.Message{
			Kind: pusher.MessageKindWyreAccountUpdated,
			At:   now,
		}

		// todo: tx
		ourWyreAccounts, err := s.Db.GetWyreAccounts(ctx, nil, userID)
		if err != nil {
			log.Printf("failure getting our wyre accounts: %#v", err)
			return nil, status.Errorf(codes.Unknown, "hook failed")
		}

		var ourWyreAccount *account.Account
		for _, wa := range ourWyreAccounts {
			if wa.ID == account.ID(objectID) {
				ourWyreAccount = wa
				break
			}
		}
		if ourWyreAccount == nil {
			log.Printf("wyre account not found: %s\n", objectID)
			return nil, status.Errorf(codes.FailedPrecondition, "hook failed")
		}

		theirAccount, err := s.Wyre.GetAccount(ourWyreAccount.SecretKey, wyre.AccountID(ourWyreAccount.ID))
		if err != nil {
			log.Printf("failure getting wyre account from them: %#v", err)
			return nil, status.Errorf(codes.Unknown, "hook failed")
		}

		now := time.Now()

		ourWyreAccount.Status = theirAccount.Status
		ourWyreAccount.UpdatedAt = &now

		err = s.Db.SaveWyreAccount(ctx, nil, userID, ourWyreAccount)
		if err != nil {
			log.Printf("failure saving our wyre account: %#v\n", err)
			return nil, status.Errorf(codes.Unknown, "hook failed")
		}

		profile, err := s.Db.GetAllProfileData(ctx, nil, userID)
		if err != nil {
			log.Printf("failure getting our profile data: %#v\n", err)
			return nil, status.Errorf(codes.Unknown, "hook failed")
		}

		for _, pf := range theirAccount.ProfileFields {
			newStatus := WyreProfileFieldStatusToProfileDataStatus(pf.Status)
			switch pf.FieldID {
			case string(wyre.ProfileFieldIDIndividualLegalName):
				for _, legalName := range profile.FilterKindLegalName() {
					log.Printf("updating legal name status to: %s", newStatus)

					now = time.Now()
					legalName.Status = newStatus
					legalName.UpdatedAt = &now
					legalName.Note = pf.Note

					_, err := s.Db.SaveProfileData(ctx, nil, userID, legalName)
					if err != nil {
						log.Printf("failure saving legal name profile data: %#v\n", err)
						return nil, status.Errorf(codes.Unknown, "hook failed")
					}

					break
				}
			case string(wyre.ProfileFieldIDIndividualEmail):
				for _, email := range profile.FilterKindEmail() {
					if email.Email != pf.Value {
						continue
					}

					log.Printf("updating email status for %s to: %s", email.Email, newStatus)

					now = time.Now()
					email.Status = newStatus
					email.UpdatedAt = &now
					email.Note = pf.Note

					_, err := s.Db.SaveProfileData(ctx, nil, userID, email)
					if err != nil {
						log.Printf("failure saving email profile data: %#v\n", err)
						return nil, status.Errorf(codes.Unknown, "hook failed")
					}

					break
				}
			case string(wyre.ProfileFieldIDIndividualCellphoneNumber):
				for _, phone := range profile.FilterKindPhone() {
					if phone.Phone != pf.Value {
						continue
					}

					log.Printf("updating phone status for %s to: %s", phone.Phone, newStatus)

					now = time.Now()
					phone.Status = newStatus
					phone.UpdatedAt = &now
					phone.Note = pf.Note

					_, err := s.Db.SaveProfileData(ctx, nil, userID, phone)
					if err != nil {
						log.Printf("failure saving phone profile data: %#v\n", err)
						return nil, status.Errorf(codes.Unknown, "hook failed")
					}

					break
				}
			case string(wyre.ProfileFieldIDIndividualDateOfBirth):
				for _, dob := range profile.FilterKindDateOfBirth() {
					log.Printf("updating dob status to: %s", newStatus)

					now = time.Now()
					dob.Status = newStatus
					dob.UpdatedAt = &now
					dob.Note = pf.Note

					_, err := s.Db.SaveProfileData(ctx, nil, userID, dob)
					if err != nil {
						log.Printf("failure saving dob profile data: %#v\n", err)
						return nil, status.Errorf(codes.Unknown, "hook failed")
					}

					break
				}
			case string(wyre.ProfileFieldIDIndividualSSN):
				for _, ssn := range profile.FilterKindSSN() {
					log.Printf("updating ssn status to: %s", newStatus)

					now = time.Now()
					ssn.Status = newStatus
					ssn.UpdatedAt = &now
					ssn.Note = pf.Note

					_, err := s.Db.SaveProfileData(ctx, nil, userID, ssn)
					if err != nil {
						log.Printf("failure saving ssn profile data: %#v\n", err)
						return nil, status.Errorf(codes.Unknown, "hook failed")
					}

					break
				}
			case string(wyre.ProfileFieldIDIndividualResidenceAddress):
				for _, addr := range profile.FilterKindAddress() {
					log.Printf("updating address status to: %s", newStatus)

					now = time.Now()
					addr.Status = newStatus
					addr.UpdatedAt = &now
					addr.Note = pf.Note

					_, err := s.Db.SaveProfileData(ctx, nil, userID, addr)
					if err != nil {
						log.Printf("failure saving address profile data: %#v\n", err)
						return nil, status.Errorf(codes.Unknown, "hook failed")
					}

					break
				}
			case string(wyre.ProfileFieldIDIndividualGovernmentID):
				for _, govtid := range profile.FilterKindUSGovernmentIDDoc() {
					log.Printf("updating government id status to: %s", newStatus)

					now = time.Now()
					govtid.Status = newStatus
					govtid.UpdatedAt = &now
					govtid.Note = pf.Note

					_, err := s.Db.SaveProfileData(ctx, nil, userID, govtid)
					if err != nil {
						log.Printf("failure saving government id profile data: %#v\n", err)
						return nil, status.Errorf(codes.Unknown, "hook failed")
					}

					break
				}
			default:
				log.Printf("Unhandled profile field id on webhook update %s\n", pf.FieldID)
			}
		}
	case "paymentmethod":
		msg = &pusher.Message{
			Kind: pusher.MessageKindWyrePaymentMethodsUpdated,
			IDs:  []string{objectID},
			At:   now,
		}

		ourWyreAccounts, err := s.Db.GetWyreAccounts(ctx, nil, userID)
		if err != nil {
			log.Printf("failure getting our wyre accounts: %#v", err)
			return nil, status.Errorf(codes.Unknown, "hook failed")
		}
		if len(ourWyreAccounts) == 0 {
			log.Printf("wyre account not found for user")
			return nil, status.Errorf(codes.FailedPrecondition, "hook failed")
		}
		ourWyreAccount := ourWyreAccounts[0]

		ourWyrePaymentMethods, err := s.Db.GetWyrePaymentMethods(ctx, nil, userID, ourWyreAccount.ID)
		if err != nil {
			log.Printf("failure getting our wyre payment methods: %#v", err)
			return nil, status.Errorf(codes.FailedPrecondition, "hook failed")
		}

		var ourWyrePaymentMethod *paymentmethod.PaymentMethod
		for _, pm := range ourWyrePaymentMethods {
			if pm.ID == paymentmethod.ID(objectID) {
				ourWyrePaymentMethod = pm
				break
			}
		}
		if ourWyrePaymentMethod == nil {
			log.Printf("wyre payment method not found for user")
			return nil, status.Errorf(codes.FailedPrecondition, "hook failed")
		}

		theirPaymentMethod, err := s.Wyre.GetPaymentMethod(ourWyreAccount.SecretKey, wyre.PaymentMethodID(ourWyrePaymentMethod.ID))
		if err != nil {
			log.Printf("failure getting wyre payment method from them: %#v", err)
			return nil, status.Errorf(codes.Unknown, "hook failed")
		}

		now := time.Now()

		ourWyrePaymentMethod.Status = theirPaymentMethod.Status
		ourWyrePaymentMethod.UpdatedAt = now

		err = s.Db.SaveWyrePaymentMethod(ctx, nil, userID, ourWyreAccount.ID, ourWyrePaymentMethod)
		if err != nil {
			log.Printf("failure saving our wyre payment method: %#v", err)
			return nil, status.Errorf(codes.Unknown, "hook failed")
		}
	case "transfer":
		swUser, err := s.Db.GetUserByID(ctx, nil, userID)

		accounts, err := s.Db.GetWyreAccounts(ctx, nil, swUser.ID)
		if err != nil {
			log.Printf("Error retrieving Wyre accounts from DB for user %s", swUser.ID)
			return nil, status.Error(codes.Internal, "hook failed")
		}

		if len(accounts) <= 0 {
			log.Printf("wyre account not found for user")
			return nil, status.Error(codes.Unknown, "hook failed")
		}

		userAccount := accounts[0]

		// Make sure it exists in our db
		_, err = s.Db.GetTransactionByExternalId(ctx, nil, swUser.ID, transaction.ExternalID(objectID))

		if err != nil {
			log.Printf("Error retrieving transaction %s", objectID)
			return nil, status.Error(codes.Internal, "hook failed")
		}

		txn, err := s.Wyre.GetTransfer(userAccount.SecretKey, objectID)
		if err != nil {
			return nil, status.Error(codes.Unknown, "hook failed")
		}

		// This should never happen.
		// We don't want to send email when status is the same
		if txn.Status == "PENDING" {
			break
		}

		log.Printf("sending email for transaction status update")
		emailMsg, err := generateTransactionStatusMessage(mail.NewEmail("Customer", *swUser.Email), txn)

		if err != nil {
			return nil, err
		}

		_, err = s.SendEmail.Send(emailMsg)
		if err != nil {
			log.Printf("Error sending email")
			return nil, err
		}

		msg = &pusher.Message{
			Kind: pusher.MessageKindWyreTransferUpdated,
			IDs:  []string{objectID},
			At:   now,
		}
	default:
		log.Printf("UNIMPLEMENTED PHP WEBHOOK: %s %s", userID, req.Trigger)
		return &emptypb.Empty{}, nil
	}

	if msg != nil {
		err := s.Pusher.Send(userID, msg)
		if err != nil {
			log.Printf("Sending Pusher notification failed")
			return nil, err
		}
	}

	return &emptypb.Empty{}, nil
}

func WyreProfileFieldStatusToProfileDataStatus(wyreProfileFieldStatus string) common.ProfileDataStatus {
	switch wyreProfileFieldStatus {
	case "OPEN":
		return common.StatusInvalid
	case "PENDING":
		return common.StatusPending
	case "APPROVED":
		return common.StatusApproved
	}

	// should never get here
	panic("wyreProfileFieldStatus unknown when WyreProfileFieldStatusToProfileDataStatus(...) was called")
}

func (s *Server) WyreGetPaymentMethods(ctx context.Context, _ *emptypb.Empty) (*proto.WyrePaymentMethods, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
	}

	var wyreAccountID account.ID
	{
		accounts, err := s.Db.GetWyreAccounts(ctx, nil, u.ID)
		if err != nil {
			return nil, err
		}
		if len(accounts) > 0 {
			wyreAccountID = accounts[0].ID
		}

	}

	var out []*proto.WyrePaymentMethod

	var pms []*paymentmethod.PaymentMethod
	if wyreAccountID != "" {
		pms, err = s.Db.GetWyrePaymentMethods(ctx, nil, u.ID, wyreAccountID)
		if err != nil {
			return nil, err
		}

		for _, pm := range pms {
			out = append(out, &proto.WyrePaymentMethod{
				LifecycleStatus:       proto.LifecycleStatus_L_CREATED,
				Id:                    string(pm.ID),
				Status:                pm.Status,
				Name:                  pm.Name,
				Last4:                 pm.Last4,
				ChargeableCurrencies:  pm.ChargeableCurrencies,
				DepositableCurrencies: pm.DepositableCurrencies,
			})
		}
	}

	pitems, err := s.Db.GetAllPlaidItems(ctx, nil, u.ID)
	if err != nil {
		return nil, err
	}

	for _, plaidItem := range pitems {
		pmCreated := false
		for _, pm := range pms {
			if item.ID(pm.PlaidItemID) == plaidItem.ID {
				pmCreated = true
				break
			}
		}

		if !pmCreated {
			for _, account := range plaidItem.Accounts {
				out = append(out, &proto.WyrePaymentMethod{
					LifecycleStatus: proto.LifecycleStatus_L_PENDING,
					Name:            fmt.Sprintf("%s (%s)", account.Name, plaidItem.Institution.Name),
					Last4:           account.Mask,
				})
			}
		}
	}

	return &proto.WyrePaymentMethods{
		PaymentMethods: out,
	}, nil
}

func (s *Server) WyreCreateTransfer(ctx context.Context, req *proto.WyreCreateTransferRequest) (*proto.WyreTransferDetail, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
	}

	var wyreAccount *account.Account
	{
		accounts, err := s.Db.GetWyreAccounts(ctx, nil, u.ID)
		if err != nil {
			return nil, err
		}
		if len(accounts) > 0 {
			wyreAccount = accounts[0]
		}
	}
	if wyreAccount == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "you must have a wyre account to create a transfer")
	}

	pms, err := s.Db.GetWyrePaymentMethods(ctx, nil, u.ID, wyreAccount.ID)
	if err != nil {
		return nil, err
	}

	var source *paymentmethod.PaymentMethod
	for _, pm := range pms {
		if pm.ID == paymentmethod.ID(req.Source) {
			source = pm
		}
	}
	if source == nil {
		return nil, status.Errorf(codes.InvalidArgument, "source %s is not a payment method ID belonging to this user", req.Source)
	}

	// TODO: more validation
	wyreReq := wyre.CreateTransferRequest{
		Source:         "paymentmethod:" + string(source.ID),
		SourceCurrency: "USD",
		Dest:           req.Dest,
		DestCurrency:   req.DestCurrency,
		Message:        "TODO",

		DestAmount:   req.GetDestAmount(),
		SourceAmount: req.GetSourceAmount(),
	}.WithDefaults()

	t, err := s.Wyre.CreateTransfer(wyreAccount.SecretKey, wyreReq)
	if err != nil {
		var wyreErr *wyre.APIError
		if errors.As(err, &wyreErr) {
			return nil, status.Errorf(codes.Unknown, wyreErr.Message)
		}

		log.Printf("unknown wyre error: %#v", err)
		return nil, status.Error(codes.Unknown, "Unknown error while contacting wyre.")
	}

	trx := transaction.Transaction{
		Kind: transaction.KindACH,
	}.WithDefaults().EnrichWithWyreTransferDetail(t)
	trx.Status = transaction.StatusQuoted
	err = s.Db.SaveTransaction(ctx, nil, u.ID, &trx)
	if err != nil {
		return nil, err
	}

	return wyre.WyreTransferDetailToProto(t), nil
}

func (s *Server) WyreConfirmTransfer(ctx context.Context, req *proto.WyreConfirmTransferRequest) (*proto.WyreTransferDetail, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
	}

	if req.TransferId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "a valid transfer ID is required")
	}

	var wyreAccount *account.Account
	{
		accounts, err := s.Db.GetWyreAccounts(ctx, nil, u.ID)
		if err != nil {
			return nil, err
		}
		if len(accounts) > 0 {
			wyreAccount = accounts[0]
		}
	}
	if wyreAccount == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "you must have a wyre account to confirm a transfer")
	}

	t, err := s.Wyre.ConfirmTransfer(wyreAccount.SecretKey, wyre.ConfirmTransferRequest{
		TransferId: req.TransferId,
	})
	if err != nil {
		return nil, err
	}

	err = s.Db.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		existingTrx, err := s.Db.GetTransactionByExternalId(ctx, tx, u.ID, transaction.ExternalID(req.TransferId))
		if err != nil {
			return err
		}

		if existingTrx == nil {
			//return status.Errorf(codes.NotFound, "existing transaction not found")
			log.Println("existing transaction not found during WyreConfirmTransfer; cannot update transaction")
			return nil
		}

		trx := existingTrx.EnrichWithWyreTransferDetail(t)
		trx.Status = transaction.StatusConfirmed

		err = s.Db.SaveTransaction(ctx, tx, u.ID, &trx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	hookResponse, err := s.Wyre.SubscribeWebhook(wyreAccount.SecretKey, "transfer:"+string(t.ID), string(s.APIHost)+"/wyre/hooks/"+string(u.ID))
	if err != nil {
		return nil, err
	}
	log.Printf("hook response from wyre: %#v", hookResponse)

	// send email
	msg, err := generateTransferMessage(mail.NewEmail("Customer", *u.Email), t)
	if err != nil {
		return nil, err
	}

	_, err = s.SendEmail.Send(msg)
	if err != nil {
		return nil, err
	}

	// TODO: store info in db about xfer
	fmt.Printf("Wyre transfer confirmation response: %#v", t)

	return wyre.WyreTransferDetailToProto(t), nil
}

func (s *Server) WidgetGetShortUrl(ctx context.Context, req *proto.SnapWidgetConfig) (*proto.WidgetGetShortUrlResponse, error) {
	shortID, err := shortid.Generate()
	if err != nil {
		return nil, err
	}

	var wallets []gotoconfig.SnapWidgetWallet
	for _, reqWallet := range req.Wallets {
		wallets = append(wallets, gotoconfig.SnapWidgetWallet{
			Asset:   reqWallet.Asset,
			Address: reqWallet.Address,
		})
	}

	swc := gotoconfig.SnapWidgetConfig{
		AppName: req.AppName,
		Wallets: wallets,
		Intent:  req.Intent,
		Focus:   req.Focus,
		Theme:   req.Theme,
	}

	id, err := swc.GetID()
	if err != nil {
		return nil, err
	}

	g := gotoconfig.Config{
		ID:      id,
		ShortID: gotoconfig.ShortID(shortID),
		Config:  &swc,
	}

	if req.Product != nil {
		swc.Product = &gotoconfig.SnapWidgetProduct{
			ImageURL:           req.Product.Image_URL,
			VideoURL:           req.Product.Video_URL,
			DestinationAmount:  req.Product.DestinationAmount,
			DestinationTicker:  req.Product.DestinationTicker,
			DestinationAddress: req.Product.DestinationAddress,
			Title:              req.Product.Title,
			Author:             req.Product.Author,
		}
	}

	shortid, err := s.Db.SaveGotoConfig(ctx, &g)
	if err != nil {
		return nil, err
	}

	return &proto.WidgetGetShortUrlResponse{
		Url: fmt.Sprintf("%s/g/%s", s.APIHost, shortid),
	}, nil
}

func (s *Server) Goto(ctx context.Context, req *proto.GotoRequest) (*proto.GotoResponse, error) {
	// lookup by shortid
	g, err := s.Db.GetGotoConfigByShortID(ctx, gotoconfig.ShortID(req.Id))
	if err != nil {
		return nil, err
	}

	if g == nil {
		return nil, status.Errorf(codes.NotFound, "goto ID not found")
	}

	configJsonBytes, err := json.Marshal(g.Config)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("config", string(configJsonBytes))
	params.Add("ts", fmt.Sprintf("%d", time.Now().Unix()))

	return &proto.GotoResponse{
		Location: string(s.WebHost) + "/widget/?" + params.Encode(),
	}, nil

}

func (s *Server) WyreGetTransfers(ctx context.Context, req *proto.WyreGetTransfersRequest) (*proto.WyreTransfers, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
	}

	var wyreAccount *account.Account
	{
		accounts, err := s.Db.GetWyreAccounts(ctx, nil, u.ID)
		if err != nil {
			return nil, err
		}
		if len(accounts) > 0 {
			wyreAccount = accounts[0]
		}
	}

	if wyreAccount == nil {
		return &proto.WyreTransfers{}, nil
	}

	history, err := s.Wyre.GetTransferHistory(wyreAccount.SecretKey, req.Page*30, 30)
	if err != nil {
		return nil, err
	}

	var out []*proto.WyreTransfer
	for _, transfer := range history.Transfers {
		out = append(out, wyre.WyreTransferToProto(&transfer))
	}

	return &proto.WyreTransfers{
		Transfers: out,
	}, nil
}

func (s *Server) WyreGetTransfer(ctx context.Context, req *proto.WyreGetTransferRequest) (*proto.WyreTransferDetail, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
	}

	var wyreAccount *account.Account
	{
		accounts, err := s.Db.GetWyreAccounts(ctx, nil, u.ID)
		if err != nil {
			return nil, err
		}
		if len(accounts) > 0 {
			wyreAccount = accounts[0]
		}
	}

	if wyreAccount == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "wyre account must exist to retrieve a transfer")
	}

	t, err := s.Wyre.GetTransfer(wyreAccount.SecretKey, req.TransferId)
	if err != nil {
		return nil, err
	}

	return wyre.WyreTransferDetailToProto(t), nil
}

func (s *Server) WyreCreateDebitCardQuote(ctx context.Context, req *proto.WyreCreateDebitCardQuoteRequest) (*proto.WyreCreateDebitCardQuoteResponse, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
	}

	var dest string
	// Wyre only supports bitcoin or erc20 but expects this prefix
	if strings.ToLower(req.DestCurrency) == "btc" {
		dest = "bitcoin:" + req.Dest
	} else {
		dest = "ethereum:" + req.Dest
	}

	reqData := wyre.CreateWalletOrderReservationRequest{
		Country:            req.Country,
		PaymentMethod:      "debit-card",
		SourceCurrency:     req.SourceCurrency,
		DestCurrency:       req.DestCurrency,
		SourceAmount:       req.SourceAmount,
		LockFields:         req.LockFields,
		Dest:               dest,
		AmountIncludesFees: &req.AmountIncludesFees,
	}

	if req.SourceAmount > 0 {
		reqData.SourceAmount = req.SourceAmount
		reqData.DestAmount = 0
	} else {
		reqData.SourceAmount = 0
		reqData.DestAmount = req.DestAmount
	}

	// Create the order reservation
	createReservationResponse, err := s.Wyre.CreateWalletOrderReservation(reqData)

	if err != nil {
		return nil, err
	}

	trx := transaction.Transaction{
		Kind: transaction.KindDebit,
	}.WithDefaults().EnrichWithCreateWalletOrderReservationResponse(createReservationResponse)
	trx.Status = transaction.StatusQuoted

	// Get the order reservation details because why would they return them in the previous call? :(
	reservationResponse, err := s.Wyre.GetWalletOrderReservation(wyre.GetWalletOrderReservationRequest{
		ReservationID: createReservationResponse.Reservation,
	})
	if err != nil {
		return nil, err
	}

	trx = trx.EnrichWithWalletOrderReservation(reservationResponse)

	err = s.Db.SaveTransaction(ctx, nil, u.ID, &trx)
	if err != nil {
		return nil, err
	}
	return &proto.WyreCreateDebitCardQuoteResponse{
		ReservationId: createReservationResponse.Reservation,
		Quote: &proto.WyreWalletOrderReservationQuote{
			// Return non prefixed dest
			Dest:                    req.Dest,
			ExchangeRate:            reservationResponse.Quote.ExchangeRate,
			DestCurrency:            reservationResponse.Quote.DestCurrency,
			SourceCurrency:          reservationResponse.Quote.SourceCurrency,
			Fees:                    reservationResponse.Quote.Fees,
			SourceAmount:            reservationResponse.Quote.SourceAmount,
			DestAmount:              reservationResponse.Quote.DestAmount,
			SourceAmountWithoutFees: reservationResponse.Quote.SourceAmountWithoutFees,
			// Wyre doesn't return this for some reason but 10 minutes is the amount of time
			// Remove a minute for latency and whatnot
			ExpiresAt: time.Now().Add(time.Minute * 9).Format(time.RFC3339),
		},
	}, nil
}

func (s *Server) WyreConfirmDebitCardQuote(ctx context.Context, req *proto.WyreConfirmDebitCardQuoteRequest) (*proto.WyreConfirmDebitCardQuoteResponse, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)

	if err != nil {
		return nil, err
	}

	var dest string
	// Wyre only supports bitcoin or erc20 but expects this prefix
	if strings.ToLower(req.DestCurrency) == "btc" {
		dest = "bitcoin:" + req.Dest
	} else {
		dest = "ethereum:" + req.Dest
	}

	card := req.Card

	orderRequest := wyre.CreateWalletOrderRequest{
		ReservationID:  req.ReservationId,
		SourceCurrency: req.SourceCurrency,
		PurchaseAmount: req.SourceAmount,
		DestCurrency:   req.DestCurrency,
		SourceAmount:   req.SourceAmount,
		Dest:           dest,
		FirstName:      card.FirstName,
		LastName:       card.LastName,
		Email:          *u.Email,
		PhoneNumber:    card.PhoneNumber,
		ReferenceID:    "crypto_moon_lambo",
		Address: wyre.WalletOrderAddress{
			Street1:    card.Address.Street_1,
			Street2:    card.Address.Street_2,
			City:       card.Address.City,
			State:      card.Address.State,
			PostalCode: card.Address.PostalCode,
			Country:    card.Address.Country,
		},
		DebitCard: wyre.WalletOrderDebitCard{
			Number:           card.Number,
			ExpirationMonth:  card.ExpirationMonth,
			ExpirationYear:   card.ExpirationYear,
			VerificationCode: card.VerificationCode,
		},
	}

	// Create the order
	orderResponse, err := s.Wyre.CreateWalletOrder(orderRequest)
	if err != nil {
		return nil, err
	}

	err = s.Db.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		existingTrx, err := s.Db.GetTransactionByExternalId(ctx, tx, u.ID, transaction.ExternalID(req.ReservationId))
		if err != nil {
			return err
		}

		if existingTrx == nil {
			//return status.Errorf(codes.NotFound, "existing transaction not found")
			log.Println("existing transaction not found during WyreConfirmDebitCardQuote; cannot update transaction")
			return nil
		}

		trx := *existingTrx
		trx = trx.EnrichWithCreateWalletOrderRequest(&orderRequest)
		trx = trx.EnrichWithWalletOrder(orderResponse)

		trx.Status = transaction.StatusConfirmed

		err = s.Db.SaveTransaction(ctx, tx, u.ID, &trx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &proto.WyreConfirmDebitCardQuoteResponse{
		OrderId:    orderResponse.ID,
		Status:     orderResponse.Status,
		TransferId: orderResponse.TransferID,
	}, nil
}

func (s *Server) WyreGetDebitCardAuthorizations(ctx context.Context, req *proto.WyreGetDebitCardOrderAuthorizationsRequest) (*proto.WyreGetDebitCardOrderAuthorizationsResponse, error) {
	res, err := s.Wyre.GetWalletOrderAuthorizations(wyre.GetWalletOrderAuthorizationsRequest{
		OrderID: req.OrderId,
	})

	if err != nil {
		return nil, err
	}

	return &proto.WyreGetDebitCardOrderAuthorizationsResponse{
		WalletOrderId: res.WalletOrderID,
		SmsNeeded:     res.SMSNeeded,
		Card2FaNeeded: res.Card2faNeeded,
	}, nil
}

func (s *Server) WyreSubmitDebitCardAuthorizations(ctx context.Context, req *proto.WyreSubmitDebitCardOrderAuthorizationsRequest) (*proto.WyreSubmitDebitCardOrderAuthorizationsResponse, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
	}

	var verificationType string

	if req.Card_2FaCode != "" && req.Sms_2FaCode == "" {
		verificationType = "CARD"
	}
	if req.Card_2FaCode == "" && req.Sms_2FaCode != "" {
		verificationType = "SMS"
	}
	if req.Card_2FaCode != "" && req.Sms_2FaCode != "" {
		verificationType = "ALL"
	}

	res, err := s.Wyre.SubmitWalletOrderAuthorizations(wyre.SubmitWalletOrderAuthorizationsRequest{
		WalletOrderID: wyre.WalletOrderID(req.OrderId),
		Type:          verificationType,
		Reservation:   req.ReservationId,
		SMS:           req.Sms_2FaCode,
		Card2fa:       req.Card_2FaCode,
	})

	if err != nil {
		return nil, err
	}

	// Send email
	msg, err := generateDebitCardTransactionMessage(mail.NewEmail("Customer", *u.Email), &wyre.WalletOrder{ID: req.OrderId})
	if err != nil {
		return nil, err
	}

	_, err = s.SendEmail.Send(msg)
	if err != nil {
		return nil, err
	}

	return &proto.WyreSubmitDebitCardOrderAuthorizationsResponse{
		Success: res.Success,
	}, nil
}

func (s *Server) Geo(ctx context.Context, _ *emptypb.Empty) (*proto.GeoResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	vals := md.Get("cf-ipcountry")
	val := ""
	if len(vals) > 0 {
		val = vals[0]
	}
	log.Printf("%#v\n", md)
	return &proto.GeoResponse{
		Country: val,
	}, nil
}

func (s *Server) GetTransactions(ctx context.Context, req *proto.GetTransactionsRequest) (*proto.Transactions, error) {
	u, err := RequireUserFromIncomingContext(ctx, s.Db)
	if err != nil {
		return nil, err
	}

	// todo: pagination
	transactions, err := s.Db.GetTransactions(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	return transactions.AsProto(), nil
}
