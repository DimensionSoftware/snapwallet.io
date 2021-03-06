package ssn

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
)

// ProfileDataDateOfBirth thkke date of birth for a user
type ProfileDataDateOfBirth struct {
	ID     common.ProfileDataID
	Status common.ProfileDataStatus
	// indicates an individuals date of birth which is a string of the format YYYY-MM-DD
	DateOfBirth string
	CreatedAt   time.Time
	SealedAt    *time.Time
}
