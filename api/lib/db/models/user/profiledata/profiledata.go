package unmarshal

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/address"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/encryption"
)

func unmarshal(kind common.ProfileDataKind, data []byte) (interface{}, error) {
	switch kind {
	case common.KindAddress:
		var out address.ProfileDataAddress
		err := json.Unmarshal(data, &out)
		if err != nil {
			return nil, err
		}

		return out, nil
	}

	return nil, fmt.Errorf("ProfileDataKind: %s is not implemented yet", kind)
}

// EncryptedProfileData is a generic container store encrypted ProfileData
type EncryptedProfileData struct {
	ID                common.ProfileDataID     `firestore:"id"`
	Kind              common.ProfileDataKind   `firestore:"kind"`
	Status            common.ProfileDataStatus `firestore:"status"`
	CreatedAt         time.Time                `firestore:"createdAt"`
	DataEncryptionKey []byte                   `firestore:"DEK"`
	EncryptedData     []byte                   `firestore:"encryptedData"`
	SealedAt          *time.Time               `firestore:"sealedAt,omitempty"`
}

func (encryptedProfileData EncryptedProfileData) decrypt(m *encryption.Manager, userID user.ID) ([]byte, error) {
	dekH, err := encryption.ParseAndDecryptKeyBytes(encryptedProfileData.DataEncryptionKey, m.Encryptor)
	if err != nil {
		return nil, err
	}

	dek := encryption.NewEncryptor(dekH)
	decrypted, err := dek.Decrypt(encryptedProfileData.EncryptedData, []byte(userID))
	if err != nil {
		return nil, err
	}

	return decrypted, nil

}

// DecryptAndUnmarshal ...
func (encryptedProfileData EncryptedProfileData) DecryptAndUnmarshal(m *encryption.Manager, userID user.ID) (*interface{}, error) {
	raw, err := encryptedProfileData.decrypt(m, userID)
	if err != nil {
		return nil, err
	}

	out, err := unmarshal(encryptedProfileData.Kind, raw)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
