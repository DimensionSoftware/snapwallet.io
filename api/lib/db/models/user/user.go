package user

import (
	"time"

	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/hashing"
)

// EncryptedUser represents a user registered with our system where PII is encrypted at rest
type EncryptedUser struct {
	ID                string     `firestore:"id"`
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
	ID              string
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

	var encEmailBytes *[]byte
	var emailHash *[]byte
	if u.Email != nil {
		b := []byte(*u.Email)

		h := hashing.Hash(b)
		emailHash = &h

		encrypted, err := dek.Encrypt(b, m.AdditionalData)
		if err != nil {
			return nil, err
		}
		encEmailBytes = &encrypted
	}

	var encPhoneBytes *[]byte
	var phoneHash *[]byte
	if u.Phone != nil {
		b := []byte(*u.Phone)

		h := hashing.Hash(b)
		phoneHash = &h

		encrypted, err := dek.Encrypt(b, m.AdditionalData)
		if err != nil {
			return nil, err
		}

		encPhoneBytes = &encrypted
	}

	return &EncryptedUser{
		ID:                u.ID,
		DataEncryptionKey: encryption.GetEncryptedKeyBytes(dekH, m.Encryptor),
		EmailHash:         emailHash,
		EmailEncrypted:    encEmailBytes,
		EmailVerifiedAt:   u.EmailVerifiedAt,
		PhoneHash:         phoneHash,
		PhoneEncrypted:    encPhoneBytes,
		PhoneVerifiedAt:   u.PhoneVerifiedAt,
		CreatedAt:         u.CreatedAt,
	}, nil
}
