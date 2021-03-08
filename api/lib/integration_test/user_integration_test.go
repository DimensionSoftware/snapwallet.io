package integration_test

import (
	"context"
	"testing"

	"github.com/khoerling/flux/api/lib/db/models/onetimepasscode"
	"github.com/khoerling/flux/api/lib/wire"
	"syreclabs.com/go/faker"

	"github.com/stretchr/testify/assert"
)

func Test_User_Lifecycle(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	s, err := wire.InitializeServer()
	if err != nil {
		panic(err)
	}

	email := faker.Internet().SafeEmail()
	u, err := s.Db.GetOrCreateUser(ctx, onetimepasscode.LoginKindEmail, email)
	a.NoError(err)

	u2, err := s.Db.GetOrCreateUser(ctx, onetimepasscode.LoginKindEmail, email)
	a.NoError(err)

	a.Equal(u.ID, u2.ID, "the same email should correlate with the same user id on second visit")
}
