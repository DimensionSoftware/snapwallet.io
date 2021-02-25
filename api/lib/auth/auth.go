package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/xid"
)

// Claims represents the data we want to put on the jwt
// https://tools.ietf.org/html/rfc7519
type Claims struct {
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
	Jti string `json:"jti"`

	/*
		The "sub" (subject) claim identifies the principal that is the
		subject of the JWT.  The claims in a JWT are normally statements
		about the subject.  The subject value MUST either be scoped to be
		locally unique in the context of the issuer or be globally unique.
		The processing of this claim is generally application specific.  The
		"sub" value is a case-sensitive string containing a StringOrURI
		value.  Use of this claim is OPTIONAL.
	*/
	Sub string `json:"sub"`

	/*
		The "iat" (issued at) claim identifies the time at which the JWT was
		issued.  This claim can be used to determine the age of the JWT.  Its
		value MUST be a number containing a NumericDate value.  Use of this
		claim is OPTIONAL.
	*/
	Iat time.Time `json:"iat"`

	/*
		The "exp" (expiration time) claim identifies the expiration time on
		or after which the JWT MUST NOT be accepted for processing.  The
		processing of the "exp" claim requires that the current date/time
		MUST be before the expiration date/time listed in the "exp" claim.
	*/
	Exp time.Time `json:"exp"`
}

// Valid checks the claims for validity
func (Claims) Valid() error {
	return nil
}

// JwtPrivateKey represents the private key for signing the jwt
type JwtPrivateKey = *rsa.PrivateKey

// JwtSigner manages the signing of our jwt
type JwtSigner struct {
	PrivateKey JwtPrivateKey
}

// Sign signs claims into a jwt token returned as a string
func (signer JwtSigner) Sign(claims Claims) (string, error) {
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(signer.PrivateKey)
}

const jwtPrivateKeyEnvVarName = "JWT_PRIVATE_PEM_BASE64"

// ParseBase64PrivatePEM converts a base64 pem encoded certificate to DER (binary) rsa format
func ParseBase64PrivatePEM(privatePEMbase64 string) (*rsa.PrivateKey, error) {
	privatePEM, err := base64.StdEncoding.DecodeString(privatePEMbase64)
	if err != nil {
		return nil, fmt.Errorf("base64 decoding failed from %s", jwtPrivateKeyEnvVarName)
	}

	block, _ := pem.Decode(privatePEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key from %s", jwtPrivateKeyEnvVarName)
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// ProvideJwtPrivateKey returns a JwtPrivateKey, and an error if not provided by env vars
func ProvideJwtPrivateKey() (JwtPrivateKey, error) {
	privatePEMbase64 := os.Getenv(jwtPrivateKeyEnvVarName)

	if privatePEMbase64 == "" {
		return nil, fmt.Errorf("you must set %s", jwtPrivateKeyEnvVarName)
	}

	priv, err := ParseBase64PrivatePEM(privatePEMbase64)
	if err != nil {
		return nil, err
	}

	return JwtPrivateKey(priv), nil
}

var key = []byte(`
-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`)

// NewClaims instantiates the claims object to be signed
func NewClaims(userID string) Claims {
	return Claims{
		Jti: xid.New().String(),
		Sub: userID,
		Iat: time.Now(),
		Exp: time.Now().Add(24 * time.Hour),
	}
}
