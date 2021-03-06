package legalname

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
)

// ProfileDataLegalName the legal name of a user
type ProfileDataLegalName struct {
	ID        common.ProfileDataID
	Status    common.ProfileDataStatus
	LegalName string
	CreatedAt time.Time
	SealedAt  *time.Time
}
