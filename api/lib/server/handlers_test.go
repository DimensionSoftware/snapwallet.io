package server_test

import (
	"context"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/khoerling/flux/api/lib/db/mock_db"
	wyre_account "github.com/khoerling/flux/api/lib/db/models/user/wyre/account"
	"github.com/khoerling/flux/api/lib/db/models/user/wyre/paymentmethod"
	"github.com/khoerling/flux/api/lib/db/test_utils"
	"github.com/khoerling/flux/api/lib/integration_t_manager/wire"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	"github.com/khoerling/flux/api/lib/integrations/wyre/mock_wyre"
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
			//Skip("WIP")
			now := time.Now()

			mockDb := s.Db.(*mock_db.MockDb)
			mockWyre := s.Wyre.(*mock_wyre.MockClientInterface)
			user := test_utils.GenFakeUser()
			//wyreAccount := test_utils.GenFakeWyreAccount()
			wyreAccount := &wyre_account.Account{
				ID:        wyre_account.ID("AC_" + shortuuid.New()),
				APIKey:    "key",
				SecretKey: "secret",
				Status:    "APPROVED",
				CreatedAt: now,
				UpdatedAt: &now,
			}
			wyrePaymentMethod := &paymentmethod.PaymentMethod{
				ID:                    paymentmethod.ID("PA_" + shortuuid.New()),
				Status:                "ACTIVE",
				Name:                  "Fake Checking 1111",
				Last4:                 "1111",
				ChargeableCurrencies:  []string{"USD"},
				DepositableCurrencies: []string{"USD"},
				CreatedAt:             now.Add(time.Minute),
				UpdatedAt:             now.Add(time.Minute),
			}
			ctx := metadata.NewIncomingContext(ctx, map[string][]string{"user-id": {string(user.ID)}})

			mockDb.EXPECT().GetUserByID(ctx, gomock.Any(), user.ID).Return(user, nil)
			mockDb.EXPECT().GetWyreAccounts(ctx, gomock.Any(), user.ID).Return(
				[]*wyre_account.Account{wyreAccount},
				nil,
			)
			mockDb.EXPECT().GetWyrePaymentMethods(ctx, gomock.Any(), user.ID, wyreAccount.ID).Return(
				[]*paymentmethod.PaymentMethod{wyrePaymentMethod},
				nil,
			)

			req := &protocol.WyreCreateTransferRequest{
				Source:       wyrePaymentMethod.ID.String(),
				Amount:       &protocol.WyreCreateTransferRequest_SourceAmount{SourceAmount: 500},
				DestCurrency: "ETH",
				// TODO: proves we need more validation
				Dest: "yomamma",
			}
			expectedWyreReq := wyre.CreateTransferRequest{
				SourceCurrency: "USD",
				Source:         "paymentmethod:" + req.Source,
				SourceAmount:   500,
				DestCurrency:   req.DestCurrency,
				// TODO: proves we need more validation
				Dest:    req.Dest,
				Message: "TODO",
			}.WithDefaults()
			mockWyre.EXPECT().CreateTransfer(wyreAccount.SecretKey, expectedWyreReq).Return(&wyre.TransferDetail{
				Source:         expectedWyreReq.Source,
				SourceCurrency: "USD",
				Dest:           expectedWyreReq.Dest,
				DestCurrency:   expectedWyreReq.DestCurrency,
				DestAmount:     1,
				SourceAmount:   expectedWyreReq.SourceAmount,
				// TODO: need message functionality
				Message: "TODO",
			}, nil)

			resp, err := s.WyreCreateTransfer(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp).To(Equal(&protocol.WyreTransferDetail{}))
		})
	})
})
