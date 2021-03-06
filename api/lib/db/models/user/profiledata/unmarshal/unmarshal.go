package unmarshal

import (
	"encoding/json"
	"fmt"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/address"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/encryption"
)

// Unmarshal ...
func Unmarshal(kind common.ProfileDataKind, data []byte) (interface{}, error) {
	switch kind {
	case common.KindAddress:
		var out address.ProfileDataAddress
		err := json.Unmarshal(data, &out)
		if err != nil {
			return nil, err
		}

		return out, nil
	}

	return nil, fmt.Errorf("ProfileDataKind: %s is not implemented yet", kind)
}

// DecryptAndUnmarshal ...
func DecryptAndUnmarshal(m *encryption.Manager, userID user.ID, data common.EncryptedProfileData) (*interface{}, error) {
	raw, err := data.Decrypt(m, userID)
	if err != nil {
		return nil, err
	}

	out, err := Unmarshal(data.Kind, raw)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
