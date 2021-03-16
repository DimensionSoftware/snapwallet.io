package common

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/file"
	"github.com/khoerling/flux/api/lib/encryption"
	proto "github.com/khoerling/flux/api/lib/protocol"
)

// ProfileDataID the id of a db stored ProfileData item
type ProfileDataID string

// ProfileDataStatus the status of a db stored ProfileData item
type ProfileDataStatus string

// ToProfileDataItemStatus converts the profile data status to a protocol.ProfileDataItemStatus
func (s ProfileDataStatus) ToProfileDataItemStatus() proto.ProfileDataItemStatus {
	switch s {
	case StatusReceived:
		return proto.ProfileDataItemStatus_S_RECEIVED
	case StatusPending:
		return proto.ProfileDataItemStatus_S_PENDING
	case StatusInvalid:
		return proto.ProfileDataItemStatus_S_INVALID
	case StatusApproved:
		return proto.ProfileDataItemStatus_S_APPROVED
	}

	return proto.ProfileDataItemStatus_S_UNKNOWN
}

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

// ToProfileDataItemKind converts the profile data kind to a protocol.ProfileDataItemKind
func (k ProfileDataKind) ToProfileDataItemKind() proto.ProfileDataItemKind {
	switch k {
	case KindLegalName:
		return proto.ProfileDataItemKind_K_LEGAL_NAME
	case KindPhone:
		return proto.ProfileDataItemKind_K_PHONE
	case KindEmail:
		return proto.ProfileDataItemKind_K_EMAIL
	case KindAddress:
		return proto.ProfileDataItemKind_K_ADDRESS
	case KindDateOfBirth:
		return proto.ProfileDataItemKind_K_DATE_OF_BIRTH
	case KindUSSSN:
		return proto.ProfileDataItemKind_K_US_SSN
	case KindUSGovernmentIDDoc:
		return proto.ProfileDataItemKind_K_US_GOVERNMENT_ID_DOC
	case KindProofOfAddressDoc:
		return proto.ProfileDataItemKind_K_PROOF_OF_ADDRESS_DOC
	}

	return proto.ProfileDataItemKind_K_UNKNOWN
}

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
	KindUSSSN ProfileDataKind = "US_SSN"
	// KindGovernmentID signifies an individuals' government id in a ProfileDataGovernmentID object
	KindUSGovernmentIDDoc ProfileDataKind = "US_GOVERNMENT_ID_DOC"
	KindProofOfAddressDoc ProfileDataKind = "PROOF_OF_ADDRESS_DOC"
)

var ProfileDataKinds = []ProfileDataKind{
	KindLegalName,
	KindPhone,
	KindEmail,
	KindAddress,
	KindDateOfBirth,
	KindUSSSN,
	KindUSGovernmentIDDoc,
}

// in the future it will be a subset of profiledatakinds
//var ProfileDataRequiredForWyre = ProfileDataKinds
var ProfileDataRequiredForWyre = []ProfileDataKind{
	KindLegalName,
	KindDateOfBirth,
	KindUSSSN,
}

// EncryptedProfileData is a generic container store encrypted ProfileData
type EncryptedProfileData struct {
	ID                ProfileDataID     `firestore:"id"`
	Kind              ProfileDataKind   `firestore:"kind"`
	SubKind           *string           `firestore:"subKind,omitempty"`
	Status            ProfileDataStatus `firestore:"status"`
	CreatedAt         time.Time         `firestore:"createdAt"`
	DataEncryptionKey *[]byte           `firestore:"DEK,omitempty"`
	EncryptedData     *[]byte           `firestore:"encryptedData,omitempty"`
	FileIDs           *[]file.ID        `firestore:"fileIds,omitempty"`
	UpdatedAt         *time.Time        `firestore:"updatedAt,omitempty"`
	SealedAt          *time.Time        `firestore:"sealedAt,omitempty"`
}

// Decrypt decrypts a type
func (encryptedProfileData EncryptedProfileData) Decrypt(m *encryption.Manager, userID user.ID) (*[]byte, error) {
	if encryptedProfileData.DataEncryptionKey == nil {
		return nil, nil
	}

	if encryptedProfileData.EncryptedData == nil {
		return nil, nil
	}

	dekH, err := encryption.ParseAndDecryptKeyBytes(*encryptedProfileData.DataEncryptionKey, m.Encryptor)
	if err != nil {
		return nil, err
	}
	dek := encryption.NewEncryptor(dekH)

	decrypted, err := dek.Decrypt(*encryptedProfileData.EncryptedData, []byte(userID))
	if err != nil {
		return nil, err
	}

	return &decrypted, nil

}
