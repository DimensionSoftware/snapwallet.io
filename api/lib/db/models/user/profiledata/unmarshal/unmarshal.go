package unmarshal

import (
	"encoding/json"
	"fmt"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/address"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/dateofbirth"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/email"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/legalname"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/phone"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/proofofaddress"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/ssn"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/usgovernmentid"
	"github.com/khoerling/flux/api/lib/encryption"
)

// Unmarshal ...
func Unmarshal(pdata *common.EncryptedProfileData, clear []byte, note string) (profiledata.ProfileData, error) {
	switch pdata.Kind {
	case common.KindLegalName:
		return &legalname.ProfileDataLegalName{
			CommonProfileData: common.CommonProfileData{
				ID:        pdata.ID,
				Status:    pdata.Status,
				CreatedAt: pdata.CreatedAt,
				UpdatedAt: pdata.UpdatedAt,
				SealedAt:  pdata.SealedAt,
				Note:      note,
			},
			LegalName: string(clear),
		}, nil
	case common.KindEmail:
		return &email.ProfileDataEmail{
			CommonProfileData: common.CommonProfileData{
				ID:        pdata.ID,
				Status:    pdata.Status,
				CreatedAt: pdata.CreatedAt,
				UpdatedAt: pdata.UpdatedAt,
				SealedAt:  pdata.SealedAt,
				Note:      note,
			},
			Email: string(clear),
		}, nil
	case common.KindPhone:
		return &phone.ProfileDataPhone{
			CommonProfileData: common.CommonProfileData{
				ID:        pdata.ID,
				Status:    pdata.Status,
				CreatedAt: pdata.CreatedAt,
				UpdatedAt: pdata.UpdatedAt,
				SealedAt:  pdata.SealedAt,
				Note:      note,
			},
			Phone: string(clear),
		}, nil
	case common.KindDateOfBirth:
		return &dateofbirth.ProfileDataDateOfBirth{
			CommonProfileData: common.CommonProfileData{
				ID:        pdata.ID,
				Status:    pdata.Status,
				CreatedAt: pdata.CreatedAt,
				UpdatedAt: pdata.UpdatedAt,
				SealedAt:  pdata.SealedAt,
				Note:      note,
			},
			DateOfBirth: string(clear),
		}, nil
	case common.KindUSSSN:
		return &ssn.ProfileDataSSN{
			CommonProfileData: common.CommonProfileData{
				ID:        pdata.ID,
				Status:    pdata.Status,
				CreatedAt: pdata.CreatedAt,
				UpdatedAt: pdata.UpdatedAt,
				SealedAt:  pdata.SealedAt,
				Note:      note,
			},
			SSN: string(clear),
		}, nil
	case common.KindAddress:
		var out address.ProfileDataAddressPIIData

		err := json.Unmarshal(clear, &out)
		if err != nil {
			return nil, err
		}

		return &address.ProfileDataAddress{
			CommonProfileData: common.CommonProfileData{
				ID:        pdata.ID,
				Status:    pdata.Status,
				CreatedAt: pdata.CreatedAt,
				UpdatedAt: pdata.UpdatedAt,
				SealedAt:  pdata.SealedAt,
				Note:      note,
			},
			Street1:    out.Street1,
			Street2:    out.Street2,
			City:       out.City,
			State:      out.State,
			PostalCode: out.PostalCode,
			Country:    out.Country,
		}, nil
	case common.KindProofOfAddressDoc:
		return &proofofaddress.ProfileDataProofOfAddressDoc{
			CommonProfileData: common.CommonProfileData{
				ID:        pdata.ID,
				Status:    pdata.Status,
				CreatedAt: pdata.CreatedAt,
				UpdatedAt: pdata.UpdatedAt,
				SealedAt:  pdata.SealedAt,
				Note:      note,
			},
			FileIDs: *pdata.FileIDs,
		}, nil
	case common.KindUSGovernmentIDDoc:
		return &usgovernmentid.ProfileDataUSGovernmentIDDoc{
			CommonProfileData: common.CommonProfileData{
				ID:        pdata.ID,
				Status:    pdata.Status,
				CreatedAt: pdata.CreatedAt,
				UpdatedAt: pdata.UpdatedAt,
				SealedAt:  pdata.SealedAt,
				Note:      note,
			},
			GovernmentIDKind: usgovernmentid.Kind(*pdata.SubKind),
			FileIDs:          *pdata.FileIDs,
		}, nil
	}

	return nil, fmt.Errorf("ProfileDataKind: %s is not implemented yet", pdata.Kind)
}

// DecryptAndUnmarshal ...
func DecryptAndUnmarshal(m *encryption.Manager, userID user.ID, data common.EncryptedProfileData) (*profiledata.ProfileData, error) {
	maybeData, maybeNote, err := data.Decrypt(m, userID)
	if err != nil {
		return nil, err
	}

	var (
		clear []byte
		note  string
	)
	if maybeData != nil {
		clear = *maybeData
	}
	if maybeNote != nil {
		note = string(*maybeNote)
	}

	out, err := Unmarshal(&data, clear, note)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
