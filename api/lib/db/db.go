package db

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/khoerling/flux/api/lib/db/models"
)

// Db represents the application interface for accessing the database
type Db struct {
	Firestore *firestore.Client
}

// CreateOneTimePasscode stores a record of a one-time-password request for verification later
func (db Db) CreateOneTimePasscode(ctx context.Context, emailOrPhone string, kind models.OneTimePasscodeLoginKind) (*models.OneTimePasscode, error) {
	code, err := sixRandomDigits()
	if err != nil {
		return nil, err
	}

	otp := models.OneTimePasscode{
		EmailOrPhone: emailOrPhone,
		Kind:         kind,
		Code:         code,
		CreatedAt:    time.Now(),
	}

	_, _, err = db.Firestore.Collection("one-time-passcodes").Add(ctx, &otp)
	if err != nil {
		return nil, err
	}

	return &otp, nil
}

func sixRandomDigits() (string, error) {
	max := big.NewInt(999999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", n.Int64()), nil
}
