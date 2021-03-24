package wyre

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/plaid/item"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/db/models/user/wyre/account"
	wyre_model "github.com/khoerling/flux/api/lib/db/models/user/wyre/account"
	"github.com/khoerling/flux/api/lib/db/models/user/wyre/paymentmethod"
	"github.com/lithammer/shortuuid/v3"
	"github.com/plaid/plaid-go/plaid"
)

const apiHostEnvVarName = "API_HOST"

type APIHost string

type Manager struct {
	APIHost APIHost
	Wyre    *Client
	Db      *db.Db
	Plaid   *plaid.Client
}

// ProvideAPIHost ...
func ProvideAPIHost() (APIHost, error) {
	apiHost := os.Getenv(apiHostEnvVarName)
	if apiHost == "" {
		return "", fmt.Errorf("you must set %s", apiHost)
	}

	log.Println("🚨 API Host for webhooks set to: ", apiHost)

	return APIHost(apiHost), nil
}

func (m Manager) CreatePaymentMethod(ctx context.Context, userID user.ID, wyreAccountID account.ID, plaidAccessToken string, plaidItemID string, plaidAccountID string) (*paymentmethod.PaymentMethod, error) {
	wyreAccounts, err := m.Db.GetWyreAccounts(ctx, nil, userID)
	if err != nil {
		return nil, err
	}
	var wyreAccount *account.Account
	for _, wyreAcc := range wyreAccounts {
		if wyreAcc.ID == wyreAccountID {
			wyreAccount = wyreAcc
		}
	}
	if wyreAccount == nil {
		return nil, fmt.Errorf("the wyreAccountID doesn't exist or is not associated with this user")
	}

	existingPm, err := m.Db.GetWyrePaymentMethodByPlaidAccountID(ctx, userID, wyreAccountID, plaidAccountID)
	if err != nil {
		return nil, err
	}
	if existingPm != nil {
		return nil, nil
	}

	resp, err := m.Plaid.CreateProcessorToken(plaidAccessToken, plaidAccountID, "wyre")
	if err != nil {
		return nil, err
	}

	wyrePm, err := m.Wyre.CreatePaymentMethod(wyreAccount.SecretKey, CreatePaymentMethodRequest{
		PlaidProcessorToken: resp.ProcessorToken,
	}.WithDefaults())
	if err != nil {
		return nil, err
	}

	pm := paymentmethod.PaymentMethod{
		ID:                    paymentmethod.ID(wyrePm.ID),
		PlaidItemID:           plaidItemID,
		PlaidAccountID:        plaidAccountID,
		Status:                wyrePm.Status,
		Name:                  wyrePm.Name,
		Last4:                 wyrePm.Last4Digits,
		ChargeableCurrencies:  wyrePm.ChargeableCurrencies,
		DepositableCurrencies: wyrePm.DepositableCurrencies,
		UpdatedAt:             time.Now(),
		CreatedAt:             time.Now(),
	}
	m.Db.SaveWyrePaymentMethod(ctx, nil, userID, wyreAccountID, &pm)
	if err != nil {
		return nil, err
	}

	return &pm, nil
}

func (m Manager) CreatePaymentMethodsFromPlaidItem(ctx context.Context, userID user.ID, wyreAccountID account.ID, item *item.Item) ([]*paymentmethod.PaymentMethod, error) {
	var out []*paymentmethod.PaymentMethod

	for _, accountID := range item.AccountIDs {
		pm, err := m.CreatePaymentMethod(ctx, userID, wyreAccountID, item.AccessToken, string(item.ID), accountID)
		if err != nil {
			return nil, err
		}
		if pm != nil {
			out = append(out, pm)
		}
	}

	return out, nil
}

func (m Manager) CreatePaymentMethodsFromPlaidItems(ctx context.Context, userID user.ID, wyreAccountID account.ID, items []*item.Item) ([]*paymentmethod.PaymentMethod, error) {
	var out []*paymentmethod.PaymentMethod

	for _, item := range items {
		pms, err := m.CreatePaymentMethodsFromPlaidItem(ctx, userID, wyreAccountID, item)
		if err != nil {
			return nil, err
		}

		for _, pm := range pms {
			out = append(out, pm)
		}
	}

	return out, nil
}

// converts pdata into wyre format
func selectWyreProfileFields(profile profiledata.ProfileDatas) ([]ProfileField, profiledata.ProfileDatas) {
	var fields []ProfileField
	var selected profiledata.ProfileDatas

	if legalNames := profile.FilterStatus(common.StatusReceived).FilterKindLegalName(); len(legalNames) > 0 {
		selected = append(selected, legalNames[0])
		fields = append(fields, ProfileField{
			FieldID: ProfileFieldIDIndividualLegalName,
			Value:   legalNames[0].LegalName,
		})
	}

	if phones := profile.FilterStatus(common.StatusReceived).FilterKindPhone(); len(phones) > 0 {
		selected = append(selected, phones[0])
		fields = append(fields, ProfileField{
			FieldID: ProfileFieldIDIndividualCellphoneNumber,
			Value:   phones[0].Phone,
		})
	}

	if emails := profile.FilterStatus(common.StatusReceived).FilterKindEmail(); len(emails) > 0 {
		selected = append(selected, emails[0])
		fields = append(fields, ProfileField{
			FieldID: ProfileFieldIDIndividualEmail,
			Value:   emails[0].Email,
		})
	}

	if dobs := profile.FilterStatus(common.StatusReceived).FilterKindDateOfBirth(); len(dobs) > 0 {
		selected = append(selected, dobs[0])
		fields = append(fields, ProfileField{
			FieldID: ProfileFieldIDIndividualDateOfBirth,
			Value:   dobs[0].DateOfBirth,
		})
	}

	if ssns := profile.FilterStatus(common.StatusReceived).FilterKindSSN(); len(ssns) > 0 {
		selected = append(selected, ssns[0])
		fields = append(fields, ProfileField{
			FieldID: ProfileFieldIDIndividualSSN,
			Value:   ssns[0].SSN,
		})
	}

	if addrs := profile.FilterStatus(common.StatusReceived).FilterKindAddress(); len(addrs) > 0 {
		address := addrs[0]
		selected = append(selected, address)
		fields = append(fields,
			ProfileField{
				FieldID: ProfileFieldIDIndividualResidenceAddress,
				Value: ProfileFieldAddress{
					Street1:    address.Street1,
					Street2:    address.Street2,
					City:       address.City,
					State:      address.State,
					PostalCode: address.PostalCode,
					Country:    address.Country,
				},
			},
		)
	}

	return fields, selected
}

func (m Manager) CreateAccount(ctx context.Context, userID user.ID, profile profiledata.ProfileDatas) (*wyre_model.Account, error) {
	now := time.Now()
	t := true

	if !profile.HasWyreAccountPreconditionsMet() {
		return nil, fmt.Errorf("Profile data is not complete enough to submit to Wyre (preconditions are unmet)")
	}

	/*
		secretKey := GenerateSecretKey(35)
		wyreAuthTokenResp, err := m.Wyre.SubmitAuthToken(secretKey)
		if err != nil {
			return nil, err
		}
		fmt.Printf("wyreAuthTokenResp: %#v", wyreAuthTokenResp)
	*/

	fields, selected := selectWyreProfileFields(profile)

	wyreAccountResp, err := m.Wyre.CreateAccount(m.Wyre.config.WyreSecretKey, CreateAccountRequest{
		SubAccount:        &t,
		DisableEmail:      &t,
		ReferrerAccountID: &m.Wyre.config.WyreAccountID,
		ProfileFields:     fields,
	}.WithDefaults())
	if err != nil {
		return nil, err
	}

	accountAPIKey, err := m.Wyre.CreateAPIKey(
		m.Wyre.config.WyreSecretKey,
		wyreAccountResp.ID,
		CreateAPIKeyRequest{
			Description: fmt.Sprintf("snapwallet.io user %s", userID),
			Type:        "FULL",
		})
	if err != nil {
		return nil, err
	}

	modifiedProfile := selected.SetStatuses(common.StatusPending)

	// todo, can't create account if they already have one

	account := wyre_model.Account{
		ID:        wyre_model.ID(wyreAccountResp.ID),
		APIKey:    accountAPIKey.APIKey,
		SecretKey: accountAPIKey.SecretKey,
		Status:    wyreAccountResp.Status,
		CreatedAt: now,
	}

	err = m.Db.SaveWyreAccount(ctx, nil, userID, &account)
	if err != nil {
		return nil, err
	}

	//TODO: use tx
	//TODO:  upload 2 docs proof of address, govt id
	_, err = m.Db.SaveProfileDatas(ctx, nil, userID, modifiedProfile)
	if err != nil {
		return nil, err
	}

	hookResponse, err := m.Wyre.SubscribeWebhook(accountAPIKey.SecretKey, "account:"+string(account.ID), string(m.APIHost)+"/wyre/hooks/"+string(userID))
	if err != nil {
		return nil, err
	}
	log.Printf("hook response from wyre: %#v", hookResponse)

	return &account, nil
}

// GenerateSecretKey ...
func GenerateSecretKey(n int) string {
	return (shortuuid.New() + shortuuid.New())[:n]
}
