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

// ProfileField represents PII data which is used during the create account process
type ProfileField struct {
	FieldID ProfileFieldID `json:"fieldId"`
	Value   interface{}    `json:"value"`
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
	// The type of account, currently INDIVIDUAL is the only supported value
	Type string `json:"type"`
	// The country of the account holder. For individuals this is the country of residence. (Currently we only support US accounts)
	Country string `json:"country"`
	// An array of the Fields submitted at the time of Account creation. You can submit as many or as few fields as you need at the time of Account creation
	ProfileFields []ProfileField `json:"profileFields"`
	// Supply your own Account ID when creating noncustodial accounts. This field is used to track which account referred the new account into our system
	ReferrerAccountID *string `json:"referrerAccountId,omitempty"`
	// When true, the newly created account will be a custodial subaccount owner by the caller. Otherwise, the account will be a standalone non-custodial account.
	SubAccount *bool `json:"subaccount,omitempty"`
	// If true prevents all outbound emails to the account.
	DisableEmail *bool `json:"disableEmail,omitempty"`
}

// WithDefaults provides default values for CreateAccountRequest
func (req CreateAccountRequest) WithDefaults() CreateAccountRequest {
	newReq := req

	if req.Country == "" {
		newReq.Country = "US" // only supported country currently
	}

	if req.Type == "" {
		newReq.Type = "INDIVIDUAL" // only supported type currently
	}

	return newReq
}

// ProfileFieldID field ID for create account request
type ProfileFieldID string

const (
	// ProfileFieldIDIndividualLegalName indicates the value is a legal name string
	ProfileFieldIDIndividualLegalName ProfileFieldID = "individualLegalName"
	// ProfileFieldIDIndividualEmail indicates the value is an email string
	ProfileFieldIDIndividualEmail ProfileFieldID = "individualEmail"
	// ProfileFieldIDIndividualCellphoneNumber indicates the value is a phone string
	ProfileFieldIDIndividualCellphoneNumber ProfileFieldID = "individualCellphoneNumber"
	// ProfileFieldIDIndividualDateOfBirth indiciates the value is a date of birth string
	ProfileFieldIDIndividualDateOfBirth ProfileFieldID = "individualDateOfBirth"
	// ProfileFieldIDIndividualSSN indicates the value is an ssn string
	ProfileFieldIDIndividualSSN ProfileFieldID = "individualSsn"
	// ProfileFieldIDIndividualResidenceAddress indicates the value is a ProfileFieldAddress object
	ProfileFieldIDIndividualResidenceAddress ProfileFieldID = "individualResidenceAddress"
	// ProfileFieldIDIndividualGovernmentID indicates the value is an uploaded document
	ProfileFieldIDIndividualGovernmentID ProfileFieldID = "individualGovernmentId"
	// ProfileFieldIDIndividualProofOfAddress indicates the value is an uploaded document
	ProfileFieldIDIndividualProofOfAddress ProfileFieldID = "individualProofOfAddress"
	// ProfileFieldIDIndividualACHAuthorizationForm indicates the value is an uploaded document
	ProfileFieldIDIndividualACHAuthorizationForm ProfileFieldID = "individualAchAuthorizationForm"
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
}

// ProvideWyreConfig provides the config necessary to connect to the wyre api
func ProvideWyreConfig() (*Config, error) {

	wyreAPIKey := os.Getenv(wyreAPIKeyEnvVarName)
	if wyreAPIKey == "" {
		return nil, fmt.Errorf("you must set %s", wyreAPIKeyEnvVarName)
	}
	wyreSecretKey := os.Getenv(wyreSecretKeyEnvVarName)
	if wyreSecretKey == "" {
		return nil, fmt.Errorf("you must set %s", wyreSecretKeyEnvVarName)
	}
	wyreAccountID := os.Getenv(wyreAccountIDEnvVarName)
	if wyreAccountID == "" {
		return nil, fmt.Errorf("you must set %s", wyreAccountIDEnvVarName)
	}

	return &Config{
		EnableProduction: false,
		WyreAPIKey:       wyreAPIKey,
		WyreSecretKey:    wyreSecretKey,
		WyreAccountID:    wyreAccountID,
	}, nil
}

// NewClient instantiates a new Client
func NewClient(config *Config) *Client {
	resty := resty.New()

	if config.EnableProduction {
		log.Println("🚨 Production Wyre API is activated")
		resty.SetHostURL(wyreProductionAPIEndpoint)
	} else {
		log.Println("🧪 Test Wyre API is activated")
		resty.SetHostURL(wyreTestAPIEndpoint)
	}

	return &Client{
		http:   resty,
		config: *config,
	}
}

// CreateAccount creates an an account in the wyre system
// https://docs.sendwyre.com/docs/create-account
// POST https://api.sendwyre.com/v3/accounts
func (c Client) CreateAccount(token string, req CreateAccountRequest) (*Account, error) {
	resp, err := c.http.R().
		SetHeader("Authorization", "Bearer "+token).
		SetError(APIError{}).
		SetResult(Account{}).
		SetBody(req).
		EnableTrace().
		Post("/v3/accounts")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, resp.Error().(*APIError)
	}

	return resp.Result().(*Account), nil
}

// GetAccount gets an an account from the wyre system
// https://docs.sendwyre.com/docs/get-account
// GET https://api.sendwyre.com/v3/accounts/:accountId
func (c Client) GetAccount(token string, accountID string) (*Account, error) {
	resp, err := c.http.R().
		SetHeader("Authorization", "Bearer "+token).
		SetError(APIError{}).
		SetResult(Account{}).
		EnableTrace().
		SetPathParam("accountID", accountID).
		Get("/v3/accounts/{accountID}")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, resp.Error().(*APIError)
	}

	return resp.Result().(*Account), nil
}

// SubscribeWebhook creates a subscription
// Receive HTTP webhooks when subscribed objects are updated
// https://docs.sendwyre.com/docs/subscribe-webhook
// POST https://api.sendwyre.com/v3/subscriptions
func (c Client) SubscribeWebhook(token string, subscribeTo string, notifyTarget string) (*SubscribeWebhookResponse, error) {
	req := map[string]string{
		"subscribeTo":  subscribeTo,
		"notifyTarget": notifyTarget,
	}

	resp, err := c.http.R().
		SetHeader("Authorization", "Bearer "+token).
		SetError(APIError{}).
		SetResult(SubscribeWebhookResponse{}).
		SetBody(req).
		EnableTrace().
		Post("/v3/subscriptions")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, resp.Error().(*APIError)
	}

	return resp.Result().(*SubscribeWebhookResponse), nil
}

type SubscribeWebhookResponse struct {
	// "id": "JF4DQ2NE1",
	ID string `json:"id"`
	// "subscribed": "account:AC-F930QD8A2RRR",
	Subscribed string `json:"subscribed"`
	// "notifyTarget": "https://www.potatoes.com/webhook1",
	NotifyTarget string `json:"notifyTarget"`
	// "createdAt": 1548368619000,
	CreatedAt int64 `json:"createdAt"`
	// "failure": null,
	Failure *interface{} `json:"failure,omitempty"`
	// "failCount": 0
	FailCount int32 `json:"failCount"`
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
	ID                    string   `json:"id"`
	Owner                 string   `json:"owner"`
	Name                  string   `json:"name"`
	Last4Digits           string   `json:"last4Digits"`
	Status                string   `json:"status"`
	CountryCode           string   `json:"countryCode"`
	DefaultCurrency       string   `json:"defaultCurrency"`
	ChargeableCurrencies  []string `json:"chargeableCurrencies"`
	DepositableCurrencies []string `json:"depositableCurrencies"`
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
func (c Client) CreatePaymentMethod(token string, req CreatePaymentMethodRequest) (*PaymentMethod, error) {
	resp, err := c.http.R().
		SetHeader("Authorization", "Bearer "+token).
		SetBody(req).
		SetResult(PaymentMethod{}).
		SetError(APIError{}).
		EnableTrace().
		Post("/v2/paymentMethods")
	if err != nil {
		return nil, err
	}

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
	ProfileFields     []ProfileData   `json:"profileFields"`
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

// SubmitAuthToken
// https://docs.sendwyre.com/docs/initialize-auth-token
// POST https://api.sendwyre.com/v2/sessions/auth/key
// secretKey: A 25-35 character randomly generated string to use as the key. Any valid JSON string without newlines is acceptable
func (c Client) SubmitAuthToken(secretKey string) (*SubmitAuthTokenResponse, error) {
	req := map[string]string{
		"secretKey": secretKey,
	}

	resp, err := c.http.R().
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		SetResult(SubmitAuthTokenResponse{}).
		SetError(APIError{}).
		EnableTrace().
		Post("/v2/sessions/auth/key")

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, resp.Error().(*APIError)
	}

	return resp.Result().(*SubmitAuthTokenResponse), nil
}

type SubmitAuthTokenResponse struct {
	APIKey          string      `json:"apiKey"`
	AuthenticatedAs interface{} `json:"authenticatedAs"`
}
