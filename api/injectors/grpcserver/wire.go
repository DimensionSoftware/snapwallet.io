package wire

import (
	"github.com/google/wire"
	"github.com/khoerling/flux/api/lib/auth"
	"github.com/khoerling/flux/api/lib/config"
	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/firebase_db"
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/filemanager"
	"github.com/khoerling/flux/api/lib/integrations/cloudstorage"
	"github.com/khoerling/flux/api/lib/integrations/firestore"
	"github.com/khoerling/flux/api/lib/integrations/plaid"
	"github.com/khoerling/flux/api/lib/integrations/pubsub"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
	"github.com/khoerling/flux/api/lib/integrations/sendemail"
	"github.com/khoerling/flux/api/lib/integrations/sendemail/sendgrid"
	"github.com/khoerling/flux/api/lib/integrations/twilio"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	"github.com/khoerling/flux/api/lib/integrations/wyremanager"
	"github.com/khoerling/flux/api/lib/jobmanager"
	"github.com/khoerling/flux/api/lib/jobpublisher"
	"github.com/khoerling/flux/api/lib/remedymanager"
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
		wire.Bind(new(jobmanager.IJobPublisher), new(jobpublisher.PubSubPublisher)),
		wire.Struct(new(jobpublisher.PubSubPublisher), "*"),
		wire.Struct(new(remedymanager.Manager), "*"),
		wire.Bind(new(sendemail.SendEmail), new(sendgrid.Client)),
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
		wire.Bind(new(db.Db), new(firebase_db.Db)),
		wire.Struct(new(firebase_db.Db), "*"),
		config.ProvideAPIHost,
		config.ProvideWebHost,
		wire.Struct(new(wyremanager.Manager), "*"),
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
		wire.Bind(new(jobmanager.IJobPublisher), new(jobpublisher.InProcessPublisher)),
		wire.Struct(new(jobpublisher.InProcessPublisher), "*"),
		wire.Struct(new(remedymanager.Manager), "*"),
		wire.Bind(new(sendemail.SendEmail), new(sendgrid.Client)),
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
		wire.Bind(new(db.Db), new(firebase_db.Db)),
		wire.Struct(new(firebase_db.Db), "*"),
		config.ProvideAPIHost,
		config.ProvideWebHost,
		wire.Struct(new(wyremanager.Manager), "*"),
		wire.Struct(new(pusher.Manager), "*"),
		pusher.ProviderPusherConfig,
		pusher.ProvidePusherClient,
	)
	return server.Server{}, nil
}
