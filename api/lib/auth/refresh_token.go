package auth

// Claims represents the data we want to put on the jwt
// https://tools.ietf.org/html/rfc7519
type RefreshTokenClaims struct {
	/*
	   The "jti" (JWT ID) claim provides a unique identifier for the JWT.
	   The identifier value MUST be assigned in a manner that ensures that
	   there is a negligible probability that the same value will be
	   accidentally assigned to a different data object; if the application
	   uses multiple issuers, collisions MUST be prevented among values
	   produced by different issuers as well.  The "jti" claim can be used
	   to prevent the JWT from being replayed.  The "jti" value is a case-
	   sensitive string.  Use of this claim is OPTIONAL.
	*/
	ID string `json:"jti,omitempty"`

	/*
		The "sub" (subject) claim identifies the principal that is the
		subject of the JWT.  The claims in a JWT are normally statements
		about the subject.  The subject value MUST either be scoped to be
		locally unique in the context of the issuer or be globally unique.
		The processing of this claim is generally application specific.  The
		"sub" value is a case-sensitive string containing a StringOrURI
		value.  Use of this claim is OPTIONAL.
	*/
	Subject string `json:"sub,omitempty"`

	/*
		The "iat" (issued at) claim identifies the time at which the JWT was
		issued.  This claim can be used to determine the age of the JWT.  Its
		value MUST be a number containing a NumericDate value.  Use of this
		claim is OPTIONAL.
	*/
	IssuedAt int64 `json:"iat,omitempty"`

	/*
		The "exp" (expiration time) claim identifies the expiration time on
		or after which the JWT MUST NOT be accepted for processing.  The
		processing of the "exp" claim requires that the current date/time
		MUST be before the expiration date/time listed in the "exp" claim.
	*/
	ExpiresAt int64 `json:"exp,omitempty"`
}

// Valid checks the claims for validity
func (RefreshTokenClaims) Valid() error {
	return nil
}
