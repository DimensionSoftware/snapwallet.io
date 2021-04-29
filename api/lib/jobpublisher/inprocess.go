package jobpublisher

import (
	"context"
	"log"

	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/models/job"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	"github.com/khoerling/flux/api/lib/jobmanager"
	"github.com/khoerling/flux/api/lib/jobs"
)

type InProcessPublisher struct {
	*db.Db
	Pusher      *pusher.Manager
	WyreManager *wyre.Manager
}

func (pub InProcessPublisher) PublishJob(ctx context.Context, j *job.Job) error {
	err := pub.Db.SaveJob(ctx, nil, j)
	if err != nil {
		return err
	}

	go func() {
		log.Printf("Job started locally: %#v\n", j)

		m := jobmanager.Manager{
			Db:           pub.Db,
			Pusher:       pub.Pusher,
			WyreManager:  pub.WyreManager,
			JobPublisher: pub,
		}

		err := jobs.RunSnapJob(context.Background(), m, j)
		if err != nil {
			log.Printf("JOB FAILURE: %#v\n", err)
		}
	}()

	return nil
}

func (pub InProcessPublisher) MarkJobDone(ctx context.Context, j *job.Job) error {
	j.Status = job.StatusDone

	return pub.Db.SaveJob(ctx, nil, j)
}

/*
	MarkJobDone(context.Context, *job.Job) error
*/
