package firebase_db_test

import (
	"context"

	"github.com/khoerling/flux/api/lib/db/models/gotoconfig"
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
				Wallets: []gotoconfig.SnapWidgetWallet{},
				Intent:  "buy",
				Focus:   true,
				Theme:   map[string]string{},
				Product: &gotoconfig.SnapWidgetProduct{
					VideoURL: "http://site.com/video.mp4",
					Title:    "snuggy socks",
				},
			}

			id, err = config.GetID()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("new ID returns ShortID", func() {
			returnedShortID, err := testManager.SaveGotoConfig(context.Background(), &gotoconfig.Config{
				ID:      id,
				ShortID: shortID,
				Config:  config,
			})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(returnedShortID).ShouldNot(BeEmpty())
			Expect(returnedShortID).Should(Equal(shortID))
		})

		It("existing ID returns same ShortID instead of new one", func() {
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

	Context("GetGotoConfigByShortID", func() {
		It("returns nil when not found", func() {
			config, err := testManager.GetGotoConfigByShortID(context.Background(), gotoconfig.NewShortID())
			Expect(err).ShouldNot(HaveOccurred())
			Expect(config).Should(BeNil())
		})
		It("returns data when present in database", func() {
			shortID := gotoconfig.NewShortID()
			config := gotoconfig.SnapWidgetConfig{
				AppName: string(shortID),
				Wallets: []gotoconfig.SnapWidgetWallet{},
				Intent:  "buy",
				Focus:   true,
				Theme:   map[string]string{},
				Product: &gotoconfig.SnapWidgetProduct{
					VideoURL: "http://site.com/video.mp4",
					Title:    "snuggy socks",
				},
			}

			id, err := config.GetID()
			Expect(err).ShouldNot(HaveOccurred())

			gotoConfig := gotoconfig.Config{
				ID:      id,
				ShortID: shortID,
				Config:  config,
			}

			_, err = testManager.SaveGotoConfig(context.Background(), &gotoConfig)
			Expect(err).ShouldNot(HaveOccurred())

			conf, err := testManager.GetGotoConfigByShortID(context.Background(), shortID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(conf).ShouldNot(BeNil())

			Expect(*conf).Should(Equal(gotoConfig))
		})
	})

})
