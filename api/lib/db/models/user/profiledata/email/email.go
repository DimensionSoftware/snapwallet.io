package ssn

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
)

// ProfileDataEmail the email of a user
type ProfileDataEmail struct {
	ID        profiledata.ID     `firestore:"id"`
	Email     string             `firestore:"email"`
	Status    profiledata.Status `firestore:"status"`
	CreatedAt time.Time          `firestore:"createdAt"`
	SealedAt  *time.Time         `firestore:"sealedAt"`
}
