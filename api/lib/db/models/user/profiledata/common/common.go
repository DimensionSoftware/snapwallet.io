package common

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/encryption"
)

// ProfileDataID the id of a db stored ProfileData item
type ProfileDataID string

// ProfileDataStatus the status of a db stored ProfileData item
type ProfileDataStatus string

const (
	// StatusReceived the information was received by the user; this is the initial state; the user is allowed to modify this information up until submission
	StatusReceived ProfileDataStatus = "RECEIVED"
	// StatusPending the information is awaiting approval from a partner; when in this status the data is sealed
	StatusPending ProfileDataStatus = "PENDING"
	// StatusInvalid the information is invalid. This data item should be converted into a remediation; when in this status the data is sealed
	StatusInvalid ProfileDataStatus = "INVALID"
	// StatusApproved the information is approved by at least on partner; when in this status the data is sealed
	StatusApproved ProfileDataStatus = "APPROVED"
)

// ProfileDataKind the kind of ProfileData
type ProfileDataKind string

const (
	// KindLegalName signifies an individuals' legal name in a ProfileDataLegalName object
	KindLegalName ProfileDataKind = "LEGAL_NAME"
	// KindPhone signifies an individuals' phone number in a ProfileDataPhone object
	KindPhone ProfileDataKind = "PHONE"
	// KindEmail signifies an individuals' email address in a ProfileDataEmail object
	KindEmail ProfileDataKind = "EMAIL"
	// KindAddress signifies a physical address in a ProfileDataAddress object
	KindAddress ProfileDataKind = "ADDRESS"
	// KindDateOfBirth signifies an individuals' date of birth in a ProfileDataDateOfBirth object
	KindDateOfBirth ProfileDataKind = "DATE_OF_BIRTH"
	// KindSSN signifies an individuals' U.S. social security number in a ProfileDataSSN object
	KindSSN ProfileDataKind = "SSN"
)

// EncryptedProfileData is a generic container store encrypted ProfileData
type EncryptedProfileData struct {
	ID                ProfileDataID     `firestore:"id"`
	Kind              ProfileDataKind   `firestore:"kind"`
	Status            ProfileDataStatus `firestore:"status"`
	DataEncryptionKey []byte            `firestore:"DEK"`
	EncryptedData     []byte            `firestore:"encryptedData"`
	CreatedAt         time.Time         `firestore:"createdAt"`
	UpdatedAt         *time.Time        `firestore:"updatedAt,omitempty"`
	SealedAt          *time.Time        `firestore:"sealedAt,omitempty"`
}

// Decrypt decrypts a type
func (encryptedProfileData EncryptedProfileData) Decrypt(m *encryption.Manager, userID user.ID) ([]byte, error) {
	dekH, err := encryption.ParseAndDecryptKeyBytes(encryptedProfileData.DataEncryptionKey, m.Encryptor)
	if err != nil {
		return nil, err
	}

	dek := encryption.NewEncryptor(dekH)
	decrypted, err := dek.Decrypt(encryptedProfileData.EncryptedData, []byte(userID))
	if err != nil {
		return nil, err
	}

	return decrypted, nil

}
