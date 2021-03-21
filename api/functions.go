package jobrunner

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/khoerling/flux/api/cmd/jobrunner/wire"
	"github.com/khoerling/flux/api/lib/db/models/job"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
)

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// RunSnapJob consumes a Pub/Sub message.
func RunSnapJob(ctx context.Context, msg PubSubMessage) error {
	m, err := wire.InitializeJobManager()
	if err != nil {
		log.Println("error during initialization: ", err)
		return nil
	}

	var j job.Job
	if err := json.Unmarshal(msg.Data, &j); err != nil {
		log.Println("unmarshaling message failed: ", err)
		return nil
	}

	if j.Kind == job.KindCreateWyreAccountForUser {
		if len(j.RelatedIDs) == 0 {
			log.Println("error: relatedIDs can't be empty")
			return nil
		}

		userID := user.ID(j.RelatedIDs[0])
		log.Println("creating wyre account")

		pdata, err := m.Db.GetAllProfileData(ctx, nil, userID)
		if err != nil {
			log.Println("error while getting profile data: ", err)
			return nil
		}

		_, err = m.WyreManager.CreateAccount(ctx, userID, pdata)
		if err != nil {
			log.Println("error while creating wyre account: ", err)
			// retryable once i work out issues?
			return nil
		}

		err = m.Pusher.Send(userID, &pusher.Message{
			Kind: pusher.MessageKindWyreAccountUpdated,
			At:   time.Now(),
		})
		if err != nil {
			log.Println("error while trying to send via pusher.io to user: ", err)
			return nil
		}

		return nil
	}

	log.Printf("error: unsupported job kind: %s\n", j.Kind)
	return nil
}
