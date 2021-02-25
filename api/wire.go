package wire

import (
	"github.com/google/wire"
	"github.com/khoerling/flux/api/lib/auth"
	"github.com/khoerling/flux/api/lib/integrations/firestore"
	"github.com/khoerling/flux/api/lib/integrations/plaid"
	"github.com/khoerling/flux/api/lib/integrations/sendgrid"
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
		firestore.ProvideFirestoreProjectID,
		firestore.ProvideFirestore,
		wyre.NewClient,
		plaid.ProvideClientOptions,
		vendorplaid.NewClient,
		auth.ProvideJwtPrivateKey,
		wire.Struct(new(auth.JwtSigner), "*"),
	)
	return server.Server{}, nil
}
