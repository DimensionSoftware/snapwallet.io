package protocol

import (
	v "github.com/Boostport/address"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Validate checks for required fields on a WyreAddBankPaymentMethodRequest
func (req *PlaidConnectBankAccountsRequest) Validate() error {
	if req.PlaidPublicToken == "" {
		return status.Errorf(codes.InvalidArgument, "plaid public token must be set")
	}

	if len(req.PlaidAccountIds) < 1 {
		return status.Errorf(codes.InvalidArgument, "plaid account ids must be set")
	}

	for _, plaidAccountID := range req.PlaidAccountIds {
		if plaidAccountID == "" {
			return status.Errorf(codes.InvalidArgument, "plaid account ids must be non-empty strings")
		}
	}

	return nil
}

// Validate checks SaveProfileDataRequest for validation errors
func (req *SaveProfileDataRequest) Validate() error {
	atLeastOneItemSetOnRequest := false

	if req.Address != nil {
		atLeastOneItemSetOnRequest = true
		_, err := v.NewValid(
			v.WithStreetAddress([]string{req.Address.Street_1, req.Address.Street_2}),
			v.WithLocality(req.Address.City),
			v.WithAdministrativeArea(req.Address.State),
			v.WithPostCode(req.Address.PostalCode),
			v.WithCountry(req.Address.Country), // Must be an ISO 3166-1 country code
		)
		return err
	}

	if !atLeastOneItemSetOnRequest {
		return status.Errorf(codes.InvalidArgument, "at least one profile field needs to be set")
	}

	return nil
}
