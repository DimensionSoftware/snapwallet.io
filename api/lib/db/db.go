//

package db

import (
	"context"

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
	"github.com/khoerling/flux/api/lib/db/models/user/transaction"
	"github.com/khoerling/flux/api/lib/db/models/user/wyre/account"
	"github.com/khoerling/flux/api/lib/db/models/user/wyre/paymentmethod"
)

// Db ...
type Db interface {
	RunTransaction(context.Context, func(ctx context.Context, tx *firestore.Transaction) error) error
	// will not save if item is already existing; returns short id of immutable first item
	SaveGotoConfig(ctx context.Context, g *gotoconfig.Config) (gotoconfig.ShortID, error)
	GetGotoConfigByShortID(ctx context.Context, shortID gotoconfig.ShortID) (*gotoconfig.Config, error)
	// CreateOneTimePasscode stores a record of a one-time-password request for verification later
	CreateOneTimePasscode(ctx context.Context, emailOrPhone string, kind onetimepasscode.LoginKind) (*onetimepasscode.OneTimePasscode, error)
	// SaveUser saves a user object (upsert/put semantics)
	SaveUser(ctx context.Context, tx *firestore.Transaction, u *user.User) error
	SaveJob(ctx context.Context, tx *firestore.Transaction, j *job.Job) error
	SaveTransaction(ctx context.Context, tx *firestore.Transaction, userID user.ID, transaction *transaction.Transaction) error
	GetTransactions(ctx context.Context, userID user.ID) (*transaction.Transactions, error)
	GetTransactionByExternalId(ctx context.Context, tx *firestore.Transaction, userID user.ID, externalID transaction.ExternalID) (*transaction.Transaction, error)
	GetJobByKindAndStatusAndRelatedId(ctx context.Context, kind job.Kind, status job.Status, relatedID string) (*job.Job, error)
	// SaveFileMetadata saves file metadata
	SaveFileMetadata(ctx context.Context, userID user.ID, md *file.Metadata) error
	// GetFileMetadata gets file metadata
	GetFileMetadata(ctx context.Context, userID user.ID, fileID file.ID) (*file.Metadata, error)
	// GetOrCreateUser creates a user object
	GetOrCreateUser(ctx context.Context, loginKind onetimepasscode.LoginKind, emailOrPhone string) (*user.User, error)
	// GetUserByID gets a user object by id
	GetUserByID(ctx context.Context, tx *firestore.Transaction, userID user.ID) (*user.User, error)
	// SavePlaidItem ...
	SavePlaidItem(ctx context.Context, userID user.ID, item *item.Item) error
	// SaveWyreAccount ...
	SaveWyreAccount(ctx context.Context, tx *firestore.Transaction, userID user.ID, account *account.Account) error
	// GetWyreAccounts ...
	GetWyreAccounts(ctx context.Context, tx *firestore.Transaction, userID user.ID) ([]*account.Account, error)
	GetWyrePaymentMethods(ctx context.Context, tx *firestore.Transaction, userID user.ID, wyreAccountID account.ID) ([]*paymentmethod.PaymentMethod, error)
	// GetWyrePaymentMethodByPlaidAccountID ...
	GetWyrePaymentMethodByPlaidAccountID(ctx context.Context, userID user.ID, wyreAccountID account.ID, plaidAccountID string) (*paymentmethod.PaymentMethod, error)
	UpdateEmail(ctx context.Context, userID user.ID, newEmail string) error
	UpdatePhone(ctx context.Context, userID user.ID, newPhone string) error
	// GetUserByEmailOrPhone will return a user if one is found matching the input by email or phone
	GetUserByEmailOrPhone(ctx context.Context, tx *firestore.Transaction, emailOrPhone string) (*user.User, error)
	// SaveProfileData ...
	SaveProfileData(ctx context.Context, tx *firestore.Transaction, userID user.ID, pdata profiledata.ProfileData) (common.ProfileDataID, error)
	// SaveProfileDatas ...
	SaveProfileDatas(ctx context.Context, passedTx *firestore.Transaction, userID user.ID, pdatas profiledata.ProfileDatas) ([]common.ProfileDataID, error)
	// GetAllPlaidItems ...
	GetAllPlaidItems(ctx context.Context, tx *firestore.Transaction, userID user.ID) ([]*item.Item, error)
	// GetAllProfileData ...
	GetAllProfileData(ctx context.Context, tx *firestore.Transaction, userID user.ID) (profiledata.ProfileDatas, error)
	SaveUsedRefreshToken(ctx context.Context, tx *firestore.Transaction, urt *usedrefreshtoken.UsedRefreshToken) error
	SaveWyrePaymentMethod(ctx context.Context, tx *firestore.Transaction, userID user.ID, wyreAccountID account.ID, wpm *paymentmethod.PaymentMethod) error
	GetUsedRefreshToken(ctx context.Context, tx *firestore.Transaction, id string) (*usedrefreshtoken.UsedRefreshToken, error)
	// AckOneTimePasscode tries to find the OneTimePasscode object matching the loginValue and code
	//   If it is found then the ack is successful and the passcode is destroyed in the database
	//   after being destroyed the original passcode object will be returned.
	//
	//   If it is not found then (nil, nil) will be returned so the caller can decide the error handling strategy
	//   for not found (invalid) codes (Not Acked)
	//
	//   If there is another type of error then (nil, error) will be returned.
	AckOneTimePasscode(ctx context.Context, loginValue string, code string) (*onetimepasscode.OneTimePasscode, error)
	// HasPlaidItems returns true if the user has plaid items
	HasPlaidItems(ctx context.Context, userID user.ID) (bool, error)
	// CleanAgedPasscodes cleanup aged passcodes; there is no security risk to not running this but it saves on db storage bills
	// returns num of passcodes deleted which were old, or an error if shit goes bad
	CleanAgedPasscodes(ctx context.Context) (int, error)
	// GetUserByID gets a user object by id
	GetUserByWyreAccountID(ctx context.Context, wyreAccountID account.ID) (*user.User, error)
}
