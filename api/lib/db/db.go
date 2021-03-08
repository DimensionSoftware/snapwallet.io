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
	"google.golang.org/api/iterator"
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
			u = user.User{
				Phone:           &emailOrPhone,
				PhoneVerifiedAt: &now,
			}.WithDefaults(now)
		} else {
			u = user.User{
				Email:           &emailOrPhone,
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

// AckOneTimePasscode tries to find the OneTimePasscode object matching the loginValue and code
//   If it is found then the ack is successful and the passcode is destroyed in the database
//   after being destroyed the original passcode object will be returned.
//
//   If it is not found then (nil, nil) will be returned so the caller can decide the error handling strategy
//   for not found (invalid) codes (Not Acked)
//
//   If there is another type of error then (nil, error) will be returned.
func (db Db) AckOneTimePasscode(ctx context.Context, loginValue string, code string) (*onetimepasscode.OneTimePasscode, error) {
	if loginValue == "" {
		// no matching otp found
		return nil, nil
	}

	if code == "" {
		// no matching otp found
		return nil, nil
	}

	passcodes := db.Firestore.Collection("one-time-passcodes").
		Where("emailOrPhone", "==", loginValue).
		Where("code", "==", code).
		Where("createdAt", ">", time.Now().Add(-10*time.Minute)).
		Documents(ctx)

	snap, err := passcodes.Next()
	if err == iterator.Done {
		// no matching otp found
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// can only be used 'once'
	_, err = snap.Ref.Delete(ctx)
	if err != nil {
		return nil, err
	}

	var passcode onetimepasscode.OneTimePasscode
	err = snap.DataTo(&passcode)
	if err != nil {
		return nil, err
	}

	// put this here for now ;)
	cleaned, err := db.CleanAgedPasscodes(ctx)
	if err != nil {
		return nil, err
	}
	if cleaned != 0 {
		log.Printf("Cleaned %d aged passcodes", cleaned)
	}

	// match found & unmarshalled
	return &passcode, nil
}

// CleanAgedPasscodes cleanup aged passcodes; there is no security risk to not running this but it saves on db storage bills
// returns num of passcodes deleted which were old, or an error if shit goes bad
func (db Db) CleanAgedPasscodes(ctx context.Context) (int, error) {
	passcodes := db.Firestore.Collection("one-time-passcodes").
		Where("createdAt", "<", time.Now().Add(-10*time.Minute)).
		Documents(ctx)

	batch := db.Firestore.Batch()
	numDeleted := 0
	for {
		doc, err := passcodes.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, err
		}

		batch.Delete(doc.Ref)
		numDeleted++
	}

	if numDeleted == 0 {
		return 0, nil
	}

	_, err := batch.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return numDeleted, nil

}

/*

snap, err := passcodes.Next()
if err == iterator.Done {
return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedOTP(onetimepasscode.LoginKindPhone))
}
if err != nil {
log.Println(err)
return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
}

var passcode onetimepasscode.OneTimePasscode
err = snap.DataTo(&passcode)
if err != nil {
log.Println(err)
return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
}
*/
