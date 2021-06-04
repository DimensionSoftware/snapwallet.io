// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package wire

import (
	"github.com/khoerling/flux/api/lib/config"
	"github.com/khoerling/flux/api/lib/db/firebase_db"
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/filemanager"
	"github.com/khoerling/flux/api/lib/integrations/cloudstorage"
	"github.com/khoerling/flux/api/lib/integrations/firestore"
	"github.com/khoerling/flux/api/lib/integrations/plaid"
	"github.com/khoerling/flux/api/lib/integrations/pubsub"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	"github.com/khoerling/flux/api/lib/integrations/wyremanager"
	"github.com/khoerling/flux/api/lib/jobmanager"
	"github.com/khoerling/flux/api/lib/jobpublisher"
	plaid2 "github.com/plaid/plaid-go/plaid"
)

// Injectors from wire.go:

// InitializeServer creates the main server container
func InitializeJobManager() (jobmanager.Manager, error) {
	fireProjectID, err := firestore.ProvideFirestoreProjectID()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	client, err := firestore.ProvideFirestore(fireProjectID)
	if err != nil {
		return jobmanager.Manager{}, err
	}
	encryptionConfig, err := encryption.ProvideConfig()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	manager, err := encryption.NewManager(encryptionConfig)
	if err != nil {
		return jobmanager.Manager{}, err
	}
	db := firebase_db.Db{
		Firestore:         client,
		EncryptionManager: manager,
	}
	pusherConfig, err := pusher.ProviderPusherConfig()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	pusherClient := pusher.ProvidePusherClient(pusherConfig)
	pusherManager := &pusher.Manager{
		Pusher: pusherClient,
	}
	apiHost, err := config.ProvideAPIHost()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	wyreConfig, err := wyre.ProvideWyreConfig()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	wyreClient := wyre.NewClient(wyreConfig)
	clientOptions, err := plaid.ProvideClientOptions()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	plaidClient, err := plaid2.NewClient(clientOptions)
	if err != nil {
		return jobmanager.Manager{}, err
	}
	bucketHandle, err := cloudstorage.ProvideBucket()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	filemanagerManager := &filemanager.Manager{
		BucketHandle:      bucketHandle,
		Db:                db,
		EncryptionManager: manager,
	}
	wyremanagerManager := &wyremanager.Manager{
		APIHost:     apiHost,
		Wyre:        wyreClient,
		Db:          db,
		Plaid:       plaidClient,
		FileManager: filemanagerManager,
	}
	pubsubClient, err := pubsub.ProvideClient(fireProjectID)
	if err != nil {
		return jobmanager.Manager{}, err
	}
	pubsubManager := &pubsub.Manager{
		PubSub: pubsubClient,
	}
	pubSubPublisher := jobpublisher.PubSubPublisher{
		Db:     db,
		PubSub: pubsubManager,
	}
	jobmanagerManager := jobmanager.Manager{
		Db:           db,
		Pusher:       pusherManager,
		WyreManager:  wyremanagerManager,
		JobPublisher: pubSubPublisher,
	}
	return jobmanagerManager, nil
}

func InitializeDevJobManager() (jobmanager.Manager, error) {
	fireProjectID, err := firestore.ProvideFirestoreProjectID()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	client, err := firestore.ProvideFirestore(fireProjectID)
	if err != nil {
		return jobmanager.Manager{}, err
	}
	encryptionConfig, err := encryption.ProvideConfig()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	manager, err := encryption.NewManager(encryptionConfig)
	if err != nil {
		return jobmanager.Manager{}, err
	}
	db := firebase_db.Db{
		Firestore:         client,
		EncryptionManager: manager,
	}
	pusherConfig, err := pusher.ProviderPusherConfig()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	pusherClient := pusher.ProvidePusherClient(pusherConfig)
	pusherManager := &pusher.Manager{
		Pusher: pusherClient,
	}
	apiHost, err := config.ProvideAPIHost()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	wyreConfig, err := wyre.ProvideWyreConfig()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	wyreClient := wyre.NewClient(wyreConfig)
	clientOptions, err := plaid.ProvideClientOptions()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	plaidClient, err := plaid2.NewClient(clientOptions)
	if err != nil {
		return jobmanager.Manager{}, err
	}
	bucketHandle, err := cloudstorage.ProvideBucket()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	filemanagerManager := &filemanager.Manager{
		BucketHandle:      bucketHandle,
		Db:                db,
		EncryptionManager: manager,
	}
	wyremanagerManager := &wyremanager.Manager{
		APIHost:     apiHost,
		Wyre:        wyreClient,
		Db:          db,
		Plaid:       plaidClient,
		FileManager: filemanagerManager,
	}
	inProcessPublisher := jobpublisher.InProcessPublisher{
		Db:          db,
		Pusher:      pusherManager,
		WyreManager: wyremanagerManager,
	}
	jobmanagerManager := jobmanager.Manager{
		Db:           db,
		Pusher:       pusherManager,
		WyreManager:  wyremanagerManager,
		JobPublisher: inProcessPublisher,
	}
	return jobmanagerManager, nil
}
