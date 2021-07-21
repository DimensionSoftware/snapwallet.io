// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package wire

import (
	"github.com/khoerling/flux/api/lib/auth"
	"github.com/khoerling/flux/api/lib/config"
	"github.com/khoerling/flux/api/lib/db/firebase_db"
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/filemanager"
	"github.com/khoerling/flux/api/lib/integrations/cloudstorage"
	"github.com/khoerling/flux/api/lib/integrations/firestore"
	"github.com/khoerling/flux/api/lib/integrations/plaid"
	"github.com/khoerling/flux/api/lib/integrations/pubsub"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
	"github.com/khoerling/flux/api/lib/integrations/sendgrid"
	"github.com/khoerling/flux/api/lib/integrations/twilio"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	"github.com/khoerling/flux/api/lib/integrations/wyremanager"
	"github.com/khoerling/flux/api/lib/jobpublisher"
	"github.com/khoerling/flux/api/lib/remedymanager"
	"github.com/khoerling/flux/api/lib/server"
	plaid2 "github.com/plaid/plaid-go/plaid"
)

// Injectors from wire.go:

// InitializeServer creates the main server container
func InitializeServer() (server.Server, error) {
	privateKey, err := auth.ProvideJwtPrivateKey()
	if err != nil {
		return server.Server{}, err
	}
	publicKey := auth.ProvideJwtPublicKey(privateKey)
	fireProjectID, err := firestore.ProvideFirestoreProjectID()
	if err != nil {
		return server.Server{}, err
	}
	client, err := firestore.ProvideFirestore(fireProjectID)
	if err != nil {
		return server.Server{}, err
	}
	encryptionConfig, err := encryption.ProvideConfig()
	if err != nil {
		return server.Server{}, err
	}
	manager, err := encryption.NewManager(encryptionConfig)
	if err != nil {
		return server.Server{}, err
	}
	db := firebase_db.Db{
		Firestore:         client,
		EncryptionManager: manager,
	}
	jwtVerifier := &auth.JwtVerifier{
		PublicKey: publicKey,
		Db:        db,
	}
	grpcServer := server.ProvideGrpcServer(jwtVerifier)
	sendAPIKey, err := sendgrid.ProvideSendClientAPIKey()
	if err != nil {
		return server.Server{}, err
	}
	sendgridClient := sendgrid.ProvideSendClient(sendAPIKey)
	twilioConfig, err := twilio.ProvideTwilioConfig()
	if err != nil {
		return server.Server{}, err
	}
	gotwilioTwilio := twilio.ProvideTwilio(twilioConfig)
	bucketHandle, err := cloudstorage.ProvideBucket()
	if err != nil {
		return server.Server{}, err
	}
	filemanagerManager := &filemanager.Manager{
		BucketHandle:      bucketHandle,
		Db:                db,
		EncryptionManager: manager,
	}
	wyreConfig, err := wyre.ProvideWyreConfig()
	if err != nil {
		return server.Server{}, err
	}
	wyreClient := wyre.NewClient(wyreConfig)
	apiHost, err := config.ProvideAPIHost()
	if err != nil {
		return server.Server{}, err
	}
	clientOptions, err := plaid.ProvideClientOptions()
	if err != nil {
		return server.Server{}, err
	}
	plaidClient, err := plaid2.NewClient(clientOptions)
	if err != nil {
		return server.Server{}, err
	}
	wyremanagerManager := &wyremanager.Manager{
		APIHost:     apiHost,
		Wyre:        wyreClient,
		Db:          db,
		Plaid:       plaidClient,
		FileManager: filemanagerManager,
	}
	jwtSigner := &auth.JwtSigner{
		PrivateKey: privateKey,
	}
	authManager := &auth.Manager{
		JwtSigner:   jwtSigner,
		JwtVerifier: jwtVerifier,
		Db:          db,
	}
	pusherConfig, err := pusher.ProviderPusherConfig()
	if err != nil {
		return server.Server{}, err
	}
	pusherClient := pusher.ProvidePusherClient(pusherConfig)
	pusherManager := &pusher.Manager{
		Pusher: pusherClient,
	}
	pubsubClient, err := pubsub.ProvideClient(fireProjectID)
	if err != nil {
		return server.Server{}, err
	}
	pubsubManager := &pubsub.Manager{
		PubSub: pubsubClient,
	}
	pubSubPublisher := jobpublisher.PubSubPublisher{
		Db:     db,
		PubSub: pubsubManager,
	}
	remedymanagerManager := &remedymanager.Manager{
		Db: db,
	}
	webHost, err := config.ProvideWebHost()
	if err != nil {
		return server.Server{}, err
	}
	serverServer := server.Server{
		GrpcServer:    grpcServer,
		Sendgrid:      sendgridClient,
		Twilio:        gotwilioTwilio,
		TwilioConfig:  twilioConfig,
		FileManager:   filemanagerManager,
		Db:            db,
		Wyre:          wyreClient,
		WyreManager:   wyremanagerManager,
		Plaid:         plaidClient,
		JwtSigner:     jwtSigner,
		JwtVerifier:   jwtVerifier,
		AuthManager:   authManager,
		Pusher:        pusherManager,
		JobPublisher:  pubSubPublisher,
		RemedyManager: remedymanagerManager,
		APIHost:       apiHost,
		WebHost:       webHost,
	}
	return serverServer, nil
}

func InitializeDevServer() (server.Server, error) {
	privateKey, err := auth.ProvideJwtPrivateKey()
	if err != nil {
		return server.Server{}, err
	}
	publicKey := auth.ProvideJwtPublicKey(privateKey)
	fireProjectID, err := firestore.ProvideFirestoreProjectID()
	if err != nil {
		return server.Server{}, err
	}
	client, err := firestore.ProvideFirestore(fireProjectID)
	if err != nil {
		return server.Server{}, err
	}
	encryptionConfig, err := encryption.ProvideConfig()
	if err != nil {
		return server.Server{}, err
	}
	manager, err := encryption.NewManager(encryptionConfig)
	if err != nil {
		return server.Server{}, err
	}
	db := firebase_db.Db{
		Firestore:         client,
		EncryptionManager: manager,
	}
	jwtVerifier := &auth.JwtVerifier{
		PublicKey: publicKey,
		Db:        db,
	}
	grpcServer := server.ProvideGrpcServer(jwtVerifier)
	sendAPIKey, err := sendgrid.ProvideSendClientAPIKey()
	if err != nil {
		return server.Server{}, err
	}
	sendgridClient := sendgrid.ProvideSendClient(sendAPIKey)
	twilioConfig, err := twilio.ProvideTwilioConfig()
	if err != nil {
		return server.Server{}, err
	}
	gotwilioTwilio := twilio.ProvideTwilio(twilioConfig)
	bucketHandle, err := cloudstorage.ProvideBucket()
	if err != nil {
		return server.Server{}, err
	}
	filemanagerManager := &filemanager.Manager{
		BucketHandle:      bucketHandle,
		Db:                db,
		EncryptionManager: manager,
	}
	wyreConfig, err := wyre.ProvideWyreConfig()
	if err != nil {
		return server.Server{}, err
	}
	wyreClient := wyre.NewClient(wyreConfig)
	apiHost, err := config.ProvideAPIHost()
	if err != nil {
		return server.Server{}, err
	}
	clientOptions, err := plaid.ProvideClientOptions()
	if err != nil {
		return server.Server{}, err
	}
	plaidClient, err := plaid2.NewClient(clientOptions)
	if err != nil {
		return server.Server{}, err
	}
	wyremanagerManager := &wyremanager.Manager{
		APIHost:     apiHost,
		Wyre:        wyreClient,
		Db:          db,
		Plaid:       plaidClient,
		FileManager: filemanagerManager,
	}
	jwtSigner := &auth.JwtSigner{
		PrivateKey: privateKey,
	}
	authManager := &auth.Manager{
		JwtSigner:   jwtSigner,
		JwtVerifier: jwtVerifier,
		Db:          db,
	}
	pusherConfig, err := pusher.ProviderPusherConfig()
	if err != nil {
		return server.Server{}, err
	}
	pusherClient := pusher.ProvidePusherClient(pusherConfig)
	pusherManager := &pusher.Manager{
		Pusher: pusherClient,
	}
	inProcessPublisher := jobpublisher.InProcessPublisher{
		Db:          db,
		Pusher:      pusherManager,
		WyreManager: wyremanagerManager,
	}
	remedymanagerManager := &remedymanager.Manager{
		Db: db,
	}
	webHost, err := config.ProvideWebHost()
	if err != nil {
		return server.Server{}, err
	}
	serverServer := server.Server{
		GrpcServer:    grpcServer,
		Sendgrid:      sendgridClient,
		Twilio:        gotwilioTwilio,
		TwilioConfig:  twilioConfig,
		FileManager:   filemanagerManager,
		Db:            db,
		Wyre:          wyreClient,
		WyreManager:   wyremanagerManager,
		Plaid:         plaidClient,
		JwtSigner:     jwtSigner,
		JwtVerifier:   jwtVerifier,
		AuthManager:   authManager,
		Pusher:        pusherManager,
		JobPublisher:  inProcessPublisher,
		RemedyManager: remedymanagerManager,
		APIHost:       apiHost,
		WebHost:       webHost,
	}
	return serverServer, nil
}
