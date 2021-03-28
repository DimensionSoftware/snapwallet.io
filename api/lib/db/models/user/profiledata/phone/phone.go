package phone

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/encryption"
	proto "github.com/khoerling/flux/api/lib/protocol"
)

// ProfileDataPhone the phone number of a user
type ProfileDataPhone struct {
	common.CommonProfileData
	Phone string
}

// Kind the kind of profile data
func (pdata ProfileDataPhone) Kind() common.ProfileDataKind {
	return common.KindPhone
}

// GetStatus get the status of the profile data
func (pdata ProfileDataPhone) GetStatus() common.ProfileDataStatus {
	return pdata.Status
}

// SetStatus set the status of the profile data
func (pdata ProfileDataPhone) SetStatus(newStatus common.ProfileDataStatus) {
	pdata.Status = newStatus
}

// GetProfileDataItemInfo converts the profile data to a ProfileDataItemInfo for protocol usage
func (pdata ProfileDataPhone) GetProfileDataItemInfo() *proto.ProfileDataItemInfo {
	info := proto.ProfileDataItemInfo{
		Id:        string(pdata.ID),
		Kind:      pdata.Kind().ToProfileDataItemKind(),
		Status:    pdata.Status.ToProfileDataItemStatus(),
		CreatedAt: pdata.CreatedAt.Format(time.RFC3339),
		Length:    int32(len(pdata.Phone)),
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
func (pdata ProfileDataPhone) Encrypt(m *encryption.Manager, userID user.ID) (*common.EncryptedProfileData, error) {
	dekH := encryption.NewDEK()
	dek := encryption.NewEncryptor(dekH)

	piiData := []byte(pdata.Phone)

	encryptedData, err := dek.Encrypt(piiData, []byte(userID))
	if err != nil {
		return nil, err
	}

	out := common.EncryptedProfileData{
		ID:                pdata.ID,
		Kind:              pdata.Kind(),
		Status:            pdata.Status,
		CreatedAt:         pdata.CreatedAt,
		UpdatedAt:         pdata.UpdatedAt,
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

type ProfileDataPhones []*ProfileDataPhone

func (pdatas ProfileDataPhones) FindByPhone(phone string) *ProfileDataPhone {
	for _, pdata := range pdatas {
		if pdata.Phone == phone {
			return pdata
		}
	}
	return nil
}
