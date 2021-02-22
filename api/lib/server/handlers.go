package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"

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

	for rateName, rate := range *wyreRates {
		rates[rateName] = &proto.PricingRate{
			Rate: rate,
		}
	}

	return &resp, nil
}
