package db

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
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

// GetOrCreateUser creates a user object
func (db Db) GetOrCreateUser(ctx context.Context, loginKind onetimepasscode.LoginKind, emailOrPhone string) (*user.User, error) {
	u, err := db.GetUserByEmailOrPhone(ctx, emailOrPhone)
	if err != nil {
		return nil, err
	}
	if u != nil {
		log.Printf("User found: %v", u)
		return u, nil
	}

	if loginKind == onetimepasscode.LoginKindPhone {
		u, err = db.CreateUser(ctx, "", emailOrPhone)
	} else {
		u, err = db.CreateUser(ctx, emailOrPhone, "")
	}
	if err != nil {
		return nil, err
	}

	log.Printf("User not found; created: %v", u)
	return u, nil
}

func userFromSnapshot(snap *firestore.DocumentSnapshot) user.User {
	data := snap.Data()

	return user.User{
		ID:        data["id"].(string),
		Email:     data["email"].(string),
		Phone:     data["phone"].(string),
		CreatedAt: data["createdAt"].(time.Time),
	}
}

// GetUserByEmailOrPhone will return a user if one is found matching the input by email or phone
func (db Db) GetUserByEmailOrPhone(ctx context.Context, emailOrPhone string) (*user.User, error) {
	if emailOrPhone == "" {
		return nil, nil
	}

	users, err := db.Firestore.Collection("users").
		Where("email", "==", emailOrPhone).
		Limit(1).
		Documents(ctx).
		GetAll()
	if err != nil {
		return nil, err
	}
	if len(users) == 1 {
		u := userFromSnapshot(users[0])
		return &u, nil
	}

	users, err = db.Firestore.Collection("users").
		Where("phone", "==", emailOrPhone).
		Limit(1).
		Documents(ctx).
		GetAll()
	if err != nil {
		return nil, err
	}
	if len(users) == 1 {
		u := userFromSnapshot(users[0])
		return &u, nil
	}

	return nil, nil
}

func sixRandomDigits() (string, error) {
	max := big.NewInt(999999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", n.Int64()), nil
}
