package protocol

import (
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
