package item

import (
	"encoding/json"
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/encryption"
)

// ID ...
type ID string

// AccountID ...
type AccountID string

// InstitutionID ...
type InstitutionID string

// EncryptedItem ...
type EncryptedItem struct {
	ID                   ID        `firestore:"id"`
	DataEncryptionKey    []byte    `firestore:"DEK"`
	AccessTokenEncrypted []byte    `firestore:"accessTokenEncrypted"`
	InstitutionEncrypted []byte    `firestore:"institutionEncrypted"`
	AccountsEncrypted    []byte    `firestore:"accountsEncrypted"`
	CreatedAt            time.Time `firestore:"createdAt"`
}

// Item storage ...
// https://plaid.com/docs/api/items/#item-get-response-item
type Item struct {
	ID          ID
	AccessToken string
	Institution Institution
	Accounts    []Account
	CreatedAt   time.Time
}

// Institution ...
type Institution struct {
	ID   InstitutionID `json:"id"`
	Name string        `json:"name"`
}

// Account ...
type Account struct {
	ID      AccountID `json:"id"`
	Name    string    `json:"name"`
	Mask    string    `json:"mask"`
	Type    string    `json:"type"`
	SubType string    `json:"subType"`
}

// Decrypt ...
func (enc *EncryptedItem) Decrypt(m *encryption.Manager, userID user.ID) (*Item, error) {
	dekH, err := encryption.ParseAndDecryptKeyBytes(enc.DataEncryptionKey, m.Encryptor)
	if err != nil {
		return nil, err
	}
	dek := encryption.NewEncryptor(dekH)

	accessToken, err := encryption.DecryptStringIfNonNil(dek, []byte(userID), &enc.AccessTokenEncrypted)
	if err != nil {
		return nil, err
	}

	institutionBytes, err := dek.Decrypt(enc.InstitutionEncrypted, []byte(userID))
	if err != nil {
		return nil, err
	}

	var institution Institution
	err = json.Unmarshal(institutionBytes, &institution)
	if err != nil {
		return nil, err
	}

	accountsBytes, err := dek.Decrypt(enc.AccountsEncrypted, []byte(userID))
	if err != nil {
		return nil, err
	}

	var accounts []Account
	err = json.Unmarshal(accountsBytes, &accounts)
	if err != nil {
		return nil, err
	}

	return &Item{
		ID:          enc.ID,
		AccessToken: *accessToken,
		Institution: institution,
		Accounts:    accounts,
		CreatedAt:   enc.CreatedAt,
	}, nil

}

// Encrypt ...
func (item *Item) Encrypt(m *encryption.Manager, userID user.ID) (*EncryptedItem, error) {
	dekH := encryption.NewDEK()
	dek := encryption.NewEncryptor(dekH)

	accessTokenEncrypted, err := encryption.EncryptStringIfNonNil(dek, []byte(userID), &item.AccessToken)
	if err != nil {
		return nil, err
	}

	institutionJSON, err := json.Marshal(&item.Institution)
	if err != nil {
		return nil, err
	}
	institutionEncrypted, err := dek.Encrypt(institutionJSON, []byte(userID))
	if err != nil {
		return nil, err
	}

	accountsJSON, err := json.Marshal(&item.Accounts)
	if err != nil {
		return nil, err
	}
	accountsEncrypted, err := dek.Encrypt(accountsJSON, []byte(userID))
	if err != nil {
		return nil, err
	}

	return &EncryptedItem{
		ID:                   item.ID,
		DataEncryptionKey:    *encryption.GetEncryptedKeyBytes(dekH, m.Encryptor),
		AccessTokenEncrypted: *accessTokenEncrypted,
		InstitutionEncrypted: institutionEncrypted,
		AccountsEncrypted:    accountsEncrypted,
		CreatedAt:            item.CreatedAt,
	}, nil
}
