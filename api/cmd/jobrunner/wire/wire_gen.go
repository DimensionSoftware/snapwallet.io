// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package wire

import (
	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/integrations/firestore"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	"github.com/khoerling/flux/api/lib/jobmanager"
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
	config, err := encryption.ProvideConfig()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	manager, err := encryption.NewManager(config)
	if err != nil {
		return jobmanager.Manager{}, err
	}
	dbDb := &db.Db{
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
	apiHost, err := wyre.ProvideAPIHost()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	wyreConfig, err := wyre.ProvideWyreConfig()
	if err != nil {
		return jobmanager.Manager{}, err
	}
	wyreClient := wyre.NewClient(wyreConfig)
	wyreManager := &wyre.Manager{
		APIHost: apiHost,
		Wyre:    wyreClient,
		Db:      dbDb,
	}
	jobmanagerManager := jobmanager.Manager{
		Db:          dbDb,
		Pusher:      pusherManager,
		WyreManager: wyreManager,
	}
	return jobmanagerManager, nil
}
