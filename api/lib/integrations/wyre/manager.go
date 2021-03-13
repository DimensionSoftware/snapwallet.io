package wyre

import (
	"fmt"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
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

	if len(profile) != len(common.ProfileDataRequiredForWyre) {
		return fmt.Errorf(
			"Number of profile data items necessary for wyre is supposed to be %d but received %d",
			len(common.ProfileDataRequiredForWyre),
			len(profile),
		)
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
	// TODO: upload docs too, all passed profile items must be uploaded

	return nil
}
