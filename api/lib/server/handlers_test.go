package server_test

import (
	"context"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/khoerling/flux/api/lib/db/mock_db"
	wyre_account "github.com/khoerling/flux/api/lib/db/models/user/wyre/account"
	"github.com/khoerling/flux/api/lib/db/test_utils"
	"github.com/khoerling/flux/api/lib/integration_t_manager/wire"
	"github.com/khoerling/flux/api/lib/protocol"
	"github.com/khoerling/flux/api/lib/server"
	"github.com/lithammer/shortuuid/v3"
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
			now := time.Now()

			mockDb := s.Db.(*mock_db.MockDb)
			//mockWyre := s.Wyre.(*mock_wyre.MockClientInterface)
			user := test_utils.GenFakeUser()
			//wyreAccount := test_utils.GenFakeWyreAccount()
			wyreAccount := &wyre_account.Account{
				ID:        wyre_account.ID(shortuuid.New()),
				APIKey:    "key",
				SecretKey: "secret",
				Status:    "APPROVED",
				CreatedAt: now,
				UpdatedAt: &now,
			}
			ctx := metadata.NewIncomingContext(ctx, map[string][]string{"user-id": {string(user.ID)}})

			mockDb.EXPECT().GetUserByID(ctx, gomock.Any(), user.ID).Return(user, nil)
			mockDb.EXPECT().GetWyreAccounts(ctx, gomock.Any(), user.ID).Return(
				[]*wyre_account.Account{wyreAccount},
				nil,
			)

			req := &protocol.WyreCreateTransferRequest{}
			resp, err := s.WyreCreateTransfer(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp).To(Equal(&protocol.WyreTransferDetail{}))
		})
	})
})
