package jobpublisher

import (
	"context"

	"github.com/khoerling/flux/api/lib/db/models/job"
	"github.com/khoerling/flux/api/lib/integrations/pubsub"
)

type PubSubPublisher struct {
	PubSub *pubsub.Manager
}

func (pub PubSubPublisher) PublishJob(ctx context.Context, j *job.Job) error {
	return pub.PubSub.SendJob(ctx, j)
}
