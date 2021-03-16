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
	"github.com/khoerling/flux/api/lib/encryption"
)

// Unmarshal ...
func Unmarshal(pdata *common.EncryptedProfileData, clear []byte) (profiledata.ProfileData, error) {
	switch pdata.Kind {
	case common.KindLegalName:
		return &legalname.ProfileDataLegalName{
			ID:        pdata.ID,
			Status:    pdata.Status,
			LegalName: string(clear),
			CreatedAt: pdata.CreatedAt,
			UpdatedAt: pdata.UpdatedAt,
			SealedAt:  pdata.SealedAt,
		}, nil
	case common.KindEmail:
		return &email.ProfileDataEmail{
			ID:        pdata.ID,
			Status:    pdata.Status,
			Email:     string(clear),
			CreatedAt: pdata.CreatedAt,
			UpdatedAt: pdata.UpdatedAt,
			SealedAt:  pdata.SealedAt,
		}, nil
	case common.KindPhone:
		return &phone.ProfileDataPhone{
			ID:        pdata.ID,
			Status:    pdata.Status,
			Phone:     string(clear),
			CreatedAt: pdata.CreatedAt,
			UpdatedAt: pdata.UpdatedAt,
			SealedAt:  pdata.SealedAt,
		}, nil
	case common.KindDateOfBirth:
		return &dateofbirth.ProfileDataDateOfBirth{
			ID:          pdata.ID,
			Status:      pdata.Status,
			DateOfBirth: string(clear),
			CreatedAt:   pdata.CreatedAt,
			UpdatedAt:   pdata.UpdatedAt,
			SealedAt:    pdata.SealedAt,
		}, nil
	case common.KindUSSSN:
		return &ssn.ProfileDataSSN{
			ID:        pdata.ID,
			Status:    pdata.Status,
			SSN:       string(clear),
			CreatedAt: pdata.CreatedAt,
			UpdatedAt: pdata.UpdatedAt,
			SealedAt:  pdata.SealedAt,
		}, nil
	case common.KindAddress:
		var out address.ProfileDataAddressPIIData

		err := json.Unmarshal(clear, &out)
		if err != nil {
			return nil, err
		}

		return &address.ProfileDataAddress{
			ID:         pdata.ID,
			Status:     pdata.Status,
			Street1:    out.Street1,
			Street2:    out.Street2,
			City:       out.City,
			State:      out.State,
			PostalCode: out.PostalCode,
			Country:    out.Country,
			CreatedAt:  pdata.CreatedAt,
			UpdatedAt:  pdata.UpdatedAt,
			SealedAt:   pdata.SealedAt,
		}, nil
	case common.KindProofOfAddressDoc:
		return &proofofaddress.ProfileDataProofOfAddressDoc{
			ID:        pdata.ID,
			Status:    pdata.Status,
			FileIDs:   *pdata.FileIDs,
			CreatedAt: pdata.CreatedAt,
			UpdatedAt: pdata.UpdatedAt,
			SealedAt:  pdata.SealedAt,
		}, nil
	}

	return nil, fmt.Errorf("ProfileDataKind: %s is not implemented yet", pdata.Kind)
}

// DecryptAndUnmarshal ...
func DecryptAndUnmarshal(m *encryption.Manager, userID user.ID, data common.EncryptedProfileData) (*profiledata.ProfileData, error) {
	maybeClear, err := data.Decrypt(m, userID)
	if err != nil {
		return nil, err
	}

	var clear []byte
	if maybeClear != nil {
		clear = *maybeClear
	}

	out, err := Unmarshal(&data, clear)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
