package address

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
)

// ProfileDataAddress an address for a user
type ProfileDataAddress struct {
	ID         profiledata.ID `json:"id"`
	Street1    string         `json:"street1"`
	Street2    string         `json:"street2"`
	City       string         `json:"city"`
	State      string         `json:"state"`
	PostalCode string         `json:"postalCode"`
	Country    string         `json:"country"`
	CreatedAt  time.Time
	SealedAt   *time.Time `firestore:"sealedAt"`
}
