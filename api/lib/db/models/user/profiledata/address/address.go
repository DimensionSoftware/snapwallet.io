package address

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
)

// ProfileDataAddress an address for a user
type ProfileDataAddress struct {
	ID         profiledata.ID `firestore:"id"`
	Street1    string         `firestore:"street1"`
	Street2    string         `firestore:"street2"`
	City       string         `firestore:"city"`
	State      string         `firestore:"state"`
	PostalCode string         `firestore:"postalCode"`
	Country    string         `firestore:"country"`
	CreatedAt  time.Time      `firestore:"createdAt"`
	SealedAt   *time.Time     `firestore:"sealedAt"`
}
