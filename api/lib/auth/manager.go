package auth

import (
	"time"

	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/models/user"
	proto "github.com/khoerling/flux/api/lib/protocol"
)

type Manager struct {
	*JwtSigner
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

/*
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

*/
