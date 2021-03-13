package governmentid

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/file"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/encryption"
	proto "github.com/khoerling/flux/api/lib/protocol"
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

// ToProfileDataItemKind converts the profile data kind to a protocol.ProfileDataItemKind
func KindFromGovernmentIdDocKind(k proto.GovernmentIdDocumentInputKind) Kind {
	switch k {
	case proto.GovernmentIdDocumentInputKind_GI_US_DRIVING_LICENSE:
		return KindUSDrivingLicense
	case proto.GovernmentIdDocumentInputKind_GI_US_GOVERNMENT_ID:
		return KindUSGovernmentID
	case proto.GovernmentIdDocumentInputKind_GI_US_PASSPORT_CARD:
		return KindUSPassportCard
	case proto.GovernmentIdDocumentInputKind_GI_US_PASSPORT:
		return KindUSPassport
	}
	// should never get here
	panic("proto.GovernmentIdDocumentInputKind unknown when KindFromGovernmentIdDocKind(...) called")
}

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
	ID               common.ProfileDataID
	Status           common.ProfileDataStatus
	GovernmentIDKind Kind
	FileIDs          []file.ID
	CreatedAt        time.Time
	UpdatedAt        *time.Time
	SealedAt         *time.Time
}

// Encrypt ...
func (pdata ProfileDataGovernmentID) Encrypt(m *encryption.Manager, userID user.ID) (*common.EncryptedProfileData, error) {

	return &common.EncryptedProfileData{
		ID:        pdata.ID,
		Kind:      common.KindGovernmentID,
		SubKind:   (*string)(&pdata.GovernmentIDKind),
		Status:    pdata.Status,
		FileIDs:   &pdata.FileIDs,
		CreatedAt: pdata.CreatedAt,
		UpdatedAt: pdata.UpdatedAt,
		SealedAt:  pdata.SealedAt,
	}, nil
}

// Kind the kind of profile data
func (pdata ProfileDataGovernmentID) Kind() common.ProfileDataKind {
	return common.KindLegalName
}

// GetStatus get the status of the profile data
func (pdata ProfileDataGovernmentID) GetStatus() common.ProfileDataStatus {
	return pdata.Status
}

// SetStatus set the status of the profile data
func (pdata ProfileDataGovernmentID) SetStatus(newStatus common.ProfileDataStatus) {
	pdata.Status = newStatus
}

// GetProfileDataItemInfo converts the profile data to a ProfileDataItemInfo for protocol usage
func (pdata ProfileDataGovernmentID) GetProfileDataItemInfo() *proto.ProfileDataItemInfo {
	info := proto.ProfileDataItemInfo{
		Id:        string(pdata.ID),
		Kind:      pdata.Kind().ToProfileDataItemKind(),
		Status:    pdata.Status.ToProfileDataItemStatus(),
		CreatedAt: pdata.CreatedAt.Format(time.RFC3339),
		Length:    int32(len(pdata.FileIDs)),
	}
	if pdata.UpdatedAt != nil {
		info.UpdatedAt = pdata.UpdatedAt.Format(time.RFC3339)
	}
	if pdata.SealedAt != nil {
		info.SealedAt = pdata.SealedAt.Format(time.RFC3339)
	}

	return &info
}
