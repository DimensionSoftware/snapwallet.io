package protocol

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Validate checks for required fields on a WyreAddBankPaymentMethodRequest
func (req *WyreAddBankPaymentMethodsRequest) Validate() error {
	if req.PlaidPublicToken == "" {
		return status.Errorf(codes.InvalidArgument, "plaid public token must be set")
	}

	if len(req.PlaidAccountIds) < 1 {
		return status.Errorf(codes.InvalidArgument, "account id must be set")
	}

	return nil
}
