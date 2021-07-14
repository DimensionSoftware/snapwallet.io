package server_test

import (
	"context"

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
		ctx = context.Background()
		s = &server.Server{}
	})

	Context("WyreCreateTransfer", func() {
		req := &protocol.WyreCreateTransferRequest{}

		resp, err := s.WyreCreateTransfer(ctx, req)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(resp).To(Equal(&protocol.WyreTransferDetail{}))

	})
})
