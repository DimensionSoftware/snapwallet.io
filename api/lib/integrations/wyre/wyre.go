package wyre

import "github.com/go-resty/resty/v2"

// ProfileField represents PII data which is used during the create account process
type ProfileField struct {
	FieldID string      `json:"fieldId"`
	Value   interface{} `json:"value"`
}

// ProfileData represents PII data which is stored after the account create process
type ProfileData struct {
	FieldID   string      `json:"fieldId"`
	FieldType string      `json:"fieldType"`
	Value     interface{} `json:"value"`
	Note      string      `json:"note"`
	Status    string      `json:"status"`
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

// Client is the client interface for wyre
type Client struct {
	http *resty.Client
}

// NewClient instantiates a new Client
func NewClient() Client {
	return Client{
		http: resty.New(),
	}
}

// CreateAccount creates an an account in the wyre system
// POST https://api.sendwyre.com/v3/accounts
func (c Client) CreateAccount(req CreateAccountRequest) (*Account, error) {
	resp, err := c.http.R().
		SetBody(req).
		SetResult(Account{}).
		EnableTrace().
		Post("https://api.sendwyre.com/v3/accounts")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*Account), nil
}

/*
req ex:
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

// CryptoAddresses represents crypto addresses where account deposits are held; keyed by crypto symbol
type CryptoAddresses = map[string]string

// CryptoAmounts represents crypto amounts; keyed by crypto symbol
type CryptoAmounts = map[string]string

// Account represents the response object for https://api.sendwyre.com/v3/accounts
type Account struct {
	ID                string          `json:"id"`
	Status            string          `json:"status"`
	Type              string          `json:"type"`
	Country           string          `json:"country"`
	CreatedAt         int64           `json:"createdAt"`
	DepositAddresses  CryptoAddresses `json:"depositAddresses"`
	TotalBalances     CryptoAmounts   `json:"totalBalances"`
	AvailableBalances CryptoAmounts   `json:"availableBalances"`
	ProfileData       []ProfileData   `json:"profileData"`
}

/*
resp ex:
{
  "id" : "AC-U4BWHGZDG6W",
  "status" : "PENDING",
  "type" : "INDIVIDUAL",
  "country" : "US",
  "createdAt" : 1541789972000,
  "depositAddresses" : {
    "ETH" : "0x98B031783d0efb1E65C4072C6576BaCa0736A912",
    "BTC" : "14CriXWTRoJmQdBzdikw6tEmSuwxMozWWq"
  },
  "totalBalances" : {
    "BTC" : 1.0000000,
    "ETH" : 0.1000000000000000000
  },
  "availableBalances" : {
    "BTC" : 1.0000000,
    "ETH" : 0.1000000000000000000
  },
  "profileData" : [ {
    "fieldId" : "individualCellphoneNumber",
    "fieldType" : "CELLPHONE",
    "value" : null,
    "note" : "Must be verified by user.",
    "status" : "OPEN"
  }, {
    "fieldId" : "individualEmail",
    "fieldType" : "EMAIL",
    "value" : "johnnyquest22@yolo.com",
    "note" : "Must be verified by user.",
    "status" : "OPEN"
  }, {
    "fieldId" : "individualLegalName",
    "fieldType" : "STRING",
    "value" : "Johnny Quest",
    "note" : null,
    "status" : "PENDING"
  }, {
    "fieldId" : "individualDateOfBirth",
    "fieldType" : "DATE",
    "value" : null,
    "note" : null,
    "status" : "OPEN"
  }, {
    "fieldId" : "individualSsn",
    "fieldType" : "STRING",
    "value" : null,
    "note" : null,
    "status" : "NULL"
  }, {
    "fieldId" : "individualResidenceAddress",
    "fieldType" : "ADDRESS",
    "value" : {
        "street1": "1 Market St",
        "street2": "Suite 402",
        "city": "San Francisco",
        "state": "CA",
        "postalCode": "94105",
        "country": "US"
    },
    "note" : null,
    "status" : "PENDING"
  }, {
    "fieldId" : "individualGovernmentId",
    "fieldType" : "DOCUMENT",
    "value" : [],
    "note" : null,
    "status" : "OPEN"
  }, {
    "fieldId" : "individualSourceOfFunds",
    "fieldType" : "PAYMENT_METHOD",
    "value" : null,
    "note" : "Payment method not yet submitted",
    "status" : "OPEN"
  } ]
}
*/

// CreateAccount https://docs.sendwyre.com/docs/create-account
