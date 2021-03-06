package unmarshal

import (
	"encoding/json"

	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/address"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
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
	return nil, nil
}
