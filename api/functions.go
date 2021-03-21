package jobrunner

import (
	"context"
	"encoding/json"
	"fmt"
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
		return err
	}

	var j job.Job
	if err := json.Unmarshal(msg.Data, &j); err != nil {
		return err
	}

	if j.Kind == job.KindCreateWyreAccountForUser {
		userID := user.ID(j.RelatedIDs[0])
		log.Println("creating wyre account")

		pdata, err := m.Db.GetAllProfileData(ctx, nil, userID)
		if err != nil {
			return err
		}

		_, err = m.WyreManager.CreateAccount(ctx, userID, pdata)
		if err != nil {
			return err
		}

		m.Pusher.Send(userID, &pusher.Message{
			Kind: pusher.MessageKindWyreAccountUpdated,
			At:   time.Now(),
		})

	} else {
		panic(fmt.Sprintf("unsupported job kind: %s\n", j.Kind))
	}

	return nil
}
