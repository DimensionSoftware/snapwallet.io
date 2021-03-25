package wire

import (
	"github.com/google/wire"
	"github.com/khoerling/flux/api/lib/auth"
	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/filemanager"
	"github.com/khoerling/flux/api/lib/integrations/cloudstorage"
	"github.com/khoerling/flux/api/lib/integrations/firestore"
	"github.com/khoerling/flux/api/lib/integrations/plaid"
	"github.com/khoerling/flux/api/lib/integrations/pubsub"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
	"github.com/khoerling/flux/api/lib/integrations/sendgrid"
	"github.com/khoerling/flux/api/lib/integrations/twilio"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	"github.com/khoerling/flux/api/lib/interfaces/ijobpublisher"
	"github.com/khoerling/flux/api/lib/jobmanager"
	"github.com/khoerling/flux/api/lib/jobpublisher"
	"github.com/khoerling/flux/api/lib/server"
	vendorplaid "github.com/plaid/plaid-go/plaid"
)

// wire.go

// InitializeServer creates the main server container
func InitializeServer() (server.Server, error) {
	wire.Build(
		server.ProvideGrpcServer,
		wire.Struct(new(auth.Manager), "*"),
		wire.Struct(new(server.Server), "*"),
		wire.Bind(new(ijobpublisher.JobPublisher), new(jobpublisher.PubSubPublisher)),
		wire.Struct(new(jobpublisher.PubSubPublisher), "*"),
		sendgrid.ProvideSendClientAPIKey,
		sendgrid.ProvideSendClient,
		twilio.ProvideTwilioConfig,
		twilio.ProvideTwilio,
		firestore.ProvideFirestoreProjectID,
		firestore.ProvideFirestore,
		cloudstorage.ProvideBucket,
		wire.Struct(new(filemanager.Manager), "*"),
		wyre.NewClient,
		wyre.ProvideWyreConfig,
		plaid.ProvideClientOptions,
		vendorplaid.NewClient,
		auth.ProvideJwtPrivateKey,
		auth.ProvideJwtPublicKey,
		wire.Struct(new(auth.JwtSigner), "*"),
		wire.Struct(new(auth.JwtVerifier), "*"),
		encryption.ProvideConfig,
		encryption.NewManager,
		wire.Struct(new(db.Db), "*"),
		wyre.ProvideAPIHost,
		wire.Struct(new(wyre.Manager), "*"),
		wire.Struct(new(pusher.Manager), "*"),
		pusher.ProviderPusherConfig,
		pusher.ProvidePusherClient,
		wire.Struct(new(pubsub.Manager), "*"),
		pubsub.ProvideClient,
	)
	return server.Server{}, nil
}

func InitializeDevServer() (server.Server, error) {
	wire.Build(
		server.ProvideGrpcServer,
		wire.Struct(new(auth.Manager), "*"),
		wire.Struct(new(server.Server), "*"),
		wire.Bind(new(ijobpublisher.JobPublisher), new(jobpublisher.InProcessPublisher)),
		wire.Struct(new(jobpublisher.InProcessPublisher), "*"),
		wire.Struct(new(jobmanager.Manager), "*"),
		sendgrid.ProvideSendClientAPIKey,
		sendgrid.ProvideSendClient,
		twilio.ProvideTwilioConfig,
		twilio.ProvideTwilio,
		firestore.ProvideFirestoreProjectID,
		firestore.ProvideFirestore,
		cloudstorage.ProvideBucket,
		wire.Struct(new(filemanager.Manager), "*"),
		wyre.NewClient,
		wyre.ProvideWyreConfig,
		plaid.ProvideClientOptions,
		vendorplaid.NewClient,
		auth.ProvideJwtPrivateKey,
		auth.ProvideJwtPublicKey,
		wire.Struct(new(auth.JwtSigner), "*"),
		wire.Struct(new(auth.JwtVerifier), "*"),
		encryption.ProvideConfig,
		encryption.NewManager,
		wire.Struct(new(db.Db), "*"),
		wyre.ProvideAPIHost,
		wire.Struct(new(wyre.Manager), "*"),
		wire.Struct(new(pusher.Manager), "*"),
		pusher.ProviderPusherConfig,
		pusher.ProvidePusherClient,
		wire.Struct(new(pubsub.Manager), "*"),
		pubsub.ProvideClient,
	)
	return server.Server{}, nil
}
