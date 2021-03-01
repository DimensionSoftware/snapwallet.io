package wyre

import (
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

// https://docs.sendwyre.com/docs/productiontest-environments
const wyreProductionAPIEndpoint = "https://api.sendwyre.com"
const wyreTestAPIEndpoint = "https://api.testwyre.com"

const wyreAPIKeyEnvVarName = "WYRE_API_KEY"
const wyreSecretKeyEnvVarName = "WYRE_SECRET_KEY"
const wyreAccountIDEnvVarName = "WYRE_ACCOUNT_ID"
const wyreTokenEnvVarName = "WYRE_TOKEN"

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
	http   *resty.Client
	config Config
}

// Config represents the config needed to start the client
type Config struct {
	EnableProduction bool
	WyreAPIKey       string
	WyreSecretKey    string
	WyreAccountID    string
	WyreToken        string
}

// ProvideWireConfig provides the config necessary to connect to the wyre api
func ProvideWireConfig() (*Config, error) {

	wyreAPIKey := os.Getenv(wyreAPIKeyEnvVarName)
	if wyreAPIKey == "" {
		return nil, fmt.Errorf("you must set %s", wyreAPIKeyEnvVarName)
	}
	wyreSecretKey := os.Getenv(wyreSecretKeyEnvVarName)
	if wyreSecretKeyEnvVarName == "" {
		return nil, fmt.Errorf("you must set %s", wyreAPIKeyEnvVarName)
	}
	wyreAccountID := os.Getenv(wyreAccountIDEnvVarName)
	if wyreSecretKeyEnvVarName == "" {
		return nil, fmt.Errorf("you must set %s", wyreAccountIDEnvVarName)
	}
	wyreToken := os.Getenv(wyreTokenEnvVarName)
	if wyreSecretKeyEnvVarName == "" {
		return nil, fmt.Errorf("you must set %s", wyreTokenEnvVarName)
	}

	return &Config{
		EnableProduction: false,
		WyreAPIKey:       wyreAPIKey,
		WyreSecretKey:    wyreSecretKey,
		WyreAccountID:    wyreAccountID,
		WyreToken:        wyreToken,
	}, nil
}

// NewClient instantiates a new Client
func NewClient(config *Config) Client {
	resty := resty.New()

	if config.EnableProduction {
		log.Println("🚨 Production Wyre API is activated")
		resty.SetHostURL(wyreProductionAPIEndpoint)
	} else {
		log.Println("🧪 Test Wyre API is activated")
		resty.SetHostURL(wyreTestAPIEndpoint)
	}

	return Client{
		http:   resty,
		config: *config,
	}
}

// CreateAccount creates an an account in the wyre system
// https://docs.sendwyre.com/docs/create-account
// POST https://api.sendwyre.com/v3/accounts
func (c Client) CreateAccount(req CreateAccountRequest) (*Account, error) {
	resp, err := c.http.R().
		SetBody(req).
		SetResult(Account{}).
		EnableTrace().
		Post("/v3/accounts")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*Account), nil
}

// CreatePaymentMethodRequest represents the request object for https://api.sendwyre.com/v2/paymentMethods
type CreatePaymentMethodRequest struct {
	PlaidProcessorToken string `json:"plaidProcessorToken"`
	PaymentMethodType   string `json:"paymentMethodType"`
	Country             string `json:"country"`
}

// WithDefaults provides default values for CreatePaymentMethodRequest
func (req CreatePaymentMethodRequest) WithDefaults() CreatePaymentMethodRequest {
	newReq := req
	if req.Country == "" {
		newReq.Country = "US"
	}
	if req.PaymentMethodType == "" {
		newReq.PaymentMethodType = "LOCAL_TRANSFER"
	}
	return newReq
}

// PaymentMethod represents the response object for https://api.sendwyre.com/v2/paymentMethods
type PaymentMethod struct {
	ID              string `json:"id"`
	Owner           string `json:"owner"`
	Name            string `json:"name"`
	Last4Digits     string `json:"last4Digits"`
	Status          string `json:"status"`
	CountryCode     string `json:"countryCode"`
	DefaultCurrency string `json:"defaultCurrency"`
	/*
	  "id": "PA-W7YN28ABCHT",
	  "owner": "account:AC-XX38VYXUA84",
	  "name": "Plaid Checking 0000",
	  "last4Digits": "0000",
	  "status": "PENDING",
	  "countryCode": "US",
	  "defaultCurrency": "USD",
	*/
}

// {"language":"en","compositeType":"","subType":"","errorCode":"accessDenied.invalidSession","exceptionId":"test_TQCJZP","message":"Invalid Session","type":"AccessDeniedException","transient":false}

// APIError represents the error object sent back by the api
type APIError struct {
	Language      string `json:"language"`
	Type          string `json:"type"`
	CompositeType string `json:"compositeType"`
	SubType       string `json:"subType"`
	ErrorCode     string `json:"errorCode"`
	ExceptionID   string `json:"exceptionId"`
	Message       string `json:"message"`
}

func (err APIError) Error() string {
	return fmt.Sprintf("%#v", err)
}

// CreatePaymentMethod adds a bank payment method from a plaid token to a wyre account
// https://docs.sendwyre.com/docs/ach-create-payment-method-processor-token-model
// POST https://api.sendwyre.com/v2/paymentMethods
func (c Client) CreatePaymentMethod(req CreatePaymentMethodRequest) (*PaymentMethod, error) {
	resp, err := c.http.R().
		SetHeader("Authorization", "Bearer "+c.config.WyreToken).
		SetBody(req).
		SetResult(PaymentMethod{}).
		SetError(APIError{}).
		EnableTrace().
		Post("/v2/paymentMethods")
	if err != nil {
		return nil, err
	}
	log.Printf("%#v", resp.Request)

	if resp.IsError() {
		return nil, resp.Error().(*APIError)
	}

	return resp.Result().(*PaymentMethod), nil
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
type CryptoAmounts = map[string]float64

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

// PricingRate represents rates keyed by currency symbol for a particular type of exchange
type PricingRate map[string]float32

// PricingRates represents rates across all markets
type PricingRates = map[string](PricingRate)

// PricedExchangeRates provides rates across all markets
// https://docs.sendwyre.com/docs/live-exchange-rates
// GET https://api.sendwyre.com/v3/rates
func (c Client) PricedExchangeRates() (*PricingRates, error) {
	resp, err := c.http.R().
		SetResult(PricingRates{}).
		EnableTrace().
		Get("/v2/rates?as=priced")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*PricingRates), nil
}
