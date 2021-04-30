package item

import (
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
	AccountIDs           []string  `firestore:"accountIDs"`
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
	ID   InstitutionID
	Name string
}

// Account ...
type Account struct {
	ID      AccountID
	Name    string
	Mask    string
	Type    string
	SubType string
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

	return &Item{
		ID:          enc.ID,
		AccessToken: *accessToken,
		AccountIDs:  enc.AccountIDs,
		CreatedAt:   enc.CreatedAt,
	}, nil
}

// Encrypt ...
func (u *Item) Encrypt(m *encryption.Manager, userID user.ID) (*EncryptedItem, error) {
	dekH := encryption.NewDEK()
	dek := encryption.NewEncryptor(dekH)

	accessTokenEncrypted, err := encryption.EncryptStringIfNonNil(dek, []byte(userID), &u.AccessToken)
	if err != nil {
		return nil, err
	}

	return &EncryptedItem{
		ID:                   u.ID,
		DataEncryptionKey:    *encryption.GetEncryptedKeyBytes(dekH, m.Encryptor),
		AccessTokenEncrypted: *accessTokenEncrypted,
		AccountIDs:           u.AccountIDs,
		CreatedAt:            u.CreatedAt,
	}, nil
}
