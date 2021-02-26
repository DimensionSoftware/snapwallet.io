package server

import (
	"context"
	"fmt"
	"log"
	mrand "math/rand"
	"strings"
	"time"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/grpc/metadata"

	faker "github.com/bxcodec/faker/v3"
	"github.com/khoerling/flux/api/lib/auth"
	"github.com/khoerling/flux/api/lib/db/models/onetimepasscode"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	proto "github.com/khoerling/flux/api/lib/protocol"
)

// https://api.sendwyre.com/v3/rates?as=priced

// UserData is an rpc handler
func (s *Server) UserData(ctx context.Context, in *proto.UserDataRequest) (*proto.UserDataResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("no grpc metadata")
	}
	auths := md.Get("authorization")
	if len(auths) < 1 {
		return nil, fmt.Errorf("grpc authorization empty")
	}
	auth := auths[0]
	log.Printf("Auth: %v", auth)

	in.ProtoMessage()
	log.Printf("Received: %v", in)

	httpResp := &proto.UserDataResponse{
		User: &proto.User{
			Id:    mrand.Int63(),
			Email: faker.Email(),
			Phone: faker.Phonenumber(),
			Organizations: []*proto.Organization{
				{
					Id:   mrand.Int63(),
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

	if loginKind == onetimepasscode.LoginKindPhone {
		return nil, fmt.Errorf("phone is not implemented yet")
	}

	otp, err := s.Db.CreateOneTimePasscode(ctx, loginValue, loginKind)
	if err != nil {
		return nil, err
	}

	msg := generateOtpMessage(mail.NewEmail("Customer", loginValue), otp.Code)

	_, err = s.SendgridClient.Send(msg)
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

	// if no doc then failure to login (make better later)
	passcode, err := passcodes.Next()
	if err != nil {
		return nil, fmt.Errorf("code verification failed")
	}

	_, err = passcode.Ref.Delete(ctx)
	if err != nil {
		return nil, fmt.Errorf("code verification failed")
	}

	u, err := s.Db.GetOrCreateUser(ctx, loginKind, loginValue)
	if err != nil {
		return nil, err
	}

	jwt, err := s.JwtSigner.Sign(auth.NewClaims(u.ID))
	if err != nil {
		return nil, err
	}

	return &proto.OneTimePasscodeVerifyResponse{
		Jwt: jwt,
	}, nil
}

// WyreAddBankPaymentMethod is an rpc handler
func (s *Server) WyreAddBankPaymentMethod(ctx context.Context, req *proto.WyreAddBankPaymentMethodRequest) (*proto.WyreAddBankPaymentMethodResponse, error) {
	accessToken := "a-token"
	accountID := "balderdash"

	processorTokenResp, err := s.Plaid.CreateProcessorToken(accessToken, accountID, "wyre")
	if err != nil {
		return nil, err
	}

	_, err = s.Wyre.CreatePaymentMethod(wyre.CreatePaymentMethodRequest{
		PlaidProcessorToken: processorTokenResp.ProcessorToken,
	}.WithDefaults())
	if err != nil {
		return nil, err
	}

	return &proto.WyreAddBankPaymentMethodResponse{}, nil
}

func generateOtpMessage(to *mail.Email, code string) *mail.SGMailV3 {
	from := mail.NewEmail("Ctulhu", "ctulhu@dreamcodez.cc")
	subject := "Your one time passcode for flux"
	plainTextContent := fmt.Sprintf("Your one time passcode is: %s", code)
	htmlContent := fmt.Sprintf("Your one time passcode is: <strong>%s</strong>", code)
	return mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
}
