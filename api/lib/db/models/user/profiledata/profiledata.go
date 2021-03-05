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
	KindAddress Kind = "ADDRESS"
)
