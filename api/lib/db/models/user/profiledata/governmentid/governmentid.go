package governmentid

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/file"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/encryption"
)

// Kind indicates the type of government identification the profiledata holds
type Kind string

const (
	// KindUSDrivingLicense is a United States drivers license
	KindUSDrivingLicense Kind = "US_DRIVING_LICENSE"
	// KindUSPassportCard is a United States passport card
	KindUSPassportCard Kind = "US_PASSPORT_CARD"
	// KindUSGovernmentID is a United States government ID
	KindUSGovernmentID Kind = "US_GOVERNMENT_ID"
	// KindUSPassport is a United States passport
	KindUSPassport Kind = "US_PASSPORT"
)

// FilesRequired indicates the number of files required for the ProfileDataGovernmentIDKind
// generic mapping:
// file 0: front of govt identification
// file 1: back of govt identification
// file N: reserved for future use in case some govt id types require more than 2 file uploads
func (govtIDKind Kind) FilesRequired() int {
	switch govtIDKind {
	case KindUSDrivingLicense:
		return 2
	case KindUSPassportCard:
		return 2
	case KindUSGovernmentID:
		return 2
	case KindUSPassport:
		return 1
	}
	// should never get here
	panic("governmentid.Kind unknown when FilesRequired() called")
}

// ProfileDataGovernmentID represents a government ID for a user
type ProfileDataGovernmentID struct {
	ID        common.ProfileDataID
	Status    common.ProfileDataStatus
	Kind      Kind
	FileIDs   []file.ID
	CreatedAt time.Time
	UpdatedAt *time.Time
	SealedAt  *time.Time
}

// Encrypt ...
func (pdata ProfileDataGovernmentID) Encrypt(m *encryption.Manager, userID user.ID) (*common.EncryptedProfileData, error) {

	return &common.EncryptedProfileData{
		ID:        pdata.ID,
		Kind:      common.KindGovernmentID,
		SubKind:   (*string)(&pdata.Kind),
		Status:    pdata.Status,
		Files:     &pdata.FileIDs,
		CreatedAt: pdata.CreatedAt,
		UpdatedAt: pdata.UpdatedAt,
		SealedAt:  pdata.SealedAt,
	}, nil
}
