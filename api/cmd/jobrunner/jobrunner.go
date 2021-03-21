package jobrunner

import (
	"context"
	"log"

	"github.com/khoerling/flux/api/cmd/jobrunner/wire"
)

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// RunSnapJob consumes a Pub/Sub message.
func RunSnapJob(ctx context.Context, msg PubSubMessage) error {
	_, err := wire.InitializeJobManager()
	if err != nil {
		return err
	}

	name := string(msg.Data) // Automatically decoded from base64.
	if name == "" {
		name = "World"
	}
	log.Printf("Hello, %s!", name)
	return nil
}
