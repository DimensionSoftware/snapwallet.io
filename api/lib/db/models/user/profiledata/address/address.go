package address

import (
	"encoding/json"
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	proto "github.com/khoerling/flux/api/lib/protocol"

	"github.com/khoerling/flux/api/lib/encryption"
)

// ProfileDataAddress an address for a user
type ProfileDataAddress struct {
	common.CommonProfileData
	Street1    string
	Street2    string
	City       string
	State      string
	PostalCode string
	Country    string
}

// Kind the kind of profile data
func (pdata ProfileDataAddress) Kind() common.ProfileDataKind {
	return common.KindAddress
}

// GetStatus get the status of the profile data
func (pdata ProfileDataAddress) GetStatus() common.ProfileDataStatus {
	return pdata.Status
}

// SetStatus set the status of the profile data
func (pdata ProfileDataAddress) SetStatus(newStatus common.ProfileDataStatus) {
	pdata.Status = newStatus
}

// GetProfileDataItemInfo converts the profile data to a ProfileDataItemInfo for protocol usage
func (pdata ProfileDataAddress) GetProfileDataItemInfo() *proto.ProfileDataItemInfo {
	length := 0 +
		len(pdata.Street1) +
		len(pdata.Street2) +
		len(pdata.City) +
		len(pdata.State) +
		len(pdata.PostalCode) +
		len(pdata.Country)

	info := proto.ProfileDataItemInfo{
		Id:        string(pdata.ID),
		Kind:      pdata.Kind().ToProfileDataItemKind(),
		Status:    pdata.Status.ToProfileDataItemStatus(),
		CreatedAt: pdata.CreatedAt.Format(time.RFC3339),
		Length:    int32(length),
	}
	if pdata.UpdatedAt != nil {
		info.UpdatedAt = pdata.UpdatedAt.Format(time.RFC3339)
	}
	if pdata.SealedAt != nil {
		info.SealedAt = pdata.SealedAt.Format(time.RFC3339)
	}

	return &info
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
func (pdata ProfileDataAddress) Encrypt(m *encryption.Manager, userID user.ID) (*common.EncryptedProfileData, error) {
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

	out := common.EncryptedProfileData{
		ID:                pdata.ID,
		Kind:              common.KindAddress,
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

// UnmarshalPIIData ...
func UnmarshalPIIData(data []byte, userID user.ID) (*ProfileDataAddressPIIData, error) {
	var out ProfileDataAddressPIIData
	err := json.Unmarshal(data, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
