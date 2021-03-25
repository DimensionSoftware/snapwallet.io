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
