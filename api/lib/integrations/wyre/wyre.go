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

type TransferStatusHistoryItem struct {
	ID           string `json:"id"`           // i.e. "N88AFATLRZY"
	TransferID   string `json:"transferId"`   // i.e. "TF-4F3HRUYPNFY"
	CreatedAt    int64  `json:"createdAt"`    // i.e. 1541552388000
	Type         string `json:"type"`         // i.e. "OUTGOING"
	StatusOrder  int32  `json:"statusOrder"`  // i.e. 0, 200, 400, or 5100
	StatusDetail string `json:"statusDetail"` // i.e. "Initiating Transfer"
	State        string `json:"state"`        // i.e. "INITIATED"
	FailedState  string `json:"failedState"`  // ??
}

type TransferBlockchainTx struct {
	ID            string  `json:"id"`            // i.e. "TR_XXXXXXXX1"
	CreatedAt     int64   `json:"createdAt"`     // i.e. 1541552388000
	TimeObserved  int64   `json:"timeObserved"`  // i.e. 1604344752000
	BlockTime     int64   `json:"blockTime"`     // i.e. 1604344752000
	Currency      string  `json:"currency"`      // i.e. "BTC"
	Direction     string  `json:"direction"`     // i.e. "OUTGOING",
	NetworkTxID   string  `json:"networkTxId"`   // i.e. "8ba2f644f71b2eaa5edd8e421df4a0b902e..."
	Blockhash     string  `json:"blockhash"`     // i.e. "0000000000000000000299574cd707c36..."
	Address       string  `json:"address"`       // i.e "3PZsfrHEsCooPM6BC9AXWyfxxxxxxxxxx"
	Confirmations int32   `json:"confirmations"` // i.e. 1
	Amount        float64 `json:"amount"`        // i.e. 0.00014529
	NetworkFee    float64 `json:"networkFee"`    // i.e. 0.00066185
	SourceAddress string  `json:"sourceAddress"` // i.e. "3PZsfrHEsCooPM6BC9AXWyfxxxxxxxxxx"
	TwinTxDd      string  `json:"twinTxId"`      // i.e. ??
}

type Transfer struct {
	ID             string             `json:"id"`             // i.e. "TF-4F3HRUYPNFY"
	Owner          string             `json:"owner"`          // i.e. "account:AC-WYUR7ZZ6UMU"
	CustomID       string             `json:"customId"`       // an optional custom ID to tag the transfer
	Source         string             `json:"source"`         // i.e. "account:AC-WYUR7ZZ6UMU"
	Dest           string             `json:"dest"`           // i.e. "bitcoin:14CriXWTRoJmQdBzdikw6tEmSuwxMozWWq"
	SourceAmount   float64            `json:"sourceAmount"`   // i.e. 5
	DestAmount     float64            `json:"destAmount"`     // i.e. 0.01
	SourceCurrency string             `json:"sourceCurrency"` // i.e. "USD"
	DestCurrency   string             `json:"destCurrency"`   // i.e. "BTC"
	Message        string             `json:"message"`        // i.e. "Payment for DorianNakamoto@sendwyre.com"
	Status         string             `json:"status"`         // i.e. "PENDING"
	ExchangeRate   float64            `json:"exchangeRate"`   // i.e. 499.00
	TotalFees      float64            `json:"totalFees"`      // i.e. 0.1
	Fees           map[string]float64 `json:"fees"`           // i.e. { "USD": 0.1, "BTC": 0 }
	CreatedAt      int64              `json:"createdAt"`      // i.e. 1541552388000 (epoch)
	ExpiresAt      int64              `json:"expiresAt"`      // i.e. 1541552388000 (epoch)
	CompletedAt    int64              `json:"completedAt"`    // i.e. 1541552388000 (epoch)
	CancelledAt    int64              `json:"cancelledAt"`    // i.e. 1541552388000 (epoch)

	StatusHistory []TransferStatusHistoryItem `json:"statusHistories"`
	BlockchainTx  TransferBlockchainTx        `json:"blockchainTx"`

	// not documented well... who knows if its right; we'll find out
	FailureReason      string `json:"failureReason"`
	ReversalReason     string `json:"reversalReason"`
	ReversingSubStatus string `json:"reversingSubStatus"`
	PendingSubStatus   string `json:"pendingSubStatus"`
}

type GetTransferHistoryResponse struct {
	Transfers []Transfer `json:"data"`
}

type CreateAPIKeyRequest struct {
	Description string   `json:"desc"`                  // A description of the credentials.
	Type        string   `json:"type"`                  // Can be "FULL" or "READONLY"
	IPWhitelist []string `json:"ipWhitelist,omitempty"` // List of IP addresses allowed to use the key, or empty for no restrictions.
}

type CreateAPIKeyResponse struct {
	APIKey      string   `json:"apiKey"`      // i.e. "AK-XXXX-YYYYY-ZZZZZ-QQQQQ"
	Owner       string   `json:"owner"`       // i.e. "account:AC_XYZ"
	Type        string   `json:"type"`        // i.e. "FULL"
	Description string   `json:"desc"`        // i.e. "AwesomeWallet Connection"
	SecretKey   string   `json:"secretKey"`   // i.e. "SK-ZZZZ-ZZZZ-ZZZZ-ZZZZ"
	IPWhitelist []string `json:"ipWhitelist"` // i.e. []
}

type CreateTransferRequest struct {
	Source             string  `json:"source"`                       // An SRN representing an account that the funds will be retrieved from
	Dest               string  `json:"dest"`                         // An email address, cellphone number, digital currency address or bank account to send the digital currency to. For bitcoin address use "bitcoin:[address]". Note: cellphone numbers are assumed to be a US number, for international numbers include a '+' and the country code as the prefix.
	SourceAmount       float64 `json:"sourceAmount,omitempty"`       // The amount to withdrawal from the source, in units of sourceCurrency. Only include sourceAmount OR destAmount, not both.
	DestAmount         float64 `json:"destAmount,omitempty"`         // Specifies the total amount of currency to deposit (as defined in depositCurrency). Only include sourceAmount OR destAmount, not both.
	SourceCurrency     string  `json:"sourceCurrency"`               // The currency (ISO 3166-1 alpha-3) to withdrawal from the source wallet
	DestCurrency       string  `json:"destCurrency,omitempty"`       // The currency (ISO 3166-1 alpha-3) to deposit. if not provided, the deposit will be the same as the withdrawal currency (no exchange performed)
	Message            string  `json:"message,omitempty"`            // An optional user visible message to be sent with the transaction.
	NotifyURL          string  `json:"notifyUrl,omitempty"`          // An optional url that Wyre will POST a status callback to (see Callbacks for more information)
	AutoConfirm        *bool   `json:"autoConfirm,omitempty"`        // An optional parameter to automatically confirm the transfer order.
	CustomID           string  `json:"customId,omitempty"`           // an optional custom ID to tag the transfer
	AmountIncludesFees *bool   `json:"amountIncludesFees,omitempty"` // Optional. When true, the amount indicated (source or dest) will be treated as already including the fees
	Preview            *bool   `json:"preview,omitempty"`            // creates a quote transfer object, but does not execute a real transfer.
	MuteMessages       *bool   `json:"muteMessages,omitempty"`       // When true, disables outbound emails/messages to the destination
}

type ConfirmTransferRequest struct {
	TransferID string `json:"transferID"` // The Wyre transfer identifier
}

// WithDefaults provides defaults for CreateTransferRequest
func (r CreateTransferRequest) WithDefaults() CreateTransferRequest {
	t := true
	f := false

	newR := r

	if r.AutoConfirm == nil {
		newR.AutoConfirm = &f
	}

	if r.MuteMessages == nil {
		newR.MuteMessages = &t
	}

	return newR
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

// CreateAccount creates an account in the wyre system
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

type UploadDocumentRequest struct {
	AccountID       string
	FieldID         ProfileFieldID
	DocumentType    string
	DocumentSubtype string
	MimeType        string
	Body            *[]byte
}

// UploadDocument upload an account document to wyre
// https://docs.sendwyre.com/docs/upload-document
// POST https://api.sendwyre.com/v3/accounts/:accountId/:fieldId
func (c Client) UploadDocument(token string, req UploadDocumentRequest) (*Account, error) {
	resp, err := c.http.R().
		SetAuthToken(token).
		SetError(APIError{}).
		SetResult(Account{}).
		EnableTrace().
		SetPathParam("accountID", req.AccountID).
		SetPathParam("fieldID", string(req.FieldID)).
		SetQueryParam("documentType", req.DocumentType).
		SetQueryParam("documentSubType", req.DocumentSubtype).
		SetHeader("content-type", req.MimeType).
		SetBody(*req.Body).
		Post("/v3/accounts/{accountID}/{fieldID}")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, resp.Error().(*APIError)
	}

	return resp.Result().(*Account), nil

}

// CreateAPIKey Generate a new set of API credentials for the token bearer
// https://docs.sendwyre.com/docs/create-api-key
// POST https://api.sendwyre.com/v2/apiKeys
func (c Client) CreateAPIKey(token string, masqueradeAs string, req CreateAPIKeyRequest) (*CreateAPIKeyResponse, error) {
	// to make sure we don't inadvertently create an api key linking to master key instead of intended customer key
	if masqueradeAs == "" {
		return nil, fmt.Errorf("masqueradeAs must be provided when creating an api key")
	}

	resp, err := c.http.R().
		SetHeader("Authorization", "Bearer "+token).
		SetError(APIError{}).
		SetResult(CreateAPIKeyResponse{}).
		SetBody(req).
		SetQueryParam("masqueradeAs", masqueradeAs).
		EnableTrace().
		Post("/v2/apiKeys")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, resp.Error().(*APIError)
	}

	return resp.Result().(*CreateAPIKeyResponse), nil
}

// CreateTransfer creates a transfer in the wyre system
// https://docs.sendwyre.com/docs/create-transfer
// POST https://api.sendwyre.com/v3/transfers
func (c Client) CreateTransfer(token string, req CreateTransferRequest) (*Transfer, error) {
	resp, err := c.http.R().
		SetHeader("Authorization", "Bearer "+token).
		SetError(APIError{}).
		SetResult(Transfer{}).
		SetBody(req).
		EnableTrace().
		Post("/v3/transfers")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, resp.Error().(*APIError)
	}

	return resp.Result().(*Transfer), nil
}

// CreateTransfer creates a transfer in the wyre system
// https://docs.sendwyre.com/docs/confirm-transfer
// POST https://api.sendwyre.com/v3/transfers/transferId:/confirm
func (c Client) ConfirmTransfer(token string, req ConfirmTransferRequest) (*Transfer, error) {
	reqURL := fmt.Sprintf("/v3/transfers/%s/confirm", req.TransferID)
	resp, err := c.http.R().
		SetHeader("Authorization", "Bearer "+token).
		SetError(APIError{}).
		SetResult(Transfer{}).
		SetBody(req).
		EnableTrace().
		Post(reqURL)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, resp.Error().(*APIError)
	}

	return resp.Result().(*Transfer), nil
}

// GetTransferHistory gets a history of transfers in the wyre system
// https://docs.sendwyre.com/docs/transfer-history
// GET https://api.sendwyre.com/v3/transfers
func (c Client) GetTransferHistory(token string) (*GetTransferHistoryResponse, error) {
	resp, err := c.http.R().
		SetHeader("Authorization", "Bearer "+token).
		SetError(APIError{}).
		SetResult(GetTransferHistoryResponse{}).
		EnableTrace().
		Get("/v3/transfers")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, resp.Error().(*APIError)
	}

	return resp.Result().(*GetTransferHistoryResponse), nil
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
type PricingRate map[string]float64

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
