package ssn

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
)

// ProfileDataSSN the social security number of a user
type ProfileDataSSN struct {
	ID        profiledata.ID     `firestore:"id"`
	SSN       string             `firestore:"ssn"`
	Status    profiledata.Status `firestore:"status"`
	CreatedAt time.Time          `firestore:"createdAt"`
	SealedAt  *time.Time         `firestore:"sealedAt"`
}
