package protocol

import (
	"regexp"

	v "github.com/Boostport/address"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Validate checks for required fields on a WyreAddBankPaymentMethodRequest
func (req *PlaidConnectBankAccountsRequest) Validate() error {
	if req.PlaidPublicToken == "" {
		return status.Errorf(codes.InvalidArgument, "plaid public token must be set")
	}

	if req.Institution == nil {
		return status.Errorf(codes.InvalidArgument, "plaid institution must be set")
	}
	if req.Institution.Id == "" {
		return status.Errorf(codes.InvalidArgument, "plaid institution id must be set")
	}
	if req.Institution.Name == "" {
		return status.Errorf(codes.InvalidArgument, "plaid institution name must be set")
	}

	if len(req.Accounts) < 1 {
		return status.Errorf(codes.InvalidArgument, "plaid accounts must be set")
	}

	for _, plaidAccount := range req.Accounts {
		if plaidAccount.Id == "" {
			return status.Errorf(codes.InvalidArgument, "Plaid account id must be set")
		}
		if plaidAccount.Name == "" {
			return status.Errorf(codes.InvalidArgument, "Plaid account name must be set")
		}
		if plaidAccount.Mask == "" {
			return status.Errorf(codes.InvalidArgument, "Plaid account mask must be set")
		}
		if plaidAccount.Type == "" {
			return status.Errorf(codes.InvalidArgument, "Plaid account type must be set")
		}
		if plaidAccount.SubType == "" {

			return status.Errorf(codes.InvalidArgument, "Plaid account subType must be set")
		}
	}

	return nil
}

// YYYY-MM-DD
var dobRegexp *regexp.Regexp = regexp.MustCompile("^\\d{4}-\\d{2}-\\d{2}$")

// https://www.ssa.gov/history/ssn/geocard.html
// AREA - GROUP - SERIAL
// AAA - GG - SSSS
var ssnRegexp *regexp.Regexp = regexp.MustCompile("^\\d{3}-\\d{2}-\\d{4}$")

// Validate checks SaveProfileDataRequest for validation errors
func (req *SaveProfileDataRequest) Validate() error {
	atLeastOneItemSetOnRequest := false

	if req.LegalName != "" {
		atLeastOneItemSetOnRequest = true

		if len(req.LegalName) < 3 {
			return status.Errorf(codes.InvalidArgument, "Legal name is too short")
		}
	}

	if req.DateOfBirth != "" {
		atLeastOneItemSetOnRequest = true

		if !dobRegexp.MatchString(req.DateOfBirth) {
			return status.Errorf(codes.InvalidArgument, "Date of birth must be in this format: YYYY-MM-DD")
		}
	}

	if req.Ssn != "" {
		atLeastOneItemSetOnRequest = true

		if !ssnRegexp.MatchString(req.Ssn) {
			return status.Errorf(codes.InvalidArgument, "Social security number must be in this format: NNN-NN-NNNN")
		}
	}

	if req.Address != nil {
		atLeastOneItemSetOnRequest = true

		_, err := v.NewValid(
			v.WithStreetAddress([]string{req.Address.Street_1, req.Address.Street_2}),
			v.WithLocality(req.Address.City),
			v.WithAdministrativeArea(req.Address.State),
			v.WithPostCode(req.Address.PostalCode),
			v.WithCountry(req.Address.Country), // Must be an ISO 3166-1 country code
		)
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "Address is invalid")
		}
	}

	if req.ProofOfAddressDoc != nil {
		atLeastOneItemSetOnRequest = true
	}

	if req.UsGovernmentIdDoc != nil {
		atLeastOneItemSetOnRequest = true
	}

	if !atLeastOneItemSetOnRequest {
		return status.Errorf(codes.InvalidArgument, "At least one profile field needs to be set")
	}

	return nil
}
