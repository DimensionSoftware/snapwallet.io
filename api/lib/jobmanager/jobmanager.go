package jobmanager

import (
	"context"

	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/models/job"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
)

type InnerJobPublisher interface {
	PublishJob(context.Context, *job.Job) error
}

type Manager struct {
	db.Db
	Pusher       *pusher.Manager
	WyreManager  *wyre.Manager
	JobPublisher InnerJobPublisher
}

func (m Manager) GetDb() db.Db {
	return m.Db
}

func (m Manager) GetPusher() *pusher.Manager {
	return m.Pusher
}

func (m Manager) GetWyreManager() *wyre.Manager {
	return m.WyreManager
}

func (m Manager) GetJobPublisher() InnerJobPublisher {
	return m.JobPublisher
}
