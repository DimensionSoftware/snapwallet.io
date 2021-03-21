package jobmanager

import (
	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/integrations/pubsub"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
)

type Manager struct {
	*db.Db
	Pusher      *pusher.Manager
	WyreManager *wyre.Manager
	PubSub      *pubsub.Manager
}
