package profiledata

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/address"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/dateofbirth"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/legalname"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/ssn"
	"github.com/khoerling/flux/api/lib/encryption"
	proto "github.com/khoerling/flux/api/lib/protocol"
)

// ProfileData interface which profile datas need to implement in order to save
type ProfileData interface {
	Kind() common.ProfileDataKind
	GetStatus() common.ProfileDataStatus
	GetProfileDataItemInfo() *proto.ProfileDataItemInfo
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
func (profile ProfileDatas) FilterKindLegalName() []legalname.ProfileDataLegalName {
	out := []legalname.ProfileDataLegalName{}

	for _, pdata := range profile.FilterKind(common.KindLegalName) {
		out = append(out, pdata.(legalname.ProfileDataLegalName))
	}

	return out
}

// FilterKindDateOfBirth ...
func (profile ProfileDatas) FilterKindDateOfBirth() []dateofbirth.ProfileDataDateOfBirth {
	out := []dateofbirth.ProfileDataDateOfBirth{}

	for _, pdata := range profile.FilterKind(common.KindDateOfBirth) {
		out = append(out, pdata.(dateofbirth.ProfileDataDateOfBirth))
	}

	return out
}

// FilterKindSSN ...
func (profile ProfileDatas) FilterKindSSN() []ssn.ProfileDataSSN {
	out := []ssn.ProfileDataSSN{}

	for _, pdata := range profile.FilterKind(common.KindSSN) {
		out = append(out, pdata.(ssn.ProfileDataSSN))
	}

	return out
}

// FilterKindAddress ...
func (profile ProfileDatas) FilterKindAddress() []address.ProfileDataAddress {
	out := []address.ProfileDataAddress{}

	for _, pdata := range profile.FilterKind(common.KindAddress) {
		out = append(out, pdata.(address.ProfileDataAddress))
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

// First ...
func (profile ProfileDatas) First() *ProfileData {
	if len(profile) > 0 {
		return &profile[0]
	}

	return nil
}

// GetProfileDataInfo ...
func (profile ProfileDatas) GetProfileDataInfo(u *user.User) *proto.ProfileDataInfo {
	// TODO: store phone/emailInfo data on create account submission into concrete profile data with ids and timestamps -- then it becomes immutable for record purposes
	// if there is no concrete record then their user emailInfo/phone become the main mutable source of truth
	//
	// this means we also need to do lookup logic changes, check concrete records first, then account details for phone/emailInfo

	out := []*proto.ProfileDataItemInfo{}

	if (u.Email != nil && *u.Email != "" && u.EmailVerifiedAt != nil && *u.EmailVerifiedAt != time.Time{}) {
		out = append(out, &proto.ProfileDataItemInfo{
			Kind:   proto.ProfileDataItemKind_K_EMAIL,
			Status: proto.ProfileDataItemStatus_S_RECEIVED,
			Length: int32(len(*u.Email)),
		})
	}

	if (u.Phone != nil && *u.Phone != "" && u.PhoneVerifiedAt != nil && *u.PhoneVerifiedAt != time.Time{}) {
		out = append(out, &proto.ProfileDataItemInfo{
			Kind:   proto.ProfileDataItemKind_K_PHONE,
			Status: proto.ProfileDataItemStatus_S_RECEIVED,
			Length: int32(len(*u.Phone)),
		})
	}

	for _, item := range profile {
		out = append(out, item.GetProfileDataItemInfo())
	}

	return &proto.ProfileDataInfo{
		Profile: out,
	}

}
