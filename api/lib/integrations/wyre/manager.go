package wyre

import (
	"fmt"

	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
)

type Manager struct {
	Wyre *Client
}

func (m Manager) CreateAccount(profile profiledata.ProfileDatas) error {
	t := true
	f := false

	if !profile.HasWyreAccountPreconditionsMet() {
		return fmt.Errorf("Profile data is not complete enough to submit to Wyre")
	}

	_, err := m.Wyre.CreateAccount(CreateAccountRequest{
		SubAccount:   &f,
		DisableEmail: &t,
		ProfileFields: []ProfileField{
			{
				FieldID: ProfileFieldIDIndividualLegalName,
				Value:   "",
			},
			{
				FieldID: ProfileFieldIDIndividualCellphoneNumber,
				Value:   "",
			},
			{
				FieldID: ProfileFieldIDIndividualEmail,
				Value:   "",
			},
			{
				FieldID: ProfileFieldIDIndividualDateOfBirth,
				Value:   "",
			},
			{
				FieldID: ProfileFieldIDIndividualSSN,
				Value:   "",
			},
			{
				FieldID: ProfileFieldIDIndividualResidenceAddress,
				Value:   "",
			},
		},
	}.WithDefaults())
	if err != nil {
		return err
	}

	return nil
}
