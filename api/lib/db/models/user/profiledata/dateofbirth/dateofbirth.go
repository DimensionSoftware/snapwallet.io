package dateofbirth

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/encryption"
)

// ProfileDataDateOfBirth thkke date of birth for a user
type ProfileDataDateOfBirth struct {
	ID     common.ProfileDataID
	Status common.ProfileDataStatus
	// indicates an individuals date of birth which is a string of the format YYYY-MM-DD
	DateOfBirth string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	SealedAt    *time.Time
}

// Kind the kind of profile data
func (pdata ProfileDataDateOfBirth) Kind() common.ProfileDataKind {
	return common.KindDateOfBirth
}

// GetStatus get the status of the profile data
func (pdata ProfileDataDateOfBirth) GetStatus() common.ProfileDataStatus {
	return pdata.Status
}

// Encrypt ...
func (pdata ProfileDataDateOfBirth) Encrypt(m *encryption.Manager, userID user.ID) (*common.EncryptedProfileData, error) {
	dekH := encryption.NewDEK()
	dek := encryption.NewEncryptor(dekH)

	piiData := []byte(pdata.DateOfBirth)

	encryptedData, err := dek.Encrypt(piiData, []byte(userID))
	if err != nil {
		return nil, err
	}

	return &common.EncryptedProfileData{
		ID:                pdata.ID,
		Kind:              pdata.Kind(),
		Status:            pdata.Status,
		CreatedAt:         pdata.CreatedAt,
		SealedAt:          pdata.SealedAt,
		DataEncryptionKey: encryption.GetEncryptedKeyBytes(dekH, m.Encryptor),
		EncryptedData:     encryptedData,
	}, nil
}
