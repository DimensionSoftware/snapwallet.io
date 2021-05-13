package jobmanager

import (
	"context"
	"time"

	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/models/job"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
	"github.com/khoerling/flux/api/lib/integrations/wyremanager"
)

type IJobPublisher interface {
	PublishJob(context.Context, *job.Job) error
}

type Manager struct {
	*db.Db
	Pusher       *pusher.Manager
	WyreManager  *wyremanager.Manager
	JobPublisher IJobPublisher
}

func (m Manager) GetDb() *db.Db {
	return m.Db
}

func (m Manager) GetPusher() *pusher.Manager {
	return m.Pusher
}

func (m Manager) GetWyreManager() *wyremanager.Manager {
	return m.WyreManager
}

func (m Manager) GetJobPublisher() IJobPublisher {
	return m.JobPublisher
}

func (m Manager) MarkJobDone(ctx context.Context, j *job.Job) error {
	j.Status = job.StatusDone
	j.UpdatedAt = time.Now().Unix()

	return m.Db.SaveJob(ctx, nil, j)
}
