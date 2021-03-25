package jobpublisher

import (
	"context"
	"log"

	"github.com/khoerling/flux/api/lib/db/models/job"
	"github.com/khoerling/flux/api/lib/interfaces/ijobmanager"
	"github.com/khoerling/flux/api/lib/jobs"
)

type InProcessPublisher struct {
	JobManager ijobmanager.JobManager
}

func (pub InProcessPublisher) PublishJob(_ context.Context, j *job.Job) error {
	go func() {
		log.Printf("Job started locally: %#v\n", j)

		err := jobs.RunSnapJob(context.Background(), pub.JobManager, j)
		if err != nil {
			log.Printf("JOB FAILURE: %#v\n", err)
		}
	}()

	return nil
}
