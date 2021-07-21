package server_test

import (
	"context"

	"github.com/khoerling/flux/api/lib/integration_t_manager/wire"
	"github.com/khoerling/flux/api/lib/protocol"
	"github.com/khoerling/flux/api/lib/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
			req := &protocol.WyreCreateTransferRequest{}
			resp, err := s.WyreCreateTransfer(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp).To(Equal(&protocol.WyreTransferDetail{}))
		})

	})
})
