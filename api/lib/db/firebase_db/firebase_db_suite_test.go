package firebase_db_test

import (
	"testing"

	"github.com/khoerling/flux/api/lib/integration_t_manager"
	"github.com/khoerling/flux/api/lib/integration_t_manager/wire"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var testManager integration_t_manager.Manager

func TestFirebaseDb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FirebaseDb Suite")
}

var _ = BeforeSuite(func() {
	m, err := wire.InitializeTestManager()
	if err != nil {
		panic(err)
	}
	testManager = m
})
