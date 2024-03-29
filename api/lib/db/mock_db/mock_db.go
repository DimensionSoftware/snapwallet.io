// Code generated by MockGen. DO NOT EDIT.
// Source: lib/db/db.go

// Package mock_db is a generated GoMock package.
package mock_db

import (
	firestore "cloud.google.com/go/firestore"
	context "context"
	gomock "github.com/golang/mock/gomock"
	gotoconfig "github.com/khoerling/flux/api/lib/db/models/gotoconfig"
	job "github.com/khoerling/flux/api/lib/db/models/job"
	onetimepasscode "github.com/khoerling/flux/api/lib/db/models/onetimepasscode"
	usedrefreshtoken "github.com/khoerling/flux/api/lib/db/models/usedrefreshtoken"
	user "github.com/khoerling/flux/api/lib/db/models/user"
	file "github.com/khoerling/flux/api/lib/db/models/user/file"
	item "github.com/khoerling/flux/api/lib/db/models/user/plaid/item"
	profiledata "github.com/khoerling/flux/api/lib/db/models/user/profiledata"
	common "github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	transaction "github.com/khoerling/flux/api/lib/db/models/user/transaction"
	account "github.com/khoerling/flux/api/lib/db/models/user/wyre/account"
	paymentmethod "github.com/khoerling/flux/api/lib/db/models/user/wyre/paymentmethod"
	reflect "reflect"
)

// MockDb is a mock of Db interface
type MockDb struct {
	ctrl     *gomock.Controller
	recorder *MockDbMockRecorder
}

// MockDbMockRecorder is the mock recorder for MockDb
type MockDbMockRecorder struct {
	mock *MockDb
}

// NewMockDb creates a new mock instance
func NewMockDb(ctrl *gomock.Controller) *MockDb {
	mock := &MockDb{ctrl: ctrl}
	mock.recorder = &MockDbMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDb) EXPECT() *MockDbMockRecorder {
	return m.recorder
}

// RunTransaction mocks base method
func (m *MockDb) RunTransaction(arg0 context.Context, arg1 func(context.Context, *firestore.Transaction) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunTransaction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTransaction indicates an expected call of RunTransaction
func (mr *MockDbMockRecorder) RunTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTransaction", reflect.TypeOf((*MockDb)(nil).RunTransaction), arg0, arg1)
}

// SaveGotoConfig mocks base method
func (m *MockDb) SaveGotoConfig(ctx context.Context, g *gotoconfig.Config) (gotoconfig.ShortID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveGotoConfig", ctx, g)
	ret0, _ := ret[0].(gotoconfig.ShortID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveGotoConfig indicates an expected call of SaveGotoConfig
func (mr *MockDbMockRecorder) SaveGotoConfig(ctx, g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveGotoConfig", reflect.TypeOf((*MockDb)(nil).SaveGotoConfig), ctx, g)
}

// GetGotoConfigByShortID mocks base method
func (m *MockDb) GetGotoConfigByShortID(ctx context.Context, shortID gotoconfig.ShortID) (*gotoconfig.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGotoConfigByShortID", ctx, shortID)
	ret0, _ := ret[0].(*gotoconfig.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGotoConfigByShortID indicates an expected call of GetGotoConfigByShortID
func (mr *MockDbMockRecorder) GetGotoConfigByShortID(ctx, shortID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGotoConfigByShortID", reflect.TypeOf((*MockDb)(nil).GetGotoConfigByShortID), ctx, shortID)
}

// CreateOneTimePasscode mocks base method
func (m *MockDb) CreateOneTimePasscode(ctx context.Context, emailOrPhone string, kind onetimepasscode.LoginKind) (*onetimepasscode.OneTimePasscode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOneTimePasscode", ctx, emailOrPhone, kind)
	ret0, _ := ret[0].(*onetimepasscode.OneTimePasscode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOneTimePasscode indicates an expected call of CreateOneTimePasscode
func (mr *MockDbMockRecorder) CreateOneTimePasscode(ctx, emailOrPhone, kind interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOneTimePasscode", reflect.TypeOf((*MockDb)(nil).CreateOneTimePasscode), ctx, emailOrPhone, kind)
}

// SaveUser mocks base method
func (m *MockDb) SaveUser(ctx context.Context, tx *firestore.Transaction, u *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUser", ctx, tx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUser indicates an expected call of SaveUser
func (mr *MockDbMockRecorder) SaveUser(ctx, tx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUser", reflect.TypeOf((*MockDb)(nil).SaveUser), ctx, tx, u)
}

// SaveJob mocks base method
func (m *MockDb) SaveJob(ctx context.Context, tx *firestore.Transaction, j *job.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveJob", ctx, tx, j)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveJob indicates an expected call of SaveJob
func (mr *MockDbMockRecorder) SaveJob(ctx, tx, j interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveJob", reflect.TypeOf((*MockDb)(nil).SaveJob), ctx, tx, j)
}

// SaveTransaction mocks base method
func (m *MockDb) SaveTransaction(ctx context.Context, tx *firestore.Transaction, userID user.ID, transaction *transaction.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTransaction", ctx, tx, userID, transaction)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTransaction indicates an expected call of SaveTransaction
func (mr *MockDbMockRecorder) SaveTransaction(ctx, tx, userID, transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTransaction", reflect.TypeOf((*MockDb)(nil).SaveTransaction), ctx, tx, userID, transaction)
}

// GetTransactions mocks base method
func (m *MockDb) GetTransactions(ctx context.Context, userID user.ID) (*transaction.Transactions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions", ctx, userID)
	ret0, _ := ret[0].(*transaction.Transactions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactions indicates an expected call of GetTransactions
func (mr *MockDbMockRecorder) GetTransactions(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockDb)(nil).GetTransactions), ctx, userID)
}

// GetTransactionByExternalId mocks base method
func (m *MockDb) GetTransactionByExternalId(ctx context.Context, tx *firestore.Transaction, userID user.ID, externalID transaction.ExternalID) (*transaction.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByExternalId", ctx, tx, userID, externalID)
	ret0, _ := ret[0].(*transaction.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByExternalId indicates an expected call of GetTransactionByExternalId
func (mr *MockDbMockRecorder) GetTransactionByExternalId(ctx, tx, userID, externalID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByExternalId", reflect.TypeOf((*MockDb)(nil).GetTransactionByExternalId), ctx, tx, userID, externalID)
}

// GetJobByKindAndStatusAndRelatedId mocks base method
func (m *MockDb) GetJobByKindAndStatusAndRelatedId(ctx context.Context, kind job.Kind, status job.Status, relatedID string) (*job.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobByKindAndStatusAndRelatedId", ctx, kind, status, relatedID)
	ret0, _ := ret[0].(*job.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobByKindAndStatusAndRelatedId indicates an expected call of GetJobByKindAndStatusAndRelatedId
func (mr *MockDbMockRecorder) GetJobByKindAndStatusAndRelatedId(ctx, kind, status, relatedID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobByKindAndStatusAndRelatedId", reflect.TypeOf((*MockDb)(nil).GetJobByKindAndStatusAndRelatedId), ctx, kind, status, relatedID)
}

// SaveFileMetadata mocks base method
func (m *MockDb) SaveFileMetadata(ctx context.Context, userID user.ID, md *file.Metadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFileMetadata", ctx, userID, md)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveFileMetadata indicates an expected call of SaveFileMetadata
func (mr *MockDbMockRecorder) SaveFileMetadata(ctx, userID, md interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFileMetadata", reflect.TypeOf((*MockDb)(nil).SaveFileMetadata), ctx, userID, md)
}

// GetFileMetadata mocks base method
func (m *MockDb) GetFileMetadata(ctx context.Context, userID user.ID, fileID file.ID) (*file.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileMetadata", ctx, userID, fileID)
	ret0, _ := ret[0].(*file.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileMetadata indicates an expected call of GetFileMetadata
func (mr *MockDbMockRecorder) GetFileMetadata(ctx, userID, fileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileMetadata", reflect.TypeOf((*MockDb)(nil).GetFileMetadata), ctx, userID, fileID)
}

// GetOrCreateUser mocks base method
func (m *MockDb) GetOrCreateUser(ctx context.Context, loginKind onetimepasscode.LoginKind, emailOrPhone string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateUser", ctx, loginKind, emailOrPhone)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrCreateUser indicates an expected call of GetOrCreateUser
func (mr *MockDbMockRecorder) GetOrCreateUser(ctx, loginKind, emailOrPhone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateUser", reflect.TypeOf((*MockDb)(nil).GetOrCreateUser), ctx, loginKind, emailOrPhone)
}

// GetUserByID mocks base method
func (m *MockDb) GetUserByID(ctx context.Context, tx *firestore.Transaction, userID user.ID) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, tx, userID)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID
func (mr *MockDbMockRecorder) GetUserByID(ctx, tx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockDb)(nil).GetUserByID), ctx, tx, userID)
}

// SavePlaidItem mocks base method
func (m *MockDb) SavePlaidItem(ctx context.Context, userID user.ID, item *item.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePlaidItem", ctx, userID, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavePlaidItem indicates an expected call of SavePlaidItem
func (mr *MockDbMockRecorder) SavePlaidItem(ctx, userID, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePlaidItem", reflect.TypeOf((*MockDb)(nil).SavePlaidItem), ctx, userID, item)
}

// SaveWyreAccount mocks base method
func (m *MockDb) SaveWyreAccount(ctx context.Context, tx *firestore.Transaction, userID user.ID, account *account.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveWyreAccount", ctx, tx, userID, account)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveWyreAccount indicates an expected call of SaveWyreAccount
func (mr *MockDbMockRecorder) SaveWyreAccount(ctx, tx, userID, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveWyreAccount", reflect.TypeOf((*MockDb)(nil).SaveWyreAccount), ctx, tx, userID, account)
}

// GetWyreAccounts mocks base method
func (m *MockDb) GetWyreAccounts(ctx context.Context, tx *firestore.Transaction, userID user.ID) ([]*account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWyreAccounts", ctx, tx, userID)
	ret0, _ := ret[0].([]*account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWyreAccounts indicates an expected call of GetWyreAccounts
func (mr *MockDbMockRecorder) GetWyreAccounts(ctx, tx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWyreAccounts", reflect.TypeOf((*MockDb)(nil).GetWyreAccounts), ctx, tx, userID)
}

// GetWyrePaymentMethods mocks base method
func (m *MockDb) GetWyrePaymentMethods(ctx context.Context, tx *firestore.Transaction, userID user.ID, wyreAccountID account.ID) ([]*paymentmethod.PaymentMethod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWyrePaymentMethods", ctx, tx, userID, wyreAccountID)
	ret0, _ := ret[0].([]*paymentmethod.PaymentMethod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWyrePaymentMethods indicates an expected call of GetWyrePaymentMethods
func (mr *MockDbMockRecorder) GetWyrePaymentMethods(ctx, tx, userID, wyreAccountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWyrePaymentMethods", reflect.TypeOf((*MockDb)(nil).GetWyrePaymentMethods), ctx, tx, userID, wyreAccountID)
}

// GetWyrePaymentMethodByPlaidAccountID mocks base method
func (m *MockDb) GetWyrePaymentMethodByPlaidAccountID(ctx context.Context, userID user.ID, wyreAccountID account.ID, plaidAccountID string) (*paymentmethod.PaymentMethod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWyrePaymentMethodByPlaidAccountID", ctx, userID, wyreAccountID, plaidAccountID)
	ret0, _ := ret[0].(*paymentmethod.PaymentMethod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWyrePaymentMethodByPlaidAccountID indicates an expected call of GetWyrePaymentMethodByPlaidAccountID
func (mr *MockDbMockRecorder) GetWyrePaymentMethodByPlaidAccountID(ctx, userID, wyreAccountID, plaidAccountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWyrePaymentMethodByPlaidAccountID", reflect.TypeOf((*MockDb)(nil).GetWyrePaymentMethodByPlaidAccountID), ctx, userID, wyreAccountID, plaidAccountID)
}

// UpdateEmail mocks base method
func (m *MockDb) UpdateEmail(ctx context.Context, userID user.ID, newEmail string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmail", ctx, userID, newEmail)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEmail indicates an expected call of UpdateEmail
func (mr *MockDbMockRecorder) UpdateEmail(ctx, userID, newEmail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmail", reflect.TypeOf((*MockDb)(nil).UpdateEmail), ctx, userID, newEmail)
}

// UpdatePhone mocks base method
func (m *MockDb) UpdatePhone(ctx context.Context, userID user.ID, newPhone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhone", ctx, userID, newPhone)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePhone indicates an expected call of UpdatePhone
func (mr *MockDbMockRecorder) UpdatePhone(ctx, userID, newPhone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhone", reflect.TypeOf((*MockDb)(nil).UpdatePhone), ctx, userID, newPhone)
}

// GetUserByEmailOrPhone mocks base method
func (m *MockDb) GetUserByEmailOrPhone(ctx context.Context, tx *firestore.Transaction, emailOrPhone string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmailOrPhone", ctx, tx, emailOrPhone)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmailOrPhone indicates an expected call of GetUserByEmailOrPhone
func (mr *MockDbMockRecorder) GetUserByEmailOrPhone(ctx, tx, emailOrPhone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmailOrPhone", reflect.TypeOf((*MockDb)(nil).GetUserByEmailOrPhone), ctx, tx, emailOrPhone)
}

// SaveProfileData mocks base method
func (m *MockDb) SaveProfileData(ctx context.Context, tx *firestore.Transaction, userID user.ID, pdata profiledata.ProfileData) (common.ProfileDataID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveProfileData", ctx, tx, userID, pdata)
	ret0, _ := ret[0].(common.ProfileDataID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveProfileData indicates an expected call of SaveProfileData
func (mr *MockDbMockRecorder) SaveProfileData(ctx, tx, userID, pdata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveProfileData", reflect.TypeOf((*MockDb)(nil).SaveProfileData), ctx, tx, userID, pdata)
}

// SaveProfileDatas mocks base method
func (m *MockDb) SaveProfileDatas(ctx context.Context, passedTx *firestore.Transaction, userID user.ID, pdatas profiledata.ProfileDatas) ([]common.ProfileDataID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveProfileDatas", ctx, passedTx, userID, pdatas)
	ret0, _ := ret[0].([]common.ProfileDataID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveProfileDatas indicates an expected call of SaveProfileDatas
func (mr *MockDbMockRecorder) SaveProfileDatas(ctx, passedTx, userID, pdatas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveProfileDatas", reflect.TypeOf((*MockDb)(nil).SaveProfileDatas), ctx, passedTx, userID, pdatas)
}

// GetAllPlaidItems mocks base method
func (m *MockDb) GetAllPlaidItems(ctx context.Context, tx *firestore.Transaction, userID user.ID) ([]*item.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPlaidItems", ctx, tx, userID)
	ret0, _ := ret[0].([]*item.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPlaidItems indicates an expected call of GetAllPlaidItems
func (mr *MockDbMockRecorder) GetAllPlaidItems(ctx, tx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPlaidItems", reflect.TypeOf((*MockDb)(nil).GetAllPlaidItems), ctx, tx, userID)
}

// GetAllProfileData mocks base method
func (m *MockDb) GetAllProfileData(ctx context.Context, tx *firestore.Transaction, userID user.ID) (profiledata.ProfileDatas, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProfileData", ctx, tx, userID)
	ret0, _ := ret[0].(profiledata.ProfileDatas)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProfileData indicates an expected call of GetAllProfileData
func (mr *MockDbMockRecorder) GetAllProfileData(ctx, tx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProfileData", reflect.TypeOf((*MockDb)(nil).GetAllProfileData), ctx, tx, userID)
}

// SaveUsedRefreshToken mocks base method
func (m *MockDb) SaveUsedRefreshToken(ctx context.Context, tx *firestore.Transaction, urt *usedrefreshtoken.UsedRefreshToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUsedRefreshToken", ctx, tx, urt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUsedRefreshToken indicates an expected call of SaveUsedRefreshToken
func (mr *MockDbMockRecorder) SaveUsedRefreshToken(ctx, tx, urt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUsedRefreshToken", reflect.TypeOf((*MockDb)(nil).SaveUsedRefreshToken), ctx, tx, urt)
}

// SaveWyrePaymentMethod mocks base method
func (m *MockDb) SaveWyrePaymentMethod(ctx context.Context, tx *firestore.Transaction, userID user.ID, wyreAccountID account.ID, wpm *paymentmethod.PaymentMethod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveWyrePaymentMethod", ctx, tx, userID, wyreAccountID, wpm)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveWyrePaymentMethod indicates an expected call of SaveWyrePaymentMethod
func (mr *MockDbMockRecorder) SaveWyrePaymentMethod(ctx, tx, userID, wyreAccountID, wpm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveWyrePaymentMethod", reflect.TypeOf((*MockDb)(nil).SaveWyrePaymentMethod), ctx, tx, userID, wyreAccountID, wpm)
}

// GetUsedRefreshToken mocks base method
func (m *MockDb) GetUsedRefreshToken(ctx context.Context, tx *firestore.Transaction, id string) (*usedrefreshtoken.UsedRefreshToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsedRefreshToken", ctx, tx, id)
	ret0, _ := ret[0].(*usedrefreshtoken.UsedRefreshToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsedRefreshToken indicates an expected call of GetUsedRefreshToken
func (mr *MockDbMockRecorder) GetUsedRefreshToken(ctx, tx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsedRefreshToken", reflect.TypeOf((*MockDb)(nil).GetUsedRefreshToken), ctx, tx, id)
}

// AckOneTimePasscode mocks base method
func (m *MockDb) AckOneTimePasscode(ctx context.Context, loginValue, code string) (*onetimepasscode.OneTimePasscode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AckOneTimePasscode", ctx, loginValue, code)
	ret0, _ := ret[0].(*onetimepasscode.OneTimePasscode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AckOneTimePasscode indicates an expected call of AckOneTimePasscode
func (mr *MockDbMockRecorder) AckOneTimePasscode(ctx, loginValue, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AckOneTimePasscode", reflect.TypeOf((*MockDb)(nil).AckOneTimePasscode), ctx, loginValue, code)
}

// HasPlaidItems mocks base method
func (m *MockDb) HasPlaidItems(ctx context.Context, userID user.ID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPlaidItems", ctx, userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasPlaidItems indicates an expected call of HasPlaidItems
func (mr *MockDbMockRecorder) HasPlaidItems(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPlaidItems", reflect.TypeOf((*MockDb)(nil).HasPlaidItems), ctx, userID)
}

// CleanAgedPasscodes mocks base method
func (m *MockDb) CleanAgedPasscodes(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanAgedPasscodes", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CleanAgedPasscodes indicates an expected call of CleanAgedPasscodes
func (mr *MockDbMockRecorder) CleanAgedPasscodes(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanAgedPasscodes", reflect.TypeOf((*MockDb)(nil).CleanAgedPasscodes), ctx)
}

// GetUserByWyreAccountID mocks base method
func (m *MockDb) GetUserByWyreAccountID(ctx context.Context, wyreAccountID account.ID) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByWyreAccountID", ctx, wyreAccountID)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByWyreAccountID indicates an expected call of GetUserByWyreAccountID
func (mr *MockDbMockRecorder) GetUserByWyreAccountID(ctx, wyreAccountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByWyreAccountID", reflect.TypeOf((*MockDb)(nil).GetUserByWyreAccountID), ctx, wyreAccountID)
}
