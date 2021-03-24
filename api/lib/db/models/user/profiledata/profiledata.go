package profiledata

import (
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/address"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/dateofbirth"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/email"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/legalname"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/phone"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/ssn"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/usgovernmentid"
	"github.com/khoerling/flux/api/lib/encryption"
	proto "github.com/khoerling/flux/api/lib/protocol"
)

// ProfileData interface which profile datas need to implement in order to save
type ProfileData interface {
	Kind() common.ProfileDataKind
	GetStatus() common.ProfileDataStatus
	GetProfileDataItemInfo() *proto.ProfileDataItemInfo
	SetStatus(common.ProfileDataStatus)
	Encrypt(m *encryption.Manager, userID user.ID) (*common.EncryptedProfileData, error)
}

// ProfileDatas all the profile data for a user (array)
type ProfileDatas []ProfileData

// WyreAccountPreconditionsMet ...
func (profile ProfileDatas) HasWyreAccountPreconditionsMet() bool {
	for _, kind := range common.ProfileDataRequiredForWyre {
		if profile.FilterKind(kind).First() == nil {
			return false
		}
	}
	return true
}

// FilterKind ...
func (profile ProfileDatas) FilterKind(kind common.ProfileDataKind) ProfileDatas {
	out := []ProfileData{}

	for _, pdata := range profile {
		if pdata.Kind() == kind {
			out = append(out, pdata)
		}
	}

	return out
}

// FilterKindLegalName ...
func (profile ProfileDatas) FilterKindLegalName() []*legalname.ProfileDataLegalName {
	out := []*legalname.ProfileDataLegalName{}

	for _, pdata := range profile.FilterKind(common.KindLegalName) {
		out = append(out, pdata.(*legalname.ProfileDataLegalName))
	}

	return out
}

// FilterKindDateOfBirth ...
func (profile ProfileDatas) FilterKindDateOfBirth() []*dateofbirth.ProfileDataDateOfBirth {
	out := []*dateofbirth.ProfileDataDateOfBirth{}

	for _, pdata := range profile.FilterKind(common.KindDateOfBirth) {
		out = append(out, pdata.(*dateofbirth.ProfileDataDateOfBirth))
	}

	return out
}

// FilterKindSSN ...
func (profile ProfileDatas) FilterKindSSN() []*ssn.ProfileDataSSN {
	out := []*ssn.ProfileDataSSN{}

	for _, pdata := range profile.FilterKind(common.KindUSSSN) {
		out = append(out, pdata.(*ssn.ProfileDataSSN))
	}

	return out
}

// FilterKindAddress ...
func (profile ProfileDatas) FilterKindAddress() []*address.ProfileDataAddress {
	out := []*address.ProfileDataAddress{}

	for _, pdata := range profile.FilterKind(common.KindAddress) {
		out = append(out, pdata.(*address.ProfileDataAddress))
	}

	return out
}

// FilterKindEmail ...
func (profile ProfileDatas) FilterKindEmail() email.ProfileDataEmails {
	out := []*email.ProfileDataEmail{}

	for _, pdata := range profile.FilterKind(common.KindEmail) {
		out = append(out, pdata.(*email.ProfileDataEmail))
	}

	return out
}

// FilterKindEmail ...
func (profile ProfileDatas) FilterKindPhone() phone.ProfileDataPhones {
	out := []*phone.ProfileDataPhone{}

	for _, pdata := range profile.FilterKind(common.KindPhone) {
		out = append(out, pdata.(*phone.ProfileDataPhone))
	}

	return out
}

// FilterKindEmail ...
func (profile ProfileDatas) FilterKindUSGovernmentIDDoc() []*usgovernmentid.ProfileDataUSGovernmentIDDoc {
	out := []*usgovernmentid.ProfileDataUSGovernmentIDDoc{}

	for _, pdata := range profile.FilterKind(common.KindUSGovernmentIDDoc) {
		out = append(out, pdata.(*usgovernmentid.ProfileDataUSGovernmentIDDoc))
	}

	return out
}

// FilterStatus ...
func (profile ProfileDatas) FilterStatus(status common.ProfileDataStatus) ProfileDatas {
	out := []ProfileData{}

	for _, pdata := range profile {
		if pdata.GetStatus() == status {
			out = append(out, pdata)
		}
	}

	return out
}

// SetStatuses ...
func (profile ProfileDatas) SetStatuses(newStatus common.ProfileDataStatus) ProfileDatas {
	out := []ProfileData{}

	for _, pdata := range profile {
		pdata.SetStatus(newStatus)
		out = append(out, pdata)
	}

	return out
}

// First ...
func (profile ProfileDatas) First() *ProfileData {
	if len(profile) > 0 {
		return &profile[0]
	}

	return nil
}

// GetProfileDataItemInfo ...
func (profile ProfileDatas) GetProfileDataItemInfo() []*proto.ProfileDataItemInfo {
	out := []*proto.ProfileDataItemInfo{}

	for _, item := range profile {
		out = append(out, item.GetProfileDataItemInfo())
	}

	return out
}
