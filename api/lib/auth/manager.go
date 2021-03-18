package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/models/usedrefreshtoken"
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

	// todo : need way to know if its a refresh token, could use separate privkey, or just sign some data in jwt
	refresh, err := m.JwtVerifier.ParseAndVerify(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	used, err := m.Db.GetUsedRefreshToken(ctx, nil, refresh.Id)
	if err != nil {
		return nil, err
	}

	if used != nil {
		used.RevokedAt = &now

		err = m.Db.SaveUsedRefreshToken(ctx, nil, used)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("TokenExchange failure: refresh token used more than once!")
	}

	material, err := m.NewTokenMaterial(now, user.ID(refresh.Subject), refresh.Id)
	if err != nil {
		return nil, err
	}

	used = &usedrefreshtoken.UsedRefreshToken{
		ID:        refresh.Id,
		Subject:   refresh.Subject,
		IssuedAt:  time.Unix(refresh.IssuedAt, 0),
		ExpiresAt: time.Unix(refresh.ExpiresAt, 0),
		UsedAt:    now,
	}

	err = m.Db.SaveUsedRefreshToken(ctx, nil, used)
	if err != nil {
		return nil, err
	}

	return material, nil

}
