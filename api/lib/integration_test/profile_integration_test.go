package integration_test

import (
	"context"
	"testing"
	"time"

	"github.com/khoerling/flux/api/lib/db/models/onetimepasscode"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/common"
	"github.com/khoerling/flux/api/lib/db/models/user/profiledata/legalname"
	"github.com/khoerling/flux/api/lib/integration_t_manager/wire"
	"syreclabs.com/go/faker"

	"github.com/lithammer/shortuuid/v3"
	"github.com/stretchr/testify/assert"
)

func Test_Profile_Lifecycle(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()

	s, err := wire.InitializeTestManager()
	if err != nil {
		panic(err)
	}

	email := faker.Internet().SafeEmail()

	u, err := s.Db.GetOrCreateUser(ctx, onetimepasscode.LoginKindEmail, email)
	a.NoError(err)

	pdata := legalname.ProfileDataLegalName{
		ID:        common.ProfileDataID(shortuuid.New()),
		Status:    common.StatusReceived,
		LegalName: "Bob Jones",
		CreatedAt: time.Now(),
	}

	returnedID, err := s.Db.SaveProfileData(ctx, nil, u.ID, pdata)
	a.NoError(err)
	a.Equal(pdata.ID, returnedID)

	profile, err := s.Db.GetAllProfileData(ctx, nil, u.ID)
	a.NoError(err)
	a.Len(profile, 2)
	pdataRetrieved := (profile.FilterKind(common.KindLegalName)[0]).(*legalname.ProfileDataLegalName)

	a.Equal(pdata.ID, pdataRetrieved.ID)
	a.Equal(pdata.LegalName, pdataRetrieved.LegalName)
	a.Equal(pdata.Status, pdataRetrieved.Status)
	a.True(pdata.CreatedAt.Equal(pdataRetrieved.CreatedAt), "pdata.CreatedAt.Equal(pdataRetrieved.CreatedAt)")
}
