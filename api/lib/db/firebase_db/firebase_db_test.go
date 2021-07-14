package firebase_db_test

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/bxcodec/faker/v3"
	"github.com/khoerling/flux/api/lib/db/models/gotoconfig"
	"github.com/khoerling/flux/api/lib/db/models/job"
	"github.com/khoerling/flux/api/lib/db/models/onetimepasscode"
	"github.com/khoerling/flux/api/lib/db/models/user"
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
			returnedShortID, err := testManager.Db.SaveGotoConfig(context.Background(), &gotoconfig.Config{
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

			returnedShortID, err := testManager.Db.SaveGotoConfig(context.Background(), &gotoconfig.Config{
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
			config, err := testManager.Db.GetGotoConfigByShortID(context.Background(), gotoconfig.NewShortID())
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

			_, err = testManager.Db.SaveGotoConfig(context.Background(), &gotoConfig)
			Expect(err).ShouldNot(HaveOccurred())

			conf, err := testManager.GetGotoConfigByShortID(context.Background(), shortID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(conf).ShouldNot(BeNil())

			Expect(*conf).Should(Equal(gotoConfig))
		})
	})
	Context("CreateOneTimePasscode", func() {
		Context("with email", func() {
			var otp *onetimepasscode.OneTimePasscode
			BeforeEach(func() {
				var err error
				otp, err = testManager.Db.CreateOneTimePasscode(context.Background(), "bob@microsoft.com", onetimepasscode.LoginKindEmail)
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("returns six digits", func() {
				Expect(otp.Code).Should(HaveLen(6))
			})

			It("should return login kind email", func() {
				Expect(otp.Kind).Should(Equal(onetimepasscode.LoginKindEmail))
			})

			It("should have a created_at timestamp which is recent", func() {
				Expect(otp.CreatedAt).Should(BeTemporally("~", time.Now(), 5*time.Second))
			})
		})
		Context("with phone", func() {
			var otp *onetimepasscode.OneTimePasscode
			BeforeEach(func() {
				var err error
				otp, err = testManager.Db.CreateOneTimePasscode(context.Background(), "818-900-3454", onetimepasscode.LoginKindPhone)
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("returns six digits", func() {
				Expect(otp.Code).Should(HaveLen(6))
			})

			It("should return login kind phone", func() {
				Expect(otp.Kind).Should(Equal(onetimepasscode.LoginKindPhone))
			})

			It("should have a created_at timestamp which is recent", func() {
				Expect(otp.CreatedAt).Should(BeTemporally("~", time.Now(), 5*time.Second))
			})
		})

	})
	Context("CreateOneTimePasscode", func() {
		Context("with email", func() {
			var otp *onetimepasscode.OneTimePasscode
			BeforeEach(func() {
				var err error
				otp, err = testManager.Db.CreateOneTimePasscode(context.Background(), "bob@microsoft.com", onetimepasscode.LoginKindEmail)
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("returns six digits", func() {
				Expect(otp.Code).Should(HaveLen(6))
			})

			It("should return login kind email", func() {
				Expect(otp.Kind).Should(Equal(onetimepasscode.LoginKindEmail))
			})

			It("should have a created_at timestamp which is recent", func() {
				Expect(otp.CreatedAt).Should(BeTemporally("~", time.Now(), 5*time.Second))
			})
		})
		Context("with phone", func() {
			var otp *onetimepasscode.OneTimePasscode
			BeforeEach(func() {
				var err error
				otp, err = testManager.Db.CreateOneTimePasscode(context.Background(), "818-900-3454", onetimepasscode.LoginKindPhone)
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("returns six digits", func() {
				Expect(otp.Code).Should(HaveLen(6))
			})

			It("should return login kind phone", func() {
				Expect(otp.Kind).Should(Equal(onetimepasscode.LoginKindPhone))
			})

			It("should have a created_at timestamp which is recent", func() {
				Expect(otp.CreatedAt).Should(BeTemporally("~", time.Now(), 5*time.Second))
			})
		})

	})

	Context("AckOneTimePasscode", func() {
		number := "818-900-3454"
		var otp *onetimepasscode.OneTimePasscode
		BeforeEach(func() {
			var err error
			otp, err = testManager.Db.CreateOneTimePasscode(context.Background(), number, onetimepasscode.LoginKindPhone)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("returns nil when not found", func() {
			returnedOTP, err := testManager.Db.AckOneTimePasscode(context.Background(), number, "333")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(returnedOTP).Should(BeNil())
		})

		It("returns item when present in the database", func() {
			returnedOTP, err := testManager.Db.AckOneTimePasscode(context.Background(), number, otp.Code)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(returnedOTP).ShouldNot(BeNil())
		})

		It("is destroyed in database after ack", func() {
			_, err := testManager.Db.AckOneTimePasscode(context.Background(), number, otp.Code)
			Expect(err).ShouldNot(HaveOccurred())

			returnedOTP, err := testManager.Db.AckOneTimePasscode(context.Background(), number, otp.Code)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(returnedOTP).Should(BeNil())
		})
	})

	Context("SaveUser", func() {
		Context("without transaction", func() {
			It("can save", func() {
				err := testManager.Db.SaveUser(context.Background(), nil, genFakeUser())
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("with transaction", func() {
			It("can save", func() {
				err := testManager.Db.RunTransaction(context.Background(), func(ctx context.Context, tx *firestore.Transaction) error {
					return testManager.Db.SaveUser(ctx, tx, genFakeUser())
				})
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})

	Context("SaveJob", func() {
		Context("without transaction", func() {
			It("can save", func() {
				err := testManager.Db.SaveJob(context.Background(), nil, genFakeJob())
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("with transaction", func() {
			It("can save", func() {
				err := testManager.Db.RunTransaction(context.Background(), func(ctx context.Context, tx *firestore.Transaction) error {
					return testManager.Db.SaveJob(ctx, tx, genFakeJob())
				})
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})

})

func genFakeUser() *user.User {
	email := faker.Email()
	phone := faker.E164PhoneNumber()

	return &user.User{
		ID:    user.ID(shortuuid.New()),
		Email: &email,
		Phone: &phone,
	}
}

func genFakeJob() *job.Job {
	now := time.Now()

	return &job.Job{
		ID:         shortuuid.New(),
		Kind:       job.KindUpdateWyreAccountForUser,
		Status:     job.StatusQueued,
		RelatedIDs: []string{"1", "2", "3"},
		CreatedAt:  now.Unix(),
		UpdatedAt:  now.Unix(),
	}
}
