package db

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/khoerling/flux/api/lib/db/models/onetimepasscode"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/rs/xid"
)

// Db represents the application interface for accessing the database
type Db struct {
	Firestore *firestore.Client
}

// CreateOneTimePasscode stores a record of a one-time-password request for verification later
func (db Db) CreateOneTimePasscode(ctx context.Context, emailOrPhone string, kind onetimepasscode.LoginKind) (*onetimepasscode.OneTimePasscode, error) {
	id := xid.New().String()

	code, err := sixRandomDigits()
	if err != nil {
		return nil, err
	}

	otp := onetimepasscode.OneTimePasscode{
		ID:           id,
		EmailOrPhone: emailOrPhone,
		Kind:         kind,
		Code:         code,
		CreatedAt:    time.Now(),
	}

	_, err = db.Firestore.Collection("one-time-passcodes").Doc(id).Set(ctx, &otp)
	if err != nil {
		return nil, err
	}

	return &otp, nil
}

// CreateUser creates a user object
func (db Db) CreateUser(ctx context.Context, email string, phone string) (*user.User, error) {
	id := xid.New().String()

	u := user.User{
		ID:        id,
		Email:     email,
		Phone:     phone,
		CreatedAt: time.Now(),
	}

	_, err := db.Firestore.Collection("users").Doc(id).Set(ctx, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func sixRandomDigits() (string, error) {
	max := big.NewInt(999999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", n.Int64()), nil
}
