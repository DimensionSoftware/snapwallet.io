package ssn

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
)

// ProfileDataPhone the phone number of a user
type ProfileDataPhone struct {
	ID        profiledata.ID     `firestore:"id"`
	Phone     string             `firestore:"phone"`
	Status    profiledata.Status `firestore:"status"`
	CreatedAt time.Time          `firestore:"createdAt"`
	SealedAt  *time.Time         `firestore:"sealedAt"`
}
