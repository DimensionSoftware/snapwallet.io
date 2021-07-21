package server_test

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/khoerling/flux/api/lib/db/mock_db"
	"github.com/khoerling/flux/api/lib/db/test_utils"
	"github.com/khoerling/flux/api/lib/integration_t_manager/wire"
	"github.com/khoerling/flux/api/lib/protocol"
	"github.com/khoerling/flux/api/lib/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/metadata"
)

var _ = Describe("Handlers", func() {
	var (
		ctx context.Context
		s   *server.Server
	)
	BeforeEach(func() {
		var err error
		ctx = context.Background()
		s, err = wire.InitializeMockServer(GinkgoT())
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("WyreCreateTransfer", func() {
		It("happy path", func() {
			Skip("WIP")

			mockDb := s.Db.(*mock_db.MockDb)
			user := test_utils.GenFakeUser()
			ctx := metadata.NewIncomingContext(ctx, map[string][]string{"user-id": {string(user.ID)}})

			mockDb.EXPECT().GetUserByID(gomock.Any(), gomock.Any(), user.ID).Return(user, nil)

			req := &protocol.WyreCreateTransferRequest{}
			resp, err := s.WyreCreateTransfer(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp).To(Equal(&protocol.WyreTransferDetail{}))
		})
	})
})
