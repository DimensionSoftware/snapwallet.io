package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/khoerling/flux/api/lib/integrations/firestore"
)

func ProvideClient(projectID firestore.FireProjectID) (*pubsub.Client, error) {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, string(projectID))
	if err != nil {
		return nil, err
	}

	return client, nil
}
