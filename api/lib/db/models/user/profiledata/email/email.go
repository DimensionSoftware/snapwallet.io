package email

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/encryption"
	proto "github.com/khoerling/flux/api/lib/protocol"
)

// ProfileDataEmail the email of a user
type ProfileDataEmail struct {
	common.CommonProfileData
	Email string
}

// Kind the kind of profile data
func (pdata ProfileDataEmail) Kind() common.ProfileDataKind {
	return common.KindEmail
}

// GetStatus get the status of the profile data
func (pdata ProfileDataEmail) GetStatus() common.ProfileDataStatus {
	return pdata.Status
}

// SetStatus set the status of the profile data
func (pdata ProfileDataEmail) SetStatus(newStatus common.ProfileDataStatus) {
	pdata.Status = newStatus
}

// GetProfileDataItemInfo converts the profile data to a ProfileDataItemInfo for protocol usage
func (pdata ProfileDataEmail) GetProfileDataItemInfo() *proto.ProfileDataItemInfo {
	info := proto.ProfileDataItemInfo{
		Id:        string(pdata.ID),
		Kind:      pdata.Kind().ToProfileDataItemKind(),
		Status:    pdata.Status.ToProfileDataItemStatus(),
		CreatedAt: pdata.CreatedAt.Format(time.RFC3339),
		Length:    int32(len(pdata.Email)),
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
func (pdata ProfileDataEmail) Encrypt(m *encryption.Manager, userID user.ID) (*common.EncryptedProfileData, error) {
	dekH := encryption.NewDEK()
	dek := encryption.NewEncryptor(dekH)

	piiData := []byte(pdata.Email)

	encryptedData, err := dek.Encrypt(piiData, []byte(userID))
	if err != nil {
		return nil, err
	}

	return &common.EncryptedProfileData{
		ID:                pdata.ID,
		Kind:              pdata.Kind(),
		Status:            pdata.Status,
		CreatedAt:         pdata.CreatedAt,
		UpdatedAt:         pdata.UpdatedAt,
		SealedAt:          pdata.SealedAt,
		DataEncryptionKey: encryption.GetEncryptedKeyBytes(dekH, m.Encryptor),
		EncryptedData:     &encryptedData,
	}, nil
}

type ProfileDataEmails []*ProfileDataEmail

func (pdatas ProfileDataEmails) FindByEmail(email string) *ProfileDataEmail {
	for _, pdata := range pdatas {
		if pdata.Email == email {
			return pdata
		}
	}
	return nil
}
