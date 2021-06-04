package auth

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/lithammer/shortuuid/v3"

	"github.com/dgrijalva/jwt-go"
	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/models/usedrefreshtoken"
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

const base64PrivateTestPEM = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKSndJQkFBS0NBZ0VBdzc1Nk5mMFVCUVlSWWJHZ2tKZGdsRk5aYmZHNDFhLzlCd2xjNXZ2R2hGSWtZeXdqCitCcUxNYUU1U3ZQR0Q4ZS9DQ1JDRGx2TW1ZUjBuVnlWNGZzV1JmcFNGYTVZcVc3RkN2TE1kSFlmU3lBNXNrSm8Kb2hhWUlGc2lOd2xQVlRyOXFudVpKd3JsaEw5ekRqMnRmV0I2VHI4S3M3Zm5VN3dXeFZQYVMxSmIyaGo4R1BMNgpZcldzektJaXBaaEo3Z3NtUEREWFdCUEpwVlg5NDQybHRZSHh4L3NjeE5LWFlNeW81elBxeUpXaG03cVdSOEsvCkJEUSs4eXJXbjU4UGRQQnFyT0dmSzlqdU1CSkZGMkZJK0RwNkFCNGxQWjFtSUtoYVdFcFJnSExoNThFaW51dFcKOWZ3cjUvQWwzeU1ueGMzUVdPeitVYVd0b3laS0NTRndmYjF3QXBURkFrVG1qNGVhYW5OVURDWmdnNnNKeDNpdApEc2pVQW5rMVNyeFNkRjd6c0VyQnBMWDNGVHFUL3RtejJzV2xaaVRNZDAwZWpxWm8zdEhTblFKeXdwaHhsMDA0CjE1dWQvSGtYSzFXQnRrR2hNcDBUY3BQTWNZWnFSaHZnbDk3dnFpRmNnWFZONzdyRERuLytRN29XWTlRdVJHVlUKV1BMaTZhREhsY1hzdXlZZTdna3J3bkJScklodU5UOXlKL1NZcnVoUExTRitoUUVJZDFXUjM5SnJJY083Y2d3TApLb1hFcUZVd21CTThiVmtFTURWUTFwV1JWejBabHEyZk9pbnl2bTYxd1kyQnltbUxGUUtrWklWbER4clEyeEpkClZiTzFidmFBamxsZFozRUhnYzRFMW9CeHRsczN4ZkQrem1sOHNpa2NIakxQS25wKzh1RFlXcC9RNjVrQ0F3RUEKQVFLQ0FnQjFTTmdwS1M2cG8rMGVRRFFZN3RycmhOVjh6dTBVL0pIN2VWeTArZjhFb2NNenVPc0VhY01sUlpqeQpsQVlFeG9acjltMnQ5TXN1NFBLT3B6OFhYRDhJUnVpUUhScjZ5bWcrR3lUdVV5aUU5eFlhL1RkOGgxVTNiZU9lClhuR3VlOTRxSEV5ejNBK1I3clNkdjg4SDVKcmtQQXZKaTFPTUZKTUFRVEgzRjYzNWpDYmhQQlZTdDlDRi9GQU8KTUtWN2dDcTB2ZjhKd3pGN1kzN3dyWnF3bXQrb0Y0b3hWSTFuYnJsMWJ1SXF0WjFUbGdVZGtrNnAwalExdDJGSgoyeUNEek5uZUpJSWUzNmc0SFhwUWUvWWc1Y2piOFRRWEt3eGFYUVZsQ1lLMEdDSExueW1EVnplUHhEejZpQ1hwCkFVTi9mY1pzd2ZQUU1CbE9QRTc0RVpVNWdoeHVJYkZKSnBPU1hYc1JubkNCYUh3cXFRV3hqK3VPSitxSTQxOEEKUjhIMTNjVVNyaG1JeFp1dCtMQjVzNVhMNktmVnpIL09FalVTYmJoWjJhd21SaHZjd3BuYUVDOUpoNHdPMUpLYwpOblVDWDlLUml6c0FBSHVNNXlOeUN1L0dUUjZOR0lhU29SdjZQSmVYY2xNMkJRUG5NdVVmRGY0cnEyQmNHemkrClM0bVcvVys4M3BUZEwwK1ZVVFJTKzM1eklEbjRyYTJVUW5DRTV2Vm9lSk9XZGFwR0N1VU02L0VJU29DajFKZFQKSzlSRUl1c2V6RjJiamhNS1lwd2xCc2dQMmFobG5wdWlySXJYN3JGTklveG9pYlpvVnA0TUplU0xpMHJXaTk2bgpmTlhjTEs1YzFiRDFGRW9VVzBpS0I2TkdXaURmejdZb2l4Z052NDBKTlBvNnNhd3VRUUtDQVFFQStTbUorWU1yCjk2TnZpdDBHRllUa2hNcjRhZEdKdmc4L1ZTeHlNV3EralZ2dm5JQ0ZCT2ZhenRqZ25TcktadjdYOWsvdGM2N3YKVk5rUjN3UlZ2K1NPQjVaMjdPRXh3ZE1PVmVLOS9NZXBRNnVOTUVaRllRc0dzM1FhUDMxa2UyRUFzbUc2YmRhSgpTTDlzdXFYWFBXWXp0T0VoS0NTZ1M3RHQ5RDZCbW9pRm1FbThjdEllUnFuWWFOL1VwSjBCL0dwbG5wVERYNTdVCkxHMHJxUTJEVWlZaVhRTER3QithdzdVa1A2RkRpOEVqQ016TVo1Zi9YUWU4MktGME1XeTBwUHhpUm5sdlY1S0QKeldrTitNWURWd01Zb3RLbGRQZkZyV0h6Q0VwOFArbk9aam1ZZEtUTVVlbVBOcSthNnN4enZ0clAzWld3TzA5MwpscEdGM1pRa0NtUU9CUUtDQVFFQXlSMm5wczlyc1F3aWRSWGpGd0I1YnhwUFdPK0I0K0lLcnlnK3YwNUM1ZkNlCjhURjFjR2svM2xtMFFoNkVDRUlyRDFiL0lhenMvQUtsdmNCdWhIMm5WM2dVdE9BNnA2bDFLVVY1L0xFN01nZVUKRW1vUDVBNnplOVhhbldZeENUOFBUQkI2clZwR3RLZlhBZy9aUU83TVI4d1JuZGkrTW8wN0MxaS9jVnQ3b1Q0VwpXUHFyYVRvUXhoTVZSTmpCaDFacHlGRjk2alhOVDJwTTNFcEVnSGt1R1JIREJYKzdFMGo1bElKZkVlY3lzVlV2ClZ2QUw4M2ZYYmE4aVNKVmM5VVo1VVdkRER3Ui81RHAyaFFWekFSRTlqR3dIVkFMTWY3TXptY3FjVm9xcmFQcUwKdTBpTlhqdUlxWi9KRTdqbVRRVXNZR1p1WURpVXExKzNvZk14aG4ySGhRS0NBUUI0S1ZaSEpNRnhmanB1NThUbApYOSt0UXcybVNzMTVBWFZ6cUNteU4xNktZY29lMmNSTTUxd1k0WE1CbVA5ZnlJdHlXSDJWaXNvVlMxSlpFWWdoCk1TSmwrbVNFUE41NE1VYjZtSTB6ZVQ3aVNidWZpbVF2TnRnV2QybXBNTm5pdnBkTmIrQXUxSVlFdFh6RVR2S00KdzlzdjRsclJOMGl1K0RicDBiTkRTWS9VTDh0WVBJa3BYd1BsSC8wM0hoazFHRUxGeGN5ck1yZjBiUG5mWDRyegpkVHloU3BJSTk3VDFxVWcvLzQvSnVHMGk4MUdvckI4VlNJUUVuV2loNVdFQktFNWsybStkOWlUbVRVSFZ0ZmtxClgyM2tLRjV2R0ljVzVPUUdIWlhxWk9HTFh6OXRFWVVRQStsellDRUFGM1hDbDFnajd1cTh6OGhHcnd1MkhwbDcKQVdIWkFvSUJBQjRlb2RnYW9sOCttUDUrQzZlTE82U0hCVEVsbFlkaWVBVXBldFE5elVrUWswMCtBZitXMDZqaQpPRnZhcEIxMGcyeGx0QW9BRXZIZkY2Rm1hMmJPUnJ5VDBFNFNjdFpmUzV1bHV4SThITWh1V0IrMTRMRmYva05pCmtMNFg2dW9lbHBUbXR1aTFaM2R3MTRPSWlobnVhWXVySlV4RmhKNmZoaU01NUZuK3dISXlrVGc4T21Xays1UWUKa1lYaERJTFBUMEpmNmdLa2toMzlwb3NyV0QwQmFRVDZJd0gwMGppZUtqcXlsN2hmcnJqZU1Cdk9FWXdoKzVLQwpzeGk1dmRWQis3TlVTcmU2RGNsSmRDeVIxMXBta1pneEtadk1XNElZZlhiZjREQVp6bmdIWmR2ampzQjVIQXVEClVhTHhneFF2M1NpR3pxcjhiaytJSTBGTzRiRHBwZDBDZ2dFQWZaalFFOGxzTmN1VUlLVjRrZE92WVBhRG9RMnUKLzYxQk9jVXNodzBCTFVZc25ydHM1OVZFMGUwc25OWmpTQlVDdUNoM1k3SkoxY0wwalJCelRXTDhROGgxc1dxVgpVcDhCV0VhSjdDWGJHSjlnalBRd1JtVXJPVmMrVittTlphTW5VYXYrbktDMGZlbWZWaml0alFCR0N1Q0RoQVhHCndodTJENlNTSmUwOUljUHpnMlFrdmdZNmNKQVdZUDIvZmFndlg3ektoVU9qOGRrQkJ3UkltUkswVUJobzdtZjgKUDN5TlFkR2wvMEhJM0w1c2RwQ2NMWEpGRlg4bENBakd6M1FWMDc0eTFlaEFldThyOGZEc28wZXJqdnVHcXp3SQp0Q3JGeHd6ck1WRjBZeVJtQXlld1JReU45MGZ2eEJxWG5SeXlzVVhCbGJEc2loYmhLVjRXZzFWNmZBPT0KLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K"

// ProvideTestJwtPrivateKey returns a test JwtPrivateKey
func ProvideTestJwtPrivateKey() JwtPrivateKey {
	priv, err := ParseBase64PrivatePEM(base64PrivateTestPEM)
	if err != nil {
		panic(err)
	}

	return JwtPrivateKey(priv)
}

// ProvideJwtPublicKey returns a JwtPublicKey, and an error if not provided by env vars
func ProvideJwtPublicKey(priv JwtPrivateKey) JwtPublicKey {
	return &priv.PublicKey
}

func NewAccessTokenClaims(now time.Time, userID user.ID, refreshTokenID string) jwt.StandardClaims {
	return jwt.StandardClaims{
		Id:        shortuuid.New(),
		Audience:  string(TokenKindAccess),
		Subject:   string(userID),
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(3 * time.Minute).Unix(),
		Issuer:    refreshTokenID,
	}
}

func NewRefreshTokenClaims(now time.Time, userID user.ID) jwt.StandardClaims {
	return jwt.StandardClaims{
		Id:        shortuuid.New(),
		Audience:  string(TokenKindRefresh),
		Subject:   string(userID),
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(20 * time.Minute).Unix(),
	}
}

type TokenKind string

const (
	TokenKindAccess  TokenKind = "ACCESS"
	TokenKindRefresh TokenKind = "REFRESH"
)

// JwtVerifier manages the verification of our jwt
type JwtVerifier struct {
	PublicKey JwtPublicKey
	db.Db
}

// ParseAndVerify parses and verifies a raw jwt token and returns the claims if successful
func (signer JwtVerifier) ParseAndVerify(ctx context.Context, expectedTokenKind TokenKind, rawToken string) (*jwt.StandardClaims, error) {
	// Create the token
	token, err := jwt.ParseWithClaims(
		rawToken,
		new(jwt.StandardClaims),
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

	if token.Claims == nil {
		return nil, fmt.Errorf("token is invalid")
	}

	claims := token.Claims.(*jwt.StandardClaims)

	if claims.Audience != string(expectedTokenKind) {
		return nil, fmt.Errorf("token is invalid")
	}

	if expectedTokenKind == TokenKindRefresh {
		now := time.Now()

		used, err := signer.Db.GetUsedRefreshToken(ctx, nil, claims.Id)
		if err != nil {
			return nil, err
		}

		if used != nil {
			used.RevokedAt = &now

			err = signer.Db.SaveUsedRefreshToken(ctx, nil, used)
			if err != nil {
				return nil, err
			}

			return nil, fmt.Errorf("Verification failure: refresh token used more than once!")
		}

		used = &usedrefreshtoken.UsedRefreshToken{
			ID:        claims.Id,
			Subject:   claims.Subject,
			IssuedAt:  time.Unix(claims.IssuedAt, 0),
			ExpiresAt: time.Unix(claims.ExpiresAt, 0),
			UsedAt:    now,
		}

		err = signer.Db.SaveUsedRefreshToken(ctx, nil, used)
		if err != nil {
			return nil, err
		}

		return claims, nil
	}

	if expectedTokenKind == TokenKindAccess {
		// refresh token issuer _must_ be known
		if claims.Issuer == "" {
			return nil, fmt.Errorf("token is invalid")
		}

		urt, err := signer.Db.GetUsedRefreshToken(ctx, nil, claims.Issuer)
		if err != nil {
			return nil, err
		}

		if urt != nil && urt.RevokedAt != nil {
			return nil, fmt.Errorf("token is invalid; revoked at: %s", urt.RevokedAt)
		}

		return claims, nil
	}

	return nil, fmt.Errorf("unhandled expected token kind: %s", expectedTokenKind)
}
