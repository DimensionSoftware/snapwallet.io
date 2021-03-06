package profiledata

import (
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/encryption"
)

// ProfileData interface which profile datas need to implement in order to save
type ProfileData interface {
	Encrypt(m *encryption.Manager, userID user.ID) (*common.EncryptedProfileData, error)
}
