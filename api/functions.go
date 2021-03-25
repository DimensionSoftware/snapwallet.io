package functions

import (
	"context"
	"encoding/json"
	"log"

	"github.com/khoerling/flux/api/cmd/jobrunner/wire"
	"github.com/khoerling/flux/api/lib/db/models/job"
	"github.com/khoerling/flux/api/lib/jobmanager"
	"github.com/khoerling/flux/api/lib/jobs"
)

var jobManager jobmanager.Manager

func init() {
	var err error

	jobManager, err = wire.InitializeJobManager()
	if err != nil {
		log.Fatalf("error during initialization: %#v", err)
	}
}

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// RunSnapJob consumes a Pub/Sub message.
func RunSnapJob(ctx context.Context, msg PubSubMessage) error {
	var j job.Job
	if err := json.Unmarshal(msg.Data, &j); err != nil {
		log.Println("unmarshaling message failed: ", err)
		return nil
	}

	if err := jobs.RunSnapJob(ctx, jobManager, &j); err != nil {
		// todo: whitelist transient errors
		log.Println("job failed: ", err)
	}

	return nil
}
