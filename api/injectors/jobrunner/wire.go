package wire

import (
	"github.com/google/wire"
	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/integrations/firestore"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	"github.com/khoerling/flux/api/lib/jobmanager"
)

// wire.go

// InitializeServer creates the main server container
func InitializeJobManager() (jobmanager.Manager, error) {
	wire.Build(
		wire.Struct(new(db.Db), "*"),
		wire.Struct(new(pusher.Manager), "*"),
		wire.Struct(new(jobmanager.Manager), "*"),
		wire.Struct(new(wyre.Manager), "*"),
		firestore.ProvideFirestoreProjectID,
		firestore.ProvideFirestore,
		encryption.ProvideConfig,
		encryption.NewManager,
		pusher.ProviderPusherConfig,
		pusher.ProvidePusherClient,
		wyre.ProvideAPIHost,
		wyre.ProvideWyreConfig,
		wyre.NewClient,
	)
	return jobmanager.Manager{}, nil
}
