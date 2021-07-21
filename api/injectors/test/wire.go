package wire

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/wire"
	"github.com/khoerling/flux/api/lib/auth"
	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/firebase_db"
	"github.com/khoerling/flux/api/lib/db/mock_db"
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/integration_t_manager"
	"github.com/khoerling/flux/api/lib/integrations/firestore"
	"github.com/khoerling/flux/api/lib/server"
)

func InitializeTestManager() (integration_t_manager.Manager, error) {
	wire.Build(
		wire.Bind(new(db.Db), new(firebase_db.Db)),
		wire.Struct(new(firebase_db.Db), "*"),
		wire.Struct(new(integration_t_manager.Manager), "*"),
		firestore.ProvideFirestoreProjectID,
		firestore.ProvideFirestore,
		encryption.ProvideConfig,
		encryption.NewManager,
	)
	return integration_t_manager.Manager{}, nil
}

/*
func InitializeMockDBTestManager(t *testing.T) integration_t_manager.Manager {
	wire.Build(
		wire.Bind(new(db.Db), new(*mock_db.MockDb)),
		wire.Struct(new(integration_t_manager.Manager), "*"),
		mock_db.NewMockDb,
		gomock.NewController,
		wire.Bind(new(gomock.TestReporter), new(*testing.T)),
	)
	return integration_t_manager.Manager{}
}
*/

func InitializeMockDBJwtVerifier(t *testing.T) auth.JwtVerifier {
	wire.Build(
		wire.Struct(new(auth.JwtVerifier), "*"),
		wire.Bind(new(db.Db), new(*mock_db.MockDb)),
		mock_db.NewMockDb,
		gomock.NewController,
		wire.Bind(new(gomock.TestReporter), new(*testing.T)),
		auth.ProvideJwtPublicKey,
		auth.ProvideTestJwtPrivateKey,
	)
	return auth.JwtVerifier{}
}

func InitializeMockDBServer(t *testing.T) server.Server {
	wire.Build(
		wire.Struct(new(server.Server), "*"),
		wire.Struct(new(auth.JwtVerifier), "*"),
		wire.Bind(new(db.Db), new(*mock_db.MockDb)),
		mock_db.NewMockDb,
		gomock.NewController,
		wire.Bind(new(gomock.TestReporter), new(*testing.T)),
		auth.ProvideJwtPublicKey,
		auth.ProvideTestJwtPrivateKey,
	)
	return server.Server{}
}
