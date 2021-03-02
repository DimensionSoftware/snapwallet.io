package server

import (
	"strings"

	"github.com/badoux/checkmail"
	"github.com/khoerling/flux/api/lib/db/models/onetimepasscode"
	"github.com/nyaruka/phonenumbers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidateAndNormalizeLogin returns a login kind, normalized version of the login
func ValidateAndNormalizeLogin(login string) (onetimepasscode.LoginKind, string, error) {
	var normalizedEmailOrPhone string
	var isPhone bool

	num, err := phonenumbers.Parse(strings.TrimSpace(login), "US")
	if err == nil {
		isPhone = true
		normalizedEmailOrPhone = phonenumbers.Format(num, phonenumbers.E164)
	} else {
		err = checkmail.ValidateFormat(login)
		if err == nil {
			normalizedEmailOrPhone = strings.TrimSpace(login)
		} else {
			return onetimepasscode.LoginKindInvalid, "", status.Errorf(codes.InvalidArgument, "a valid phone number or email is required")
		}
	}

	if isPhone {
		return onetimepasscode.LoginKindPhone, normalizedEmailOrPhone, nil
	}
	return onetimepasscode.LoginKindEmail, normalizedEmailOrPhone, nil
}
