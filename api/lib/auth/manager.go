package auth

import (
	"context"
	"time"

	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/models/user"
	proto "github.com/khoerling/flux/api/lib/protocol"
)

type Manager struct {
	*JwtSigner
	*JwtVerifier
	*db.Db
}

func (m Manager) NewTokenMaterial(now time.Time, userID user.ID, parentRefreshTokenID string) (*proto.TokenMaterial, error) {
	refresh := NewRefreshTokenClaims(now, userID)
	access := NewAccessTokenClaims(now, userID, parentRefreshTokenID)

	refreshToken, err := m.JwtSigner.Sign(refresh)
	if err != nil {
		return nil, err
	}

	accessToken, err := m.JwtSigner.Sign(access)
	if err != nil {
		return nil, err
	}

	return &proto.TokenMaterial{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil

}

func (m Manager) TokenExchange(ctx context.Context, now time.Time, refreshToken string) (*proto.TokenMaterial, error) {
	// if refresh token found in db, then access denied, and access tokens should be revoked which were birthed by that refresh token
	// if not found in db, new token material can be presented to the user

	refresh, err := m.JwtVerifier.ParseAndVerify(ctx, TokenKindRefresh, refreshToken)
	if err != nil {
		return nil, err
	}

	material, err := m.NewTokenMaterial(now, user.ID(refresh.Subject), refresh.Id)
	if err != nil {
		return nil, err
	}

	return material, nil

}
