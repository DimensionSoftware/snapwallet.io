package ssn

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/encryption"
)

// ProfileDataSSN the social security number of a user
type ProfileDataSSN struct {
	ID        common.ProfileDataID
	Status    common.ProfileDataStatus
	SSN       string
	CreatedAt time.Time
	SealedAt  *time.Time
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

	return &common.EncryptedProfileData{
		ID:                pdata.ID,
		Kind:              common.KindSSN,
		Status:            pdata.Status,
		CreatedAt:         pdata.CreatedAt,
		SealedAt:          pdata.SealedAt,
		DataEncryptionKey: encryption.GetEncryptedKeyBytes(dekH, m.Encryptor),
		EncryptedData:     encryptedData,
	}, nil
}
