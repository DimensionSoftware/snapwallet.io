package server

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	mrand "math/rand"
	"strings"
	"time"

	resty "github.com/go-resty/resty/v2"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/badoux/checkmail"
	faker "github.com/bxcodec/faker/v3"
	proto "github.com/khoerling/flux/api/lib/protocol"
	"github.com/nyaruka/phonenumbers"
)

// https://api.sendwyre.com/v3/rates?as=priced

// UserData is an rpc handler
func (s *Server) UserData(ctx context.Context, in *proto.UserDataRequest) (*proto.UserDataResponse, error) {
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

type wyrePricingRate map[string]float32
type wyrePricingRates = map[string](wyrePricingRate)

// PricingData is an rpc handler
func (s *Server) PricingData(ctx context.Context, in *proto.PricingDataRequest) (*proto.PricingDataResponse, error) {
	client := resty.New()
	pricingResp, err := client.R().
		SetResult(wyrePricingRates{}).
		EnableTrace().
		Get("https://api.sendwyre.com/v3/rates?as=priced")

	if err != nil {
		return nil, err
	}

	wyreRates := pricingResp.Result().(*wyrePricingRates)

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

// 1. firestore store otp by email or phone code and send to it
// 2. otp code entered on page

// OneTimePasscode is an rpc handler
func (s *Server) OneTimePasscode(ctx context.Context, req *proto.OneTimePasscodeRequest) (*proto.OneTimePasscodeResponse, error) {
	var normalizedEmailOrPhone string
	var isPhone bool

	num, err := phonenumbers.Parse(strings.TrimSpace(req.EmailOrPhone), "US")
	if err == nil {
		isPhone = true
		normalizedEmailOrPhone = phonenumbers.Format(num, phonenumbers.E164)
	} else {
		err = checkmail.ValidateFormat(req.EmailOrPhone)
		if err == nil {
			err = checkmail.ValidateHost(req.EmailOrPhone)
			if err == nil {
				normalizedEmailOrPhone = strings.TrimSpace(req.EmailOrPhone)
			} else {
				return nil, fmt.Errorf("a valid phone number or email is required")
			}
		} else {
			return nil, fmt.Errorf("a valid phone number or email is required")
		}
	}

	code, err := sixRandomDigits()
	if err != nil {
		return nil, err
	}

	if isPhone {
		return nil, fmt.Errorf("phone is not implemented yet")
	}

	msg := generateOtpMessage(mail.NewEmail("Matt", normalizedEmailOrPhone), code)

	_, _, err = s.Firestore.Collection("one-time-passcodes").Add(ctx, map[string]interface{}{
		"emailOrPhone": normalizedEmailOrPhone,
		"code":         code,
		"createdAt":    time.Now(),
	})
	if err != nil {
		return nil, err
	}

	_, err = s.SendgridClient.Send(msg)
	if err != nil {
		return nil, err
	}

	return &proto.OneTimePasscodeResponse{}, nil
}

// OneTimePasscodeVerify is an rpc handler
func (s *Server) OneTimePasscodeVerify(ctx context.Context, in *proto.OneTimePasscodeVerifyRequest) (*proto.OneTimePasscodeVerifyResponse, error) {
	return &proto.OneTimePasscodeVerifyResponse{}, nil
}

func generateOtpMessage(to *mail.Email, code string) *mail.SGMailV3 {
	from := mail.NewEmail("Ctulhu", "ctulhu@dreamcodez.cc")
	subject := "Your one time passcode for flux"
	plainTextContent := fmt.Sprintf("Your one time passcode is: %s", code)
	htmlContent := fmt.Sprintf("Your one time passcode is: <strong>%s</strong>", code)
	return mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
}

func sixRandomDigits() (string, error) {
	max := big.NewInt(999999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d\n", n.Int64()), nil
}
