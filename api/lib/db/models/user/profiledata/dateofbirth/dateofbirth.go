package ssn

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
)

// ProfileDataDateOfBirth the date of birth for a user
type ProfileDataDateOfBirth struct {
	ID          profiledata.ID `json:"id"`
	DateOfBirth string         `json:"phoneNumber"`
	CreatedAt   time.Time
	SealedAt    *time.Time `firestore:"sealedAt"`
}
