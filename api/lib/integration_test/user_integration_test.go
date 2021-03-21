package integration_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/khoerling/flux/api/lib/db/models/onetimepasscode"
	"github.com/khoerling/flux/api/lib/integration_t_manager/wire"
	"syreclabs.com/go/faker"

	"github.com/stretchr/testify/assert"
)

func Test_User_Lifecycle(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	s, err := wire.InitializeTestManager()
	if err != nil {
		panic(err)
	}

	email := faker.Internet().SafeEmail()

	u, err := s.Db.GetOrCreateUser(ctx, onetimepasscode.LoginKindEmail, email)
	a.NoError(err)

	u2, err := s.Db.GetOrCreateUser(ctx, onetimepasscode.LoginKindEmail, email)
	a.NoError(err)
	a.Equal(u.ID, u2.ID, "the same email should correlate with the same user id on second visit")
	a.Equal(u.Email, u2.Email)
	log.Println(u.CreatedAt, u2.CreatedAt)
	a.WithinDuration(u.CreatedAt, u2.CreatedAt, time.Second)

	u3, err := s.Db.GetUserByID(ctx, u.ID)
	a.NoError(err, "fetch user by id needs to work and should return the same results as an email lookup")
	a.Equal(u.ID, u3.ID)
	a.Equal(u.Email, u3.Email)
	a.WithinDuration(u.CreatedAt, u3.CreatedAt, time.Second)
}
