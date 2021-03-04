package server

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/plaid/plaid-go/plaid"
	"github.com/rs/xid"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	faker "github.com/bxcodec/faker/v3"
	"github.com/khoerling/flux/api/lib/auth"
	"github.com/khoerling/flux/api/lib/db/models/onetimepasscode"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/plaid/item"
	proto "github.com/khoerling/flux/api/lib/protocol"
)

// https://api.sendwyre.com/v3/rates?as=priced

// UserData is an rpc handler
func (s *Server) UserData(ctx context.Context, in *proto.UserDataRequest) (*proto.UserDataResponse, error) {
	log.Printf("Received: %v", in)

	httpResp := &proto.UserDataResponse{
		User: &proto.User{
			Id:    xid.New().String(),
			Email: faker.Email(),
			Phone: faker.Phonenumber(),
			Organizations: []*proto.Organization{
				{
					Id:   xid.New().String(),
					Name: fmt.Sprintf("%s %s Inc.", faker.LastName(), faker.Word()),
				},
			},
		},
	}

	return httpResp, nil
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

	invalidMsg := "The email code provided was not valid. Please try again."

	unknownMsg := "An unknown error occurred. Please try again later."

	passcode, err := passcodes.Next()
	if err == iterator.Done {
		return nil, status.Errorf(codes.Unauthenticated, invalidMsg)
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
		respUser.Email = *u.Email
	}

	if u.Phone != nil {
		respUser.Phone = *u.Phone
	}

	return &proto.OneTimePasscodeVerifyResponse{
		Jwt:  jwt,
		User: respUser,
	}, nil
}

// PlaidConnectBankAccounts is an rpc handler
func (s *Server) PlaidConnectBankAccounts(ctx context.Context, req *proto.PlaidConnectBankAccountsRequest) (*proto.PlaidConnectBankAccountsResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, codes.Unauthenticated.String())
	}

	vals := md.Get("user-id")

	var userID user.ID
	if len(vals) > 0 {
		userID = user.ID(vals[0])
	} else {
		return nil, status.Errorf(codes.Unauthenticated, codes.Unauthenticated.String())
	}

	if userID == "" {
		return nil, status.Errorf(codes.Unauthenticated, codes.Unauthenticated.String())
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
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, codes.Unauthenticated.String())
	}

	vals := md.Get("user-id")

	var userID string
	if len(vals) > 0 {
		userID = vals[0]
	} else {
		return nil, status.Errorf(codes.Unauthenticated, codes.Unauthenticated.String())
	}

	if userID == "" {
		return nil, status.Errorf(codes.Unauthenticated, codes.Unauthenticated.String())
	}

	log.Printf("Generating Plaid Link Token for User ID: %s", userID)

	u, err := s.Db.GetUserByID(ctx, userID)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unknown, "An unknown error ocurred; please try again.")
	}
	if u == nil {
		return nil, status.Errorf(codes.NotFound, "Your user was not found. Plaid Link token could not be created.")
	}

	plaidUserDetails := plaid.LinkTokenUser{
		ClientUserID: userID,
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
