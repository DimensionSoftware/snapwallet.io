package db

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/khoerling/flux/api/lib/db/models/onetimepasscode"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/plaid/item"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/address"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/hashing"
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
func (db Db) CreateUser(ctx context.Context, email *string, phone *string, emailVerified bool, phoneVerified bool) (*user.User, error) {
	id := user.ID(xid.New().String())

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

	encryptedUser, err := u.Encrypt(db.EncryptionManager)
	if err != nil {
		return nil, err
	}

	_, err = db.Firestore.Collection("users").Doc(string(id)).Set(ctx, encryptedUser)
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
		u, err = db.CreateUser(ctx, nil, &emailOrPhone, false, true)
	} else {
		u, err = db.CreateUser(ctx, &emailOrPhone, nil, true, false)
	}
	if err != nil {
		return nil, err
	}

	log.Printf("User not found; created: %v", u)
	return u, nil
}

// GetUserByID gets a user object by id
func (db Db) GetUserByID(ctx context.Context, id user.ID) (*user.User, error) {
	if id == "" {
		return nil, nil

	}

	snap, err := db.Firestore.Collection("users").Doc(string(id)).Get(ctx)
	if err != nil {
		return nil, err
	}

	if snap.Exists() {
		var encU user.EncryptedUser
		snap.DataTo(&encU)

		u, err := encU.Decrypt(db.EncryptionManager)
		if err != nil {
			return nil, err
		}

		return u, nil
	}

	// no user found (non-error)
	return nil, nil
}

// SavePlaidItem ...
func (db Db) SavePlaidItem(ctx context.Context, userID user.ID, itemID item.ID, accessToken string) (*item.Item, error) {
	item := item.Item{
		ID:          itemID,
		AccessToken: accessToken,
		CreatedAt:   time.Now(),
	}

	encryptedItem, err := item.Encrypt(db.EncryptionManager)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("users/%s/plaidItems", userID)

	_, err = db.Firestore.Collection(path).Doc(string(itemID)).Set(ctx, encryptedItem)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// GetUserByEmailOrPhone will return a user if one is found matching the input by email or phone
func (db Db) GetUserByEmailOrPhone(ctx context.Context, emailOrPhone string) (*user.User, error) {
	if emailOrPhone == "" {
		return nil, nil
	}

	emailOrPhoneBytes := []byte(emailOrPhone)
	emailOrPhoneHash := hashing.Hash(emailOrPhoneBytes)

	users, err := db.Firestore.Collection("users").
		Where("emailHash", "==", emailOrPhoneHash).
		Limit(1).
		Documents(ctx).
		GetAll()
	if err != nil {
		return nil, err
	}
	if len(users) == 1 {
		var encU user.EncryptedUser
		err := users[0].DataTo(&encU)
		if err != nil {
			return nil, err
		}

		u, err := encU.Decrypt(db.EncryptionManager)
		if err != nil {
			return nil, err
		}

		return u, nil
	}

	users, err = db.Firestore.Collection("users").
		Where("phoneHash", "==", emailOrPhoneHash).
		Limit(1).
		Documents(ctx).
		GetAll()
	if err != nil {
		return nil, err
	}
	if len(users) == 1 {
		var encU user.EncryptedUser
		err := users[0].DataTo(&encU)
		if err != nil {
			return nil, err
		}

		u, err := encU.Decrypt(db.EncryptionManager)
		if err != nil {
			return nil, err
		}

		return u, nil
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

// SaveProfileData ...
func (db Db) SaveProfileData(ctx context.Context, userID user.ID, pdata interface{}) (common.ProfileDataID, error) {
	profile := db.Firestore.Collection("users").Doc(string(userID)).Collection("profile")

	var out *common.EncryptedProfileData
	switch obj := pdata.(type) {
	/*
		case legalname.ProfileDataLegalName:
			encrypted, err := obj.Encrypt(db.EncryptionManager, userID)
			if err != nil {
				return "", err
			}
			out = encrypted
	*/
	case address.ProfileDataAddress:
		encrypted, err := obj.Encrypt(db.EncryptionManager, userID)
		if err != nil {
			return "", err
		}
		out = encrypted
	default:
		typeName := reflect.TypeOf(obj).String()
		return "", fmt.Errorf("cannot save profile data of unknown type: %s", typeName)
	}

	_, err := profile.Doc(string(out.ID)).Set(ctx, out)
	if err != nil {
		return "", err
	}

	return out.ID, nil
}
