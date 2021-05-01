package auth

// RPCPublicWhitelist is a whitelist of FullMethod rpc names to allow public access to without a jwt
var RPCPublicWhitelist = map[string]bool{
	"/Flux/PricingData":           true,
	"/Flux/OneTimePasscode":       true,
	"/Flux/OneTimePasscodeVerify": true,
	"/Flux/WyreWebhook":           true,
	"/Flux/TokenExchange":         true,
	"/Flux/WidgetGetShortUrl":     true,
	"/Flux/Goto":                  true,
	// TODO: remove these once debit card integration is completed
	"/Flux/WyreCreateDebitCardQuote":         true,
	"/Flux/WyreConfirmDebitCardQuote":        true,
	"/Flux/WyreGetWalletOrderAuthorizations": true,
}
