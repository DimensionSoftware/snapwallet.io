package profiledata

// ID the id of a db stored ProfileData item
type ID string

// Status the status of a db stored ProfileData item
type Status string

const (
	// StatusReceived the information was received by the user; this is the initial state; the user is allowed to modify this information up until submission
	StatusReceived Status = "RECEIVED"
	// StatusPending the information is awaiting approval from a partner; when in this status the data is sealed
	StatusPending Status = "PENDING"
	// StatusInvalid the information is invalid. This data item should be converted into a remediation; when in this status the data is sealed
	StatusInvalid Status = "INVALID"
	// StatusApproved the information is approved by at least on partner; when in this status the data is sealed
	StatusApproved Status = "APPROVED"
)

// Kind the kind of ProfileData
type Kind string

const (
	// KindLegalName signifies an individuals' legal name in a ProfileDataLegalName object
	KindLegalName Kind = "LEGAL_NAME"
	// KindPhone signifies an individuals' phone number in a ProfileDataPhone object
	KindPhone Kind = "PHONE"
	// KindEmail signifies an individuals' email address in a ProfileDataEmail object
	KindEmail Kind = "EMAIL"
	// KindAddress signifies a physical address in a ProfileDataEmail object
	KindAddress Kind = "ADDRESS"
	// KindDateOfBirth signifies an individuals' date of birth in a ProfileDataDateOfBirth object
	KindDateOfBirth Kind = "DATE_OF_BIRTH"
	// KindSSN signifies an individuals' U.S. social security number in a ProfileDataSSN object
	KindSSN Kind = "SSN"
)

// EncryptedProfileData is a generic container store encrypted ProfileData
type EncryptedProfileData struct {
	ID                ID     `firestore:"id"`
	Kind              Kind   `firestore:"kind"`
	DataEncryptionKey []byte `firestore:"DEK"`
	Status            Status `firestore:"status"`
}
