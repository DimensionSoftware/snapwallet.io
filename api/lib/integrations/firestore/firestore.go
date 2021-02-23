package firestore

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

const firestoreProjectEnvVarName = "FIRESTORE_PROJECT"

// FireProjectID is the google cloud project where the firestore database is located
type FireProjectID string

// ProvideFirestore returns
func ProvideFirestore(projectID FireProjectID) *firestore.Client {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: string(projectID)}

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		panic(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}

	return client
}

// ProvideFirestoreProjectID returns a google cloud project id
func ProvideFirestoreProjectID() FireProjectID {
	projectName := os.Getenv(firestoreProjectEnvVarName)
	if projectName == "" {
		panic(fmt.Sprintf("you must set %s", firestoreProjectEnvVarName))
	}
	return FireProjectID(projectName)
}
