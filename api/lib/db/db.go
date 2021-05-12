package db

import (
	"context"

	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/lithammer/shortuuid/v3"

	"cloud.google.com/go/firestore"
	"github.com/khoerling/flux/api/lib/db/models/gotoconfig"
	"github.com/khoerling/flux/api/lib/db/models/job"
	"github.com/khoerling/flux/api/lib/db/models/onetimepasscode"
	"github.com/khoerling/flux/api/lib/db/models/usedrefreshtoken"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/file"
	"github.com/khoerling/flux/api/lib/db/models/user/plaid/item"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/email"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/phone"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/unmarshal"
	"github.com/khoerling/flux/api/lib/db/models/user/wyre/account"
	"github.com/khoerling/flux/api/lib/db/models/user/wyre/paymentmethod"
	"github.com/khoerling/flux/api/lib/db/models/user/wyre/walletorder"
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/hashing"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Db represents the application interface for accessing the database
type Db struct {
	Firestore         *firestore.Client
	EncryptionManager *encryption.Manager
}

// will not save if item is already existing; returns short id of immutable first item
func (db Db) SaveGotoConfig(ctx context.Context, g *gotoconfig.Config) (gotoconfig.ShortID, error) {
	ref := db.Firestore.Collection("goto-configs").Doc(string(g.ID))

	var out gotoconfig.ShortID

	err := db.Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		existingDoc, err := tx.Get(ref)
		if status.Code(err) == codes.NotFound {
			err = tx.Set(ref, g)
			if err != nil {
				return err
			}

			out = g.ShortID
			return nil
		}
		if err != nil {
			return err
		}

		var existingG gotoconfig.Config
		err = existingDoc.DataTo(&existingG)
		if err != nil {
			return err
		}

		out = existingG.ShortID
		return nil
	})
	if err != nil {
		return "", err
	}

	return out, nil
}

func (db Db) GetGotoConfigByShortID(ctx context.Context, shortID gotoconfig.ShortID) (*gotoconfig.Config, error) {
	var err error

	table := db.Firestore.Collection("goto-configs")

	records, err := table.
		Where("shortID", "==", shortID).
		Limit(1).
		Documents(ctx).
		GetAll()
	if err != nil {
		return nil, err
	}
	if len(records) == 1 {
		var g gotoconfig.Config
		err := records[0].DataTo(&g)
		if err != nil {
			return nil, err
		}

		return &g, nil
	}

	return nil, nil
}

func (db Db) SaveWalletOrderForUser(ctx context.Context, userID user.ID, woID walletorder.ID) error {
	ref := db.Firestore.Collection("users").Doc(string(userID)).Collection("wyreWalletOrders").Doc(string(woID))

	wo := walletorder.WalletOrder{
		ID:        woID,
		CreatedAt: time.Now(),
	}
	_, err := ref.Set(ctx, &wo)

	return err
}

func (db Db) GetAllWalletOrdersForUser(ctx context.Context, userID user.ID) (walletorder.WalletOrders, error) {
	ref := db.Firestore.Collection("users").Doc(string(userID)).Collection("wyreWalletOrders")

	docs, err := ref.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var out []walletorder.WalletOrder

	for _, doc := range docs {
		var wo walletorder.WalletOrder

		err := doc.DataTo(&wo)
		if err != nil {
			return nil, err
		}

		out = append(out, wo)
	}

	return out, nil
}

// CreateOneTimePasscode stores a record of a one-time-password request for verification later
func (db Db) CreateOneTimePasscode(ctx context.Context, emailOrPhone string, kind onetimepasscode.LoginKind) (*onetimepasscode.OneTimePasscode, error) {
	id := shortuuid.New()

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

func (db Db) SaveJob(ctx context.Context, tx *firestore.Transaction, j *job.Job) error {
	var err error

	ref := db.Firestore.Collection("jobs").Doc(string(j.ID))
	if tx == nil {
		_, err = ref.Set(ctx, j)
	} else {
		err = tx.Set(ref, j)
	}

	return err
}

func (db Db) GetJobByKindAndStatusAndRelatedId(ctx context.Context, kind job.Kind, status job.Status, relatedID string) (*job.Job, error) {
	table := db.Firestore.Collection("jobs").
		Where("kind", "==", kind).
		Where("status", "==", status).
		Where("relatedIDs", "array-contains", relatedID).
		Documents(ctx)

	snap, err := table.Next()
	if err == iterator.Done {
		// not found
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var job job.Job
	err = snap.DataTo(&job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

// SaveFileMetadata saves file metadata
func (db Db) SaveFileMetadata(ctx context.Context, userID user.ID, md *file.Metadata) error {
	ref := db.Firestore.Collection("users").Doc(string(userID)).Collection("files").Doc(string(md.ID))

	_, err := ref.Set(ctx, md)

	return err
}

// GetFileMetadata gets file metadata
func (db Db) GetFileMetadata(ctx context.Context, userID user.ID, fileID file.ID) (*file.Metadata, error) {
	ref := db.Firestore.Collection("users").Doc(string(userID)).Collection("files").Doc(string(fileID))

	snap, err := ref.Get(ctx)
	if status.Code(err) == codes.NotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var md file.Metadata
	err = snap.DataTo(&md)
	if err != nil {
		return nil, err
	}

	return &md, nil
}

// GetOrCreateUser creates a user object
func (db Db) GetOrCreateUser(ctx context.Context, loginKind onetimepasscode.LoginKind, emailOrPhone string) (*user.User, error) {
	now := time.Now()

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

		} else {

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

			err = db.SaveUser(ctx, tx, &u)
			if err != nil {
				return err
			}
			log.Printf("User not found; created: %v", u)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	log.Printf("user??? : %#v\n", u)

	err = db.Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		changed := false
		pdatas, err := db.GetAllProfileData(ctx, tx, u.ID)
		if err != nil {
			return err
		}

		if loginKind == onetimepasscode.LoginKindPhone {
			existingPdata := pdatas.FilterKindPhone().FindByPhone(emailOrPhone)

			if existingPdata == nil {
				changed = true
				pdatas = append(pdatas, phone.ProfileDataPhone{
					CommonProfileData: common.CommonProfileData{
						ID:        common.ProfileDataID(shortuuid.New()),
						Status:    common.StatusReceived,
						CreatedAt: now,
					},
					Phone: emailOrPhone,
				})
			}
		} else {
			existingPdata := pdatas.FilterKindEmail().FindByEmail(emailOrPhone)

			if existingPdata == nil {
				changed = true
				pdatas = append(pdatas, email.ProfileDataEmail{
					CommonProfileData: common.CommonProfileData{
						ID:        common.ProfileDataID(shortuuid.New()),
						Status:    common.StatusReceived,
						CreatedAt: now,
					},
					Email: emailOrPhone,
				})
			}
		}

		if changed {
			_, err = db.SaveProfileDatas(ctx, tx, u.ID, pdatas)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// GetUserByID gets a user object by id
func (db Db) GetUserByID(ctx context.Context, tx *firestore.Transaction, userID user.ID) (*user.User, error) {
	if userID == "" {
		return nil, nil

	}

	ref := db.Firestore.Collection("users").Doc(string(userID))

	var (
		snap *firestore.DocumentSnapshot
		err  error
	)
	if tx == nil {
		snap, err = ref.Get(ctx)
	} else {
		snap, err = tx.Get(ref)
	}

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
func (db Db) SavePlaidItem(ctx context.Context, userID user.ID, item *item.Item) error {
	encryptedItem, err := item.Encrypt(db.EncryptionManager, userID)
	if err != nil {
		return err
	}

	_, err = db.Firestore.Collection("users").Doc(string(userID)).Collection("plaidItems").Doc(string(item.ID)).Set(ctx, encryptedItem)
	if err != nil {
		return err
	}

	return nil
}

// SaveWyreAccount ...
func (db Db) SaveWyreAccount(ctx context.Context, tx *firestore.Transaction, userID user.ID, account *account.Account) error {
	ref := db.Firestore.Collection("users").Doc(string(userID)).Collection("wyreAccounts").Doc(string(account.ID))

	out, err := account.Encrypt(db.EncryptionManager, userID)
	if err != nil {
		return err
	}

	if tx == nil {
		_, err = ref.Set(ctx, out)
	} else {
		err = tx.Set(ref, out)
	}

	return err
}

// GetWyreAccounts ...
func (db Db) GetWyreAccounts(ctx context.Context, tx *firestore.Transaction, userID user.ID) ([]*account.Account, error) {
	ref := db.Firestore.Collection("users").Doc(string(userID)).Collection("wyreAccounts")

	var (
		snaps []*firestore.DocumentSnapshot
		err   error
	)
	if tx == nil {
		snaps, err = ref.Documents(ctx).GetAll()
	} else {
		snaps, err = tx.Documents(ref).GetAll()
	}
	if err != nil {
		return nil, err
	}

	var out []*account.Account

	for _, snap := range snaps {
		var encA account.EncryptedAccount

		err := snap.DataTo(&encA)
		if err != nil {
			return nil, err
		}

		account, err := encA.Decrypt(db.EncryptionManager, userID)
		if err != nil {
			return nil, err
		}

		out = append(out, account)
	}

	return out, nil
}

func (db Db) GetWyrePaymentMethods(ctx context.Context, tx *firestore.Transaction, userID user.ID, wyreAccountID account.ID) ([]*paymentmethod.PaymentMethod, error) {
	ref := db.Firestore.Collection("users").Doc(string(userID)).Collection("wyreAccounts").Doc(string(wyreAccountID)).Collection("wyrePaymentMethods")

	var (
		snaps []*firestore.DocumentSnapshot
		err   error
	)
	if tx == nil {
		snaps, err = ref.Documents(ctx).GetAll()
	} else {
		snaps, err = tx.Documents(ref).GetAll()
	}
	if err != nil {
		return nil, err
	}

	var out []*paymentmethod.PaymentMethod

	for _, snap := range snaps {
		var pm paymentmethod.PaymentMethod

		err := snap.DataTo(&pm)
		if err != nil {
			return nil, err
		}

		out = append(out, &pm)
	}

	return out, nil
}

// GetWyrePaymentMethodByPlaidAccountID ...
func (db Db) GetWyrePaymentMethodByPlaidAccountID(ctx context.Context, userID user.ID, wyreAccountID account.ID, plaidAccountID string) (*paymentmethod.PaymentMethod, error) {
	collection := db.Firestore.Collection("users").Doc(string(userID)).Collection("wyreAccounts").Doc(string(wyreAccountID)).Collection("wyrePaymentMethods")
	docs := collection.
		Where("plaidAccountID", "==", plaidAccountID).
		Limit(1).
		Documents(ctx)

	snap, err := docs.Next()
	if err == iterator.Done {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var pm paymentmethod.PaymentMethod
	err = snap.DataTo(&pm)
	if err != nil {
		return nil, err
	}

	return &pm, nil

}

func (db Db) UpdateEmail(ctx context.Context, userID user.ID, newEmail string) error {
	return db.Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		now := time.Now()

		// read
		u, err := db.GetUserByID(ctx, tx, userID)
		if err != nil {
			return err
		}
		if u == nil {
			return fmt.Errorf("user id not found: %s", userID)
		}

		profile, err := db.GetAllProfileData(ctx, tx, userID)
		if err != nil {
			return err
		}

		// modify / upsert
		u.Email = &newEmail
		u.EmailVerifiedAt = &now

		emails := profile.FilterStatus(common.StatusReceived).FilterKindEmail()
		if len(emails) == 0 {
			emails = append(emails, &email.ProfileDataEmail{
				CommonProfileData: common.CommonProfileData{
					ID:        common.ProfileDataID(shortuuid.New()),
					Status:    common.StatusReceived,
					CreatedAt: now,
				},
				Email: newEmail,
			})
		}
		email := emails[0]

		// save
		if err := db.SaveUser(ctx, tx, u); err != nil {
			return err
		}
		if _, err = db.SaveProfileData(ctx, tx, userID, email); err != nil {
			return err
		}

		return nil
	})
}

func (db Db) UpdatePhone(ctx context.Context, userID user.ID, newPhone string) error {
	return db.Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		now := time.Now()

		// read
		u, err := db.GetUserByID(ctx, tx, userID)
		if err != nil {
			return err
		}
		if u == nil {
			return fmt.Errorf("user id not found: %s", userID)
		}

		profile, err := db.GetAllProfileData(ctx, tx, userID)
		if err != nil {
			return err
		}

		// modify / upsert
		u.Phone = &newPhone
		u.PhoneVerifiedAt = &now

		phones := profile.FilterStatus(common.StatusReceived).FilterKindPhone()
		if len(phones) == 0 {
			phones = append(phones, &phone.ProfileDataPhone{
				CommonProfileData: common.CommonProfileData{
					ID:        common.ProfileDataID(shortuuid.New()),
					Status:    common.StatusReceived,
					CreatedAt: now,
				},
				Phone: newPhone,
			})
		}
		phone := phones[0]

		// save
		if err := db.SaveUser(ctx, tx, u); err != nil {
			return err
		}
		if _, err = db.SaveProfileData(ctx, tx, userID, phone); err != nil {
			return err
		}

		return nil
	})
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

	if tx == nil {
		_, err = ref.Set(ctx, out)
	} else {
		err = tx.Set(ref, out)
	}
	if err != nil {
		return "", err
	}

	return out.ID, nil
}

// SaveProfileDatas ...
func (db Db) SaveProfileDatas(ctx context.Context, passedTx *firestore.Transaction, userID user.ID, pdatas profiledata.ProfileDatas) ([]common.ProfileDataID, error) {
	ids := []common.ProfileDataID{}

	save := func(ctx context.Context, tx *firestore.Transaction) error {
		for _, pdata := range pdatas {
			id, err := db.SaveProfileData(ctx, tx, userID, pdata)
			if err != nil {
				return err
			}
			ids = append(ids, id)
		}
		return nil
	}

	var err error
	if passedTx == nil {
		err = db.Firestore.RunTransaction(ctx, save)
	} else {
		err = save(ctx, passedTx)
	}
	if err != nil {
		return nil, err
	}

	return ids, nil
}

// GetAllPlaidItems ...
func (db Db) GetAllPlaidItems(ctx context.Context, tx *firestore.Transaction, userID user.ID) ([]*item.Item, error) {
	ref := db.Firestore.Collection("users").Doc(string(userID)).Collection("plaidItems")

	var (
		docs []*firestore.DocumentSnapshot
		err  error
	)
	if tx == nil {
		docs, err = ref.Documents(ctx).GetAll()
	} else {
		docs, err = tx.Documents(ref).GetAll()
	}
	if err != nil {
		return nil, err
	}

	var out []*item.Item
	for _, snap := range docs {
		var enc item.EncryptedItem
		err := snap.DataTo(&enc)
		if err != nil {
			return nil, err
		}

		item, err := enc.Decrypt(db.EncryptionManager, userID)
		if err != nil {
			return nil, err
		}
		out = append(out, item)
	}

	return out, nil
}

// GetAllProfileData ...
func (db Db) GetAllProfileData(ctx context.Context, tx *firestore.Transaction, userID user.ID) (profiledata.ProfileDatas, error) {
	ref := db.Firestore.Collection("users").Doc(string(userID)).Collection("profile")

	var (
		docs []*firestore.DocumentSnapshot
		err  error
	)
	if tx == nil {
		docs, err = ref.Documents(ctx).GetAll()
	} else {
		docs, err = tx.Documents(ref).GetAll()
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

func (db Db) SaveUsedRefreshToken(ctx context.Context, tx *firestore.Transaction, urt *usedrefreshtoken.UsedRefreshToken) error {
	var err error

	ref := db.Firestore.Collection("used-refresh-tokens").Doc(urt.ID)

	if tx == nil {
		_, err = ref.Set(ctx, urt)
	} else {
		err = tx.Set(ref, urt)
	}

	return err

}

func (db Db) SaveWyrePaymentMethod(ctx context.Context, tx *firestore.Transaction, userID user.ID, wyreAccountID account.ID, wpm *paymentmethod.PaymentMethod) error {
	var err error

	ref := db.Firestore.Collection("users").Doc(string(userID)).Collection("wyreAccounts").Doc(string(wyreAccountID)).Collection("wyrePaymentMethods").Doc(string(wpm.ID))

	if tx == nil {
		_, err = ref.Set(ctx, wpm)
	} else {
		err = tx.Set(ref, wpm)
	}

	return err
}

func (db Db) GetUsedRefreshToken(ctx context.Context, tx *firestore.Transaction, id string) (*usedrefreshtoken.UsedRefreshToken, error) {
	var (
		err  error
		snap *firestore.DocumentSnapshot
	)

	ref := db.Firestore.Collection("used-refresh-tokens").Doc(id)

	if tx == nil {
		snap, err = ref.Get(ctx)
	} else {
		snap, err = tx.Get(ref)
	}
	if status.Code(err) == codes.NotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var usedRefreshToken usedrefreshtoken.UsedRefreshToken
	err = snap.DataTo(&usedRefreshToken)
	if err != nil {
		return nil, err
	}

	return &usedRefreshToken, nil
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

	var passcode onetimepasscode.OneTimePasscode
	err = snap.DataTo(&passcode)
	if err != nil {
		return nil, err
	}

	// can only be used 'once'
	_, err = snap.Ref.Delete(ctx)
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

// HasPlaidItems returns true if the user has plaid items
func (db Db) HasPlaidItems(ctx context.Context, userID user.ID) (bool, error) {
	plaidItems := db.Firestore.
		Collection("users").
		Doc(string(userID)).Collection("plaidItems").
		Limit(1).Documents(ctx)

	_, err := plaidItems.Next()
	if err == iterator.Done {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
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
