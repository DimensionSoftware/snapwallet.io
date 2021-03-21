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

func (m Manager) CreatePaymentMethod(ctx context.Context, userID user.ID, wyreAccountID account.ID, plaidAccessToken string, plaidAccountID string) (*paymentmethod.PaymentMethod, error) {
	resp, err := m.Plaid.CreateProcessorToken(plaidAccessToken, plaidAccountID, "wyre")
	if err != nil {
		return nil, err
	}

	wyrePm, err := m.Wyre.CreatePaymentMethod(CreatePaymentMethodRequest{
		PlaidProcessorToken: resp.ProcessorToken,
	}.WithDefaults())
	if err != nil {
		return nil, err
	}

	pm := paymentmethod.PaymentMethod{
		ID:                    paymentmethod.ID(wyrePm.ID),
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
		pm, err := m.CreatePaymentMethod(ctx, userID, wyreAccountID, item.AccessToken, accountID)
		if err != nil {
			return nil, err
		}

		out = append(out, pm)
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

func (m Manager) CreateAccount(ctx context.Context, userID user.ID, profile profiledata.ProfileDatas) (*wyre_model.Account, error) {
	now := time.Now()
	t := true
	f := false

	if !profile.HasWyreAccountPreconditionsMet() {
		return nil, fmt.Errorf("Profile data is not complete enough to submit to Wyre (preconditions are unmet)")
	}

	/*
		if len(profile) != len(common.ProfileDataRequiredForWyre) {
			return nil, fmt.Errorf(
				"Number of profile data items necessary for wyre is supposed to be %d but received %d",
				len(common.ProfileDataRequiredForWyre),
				len(profile),
			)
		}
	*/

	var fields []*ProfileField

	if addrs := profile.FilterKindAddress(); len(addrs) > 0 {
		address := addrs[0]
		fields = append(fields,
			&ProfileField{
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

	secretKey := GenerateSecretKey(35)
	wyreAuthTokenResp, err := m.Wyre.SubmitAuthToken(secretKey)
	if err != nil {
		return nil, err
	}
	fmt.Printf("wyreAuthTokenResp: %#v", wyreAuthTokenResp)

	wyreAccountResp, err := m.Wyre.CreateAccount(secretKey, CreateAccountRequest{
		SubAccount:        &f,
		DisableEmail:      &t,
		ReferrerAccountID: &m.Wyre.config.WyreAccountID,
		ProfileFields: []ProfileField{
			{
				FieldID: ProfileFieldIDIndividualLegalName,
				Value:   profile.FilterKindLegalName()[0].LegalName,
			},
			/*
				{
					FieldID: ProfileFieldIDIndividualCellphoneNumber,
					Value:   profile.FilterKindPhone()[0].Phone,
				},
				{
					FieldID: ProfileFieldIDIndividualEmail,
					Value:   profile.FilterKindEmail()[0].Email,
				},
			*/
			{
				FieldID: ProfileFieldIDIndividualDateOfBirth,
				Value:   profile.FilterKindDateOfBirth()[0].DateOfBirth,
			},
			{
				FieldID: ProfileFieldIDIndividualSSN,
				Value:   profile.FilterKindSSN()[0].SSN,
			},
		},
	}.WithDefaults())
	if err != nil {
		return nil, err
	}

	modifiedProfile := profile.SetStatuses(common.StatusPending)

	// todo, can't create account if they already have one

	account := wyre_model.Account{
		ID:        wyre_model.ID(wyreAccountResp.ID),
		APIKey:    wyreAuthTokenResp.APIKey,
		SecretKey: secretKey,
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

	hookResponse, err := m.Wyre.SubscribeWebhook(secretKey, "account:"+string(account.ID), string(m.APIHost)+"/wyre/hooks/"+string(userID))
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
