package cloudstorage

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

const cloudStorageBucketNameEnvVarName = "GCLOUD_BUCKET_NAME"

// ProvideBucket provides a storage client bucket ref
func ProvideBucket() (*storage.BucketHandle, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	bucketName := os.Getenv(cloudStorageBucketNameEnvVarName)
	if bucketName == "" {
		return nil, fmt.Errorf("you must set %s", cloudStorageBucketNameEnvVarName)
	}

	log.Println("🚨 Production Cloud Storage API is activated")

	return client.Bucket(bucketName), nil
}
