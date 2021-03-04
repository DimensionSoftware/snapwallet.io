package plaid

import (
	"time"

	"github.com/khoerling/flux/api/lib/encryption"
)

// EncryptedItem ...
type EncryptedItem struct {
	ID                   string    `firestore:"id"`
	DataEncryptionKey    []byte    `firestore:"DEK"`
	AccessTokenEncrypted []byte    `firestore:"accessTokenEncrypted"`
	CreatedAt            time.Time `firestore:"createdAt"`
}

// Item storage ...
// https://plaid.com/docs/api/items/#item-get-response-item
type Item struct {
	ID          string
	AccessToken string
	CreatedAt   time.Time
}

// Decrypt ...
func (enc *EncryptedItem) Decrypt(m *encryption.Manager) (*Item, error) {
	dekH, err := encryption.ParseAndDecryptKeyBytes(enc.DataEncryptionKey, m.Encryptor)
	if err != nil {
		return nil, err
	}
	dek := encryption.NewEncryptor(dekH)

	accessToken, err := encryption.DecryptStringIfNonNil(dek, m.AdditionalData, &enc.AccessTokenEncrypted)
	if err != nil {
		return nil, err
	}

	return &Item{
		ID:          enc.ID,
		AccessToken: *accessToken,
		CreatedAt:   enc.CreatedAt,
	}, nil
}

// Encrypt ...
func (u *Item) Encrypt(m *encryption.Manager) (*EncryptedItem, error) {
	dekH := encryption.NewDEK()
	dek := encryption.NewEncryptor(dekH)

	accessTokenEncrypted, err := encryption.EncryptStringIfNonNil(dek, m.AdditionalData, &u.AccessToken)
	if err != nil {
		return nil, err
	}

	return &EncryptedItem{
		ID:                   u.ID,
		DataEncryptionKey:    encryption.GetEncryptedKeyBytes(dekH, m.Encryptor),
		AccessTokenEncrypted: *accessTokenEncrypted,
		CreatedAt:            u.CreatedAt,
	}, nil
}
