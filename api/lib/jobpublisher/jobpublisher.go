package jobpublisher

import (
	"context"

	"github.com/khoerling/flux/api/lib/db/models/job"
)

type Publisher interface {
	PublishJob(context.Context, *job.Job) error
}
