package auth

// RPCPublicWhitelist is a whitelist of FullMethod rpc names to allow public access to without a jwt
var RPCPublicWhitelist = map[string]bool{
	"/Flux/PricingData":           true,
	"/Flux/OneTimePasscode":       true,
	"/Flux/OneTimePasscodeVerify": true,
	// here during testing then will remove
	"/Flux/UploadFile": true,
}
