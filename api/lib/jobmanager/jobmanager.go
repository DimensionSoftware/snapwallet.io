package jobmanager

import (
	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	"github.com/khoerling/flux/api/lib/interfaces/ijobpublisher"
)

type Manager struct {
	*db.Db
	Pusher       *pusher.Manager
	WyreManager  *wyre.Manager
	JobPublisher ijobpublisher.JobPublisher
}

func (m Manager) GetDb() *db.Db {
	return m.Db
}

func (m Manager) GetPusher() *pusher.Manager {
	return m.Pusher
}

func (m Manager) GetWyreManager() *wyre.Manager {
	return m.WyreManager
}

func (m Manager) GetJobPublisher() ijobpublisher.JobPublisher {
	return m.JobPublisher
}
