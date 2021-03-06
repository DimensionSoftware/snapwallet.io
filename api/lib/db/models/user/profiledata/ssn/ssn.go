package ssn

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
)

// ProfileDataSSN the social security number of a user
type ProfileDataSSN struct {
	ID        common.ProfileDataID
	Status    common.ProfileDataStatus
	SSN       string
	CreatedAt time.Time
	SealedAt  *time.Time
}
