package ssn

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
)

// ProfileDataPhone the phone number of a user
type ProfileDataPhone struct {
	ID        common.ProfileDataID
	Status    common.ProfileDataStatus
	Phone     string
	CreatedAt time.Time
	SealedAt  *time.Time
}
