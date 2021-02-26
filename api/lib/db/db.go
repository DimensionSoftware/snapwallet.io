package db

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/khoerling/flux/api/lib/db/models"
)

// Db represents the application interface for accessing the database
type Db struct {
	Firestore *firestore.Client
}

// CreateOneTimePasscode stores a record of a one-time-password request for verification later
func (db Db) CreateOneTimePasscode(ctx context.Context, otp models.OneTimePasscode) (string, error) {
	ref, _, err := db.Firestore.Collection("one-time-passcodes").Add(ctx, otp)
	if err != nil {
		return "", err
	}
	return ref.ID, nil
}

/*
	_, _, err = s.Firestore.Collection("one-time-passcodes").Add(ctx, map[string]interface{}{
		"emailOrPhone": loginValue,
		"kind":         loginKind,
		"code":         code,
		"createdAt":    time.Now(),
	})
	if err != nil {
		return nil, err
*/
