package user

import (
	"time"

	"github.com/google/tink/go/tink"
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/hashing"
	"github.com/rs/xid"
)

// PhoneEncrypted represents encrypted phone number
type PhoneEncrypted []byte

// Phone represents cleartext phone number
type Phone string

// Decrypt ...
func (enc *PhoneEncrypted) Decrypt(dek tink.AEAD, userID ID) (*Phone, error) {
	if enc == nil || len(*enc) == 0 {
		return nil, nil
	}

	cleartext, err := encryption.DecryptStringIfNonNil(dek, []byte(userID), (*[]byte)(enc))
	if err != nil {
		return nil, err
	}

	return (*Phone)(cleartext), nil
}

// Decrypt ...
func (enc *EmailEncrypted) Decrypt(dek tink.AEAD, userID ID) (*Email, error) {
	if enc == nil || len(*enc) == 0 {
		return nil, nil
	}

	cleartext, err := encryption.DecryptStringIfNonNil(dek, []byte(userID), (*[]byte)(enc))
	if err != nil {
		return nil, err
	}

	return (*Email)(cleartext), nil
}

// EmailEncrypted represents encrypted email
type EmailEncrypted []byte

// Email represents cleartext email
type Email string

// ID ...
type ID string

// EncryptedUser represents a user registered with our system where PII is encrypted at rest
type EncryptedUser struct {
	ID                ID              `firestore:"id"`
	DataEncryptionKey []byte          `firestore:"DEK"`
	EmailHash         *[]byte         `firestore:"emailHash,omitempty"`
	EmailEncrypted    *EmailEncrypted `firestore:"emailEncrypted,omitempty"`
	EmailVerifiedAt   *time.Time      `firestore:"emailVerifiedAt,omitempty"`
	PhoneHash         *[]byte         `firestore:"phoneHash,omitempty"`
	PhoneEncrypted    *PhoneEncrypted `firestore:"phoneEncrypted,omitempty"`
	PhoneVerifiedAt   *time.Time      `firestore:"phoneVerifiedAt,omitempty"`
	CreatedAt         time.Time       `firestore:"createdAt"`
}

// User is the decrypted user
type User struct {
	ID              ID
	Email           *Email
	EmailVerifiedAt *time.Time
	Phone           *Phone
	PhoneVerifiedAt *time.Time
	CreatedAt       time.Time
}

// WithDefaults provides defaults for User
func (u User) WithDefaults() User {
	newU := u
	if u.ID == "" {
		newU.ID = ID(xid.New().String())
	}

	if (u.CreatedAt == time.Time{}) {
		newU.CreatedAt = time.Now()
	}

	return newU
}

// Decrypt decrypts the user
func (enc *EncryptedUser) Decrypt(m *encryption.Manager, userID ID) (*User, error) {
	dekH, err := encryption.ParseAndDecryptKeyBytes(enc.DataEncryptionKey, m.Encryptor)
	if err != nil {
		return nil, err
	}
	dek := encryption.NewEncryptor(dekH)

	email, err := enc.EmailEncrypted.Decrypt(dek, userID)
	if err != nil {
		return nil, err
	}

	phone, err := enc.PhoneEncrypted.Decrypt(dek, userID)
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
func (u *User) Encrypt(m *encryption.Manager, userID ID) (*EncryptedUser, error) {
	dekH := encryption.NewDEK()
	dek := encryption.NewEncryptor(dekH)

	emailEncrypted, err := encryption.EncryptStringIfNonNil(dek, []byte(userID), (*string)(u.Email))
	if err != nil {
		return nil, err
	}

	phoneEncrypted, err := encryption.EncryptStringIfNonNil(dek, []byte(userID), (*string)(u.Phone))
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
		EmailEncrypted:    (*EmailEncrypted)(emailEncrypted),
		EmailVerifiedAt:   u.EmailVerifiedAt,
		PhoneHash:         phoneHash,
		PhoneEncrypted:    (*PhoneEncrypted)(phoneEncrypted),
		PhoneVerifiedAt:   u.PhoneVerifiedAt,
		CreatedAt:         u.CreatedAt,
	}, nil
}
