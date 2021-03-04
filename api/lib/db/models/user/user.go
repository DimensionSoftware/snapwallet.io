package user

import (
	"time"

	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/hashing"
)

// ID ...
type ID string

// EncryptedUser represents a user registered with our system where PII is encrypted at rest
type EncryptedUser struct {
	ID                ID         `firestore:"id"`
	DataEncryptionKey []byte     `firestore:"DEK"`
	EmailHash         *[]byte    `firestore:"emailHash,omitempty"`
	EmailEncrypted    *[]byte    `firestore:"emailEncrypted,omitempty"`
	EmailVerifiedAt   *time.Time `firestore:"emailVerifiedAt,omitempty"`
	PhoneHash         *[]byte    `firestore:"phoneHash,omitempty"`
	PhoneEncrypted    *[]byte    `firestore:"phoneEncrypted,omitempty"`
	PhoneVerifiedAt   *time.Time `firestore:"phoneVerifiedAt,omitempty"`
	CreatedAt         time.Time  `firestore:"createdAt"`
}

// User is the decrypted user
type User struct {
	ID              ID
	Email           *string
	EmailVerifiedAt *time.Time
	Phone           *string
	PhoneVerifiedAt *time.Time
	CreatedAt       time.Time
}

// Decrypt decrypts the user
func (enc *EncryptedUser) Decrypt(m *encryption.Manager) (*User, error) {
	dekH, err := encryption.ParseAndDecryptKeyBytes(enc.DataEncryptionKey, m.Encryptor)
	if err != nil {
		return nil, err
	}
	dek := encryption.NewEncryptor(dekH)

	email, err := encryption.DecryptStringIfNonNil(dek, m.AdditionalData, enc.EmailEncrypted)
	if err != nil {
		return nil, err
	}

	phone, err := encryption.DecryptStringIfNonNil(dek, m.AdditionalData, enc.PhoneEncrypted)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:              enc.ID,
		Email:           email,
		EmailVerifiedAt: enc.EmailVerifiedAt,
		Phone:           phone,
		PhoneVerifiedAt: enc.PhoneVerifiedAt,
		CreatedAt:       enc.CreatedAt,
	}, nil
}

// Encrypt encrypts the user
func (u *User) Encrypt(m *encryption.Manager) (*EncryptedUser, error) {
	dekH := encryption.NewDEK()
	dek := encryption.NewEncryptor(dekH)

	emailEncrypted, err := encryption.EncryptStringIfNonNil(dek, m.AdditionalData, u.Email)
	if err != nil {
		return nil, err
	}

	phoneEncrypted, err := encryption.EncryptStringIfNonNil(dek, m.AdditionalData, u.Phone)
	if err != nil {
		return nil, err
	}

	var emailHash *[]byte
	if emailEncrypted != nil {
		b := []byte(*u.Email)

		h := hashing.Hash(b)
		emailHash = &h
	}

	var phoneHash *[]byte
	if phoneEncrypted != nil {
		b := []byte(*u.Phone)

		h := hashing.Hash(b)
		phoneHash = &h
	}

	return &EncryptedUser{
		ID:                u.ID,
		DataEncryptionKey: encryption.GetEncryptedKeyBytes(dekH, m.Encryptor),
		EmailHash:         emailHash,
		EmailEncrypted:    emailEncrypted,
		EmailVerifiedAt:   u.EmailVerifiedAt,
		PhoneHash:         phoneHash,
		PhoneEncrypted:    phoneEncrypted,
		PhoneVerifiedAt:   u.PhoneVerifiedAt,
		CreatedAt:         u.CreatedAt,
	}, nil
}
