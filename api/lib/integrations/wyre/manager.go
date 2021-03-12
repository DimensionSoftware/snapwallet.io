package wyre

import (
	"fmt"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
)

type Manager struct {
	Wyre *Client
}

func (m Manager) CreateAccount(u *user.User, profile profiledata.ProfileDatas) error {
	t := true
	f := false

	if !profile.HasWyreAccountPreconditionsMet() {
		return fmt.Errorf("Profile data is not complete enough to submit to Wyre")
	}

	address := profile.FilterKindAddress()[0]

	_, err := m.Wyre.CreateAccount(CreateAccountRequest{
		SubAccount:   &f,
		DisableEmail: &t,
		ProfileFields: []ProfileField{
			{
				FieldID: ProfileFieldIDIndividualLegalName,
				Value:   profile.FilterKindLegalName()[0].LegalName,
			},
			{
				FieldID: ProfileFieldIDIndividualCellphoneNumber,
				Value:   u.Phone,
			},
			{
				FieldID: ProfileFieldIDIndividualEmail,
				Value:   u.Email,
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
		return err
	}
	// TODO: update statuses of profile data once submitted

	return nil
}
