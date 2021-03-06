package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_Validate_SaveProfileDataRequest_no_fields_should_throw_error(t *testing.T) {
	a := assert.New(t)
	req := SaveProfileDataRequest{}

	expectedErr := status.Errorf(codes.InvalidArgument, "At least one profile field needs to be set")
	a.Equal(expectedErr, req.Validate())
}

func Test_Validate_SaveProfileDataRequest_valid_legal_name_should_not_throw_error(t *testing.T) {
	a := assert.New(t)
	req := SaveProfileDataRequest{
		LegalName: "Bob Jones",
	}

	a.NoError(req.Validate())
}

func Test_Validate_SaveProfileDataRequest_invalid_legal_name_should_throw_error(t *testing.T) {
	a := assert.New(t)
	req := SaveProfileDataRequest{
		LegalName: "ME",
	}

	expectedErr := status.Errorf(codes.InvalidArgument, "Legal name is too short")
	a.Equal(expectedErr, req.Validate())
}

func Test_Validate_SaveProfileDataRequest_valid_dob_should_not_throw_error(t *testing.T) {
	a := assert.New(t)
	req := SaveProfileDataRequest{
		DateOfBirth: "2000-12-01",
	}

	a.NoError(req.Validate())
}

func Test_Validate_SaveProfileDataRequest_invalid_dob_should_throw_error(t *testing.T) {
	a := assert.New(t)
	req := SaveProfileDataRequest{
		DateOfBirth: "muhahahahaha",
	}

	expectedErr := status.Errorf(codes.InvalidArgument, "Date of birth must be in this format: YYYY-MM-DD")
	a.Equal(expectedErr, req.Validate())
}

func Test_Validate_SaveProfileDataRequest_valid_ssn_should_not_throw_error(t *testing.T) {
	a := assert.New(t)
	req := SaveProfileDataRequest{
		Ssn: "123-12-1234",
	}

	a.NoError(req.Validate())
}

func Test_Validate_SaveProfileDataRequest_invalid_ssn_should_throw_error(t *testing.T) {
	a := assert.New(t)
	req := SaveProfileDataRequest{
		Ssn: "i will rule the world",
	}

	expectedErr := status.Errorf(codes.InvalidArgument, "Social security number must be in this format: NNN-NN-NNNN")
	a.Equal(expectedErr, req.Validate())
}

func Test_Validate_SaveProfileDataRequest_valid_address_should_not_throw_error(t *testing.T) {
	a := assert.New(t)
	req := SaveProfileDataRequest{
		Address: &Address{
			Street_1:   "7840 Brimhall Rd",
			City:       "Bakersfield",
			State:      "CA",
			PostalCode: "93308",
			Country:    "US",
		},
	}
	a.NoError(req.Validate())
}

func Test_Validate_SaveProfileDataRequest_invalid_address_should_throw_error(t *testing.T) {
	a := assert.New(t)
	req := SaveProfileDataRequest{
		Address: &Address{
			Street_1:   "7840 Brimhall Rd",
			State:      "CA",
			PostalCode: "93308",
			Country:    "US",
		},
	}

	expectedErr := status.Errorf(codes.InvalidArgument, "Address is invalid")
	a.Equal(expectedErr, req.Validate())
}
