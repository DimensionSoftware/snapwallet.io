package db

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/khoerling/flux/api/lib/db/models/onetimepasscode"
	"github.com/khoerling/flux/api/lib/db/models/usedrefreshtoken"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/file"
	"github.com/khoerling/flux/api/lib/db/models/user/plaid/item"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/db/models/user/wyre/account"
	"github.com/khoerling/flux/api/lib/db/models/user/wyre/paymentmethod"
)

type Db interface {
	CreateOneTimePasscode(ctx context.Context, emailOrPhone string, kind onetimepasscode.LoginKind) (*onetimepasscode.OneTimePasscode, error)
	SaveUser(ctx context.Context, tx *firestore.Transaction, u *user.User) error
	SaveFileMetadata(ctx context.Context, userID user.ID, md *file.Metadata) error
	GetFileMetadata(ctx context.Context, userID user.ID, fileID file.ID) (*file.Metadata, error)
	GetOrCreateUser(ctx context.Context, loginKind onetimepasscode.LoginKind, emailOrPhone string) (*user.User, error)
	GetUserByID(ctx context.Context, tx *firestore.Transaction, userID user.ID) (*user.User, error)
	SavePlaidItem(ctx context.Context, userID user.ID, itemID item.ID, accessToken string, accountIDs []string) (*item.Item, error)
	SaveWyreAccount(ctx context.Context, tx *firestore.Transaction, userID user.ID, account *account.Account) error
	GetWyreAccounts(ctx context.Context, tx *firestore.Transaction, userID user.ID) ([]*account.Account, error)
	GetWyrePaymentMethods(ctx context.Context, tx *firestore.Transaction, userID user.ID, wyreAccountID account.ID) ([]*paymentmethod.PaymentMethod, error)
	GetWyrePaymentMethodByPlaidAccountID(ctx context.Context, userID user.ID, wyreAccountID account.ID, plaidAccountID string) (*paymentmethod.PaymentMethod, error)
	UpdateEmail(ctx context.Context, userID user.ID, newEmail string) error
	UpdatePhone(ctx context.Context, userID user.ID, newPhone string) error
	GetUserByEmailOrPhone(ctx context.Context, tx *firestore.Transaction, emailOrPhone string) (*user.User, error)
	SaveProfileData(ctx context.Context, tx *firestore.Transaction, userID user.ID, pdata profiledata.ProfileData) (common.ProfileDataID, error)
	SaveProfileDatas(ctx context.Context, passedTx *firestore.Transaction, userID user.ID, pdatas profiledata.ProfileDatas) ([]common.ProfileDataID, error)
	GetAllPlaidItems(ctx context.Context, tx *firestore.Transaction, userID user.ID) ([]*item.Item, error)
	GetAllProfileData(ctx context.Context, tx *firestore.Transaction, userID user.ID) (profiledata.ProfileDatas, error)
	SaveUsedRefreshToken(ctx context.Context, tx *firestore.Transaction, urt *usedrefreshtoken.UsedRefreshToken) error
	SaveWyrePaymentMethod(ctx context.Context, tx *firestore.Transaction, userID user.ID, wyreAccountID account.ID, wpm *paymentmethod.PaymentMethod) error
	GetUsedRefreshToken(ctx context.Context, tx *firestore.Transaction, id string) (*usedrefreshtoken.UsedRefreshToken, error)
	AckOneTimePasscode(ctx context.Context, loginValue string, code string) (*onetimepasscode.OneTimePasscode, error)
	HasPlaidItems(ctx context.Context, userID user.ID) (bool, error)
	CleanAgedPasscodes(ctx context.Context) (int, error)
}
