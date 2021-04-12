package wire

import (
	"github.com/google/wire"
	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/firebase_db"
	"github.com/khoerling/flux/api/lib/db/mock_db"
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/integration_t_manager"
	"github.com/khoerling/flux/api/lib/integrations/firestore"
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

func InitializeMockDBTestManager() (integration_t_manager.Manager, error) {
	wire.Build(
		wire.Bind(new(db.Db), new(*mock_db.MockDb)),
		wire.Struct(new(integration_t_manager.Manager), "*"),
		firestore.ProvideFirestoreProjectID,
		firestore.ProvideFirestore,
		encryption.ProvideConfig,
		encryption.NewManager,
		mock_db.NewMockDb,
		//gomock.NewController(),
	)
	return integration_t_manager.Manager{}, nil
}
