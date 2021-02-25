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

// ProvideFirestore returns a *firestore.Client
func ProvideFirestore(projectID FireProjectID) (*firestore.Client, error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: string(projectID)}

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// ProvideFirestoreProjectID returns a google cloud project id
func ProvideFirestoreProjectID() (FireProjectID, error) {
	projectName := os.Getenv(firestoreProjectEnvVarName)
	if projectName == "" {
		return "", fmt.Errorf("you must set %s", firestoreProjectEnvVarName)
	}
	return FireProjectID(projectName), nil
}
