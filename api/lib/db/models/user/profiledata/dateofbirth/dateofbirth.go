package ssn

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
)

// ProfileDataDateOfBirth the date of birth for a user
type ProfileDataDateOfBirth struct {
	ID          profiledata.ID `firestore:"id"`
	DateOfBirth string         `firestore:"phoneNumber"`
	CreatedAt   time.Time      `firestore:"createdAt"`
	SealedAt    *time.Time     `firestore:"sealedAt"`
}
