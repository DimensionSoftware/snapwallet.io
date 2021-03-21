package pubsub

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/khoerling/flux/api/lib/db/models/job"
)

type Manager struct {
	PubSub *pubsub.Client
}

func (m Manager) SendJob(ctx context.Context, j *job.Job) error {
	topic := m.PubSub.Topic("snap-jobs")

	data, err := json.Marshal(j)
	if err != nil {
		return err
	}

	resp := topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})
	_, err = resp.Get(ctx)
	if err != nil {
		return err
	}

	log.Printf("Job submitted: %#v\n", j)

	return nil
}
