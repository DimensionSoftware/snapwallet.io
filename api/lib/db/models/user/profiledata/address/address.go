package address

import (
	"encoding/json"
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
	"github.com/khoerling/flux/api/lib/encryption"
)

// ProfileDataAddress an address for a user
type ProfileDataAddress struct {
	ID         profiledata.ID
	Status     profiledata.Status
	Street1    string
	Street2    string
	City       string
	State      string
	PostalCode string
	Country    string
	CreatedAt  time.Time
	SealedAt   *time.Time
}

// ProfileDataAddressPIIData ...
type ProfileDataAddressPIIData struct {
	Street1    string `json:"street1"`
	Street2    string `json:"street2"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

// Encrypt ...
func (pdata ProfileDataAddress) Encrypt(m *encryption.Manager, userID user.ID) (*profiledata.EncryptedProfileData, error) {
	dekH := encryption.NewDEK()
	dek := encryption.NewEncryptor(dekH)

	piiData := ProfileDataAddressPIIData{
		Street1:    pdata.Street1,
		Street2:    pdata.Street2,
		City:       pdata.City,
		State:      pdata.State,
		PostalCode: pdata.PostalCode,
		Country:    pdata.Country,
	}

	piiJSONData, err := json.Marshal(piiData)
	if err != nil {
		return nil, err
	}

	encryptedData, err := dek.Encrypt(piiJSONData, []byte(userID))
	if err != nil {
		return nil, err
	}

	return &profiledata.EncryptedProfileData{
		ID:                pdata.ID,
		Kind:              profiledata.KindAddress,
		Status:            pdata.Status,
		CreatedAt:         pdata.CreatedAt,
		SealedAt:          pdata.SealedAt,
		DataEncryptionKey: encryption.GetEncryptedKeyBytes(dekH, m.Encryptor),
		EncryptedData:     encryptedData,
	}, nil
}
