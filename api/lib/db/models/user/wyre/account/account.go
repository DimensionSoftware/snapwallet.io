package account

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/encryption"
)

// ID ...
type ID string

// EncryptedItem ...
type EncryptedAccount struct {
	ID                 ID         `firestore:"id"`
	DataEncryptionKey  []byte     `firestore:"DEK"`
	SecretKeyEncrypted []byte     `firestore:"secretKeyEncrypted"`
	Status             string     `firestore:"status"`
	CreatedAt          time.Time  `firestore:"createdAt"`
	UpdatedAt          *time.Time `firestore:"updatedAt,omitempty"`
}

// Account ...
type Account struct {
	ID        ID
	SecretKey string
	Status    string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// Decrypt ...
func (enc EncryptedAccount) Decrypt(m *encryption.Manager, userID user.ID) (*Account, error) {
	dekH, err := encryption.ParseAndDecryptKeyBytes(enc.DataEncryptionKey, m.Encryptor)
	if err != nil {
		return nil, err
	}
	dek := encryption.NewEncryptor(dekH)

	secretKey, err := encryption.DecryptStringIfNonNil(dek, []byte(userID), &enc.SecretKeyEncrypted)
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:        enc.ID,
		SecretKey: *secretKey,
		Status:    enc.Status,
		CreatedAt: enc.CreatedAt,
		UpdatedAt: enc.UpdatedAt,
	}, nil
}

// Encrypt ...
func (account Account) Encrypt(m *encryption.Manager, userID user.ID) (*EncryptedAccount, error) {
	dekH := encryption.NewDEK()
	dek := encryption.NewEncryptor(dekH)

	secretKeyEncrypted, err := encryption.EncryptStringIfNonNil(dek, []byte(userID), &account.SecretKey)
	if err != nil {
		return nil, err
	}

	return &EncryptedAccount{
		ID:                 account.ID,
		DataEncryptionKey:  *encryption.GetEncryptedKeyBytes(dekH, m.Encryptor),
		SecretKeyEncrypted: *secretKeyEncrypted,
		Status:             account.Status,
		CreatedAt:          account.CreatedAt,
		UpdatedAt:          account.UpdatedAt,
	}, nil
}
