package ijobmanager

import (
	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	"github.com/khoerling/flux/api/lib/interfaces/ijobpublisher"
)

type JobManager interface {
	GetDb() *db.Db
	GetPusher() *pusher.Manager
	GetWyreManager() *wyre.Manager
	GetJobPublisher() ijobpublisher.JobPublisher
}
