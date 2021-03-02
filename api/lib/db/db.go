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
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/rs/xid"
)

// Db represents the application interface for accessing the database
type Db struct {
	Firestore         *firestore.Client
	EncryptionManager *encryption.Manager
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
func (db Db) CreateUser(ctx context.Context, email string, phone string, emailVerified bool, phoneVerified bool) (*user.User, error) {
	id := xid.New().String()

	now := time.Now()
	u := user.User{
		ID:        id,
		Email:     email,
		Phone:     phone,
		CreatedAt: now,
	}

	if emailVerified {
		u.EmailVerifiedAt = &now
	}

	if phoneVerified {
		u.PhoneVerifiedAt = &now
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
		log.Printf("User found: %#v", u)
		return u, nil
	}

	// first time login means that we verified them with otp
	if loginKind == onetimepasscode.LoginKindPhone {
		u, err = db.CreateUser(ctx, "", emailOrPhone, false, true)
	} else {
		u, err = db.CreateUser(ctx, emailOrPhone, "", true, false)
	}
	if err != nil {
		return nil, err
	}

	log.Printf("User not found; created: %v", u)
	return u, nil
}

// GetUserByID gets a user object by id
func (db Db) GetUserByID(ctx context.Context, id string) (*user.User, error) {
	if id == "" {
		return nil, nil

	}

	snap, err := db.Firestore.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	if snap.Exists() {
		var u user.User
		snap.DataTo(&u)
		return &u, nil
	}

	// no user found (non-error)
	return nil, nil
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
		var u user.User
		err := users[0].DataTo(&u)
		if err != nil {
			return nil, err
		}
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
		var u user.User
		err := users[0].DataTo(&u)
		if err != nil {
			return nil, err
		}
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
