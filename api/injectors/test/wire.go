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
	"github.com/khoerling/flux/api/lib/integrations/sendemail"
	"github.com/khoerling/flux/api/lib/integrations/sendemail/mock_sendemail"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	"github.com/khoerling/flux/api/lib/integrations/wyre/mock_wyre"
	"github.com/khoerling/flux/api/lib/server"
	"github.com/onsi/ginkgo"
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

func InitializeMockServer(t ginkgo.GinkgoTInterface) (server.Server, error) {
	wire.Build(
		wire.Struct(new(server.Server), "Db", "SendEmail", "Wyre"),
		wire.Bind(new(db.Db), new(*mock_db.MockDb)),
		mock_db.NewMockDb,
		wire.Bind(new(sendemail.SendEmail), new(*mock_sendemail.MockSendEmail)),
		mock_sendemail.NewMockSendEmail,
		wire.Bind(new(wyre.ClientInterface), new(*mock_wyre.MockClientInterface)),
		mock_wyre.NewMockClientInterface,
		gomock.NewController,
		wire.Bind(new(gomock.TestReporter), new(ginkgo.GinkgoTInterface)),
	)
	return server.Server{}, nil
}
