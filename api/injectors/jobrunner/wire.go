package wire

import (
	"github.com/google/wire"
	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/filemanager"
	"github.com/khoerling/flux/api/lib/integrations/cloudstorage"
	"github.com/khoerling/flux/api/lib/integrations/firestore"
	"github.com/khoerling/flux/api/lib/integrations/plaid"
	"github.com/khoerling/flux/api/lib/integrations/pubsub"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	"github.com/khoerling/flux/api/lib/jobmanager"
	vendorplaid "github.com/plaid/plaid-go/plaid"
)

// wire.go

// InitializeServer creates the main server container
func InitializeJobManager() (jobmanager.Manager, error) {
	wire.Build(
		wire.Struct(new(db.Db), "*"),
		wire.Struct(new(pusher.Manager), "*"),
		wire.Struct(new(pubsub.Manager), "*"),
		wire.Struct(new(jobmanager.Manager), "*"),
		wire.Struct(new(filemanager.Manager), "*"),
		wire.Struct(new(wyre.Manager), "*"),
		cloudstorage.ProvideBucket,
		vendorplaid.NewClient,
		plaid.ProvideClientOptions,
		firestore.ProvideFirestoreProjectID,
		firestore.ProvideFirestore,
		encryption.ProvideConfig,
		encryption.NewManager,
		pusher.ProviderPusherConfig,
		pusher.ProvidePusherClient,
		wyre.ProvideAPIHost,
		wyre.ProvideWyreConfig,
		wyre.NewClient,
		pubsub.ProvideClient,
	)
	return jobmanager.Manager{}, nil
}
