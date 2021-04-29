package jobpublisher

import (
	"context"

	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/models/job"
	"github.com/khoerling/flux/api/lib/integrations/pubsub"
)

type PubSubPublisher struct {
	Db     *db.Db
	PubSub *pubsub.Manager
}

func (pub PubSubPublisher) PublishJob(ctx context.Context, j *job.Job) error {
	err := pub.Db.SaveJob(ctx, nil, j)
	if err != nil {
		return err
	}

	return pub.PubSub.SendJob(ctx, j)
}
