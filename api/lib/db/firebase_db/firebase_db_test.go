package firebase_db_test

import (
	"context"

	"github.com/khoerling/flux/api/lib/db/models/gotoconfig"
	"github.com/lithammer/shortuuid/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/teris-io/shortid"
)

var _ = Describe("FirebaseDb", func() {

	Context("SaveGotoConfig", func() {
		It("New", func() {
			shortIDStr, err := shortid.Generate()
			if err != nil {
				panic(err)
			}
			shortID := gotoconfig.ShortID(shortIDStr)

			returnedShortID, err := testManager.SaveGotoConfig(context.Background(), &gotoconfig.Config{
				ID:      gotoconfig.ID(shortuuid.New()),
				ShortID: shortID,
				Config: gotoconfig.SnapWidgetConfig{
					AppName: shortIDStr,
				},
			})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(returnedShortID).ShouldNot(BeEmpty())
			Expect(returnedShortID).Should(Equal(shortID))
		})
	})

})
