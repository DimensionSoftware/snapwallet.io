package user

import (
	"time"

	"github.com/khoerling/flux/api/lib/encryption"
)

// EncryptedUser represents a user registered with our system where PII is encrypted at rest
type EncryptedUser struct {
	ID              string     `firestore:"id"`
	EncryptedEmail  *[]byte    `firestore:"encryptedEmail,omitempty"`
	EmailVerifiedAt *time.Time `firestore:"emailVerifiedAt,omitempty"`
	EncryptedPhone  *[]byte    `firestore:"encryptedPhone,omitempty"`
	PhoneVerifiedAt *time.Time `firestore:"phoneVerifiedAt,omitempty"`
	CreatedAt       time.Time  `firestore:"createdAt"`
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
	emailBytes, err := m.Decrypt(enc.EncryptedEmail)
	if err != nil {
		return nil, err
	}
	var email string
	if emailBytes != nil {
		email = string(*emailBytes)
	}

	phoneBytes, err := m.Decrypt(enc.EncryptedPhone)
	if err != nil {
		return nil, err
	}
	var phone string
	if phoneBytes != nil {
		email = string(*phoneBytes)
	}

	u := User{
		ID:              enc.ID,
		EmailVerifiedAt: enc.EmailVerifiedAt,
		PhoneVerifiedAt: enc.PhoneVerifiedAt,
		CreatedAt:       enc.CreatedAt,
	}

	if email != "" {
		u.Email = &email
	}

	if phone != "" {
		u.Phone = &phone
	}

	return &u, nil
}

// Encrypt encrypts the user
func (u *User) Encrypt(m *encryption.Manager) (*EncryptedUser, error) {
	var emailBytes *[]byte
	if u.Email != nil {
		b := []byte(*u.Email)
		emailBytes = &b

	}
	encEmailBytes, err := m.Encrypt(emailBytes)
	if err != nil {
		return nil, err
	}

	var phoneBytes *[]byte
	if u.Phone != nil {
		b := []byte(*u.Phone)
		phoneBytes = &b

	}
	encPhoneBytes, err := m.Encrypt(phoneBytes)
	if err != nil {
		return nil, err
	}

	return &EncryptedUser{
		ID:              u.ID,
		EncryptedEmail:  encEmailBytes,
		EmailVerifiedAt: u.EmailVerifiedAt,
		EncryptedPhone:  encPhoneBytes,
		PhoneVerifiedAt: u.PhoneVerifiedAt,
		CreatedAt:       u.CreatedAt,
	}, nil
}
