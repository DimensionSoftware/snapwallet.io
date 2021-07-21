package server_test

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/khoerling/flux/api/lib/db/mock_db"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/transaction"
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
			u := test_utils.GenFakeUser()
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
			ctx := metadata.NewIncomingContext(ctx, map[string][]string{"user-id": {string(u.ID)}})

			mockDb.EXPECT().GetUserByID(ctx, gomock.Any(), u.ID).Return(u, nil)
			mockDb.EXPECT().GetWyreAccounts(ctx, gomock.Any(), u.ID).Return(
				[]*wyre_account.Account{wyreAccount},
				nil,
			)
			mockDb.EXPECT().GetWyrePaymentMethods(ctx, gomock.Any(), u.ID, wyreAccount.ID).Return(
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
			wyreTransferDetail := &wyre.TransferDetail{
				ID:             wyre.TransferID("TF_" + shortuuid.New()),
				Source:         expectedWyreReq.Source,
				SourceCurrency: "USD",
				Dest:           expectedWyreReq.Dest,
				DestCurrency:   expectedWyreReq.DestCurrency,
				DestAmount:     1,
				SourceAmount:   expectedWyreReq.SourceAmount,
				// TODO: need message functionality
				Message: "TODO",
			}
			mockWyre.EXPECT().CreateTransfer(wyreAccount.SecretKey, expectedWyreReq).Return(wyreTransferDetail, nil)

			expectedTxn := &transaction.Transaction{
				SourceName:     fmt.Sprintf("Bank Account: %s", wyrePaymentMethod.ID),
				Source:         req.Source,
				SourceCurrency: "USD",
				DestName:       "ETH Address: yomamma",
				Dest:           "yomamma",
				DestCurrency:   "ETH",
				SourceAmount:   500,
				DestAmount:     1,
				Message:        "TODO",
				Partner:        transaction.PartnerWyre,
				Kind:           transaction.KindACH,
				Status:         transaction.StatusQuoted,
				ExternalIDs:    transaction.ExternalIDs{transaction.ExternalID(wyreTransferDetail.ID.String())},
			}
			mockDb.EXPECT().SaveTransaction(ctx, gomock.Any(), u.ID, gomock.Any()).DoAndReturn(
				func(ctx context.Context, tx *firestore.Transaction, userID user.ID, txn *transaction.Transaction) error {
					ExpectSame(expectedTxn, txn, cmpopts.IgnoreFields(transaction.Transaction{},
						"ID",
					))
					Expect(txn.ID).ToNot(BeEmpty())
					Expect(txn.ExternalIDs).ToNot(BeEmpty())
					return nil
				},
			)

			resp, err := s.WyreCreateTransfer(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp).To(Equal(&protocol.WyreTransferDetail{}))
		})
	})
})

func ExpectSame(expected interface{}, actual interface{}, opts ...cmp.Option) {
	diff := cmp.Diff(expected, actual, opts...)
	if diff != "" {
		Fail(fmt.Sprintf("%T did not match %T:\n%s", actual, expected, diff))
	}
}
