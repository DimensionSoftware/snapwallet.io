package profiledata

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
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
