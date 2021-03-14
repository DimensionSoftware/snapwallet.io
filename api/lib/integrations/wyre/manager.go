package wyre

import (
	"context"
	"fmt"
	"time"

	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/db/models/user/wyre/account"
	"github.com/lithammer/shortuuid/v3"
)

type Manager struct {
	Wyre *Client
	Db   *db.Db
}

func (m Manager) CreateAccount(ctx context.Context, userID user.ID, profile profiledata.ProfileDatas) (*account.Account, error) {
	now := time.Now()
	t := true
	f := false

	if !profile.HasWyreAccountPreconditionsMet() {
		return nil, fmt.Errorf("Profile data is not complete enough to submit to Wyre (preconditions are unmet)")
	}

	if len(profile) != len(common.ProfileDataRequiredForWyre) {
		return nil, fmt.Errorf(
			"Number of profile data items necessary for wyre is supposed to be %d but received %d",
			len(common.ProfileDataRequiredForWyre),
			len(profile),
		)
	}

	address := profile.FilterKindAddress()[0]

	secretKey := GenerateSecretKey(35)
	wyreAuthTokenResp, err := m.Wyre.SubmitAuthToken(secretKey)
	if err != nil {
		return nil, err
	}
	fmt.Printf("wyreAuthTokenResp: %#v", wyreAuthTokenResp)

	wyreAccountResp, err := m.Wyre.CreateAccount(wyreAuthTokenResp.APIKey, CreateAccountRequest{
		SubAccount:   &f,
		DisableEmail: &t,
		ProfileFields: []ProfileField{
			{
				FieldID: ProfileFieldIDIndividualLegalName,
				Value:   profile.FilterKindLegalName()[0].LegalName,
			},
			{
				FieldID: ProfileFieldIDIndividualCellphoneNumber,
				Value:   profile.FilterKindPhone()[0].Phone,
			},
			{
				FieldID: ProfileFieldIDIndividualEmail,
				Value:   profile.FilterKindEmail()[0].Email,
			},
			{
				FieldID: ProfileFieldIDIndividualDateOfBirth,
				Value:   profile.FilterKindDateOfBirth()[0].DateOfBirth,
			},
			{
				FieldID: ProfileFieldIDIndividualSSN,
				Value:   profile.FilterKindSSN()[0].SSN,
			},
			{
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
		},
	}.WithDefaults())
	if err != nil {
		return nil, err
	}

	modifiedProfile := profile.SetStatuses(common.StatusPending)

	// todo, can't create account if they already have one

	account := account.Account{
		ID:        account.ID(wyreAccountResp.ID),
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

	return &account, nil
}

// GenerateSecretKey ...
func GenerateSecretKey(n int) string {
	return (shortuuid.New() + shortuuid.New())[:n]
}
