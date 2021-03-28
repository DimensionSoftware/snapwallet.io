package remedymanager

import (
	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	proto "github.com/khoerling/flux/api/lib/protocol"
)

// manages user profile remedies
type Manager struct {
	Db *db.Db
}

func (m Manager) GetRemediationsProto(userID user.ID, profile profiledata.ProfileDatas) ([]*proto.ProfileDataItemRemediation, error) {
	var out []*proto.ProfileDataItemRemediation

	for _, remediation := range BuildSubmitRemediationsForMissing(profile) {
		out = append(out, remediation)
	}

	return out, nil
}

func BuildSubmitRemediationsForMissing(profile profiledata.ProfileDatas) []*proto.ProfileDataItemRemediation {
	var out []*proto.ProfileDataItemRemediation

	for _, kind := range common.ProfileDataRequiredForWyre {
		if len(profile.FilterKind(kind)) == 0 {
			out = append(out, &proto.ProfileDataItemRemediation{
				Kind: kind.ToProfileDataItemKind(),
				Note: "Please submit.",
			})
		}
	}

	return out
}
