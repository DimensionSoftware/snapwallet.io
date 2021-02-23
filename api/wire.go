package wire

import (
	"github.com/google/wire"
	"github.com/khoerling/flux/api/lib/integrations/firestore"
	"github.com/khoerling/flux/api/lib/integrations/sendgrid"
	"github.com/khoerling/flux/api/lib/server"
)

// wire.go

// InitializeServer creates the main server container
func InitializeServer() server.Server {
	wire.Build(
		server.ProvideServer,
		sendgrid.ProvideSendClientAPIKey,
		sendgrid.ProvideSendClient,
		firestore.ProvideFirestoreProjectID,
		firestore.ProvideFirestore,
	)
	return server.Server{}
}
