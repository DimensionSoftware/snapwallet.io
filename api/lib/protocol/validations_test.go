package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Validate_SaveProfileDataRequest_empty_fields_should_throw_error(t *testing.T) {
	a := assert.New(t)
	req := SaveProfileDataRequest{}
	a.Error(req.Validate())
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
