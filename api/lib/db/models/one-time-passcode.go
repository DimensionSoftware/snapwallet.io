package models

import "time"

// OneTimePasscodeLoginKind describes the kind of login identifier used for the one time passcode
type OneTimePasscodeLoginKind string

const (
	// OneTimePasscodeLoginKindPhone represents a phone-based login
	OneTimePasscodeLoginKindPhone OneTimePasscodeLoginKind = "PHONE"
	// OneTimePasscodeLoginKindEmail represents an email-based login
	OneTimePasscodeLoginKindEmail OneTimePasscodeLoginKind = "EMAIL"
	// OneTimePasscodeLoginKindInvalid indicates that the login is malformed or invalid
	OneTimePasscodeLoginKindInvalid OneTimePasscodeLoginKind = "INVALID"
)

// OneTimePasscode stores the transient record of a login attempt factor so it can be verified in exchange for a token
type OneTimePasscode struct {
	EmailOrPhone string                   `firestore:"emailOrPhone"`
	Kind         OneTimePasscodeLoginKind `firestore:"kind"`
	Code         string                   `firestore:"code"`
	CreatedAt    time.Time                `firestore:"createdAt"`
}
