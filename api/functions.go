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
	"github.com/khoerling/flux/api/lib/jobmanager"
	"github.com/lithammer/shortuuid"
)

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// RunSnapJob consumes a Pub/Sub message.
func RunSnapJob(ctx context.Context, msg PubSubMessage) error {
	var err error

	m, err := wire.InitializeJobManager()
	if err != nil {
		log.Println("error during initialization: ", err)
		return nil
	}

	var j job.Job
	if err = json.Unmarshal(msg.Data, &j); err != nil {
		log.Println("unmarshaling message failed: ", err)
		return nil
	}

	switch j.Kind {
	case job.KindCreateWyreAccountForUser:
		err = runCreateWyreAccountForUser(ctx, m, j)
	case job.KindCreateWyrePaymentMethodsForUser:
		err = runCreateWyrePaymentMethodsForUser(ctx, m, j)
	default:
		log.Printf("error: unsupported job kind: %s\n", j.Kind)
		return nil
	}
	if err != nil {
		// todo: whitelist transient errors
		log.Println("job failed: ", err)
	}

	return nil
}

func runCreateWyrePaymentMethodsForUser(ctx context.Context, m jobmanager.Manager, j job.Job) error {
	if len(j.RelatedIDs) == 0 {
		return fmt.Errorf("error: relatedIDs can't be empty")
	}

	userID := user.ID(j.RelatedIDs[0])

	items, err := m.Db.GetAllPlaidItems(ctx, nil, userID)
	if err != nil {
		return err
	}

	wyreAccounts, err := m.Db.GetWyreAccounts(ctx, nil, userID)
	if err != nil {
		return err
	}
	if len(wyreAccounts) == 0 {
		return fmt.Errorf("error: user %s must have wyre account to create payment methods", userID)
	}
	wyreAccount := wyreAccounts[0]

	log.Println("creating wyre payment methods (if needed)")
	pms, err := m.WyreManager.CreatePaymentMethodsFromPlaidItems(ctx, userID, wyreAccount.ID, items)
	if err != nil {
		return err
	}

	{
		var ids []string
		for _, pm := range pms {
			ids = append(ids, string(pm.ID))
		}

		err = m.Pusher.Send(userID, &pusher.Message{
			Kind: pusher.MessageKindWyrePaymentMethodUpdated,
			IDs:  ids,
			At:   time.Now(),
		})
		if err != nil {
			return fmt.Errorf("error while trying to send via pusher.io to user: %#v", err)
		}
	}

	return nil
}

func runCreateWyreAccountForUser(ctx context.Context, m jobmanager.Manager, j job.Job) error {
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

	// submit wyre create payment methods job once account is complete
	{
		now := time.Now()

		err = m.PubSub.SendJob(ctx, &job.Job{
			ID:         shortuuid.New(),
			Kind:       job.KindCreateWyrePaymentMethodsForUser,
			Status:     job.StatusQueued,
			RelatedIDs: []string{string(userID)},
			CreatedAt:  now.Unix(),
			UpdatedAt:  now.Unix(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
