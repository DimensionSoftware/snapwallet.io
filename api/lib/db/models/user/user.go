package user

import (
	"time"

	"github.com/khoerling/flux/api/lib/encryption"
)

// EncryptedUser represents a user registered with our system where PII is encrypted at rest
type EncryptedUser struct {
	ID              string     `firestore:"id"`
	EncryptedEmail  []byte     `firestore:"encryptedEmail"`
	EmailVerifiedAt *time.Time `firestore:"emailVerifiedAt,omitempty"`
	EncryptedPhone  []byte     `firestore:"encryptedPhone"`
	PhoneVerifiedAt *time.Time `firestore:"phoneVerifiedAt,omitempty"`
	CreatedAt       time.Time  `firestore:"createdAt"`
}

// User is the decrypted user
type User struct {
	ID              string
	Email           string
	EmailVerifiedAt *time.Time
	Phone           string
	PhoneVerifiedAt *time.Time
	CreatedAt       time.Time
}

// Decrypt decrypts the user
func (enc *EncryptedUser) Decrypt(m *encryption.Manager) (*User, error) {
	email, err := m.Decrypt(enc.EncryptedEmail)
	if err != nil {
		return nil, err
	}

	phone, err := m.Decrypt(enc.EncryptedPhone)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:              enc.ID,
		Email:           string(email),
		EmailVerifiedAt: enc.EmailVerifiedAt,
		Phone:           string(phone),
		PhoneVerifiedAt: enc.PhoneVerifiedAt,
		CreatedAt:       enc.CreatedAt,
	}, nil
}

// Encrypt encrypts the user
func (u *User) Encrypt(m *encryption.Manager) (*EncryptedUser, error) {
	email, err := m.Encrypt([]byte(u.Email))
	if err != nil {
		return nil, err
	}

	phone, err := m.Encrypt([]byte(u.Phone))
	if err != nil {
		return nil, err
	}

	return &EncryptedUser{
		ID:              u.ID,
		EncryptedEmail:  email,
		EmailVerifiedAt: u.EmailVerifiedAt,
		EncryptedPhone:  phone,
		PhoneVerifiedAt: u.PhoneVerifiedAt,
		CreatedAt:       u.CreatedAt,
	}, nil
}
