package ssn

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/encryption"
	proto "github.com/khoerling/flux/api/lib/protocol"
)

// ProfileDataSSN the social security number of a user
type ProfileDataSSN struct {
	common.CommonProfileData
	SSN string
}

// Kind the kind of profile data
func (pdata ProfileDataSSN) Kind() common.ProfileDataKind {
	return common.KindUSSSN
}

// SetStatus set the status of the profile data
func (pdata ProfileDataSSN) SetStatus(newStatus common.ProfileDataStatus) {
	pdata.Status = newStatus
}

// GetProfileDataItemInfo converts the profile data to a ProfileDataItemInfo for protocol usage
func (pdata ProfileDataSSN) GetProfileDataItemInfo() *proto.ProfileDataItemInfo {
	info := proto.ProfileDataItemInfo{
		Id:        string(pdata.ID),
		Kind:      pdata.Kind().ToProfileDataItemKind(),
		Status:    pdata.Status.ToProfileDataItemStatus(),
		CreatedAt: pdata.CreatedAt.Format(time.RFC3339),
		Length:    int32(len(pdata.SSN)),
	}
	if pdata.UpdatedAt != nil {
		info.UpdatedAt = pdata.UpdatedAt.Format(time.RFC3339)
	}
	if pdata.SealedAt != nil {
		info.SealedAt = pdata.SealedAt.Format(time.RFC3339)
	}

	return &info
}

// Encrypt ...
func (pdata ProfileDataSSN) Encrypt(m *encryption.Manager, userID user.ID) (*common.EncryptedProfileData, error) {
	dekH := encryption.NewDEK()
	dek := encryption.NewEncryptor(dekH)

	piiData := []byte(pdata.SSN)

	encryptedData, err := dek.Encrypt(piiData, []byte(userID))
	if err != nil {
		return nil, err
	}

	out := common.EncryptedProfileData{
		ID:                pdata.ID,
		Kind:              pdata.Kind(),
		Status:            pdata.Status,
		CreatedAt:         pdata.CreatedAt,
		SealedAt:          pdata.SealedAt,
		DataEncryptionKey: encryption.GetEncryptedKeyBytes(dekH, m.Encryptor),
		EncryptedData:     &encryptedData,
	}

	if pdata.Note != "" {
		noteData := []byte(pdata.Note)
		encryptedNote, err := dek.Encrypt(noteData, []byte(userID))
		if err != nil {
			return nil, err
		}
		out.EncryptedNote = &encryptedNote
	}

	return &out, nil
}
