package wire

import (
	"github.com/google/wire"
	"github.com/khoerling/flux/api/lib/auth"
	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/integrations/firestore"
	"github.com/khoerling/flux/api/lib/integrations/plaid"
	"github.com/khoerling/flux/api/lib/integrations/sendgrid"
	"github.com/khoerling/flux/api/lib/integrations/twilio"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	"github.com/khoerling/flux/api/lib/server"
	vendorplaid "github.com/plaid/plaid-go/plaid"
)

// wire.go

// InitializeServer creates the main server container
func InitializeServer() (server.Server, error) {
	wire.Build(
		server.ProvideServer,
		sendgrid.ProvideSendClientAPIKey,
		sendgrid.ProvideSendClient,
		twilio.ProvideTwilioConfig,
		twilio.ProvideTwilio,
		firestore.ProvideFirestoreProjectID,
		firestore.ProvideFirestore,
		wyre.NewClient,
		wyre.ProvideWyreConfig,
		plaid.ProvideClientOptions,
		vendorplaid.NewClient,
		auth.ProvideJwtPrivateKey,
		auth.ProvideJwtPublicKey,
		wire.Struct(new(auth.JwtSigner), "*"),
		wire.Struct(new(auth.JwtVerifier), "*"),
		wire.Struct(new(db.Db), "*"),
	)
	return server.Server{}, nil
}
