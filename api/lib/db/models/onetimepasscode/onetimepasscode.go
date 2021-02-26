package onetimepasscode

import "time"

// LoginKind describes the kind of login identifier used for the one time passcode
type LoginKind string

const (
	// LoginKindPhone represents a phone-based login
	LoginKindPhone LoginKind = "PHONE"
	// LoginKindEmail represents an email-based login
	LoginKindEmail LoginKind = "EMAIL"
	// LoginKindInvalid indicates that the login is malformed or invalid
	LoginKindInvalid LoginKind = "INVALID"
)

// OneTimePasscode stores the transient record of a login attempt factor so it can be verified in exchange for a token
type OneTimePasscode struct {
	ID           string    `firestore:"id"`
	EmailOrPhone string    `firestore:"emailOrPhone"`
	Kind         LoginKind `firestore:"kind"`
	Code         string    `firestore:"code"`
	CreatedAt    time.Time `firestore:"createdAt"`
}
