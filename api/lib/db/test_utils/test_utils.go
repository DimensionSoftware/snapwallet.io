package test_utils

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/khoerling/flux/api/lib/db/models/job"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/lithammer/shortuuid"
)

func GenFakeUser() *user.User {
	email := faker.Email()
	phone := faker.E164PhoneNumber()

	return &user.User{
		ID:    user.ID(shortuuid.New()),
		Email: &email,
		Phone: &phone,
	}
}

func GenFakeJob() *job.Job {
	now := time.Now()

	return &job.Job{
		ID:         shortuuid.New(),
		Kind:       job.KindUpdateWyreAccountForUser,
		Status:     job.StatusQueued,
		RelatedIDs: []string{"1", "2", "3"},
		CreatedAt:  now.Unix(),
		UpdatedAt:  now.Unix(),
	}
}
