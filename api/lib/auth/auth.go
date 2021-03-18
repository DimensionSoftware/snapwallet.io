package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/lithammer/shortuuid/v3"

	"github.com/dgrijalva/jwt-go"
	"github.com/khoerling/flux/api/lib/db/models/user"
)

// JwtPrivateKey represents the private key for signing the jwt
type JwtPrivateKey = *rsa.PrivateKey

// JwtPublicKey represents the public key for verifying the jwt
type JwtPublicKey = *rsa.PublicKey

// JwtSigner manages the signing of our jwt
type JwtSigner struct {
	PrivateKey JwtPrivateKey
}

// Sign signs claims into a jwt token returned as a string
func (signer JwtSigner) Sign(claims jwt.StandardClaims) (string, error) {
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
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
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

// ProvideJwtPublicKey returns a JwtPublicKey, and an error if not provided by env vars
func ProvideJwtPublicKey(priv JwtPrivateKey) JwtPublicKey {
	return &priv.PublicKey
}

func NewAccessTokenClaims(now time.Time, userID user.ID, refreshTokenID string) jwt.StandardClaims {
	return jwt.StandardClaims{
		Id:        shortuuid.New(),
		Subject:   string(userID),
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(1 * time.Minute).Unix(),
		Issuer:    refreshTokenID,
	}
}

func NewRefreshTokenClaims(now time.Time, userID user.ID) jwt.StandardClaims {
	return jwt.StandardClaims{
		Id:        shortuuid.New(),
		Subject:   string(userID),
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(24 * time.Hour).Unix(),
	}
}

// JwtVerifier manages the verification of our jwt
type JwtVerifier struct {
	PublicKey JwtPublicKey
}

// ParseAndVerify parses and verifies a raw jwt token and returns the claims if successful
func (signer JwtVerifier) ParseAndVerify(rawToken string) (*jwt.StandardClaims, error) {
	// Create the token
	token, err := jwt.ParseWithClaims(
		rawToken,
		jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return signer.PublicKey, nil
		},
	)

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	//claims := token.Claims.(jwt.StandardClaims)

	return token.Claims.(*jwt.StandardClaims), nil
}
