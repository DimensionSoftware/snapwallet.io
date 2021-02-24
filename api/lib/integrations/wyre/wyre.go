package wyre

// ProfileField represents PII data which is used during the create account process
type ProfileField struct {
	FieldID   string      `json:"fieldId"`
	FieldType string      `json:"fieldType,omitempty"` // is only on output, not on input
	Value     interface{} `json:"value"`
}

// ProfileFieldAddress represents the profile field with the fieldType "ADDRESS"
type ProfileFieldAddress struct {
	Street1    string `json:"street1"`
	Street2    string `json:"street2"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

// CreateAccountRequest represents the request object for https://api.sendwyre.com/v3/accounts
type CreateAccountRequest struct {
	Country       string         `json:"country"`
	SubAccount    bool           `json:"subaccount"`
	Type          string         `json:"type"`
	ProfileFields []ProfileField `json:"profileFields"`
}

// WithDefaults provides default values for CreateAccountRequest
func (CreateAccountRequest) WithDefaults() CreateAccountRequest {
	return CreateAccountRequest{
		Country:    "US",         // only supported country currently
		Type:       "INDIVIDUAL", // only supported type currently
		SubAccount: true,         // figure all accounts will be created as a sub account
	}
}

// ProfileFieldID field ID for create account request
type ProfileFieldID string

const (
	// ProfileFieldIDIndividualLegalName indicates the value is a legal name string
	ProfileFieldIDIndividualLegalName ProfileFieldID = "individualLegalName"
	// ProfileFieldIDIndividualEmail indicates the value is an email string
	ProfileFieldIDIndividualEmail ProfileFieldID = "individualEmail"
	// ProfileFieldIDIndividualResidenceAddress indicates the value is a ProfileFieldAddress object
	ProfileFieldIDIndividualResidenceAddress ProfileFieldID = "individualResidenceAddress"
)

/*
ex:
{
    "country": "US",
    "profileFields": [
        {
            "fieldId": "individualLegalName",
            "value": "YOUR_NAME"
        },
        {
            "fieldId": "individualEmail",
            "value": "YOUREMAIL@EMAIL.com"
        },
        {
            "fieldId": "individualResidenceAddress",
            "value": {
                "city": "San Francisco",
                "country": "US",
                "postalCode": "94105",
                "state": "CA",
                "street1": "1 Market St",
                "street2": "Suite 402"
            }
        }
    ],
    "subaccount": true,
    "type": "INDIVIDUAL"
}

*/

// CreateAccount https://docs.sendwyre.com/docs/create-account
