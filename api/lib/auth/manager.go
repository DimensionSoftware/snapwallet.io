package auth

import (
	"context"
	"fmt"
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

func (m Manager) NewTokenMaterial(userID user.ID) (*proto.TokenMaterial, error) {
	now := time.Now()
	refresh := NewRefreshTokenClaims(now, userID)
	access := NewAccessTokenClaims(now, userID, refresh.Id)

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

func (m Manager) TokenExchange(ctx context.Context, refreshToken string) (*proto.TokenMaterial, error) {
	// if refresh token found in db, then access denied, and access tokens should be revoked which were birthed by that refresh token
	// if not found in db, new token material can be presented to the user

	// todo : need way to know if its a refresh token, could use separate privkey, or just sign some data in jwt
	refresh, err := m.JwtVerifier.ParseAndVerify(refreshToken)
	if err != nil {
		return nil, err
	}

	used, err := m.Db.GetUsedRefreshToken(ctx, nil, refresh.Id)
	if err != nil {
		return nil, err
	}

	if used == nil {
		// TODO: save used token
		material, err := m.NewTokenMaterial(user.ID(refresh.Subject))
		if err != nil {
			return nil, err
		}

		return material, nil
	} else {
		// TODO: mark as revoked, update interceptor to use this info to block access tokens?
		return nil, fmt.Errorf("TokenExchange failure: refresh token used more than once!")
	}
}
