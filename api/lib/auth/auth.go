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
		return nil, err
	}

	block, _ := pem.Decode(privatePEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, err
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

// NewClaims instantiates the claims object to be signed
func NewClaims(userID string) Claims {
	return Claims{
		Jti: xid.New().String(),
		Sub: userID,
		Iat: time.Now(),
		Exp: time.Now().Add(24 * time.Hour),
	}
}
