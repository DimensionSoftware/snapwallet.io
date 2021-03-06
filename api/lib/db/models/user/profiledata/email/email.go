package email

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
)

// ProfileDataEmail the email of a user
type ProfileDataEmail struct {
	ID        common.ProfileDataID
	Status    common.ProfileDataStatus
	Email     string
	CreatedAt time.Time
	SealedAt  *time.Time
}
