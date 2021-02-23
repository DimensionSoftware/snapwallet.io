package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"

	resty "github.com/go-resty/resty/v2"

	faker "github.com/bxcodec/faker/v3"
	proto "github.com/khoerling/flux/api/lib/protocol"
)

// https://api.sendwyre.com/v3/rates?as=priced

// UserData is an rpc handler
func (s *Server) UserData(ctx context.Context, in *proto.UserDataRequest) (*proto.UserDataResponse, error) {
	log.Printf("Received: %v", in)

	httpResp := &proto.UserDataResponse{
		User: &proto.User{
			Id:    rand.Int63(),
			Email: faker.Email(),
			Phone: faker.Phonenumber(),
			Organizations: []*proto.Organization{
				{
					Id:   rand.Int63(),
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
func (s *Server) OneTimePasscode(ctx context.Context, in *proto.OneTimePasscodeRequest) (*proto.OneTimePasscodeResponse, error) {
	return nil, nil
}
