//

package wyre

// ClientInterface ...
type ClientInterface interface {
	// GetPaymentMethod
	// https://docs.sendwyre.com/docs/get-payment-method
	// GET https://api.sendwyre.com/v2/paymentMethod/:paymentMethodId
	GetPaymentMethod(token string, paymentMethodID PaymentMethodID) (*PaymentMethod, error)
	// CreateAccount creates an account in the wyre system
	// https://docs.sendwyre.com/docs/create-account
	// POST https://api.sendwyre.com/v3/accounts
	CreateAccount(token string, req CreateAccountRequest) (*Account, error)
	// UploadDocument upload an account document to wyre
	// https://docs.sendwyre.com/docs/upload-document
	// POST https://api.sendwyre.com/v3/accounts/:accountId/:fieldId
	UploadDocument(token string, req UploadDocumentRequest) (*Account, error)
	// CreateAPIKey Generate a new set of API credentials for the token bearer
	// https://docs.sendwyre.com/docs/create-api-key
	// POST https://api.sendwyre.com/v2/apiKeys
	CreateAPIKey(token string, masqueradeAs AccountID, req CreateAPIKeyRequest) (*CreateAPIKeyResponse, error)
	// CreateTransfer creates a transfer in the wyre system
	// https://docs.sendwyre.com/docs/create-transfer
	// POST https://api.sendwyre.com/v3/transfers
	CreateTransfer(token string, req CreateTransferRequest) (*TransferDetail, error)
	// ConfirmTransfer confirms an existing transfer in the wyre system
	// https://docs.sendwyre.com/docs/confirm-transfer
	// POST https://api.sendwyre.com/v3/transfers/transferId:/confirm
	ConfirmTransfer(token string, req ConfirmTransferRequest) (*TransferDetail, error)
	// UpdateAccount updates a user account in the wyre system with new profile fields
	// https://docs.sendwyre.com/docs/submit-account-info
	// POST https://api.sendwyre.com/v3/accounts/:accountId
	UpdateAccount(token string, accountID AccountID, req UpdateAccountRequest) (*Account, error)
	// GetTransferHistory gets a history of transfers in the wyre system
	// https://docs.sendwyre.com/docs/transfer-history
	// GET https://api.sendwyre.com/v3/transfers
	GetTransferHistory(token string, offset int64, length int64) (*GetTransferHistoryResponse, error)
	// GetTransfer a detailed transfer record from the the wyre system
	// https://docs.sendwyre.com/docs/get-transfer
	// GET https://api.sendwyre.com/v3/transfers/:transferId
	GetTransfer(token string, transferID string) (*TransferDetail, error)
	// GetAccount gets an an account from the wyre system
	// https://docs.sendwyre.com/docs/get-account
	// GET https://api.sendwyre.com/v3/accounts/:accountId
	GetAccount(token string, accountID AccountID) (*Account, error)
	// SubscribeWebhook creates a subscription
	// Receive HTTP webhooks when subscribed objects are updated
	// https://docs.sendwyre.com/docs/subscribe-webhook
	// POST https://api.sendwyre.com/v3/subscriptions
	SubscribeWebhook(token string, subscribeTo string, notifyTarget string) (*SubscribeWebhookResponse, error)
	// CreatePaymentMethod adds a bank payment method from a plaid token to a wyre account
	// https://docs.sendwyre.com/docs/ach-create-payment-method-processor-token-model
	// POST https://api.sendwyre.com/v2/paymentMethods
	CreatePaymentMethod(token string, req CreatePaymentMethodRequest) (*PaymentMethod, error)
	// CreateWyrePaymentMethod adds a bank payment method using Wyre's Plaid integration
	// https://docs.sendwyre.com/v3/docs/local_transfer-ach-create-payment-method
	// POST https://api.sendwyre.com/v2/paymentMethods
	CreateWyrePaymentMethod(token string, req CreateWyrePaymentMethodRequest) (*PaymentMethod, error)
	// CreateWalletOrderReservation creates a wallet order reservation in Wyre's system
	// NOTE: This endpoint uses centralized authentication.
	// https://docs.sendwyre.com/v3/docs/wallet-order-reservations
	// POST https://api.sendwyre.com/v3/orders/reserve
	CreateWalletOrderReservation(req CreateWalletOrderReservationRequest) (*CreateWalletOrderReservationResponse, error)
	// CreateWalletOrderReservation creates a wallet order reservation in Wyre's system
	// NOTE: This endpoint uses centralized authentication.
	// https://docs.sendwyre.com/docs/rate-locked-reservation
	// GET https://api.sendwyre.com/v3/orders/reservation/:reservationId
	GetWalletOrderReservation(req GetWalletOrderReservationRequest) (*WalletOrderReservation, error)
	// WalletOrderDetails gets wallet order details from wyre
	// NOTE: This endpoint uses centralized authentication.
	// https://docs.sendwyre.com/docs/wallet-order-details
	// GET https://api.sendwyre.com/v3/orders/:orderId
	WalletOrderDetails(woID WalletOrderID) (*WalletOrder, error)
	// CreateWalletOrder creates a wallet order in Wyre's system
	// NOTE: This endpoint uses centralized authentication.
	// https://docs.sendwyre.com/v3/docs/white-label-card-processing-api
	// POST https://api.sendwyre.com/v3/debitcard/process/partner
	CreateWalletOrder(req CreateWalletOrderRequest) (*WalletOrder, error)
	// GetWalletOrderAuthorizations retrieves required auth mechanisms for a wallet order
	// NOTE: This endpoint uses centralized authentication.
	// https://docs.sendwyre.com/v3/docs/white-label-card-processing-api
	// POST https://api.sendwyre.com/v3/debitcard/authorization/:orderId
	GetWalletOrderAuthorizations(req GetWalletOrderAuthorizationsRequest) (*WalletOrderAuthorizations, error)
	// SubmitWalletOrderAuthorizations retrieves required auth mechanisms for a wallet order
	// NOTE: This endpoint uses centralized authentication.
	// https://docs.sendwyre.com/v3/docs/authorize-card
	// POST https://api.sendwyre.com/v3/debitcard/authorize/partner request
	SubmitWalletOrderAuthorizations(req SubmitWalletOrderAuthorizationsRequest) (*WalletOrderAuthorizationsSubmissionStatus, error)
	// PricedExchangeRates provides rates across all markets
	// https://docs.sendwyre.com/docs/live-exchange-rates
	// GET https://api.sendwyre.com/v3/rates
	PricedExchangeRates() (*PricingRates, error)
	// SubmitAuthToken
	// https://docs.sendwyre.com/docs/initialize-auth-token
	// POST https://api.sendwyre.com/v2/sessions/auth/key
	// secretKey: A 25-35 character randomly generated string to use as the key. Any valid JSON string without newlines is acceptable
	SubmitAuthToken(secretKey string) (*SubmitAuthTokenResponse, error)
}

