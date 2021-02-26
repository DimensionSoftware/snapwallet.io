package server

import (
	"fmt"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/khoerling/flux/api/lib/db/models"
	"github.com/nyaruka/phonenumbers"
)

// ValidateAndNormalizeLogin returns a login kind, normalized version of the login
func ValidateAndNormalizeLogin(login string) (models.OneTimePasscodeLoginKind, string, error) {
	var normalizedEmailOrPhone string
	var isPhone bool

	num, err := phonenumbers.Parse(strings.TrimSpace(login), "US")
	if err == nil {
		isPhone = true
		normalizedEmailOrPhone = phonenumbers.Format(num, phonenumbers.E164)
	} else {
		err = checkmail.ValidateFormat(login)
		if err == nil {
			err = checkmail.ValidateHost(login)
			if err == nil {
				normalizedEmailOrPhone = strings.TrimSpace(login)
			} else {
				return models.OneTimePasscodeLoginKindInvalid, "", fmt.Errorf("a valid phone number or email is required")
			}
		} else {
			return models.OneTimePasscodeLoginKindInvalid, "", fmt.Errorf("a valid phone number or email is required")
		}
	}

	if isPhone {
		return models.OneTimePasscodeLoginKindPhone, normalizedEmailOrPhone, nil
	}
	return models.OneTimePasscodeLoginKindEmail, normalizedEmailOrPhone, nil
}
