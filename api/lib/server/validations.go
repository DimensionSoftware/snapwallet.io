package server

import (
	"fmt"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/nyaruka/phonenumbers"
)

// LoginKind indicates the type of login which was detected
type LoginKind string

const (
	// LoginKindEmail indicates the type of login is an email
	LoginKindEmail LoginKind = "EMAIL"
	// LoginKindPhone indicates the type of login is a phone number
	LoginKindPhone LoginKind = "PHONE"
	// LoginKindInvalid indicates that the login is malformed or invalid
	LoginKindInvalid LoginKind = "INVALID"
)

// ValidateAndNormalizeLogin returns a login kind, normalized version of the login
func ValidateAndNormalizeLogin(login string) (LoginKind, string, error) {
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
				return LoginKindInvalid, "", fmt.Errorf("a valid phone number or email is required")
			}
		} else {
			return LoginKindInvalid, "", fmt.Errorf("a valid phone number or email is required")
		}
	}

	if isPhone {
		return LoginKindPhone, normalizedEmailOrPhone, nil
	}
	return LoginKindEmail, normalizedEmailOrPhone, nil
}
