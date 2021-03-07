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
	"github.com/khoerling/flux/api/lib/db/models/user/plaid/item"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/unmarshal"
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/hashing"
	"github.com/rs/xid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// SaveUser saves a user object (upsert/put semantics)
func (db Db) SaveUser(ctx context.Context, tx *firestore.Transaction, u *user.User) error {
	encryptedUser, err := u.Encrypt(db.EncryptionManager, u.ID)
	if err != nil {
		return err
	}

	ref := db.Firestore.Collection("users").Doc(string(u.ID))
	if tx == nil {
		_, err = ref.Set(ctx, encryptedUser)
	} else {
		err = tx.Set(ref, encryptedUser)
	}

	return err
}

// GetOrCreateUser creates a user object
func (db Db) GetOrCreateUser(ctx context.Context, loginKind onetimepasscode.LoginKind, emailOrPhone string) (*user.User, error) {
	var u user.User

	err := db.Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		var err error

		uptr, err := db.GetUserByEmailOrPhone(ctx, tx, emailOrPhone)
		if err != nil {
			return err
		}
		if uptr != nil {
			u = *uptr
			log.Printf("User found: %#v", u)

			return nil
		}

		now := time.Now()
		// first time login means that we verified them with otp
		if loginKind == onetimepasscode.LoginKindPhone {
			phone := user.Phone(emailOrPhone)
			u = user.User{
				Phone:           &phone,
				PhoneVerifiedAt: &now,
			}.WithDefaults(now)
		} else {
			email := user.Email(emailOrPhone)
			u = user.User{
				Email:           &email,
				EmailVerifiedAt: &now,
			}.WithDefaults(now)
		}

		err = db.SaveUser(ctx, nil, &u)
		if err != nil {
			return err
		}
		log.Printf("User not found; created: %v", u)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &u, nil
}

// GetUserByID gets a user object by id
func (db Db) GetUserByID(ctx context.Context, userID user.ID) (*user.User, error) {
	if userID == "" {
		return nil, nil

	}

	snap, err := db.Firestore.Collection("users").Doc(string(userID)).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var encU user.EncryptedUser
	snap.DataTo(&encU)

	u, err := encU.Decrypt(db.EncryptionManager, userID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// SavePlaidItem ...
func (db Db) SavePlaidItem(ctx context.Context, userID user.ID, itemID item.ID, accessToken string) (*item.Item, error) {
	item := item.Item{
		ID:          itemID,
		AccessToken: accessToken,
		CreatedAt:   time.Now(),
	}

	encryptedItem, err := item.Encrypt(db.EncryptionManager, userID)
	if err != nil {
		return nil, err
	}

	_, err = db.Firestore.Collection("users").Doc(string(userID)).Collection("plaidItems").Doc(string(itemID)).Set(ctx, encryptedItem)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// GetUserByEmailOrPhone will return a user if one is found matching the input by email or phone
func (db Db) GetUserByEmailOrPhone(ctx context.Context, tx *firestore.Transaction, emailOrPhone string) (*user.User, error) {
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

		u, err := encU.Decrypt(db.EncryptionManager, encU.ID)
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

		u, err := encU.Decrypt(db.EncryptionManager, encU.ID)
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
func (db Db) SaveProfileData(ctx context.Context, tx *firestore.Transaction, userID user.ID, pdata profiledata.ProfileData) (common.ProfileDataID, error) {
	profile := db.Firestore.Collection("users").Doc(string(userID)).Collection("profile")

	out, err := pdata.Encrypt(db.EncryptionManager, userID)
	if err != nil {
		return "", err
	}
	ref := profile.Doc(string(out.ID))

	err = tx.Set(ref, out)
	if err != nil {
		return "", err
	}

	return out.ID, nil
}

// GetAllProfileData ...
func (db Db) GetAllProfileData(ctx context.Context, tx *firestore.Transaction, userID user.ID) (profiledata.ProfileDatas, error) {
	profile := db.Firestore.Collection("users").Doc(string(userID)).Collection("profile")

	var (
		docs []*firestore.DocumentSnapshot
		err  error
	)
	if tx == nil {
		docs, err = profile.Documents(ctx).GetAll()
	} else {
		docs, err = tx.Documents(profile).GetAll()
	}
	if err != nil {
		return []profiledata.ProfileData{}, err
	}

	var out []profiledata.ProfileData
	for _, snap := range docs {
		var pdata common.EncryptedProfileData
		err := snap.DataTo(&pdata)
		if err != nil {
			return []profiledata.ProfileData{}, err
		}

		decrypted, err := unmarshal.DecryptAndUnmarshal(db.EncryptionManager, userID, pdata)
		if err != nil {
			return []profiledata.ProfileData{}, err
		}
		out = append(out, *decrypted)
	}

	return out, nil
}
