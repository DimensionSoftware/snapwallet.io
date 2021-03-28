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

// FromProfileDataItemKind converts from proto.ProfileDataItemKind to a profile data kind
func FromProfileDataItemKind(pdik proto.ProfileDataItemKind) ProfileDataKind {
	switch pdik {
	case proto.ProfileDataItemKind_K_LEGAL_NAME:
		return KindLegalName
	case proto.ProfileDataItemKind_K_PHONE:
		return KindPhone
	case proto.ProfileDataItemKind_K_EMAIL:
		return KindEmail
	case proto.ProfileDataItemKind_K_ADDRESS:
		return KindAddress
	case proto.ProfileDataItemKind_K_DATE_OF_BIRTH:
		return KindDateOfBirth
	case proto.ProfileDataItemKind_K_US_SSN:
		return KindUSSSN
	case proto.ProfileDataItemKind_K_US_GOVERNMENT_ID_DOC:
		return KindUSGovernmentIDDoc
	case proto.ProfileDataItemKind_K_PROOF_OF_ADDRESS_DOC:
		return KindProofOfAddressDoc
	}
	// should never get here
	panic("proto.ProfileDataItemKind unknown when FromProfileDataItemKind(...) called")
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
	KindProofOfAddressDoc,
}

// in the future it will be a subset of profiledatakinds
//var ProfileDataRequiredForWyre = ProfileDataKinds
var ProfileDataRequiredForWyre = []ProfileDataKind{
	KindLegalName,
	KindEmail,
	KindPhone,
	KindAddress,
	KindDateOfBirth,
	KindUSSSN,
	KindUSGovernmentIDDoc,
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
	EncryptedNote     *[]byte           `firestore:"encryptedNote,omitempty"`
	FileIDs           *[]file.ID        `firestore:"fileIds,omitempty"`
	UpdatedAt         *time.Time        `firestore:"updatedAt,omitempty"`
	SealedAt          *time.Time        `firestore:"sealedAt,omitempty"`
}

type CommonProfileData struct {
	ID        ProfileDataID
	Status    ProfileDataStatus
	Note      string
	CreatedAt time.Time
	UpdatedAt *time.Time
	SealedAt  *time.Time
}

func (c CommonProfileData) GetID() ProfileDataID {
	return c.ID
}

func (c CommonProfileData) GetStatus() ProfileDataStatus {
	return c.Status
}

func (c CommonProfileData) GetNote() string {
	return c.Note
}

// Decrypt decrypts a type; returns encrypted data, note data
func (encryptedProfileData EncryptedProfileData) Decrypt(m *encryption.Manager, userID user.ID) (*[]byte, *[]byte, error) {
	if encryptedProfileData.DataEncryptionKey == nil {
		return nil, nil, nil
	}

	dekH, err := encryption.ParseAndDecryptKeyBytes(*encryptedProfileData.DataEncryptionKey, m.Encryptor)
	if err != nil {
		return nil, nil, err
	}
	dek := encryption.NewEncryptor(dekH)

	var (
		outData *[]byte
		outNote *[]byte
	)

	if encryptedProfileData.EncryptedData != nil {
		decryptedData, err := dek.Decrypt(*encryptedProfileData.EncryptedData, []byte(userID))
		if err != nil {
			return nil, nil, err
		}
		outData = &decryptedData
	}

	if encryptedProfileData.EncryptedNote != nil {
		decryptedNote, err := dek.Decrypt(*encryptedProfileData.EncryptedNote, []byte(userID))
		if err != nil {
			return nil, nil, err
		}
		outNote = &decryptedNote
	}

	return outData, outNote, nil

}
