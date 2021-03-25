package ijobpublisher

import (
	"context"

	"github.com/khoerling/flux/api/lib/db/models/job"
)

type JobPublisher interface {
	PublishJob(context.Context, *job.Job) error
}
