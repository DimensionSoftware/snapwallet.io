package legalname

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
)

// ProfileDataLegalName the legal name of a user
type ProfileDataLegalName struct {
	ID        profiledata.ID     `firestore:"id"`
	LegalName string             `firestore:"legalName"`
	Status    profiledata.Status `firestore:"status"`
	CreatedAt time.Time          `firestore:"createdAt"`
	SealedAt  *time.Time         `firestore:"sealedAt"`
}
