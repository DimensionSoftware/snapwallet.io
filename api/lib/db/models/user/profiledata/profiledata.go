package profiledata

import (
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
func (profile ProfileDatas) GetProfileDataInfo() *proto.ProfileDataInfo {
	var legalNameData *legalname.ProfileDataLegalName
	var dobData *dateofbirth.ProfileDataDateOfBirth
	var ssnData *ssn.ProfileDataSSN
	var addressData *address.ProfileDataAddress
	{
		existingProfileData := profile.FilterKind(common.KindLegalName).First()
		if existingProfileData != nil {
			legalNameData = (*existingProfileData).(*legalname.ProfileDataLegalName)
		}
	}
	{
		existingProfileData := profile.FilterKind(common.KindDateOfBirth).First()
		if existingProfileData != nil {
			dobData = (*existingProfileData).(*dateofbirth.ProfileDataDateOfBirth)
		}
	}
	{
		existingProfileData := profile.FilterKind(common.KindSSN).First()
		if existingProfileData != nil {
			ssnData = (*existingProfileData).(*ssn.ProfileDataSSN)
		}
	}
	{
		existingProfileData := profile.FilterKind(common.KindAddress).First()
		if existingProfileData != nil {
			addressData = (*existingProfileData).(*address.ProfileDataAddress)
		}
	}

	var legalNameInfo *proto.ProfileDataItemInfo
	if legalNameData != nil {
		legalNameInfo = legalNameData.GetProfileDataItemInfo()
	}

	var dobInfo *proto.ProfileDataItemInfo
	if dobData != nil {
		dobInfo = dobData.GetProfileDataItemInfo()
	}

	var ssnInfo *proto.ProfileDataItemInfo
	if ssnData != nil {
		ssnInfo = ssnData.GetProfileDataItemInfo()
	}

	var addressInfo *proto.ProfileDataItemInfo
	if addressData != nil {
		addressInfo = addressData.GetProfileDataItemInfo()
	}

	return &proto.ProfileDataInfo{
		LegalName:   legalNameInfo,
		DateOfBirth: dobInfo,
		Ssn:         ssnInfo,
		Address:     addressInfo,
		//Email:       email,
		//Phone:       phone,
	}

}
