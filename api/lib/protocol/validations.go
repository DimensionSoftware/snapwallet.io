package protocol

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Validate checks for required fields on a WyreAddBankPaymentMethodRequest
func (req *WyreAddBankPaymentMethodRequest) Validate() error {
	if req.AccessToken == "" {
		return status.Errorf(codes.InvalidArgument, "access token must be set")
	}

	if req.AccountId == "" {
		return status.Errorf(codes.InvalidArgument, "account id must be set")
	}

	return nil
}
