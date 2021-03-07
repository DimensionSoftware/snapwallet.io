package legalname

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/encryption"
)

// ProfileDataLegalName the legal name of a user
type ProfileDataLegalName struct {
	ID        common.ProfileDataID
	Status    common.ProfileDataStatus
	LegalName string
	CreatedAt time.Time
	UpdatedAt *time.Time
	SealedAt  *time.Time
}

// Kind the kind of profile data
func (pdata ProfileDataLegalName) Kind() common.ProfileDataKind {
	return common.KindLegalName
}

// GetStatus get the status of the profile data
func (pdata ProfileDataLegalName) GetStatus() common.ProfileDataStatus {
	return pdata.Status
}

// Encrypt ...
func (pdata ProfileDataLegalName) Encrypt(m *encryption.Manager, userID user.ID) (*common.EncryptedProfileData, error) {
	dekH := encryption.NewDEK()
	dek := encryption.NewEncryptor(dekH)

	piiData := []byte(pdata.LegalName)

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
		EncryptedData:     encryptedData,
	}, nil
}
