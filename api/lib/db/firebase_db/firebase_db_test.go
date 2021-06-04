package firebase_db_test

import (
	"context"

	"github.com/khoerling/flux/api/lib/db/models/gotoconfig"
	"github.com/lithammer/shortuuid/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FirebaseDb", func() {

	Context("SaveGotoConfig", func() {
		var id gotoconfig.ID
		var shortID gotoconfig.ShortID
		var config gotoconfig.SnapWidgetConfig

		BeforeEach(func() {
			var err error

			shortID = gotoconfig.NewShortID()

			config = gotoconfig.SnapWidgetConfig{
				AppName: string(shortID),
			}

			id, err = config.GetID()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("New", func() {
			returnedShortID, err := testManager.SaveGotoConfig(context.Background(), &gotoconfig.Config{
				ID:      gotoconfig.ID(shortuuid.New()),
				ShortID: shortID,
				Config:  config,
			})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(returnedShortID).ShouldNot(BeEmpty())
			Expect(returnedShortID).Should(Equal(shortID))
		})

		It("Existing", func() {
			_, err := testManager.SaveGotoConfig(context.Background(), &gotoconfig.Config{
				ID:      id,
				ShortID: shortID,
				Config:  config,
			})
			Expect(err).ShouldNot(HaveOccurred())

			returnedShortID, err := testManager.SaveGotoConfig(context.Background(), &gotoconfig.Config{
				ID:      id,
				ShortID: gotoconfig.NewShortID(),
				Config:  config,
			})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(returnedShortID).ShouldNot(BeEmpty())
			Expect(returnedShortID).Should(Equal(shortID))
		})
	})

})
