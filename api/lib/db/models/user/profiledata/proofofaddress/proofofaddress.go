package proofofaddress

import (
	"log"
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/file"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/encryption"
	proto "github.com/khoerling/flux/api/lib/protocol"
)

// ProfileDataProofOfAddress represents a proof of address document
type ProfileDataProofOfAddressDoc struct {
	common.CommonProfileData
	FileIDs []file.ID
}

// Encrypt ...
func (pdata ProfileDataProofOfAddressDoc) Encrypt(m *encryption.Manager, userID user.ID) (*common.EncryptedProfileData, error) {

	log.Printf("%#v: ", pdata)
	return &common.EncryptedProfileData{
		ID:        pdata.ID,
		Kind:      pdata.Kind(),
		Status:    pdata.Status,
		FileIDs:   &pdata.FileIDs,
		CreatedAt: pdata.CreatedAt,
		UpdatedAt: pdata.UpdatedAt,
		SealedAt:  pdata.SealedAt,
	}, nil
}

// Kind the kind of profile data
func (pdata ProfileDataProofOfAddressDoc) Kind() common.ProfileDataKind {
	return common.KindProofOfAddressDoc
}

// GetStatus get the status of the profile data
func (pdata ProfileDataProofOfAddressDoc) GetStatus() common.ProfileDataStatus {
	return pdata.Status
}

// SetStatus set the status of the profile data
func (pdata ProfileDataProofOfAddressDoc) SetStatus(newStatus common.ProfileDataStatus) {
	pdata.Status = newStatus
}

// GetProfileDataItemInfo converts the profile data to a ProfileDataItemInfo for protocol usage
func (pdata ProfileDataProofOfAddressDoc) GetProfileDataItemInfo() *proto.ProfileDataItemInfo {
	var fileIDs []string
	for _, fileID := range pdata.FileIDs {
		fileIDs = append(fileIDs, string(fileID))
	}

	info := proto.ProfileDataItemInfo{
		Id:        string(pdata.ID),
		Kind:      pdata.Kind().ToProfileDataItemKind(),
		FileIds:   fileIDs,
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
