package jobs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/khoerling/flux/api/lib/db/models/job"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
	"github.com/khoerling/flux/api/lib/jobmanager"
	"github.com/lithammer/shortuuid"
)

// RunSnapJob consumes a Pub/Sub message.
func RunSnapJob(ctx context.Context, jobManager jobmanager.Manager, j *job.Job) error {
	var err error
	switch j.Kind {
	case job.KindCreateWyreAccountForUser:
		err = runCreateWyreAccountForUser(ctx, jobManager, j)
	case job.KindCreateWyrePaymentMethodsForUser:
		err = runCreateWyrePaymentMethodsForUser(ctx, jobManager, j)
	default:
		err = fmt.Errorf("error: unsupported job kind: %s", j.Kind)
	}

	if err != nil {
		return err
	}

	return jobManager.MarkJobDone(ctx, j)
}

func runCreateWyrePaymentMethodsForUser(ctx context.Context, m jobmanager.Manager, j *job.Job) error {
	if len(j.RelatedIDs) == 0 {
		return fmt.Errorf("error: relatedIDs can't be empty")
	}

	userID := user.ID(j.RelatedIDs[0])

	items, err := m.GetDb().GetAllPlaidItems(ctx, nil, userID)
	if err != nil {
		return err
	}

	wyreAccounts, err := m.GetDb().GetWyreAccounts(ctx, nil, userID)
	if err != nil {
		return err
	}
	if len(wyreAccounts) == 0 {
		return fmt.Errorf("error: user %s must have wyre account to create payment methods", userID)
	}
	wyreAccount := wyreAccounts[0]

	log.Println("creating wyre payment methods (if needed)")
	pms, err := m.GetWyreManager().CreatePaymentMethodsFromPlaidItems(ctx, userID, wyreAccount.ID, items)
	if err != nil {
		return err
	}

	{
		var ids []string
		for _, pm := range pms {
			ids = append(ids, string(pm.ID))
		}

		err = m.GetPusher().Send(userID, &pusher.Message{
			Kind: pusher.MessageKindWyrePaymentMethodsUpdated,
			IDs:  ids,
			At:   time.Now(),
		})
		if err != nil {
			return fmt.Errorf("error while trying to send via pusher.io to user: %#v", err)
		}
	}

	return nil
}

func runCreateWyreAccountForUser(ctx context.Context, m jobmanager.Manager, j *job.Job) error {
	if len(j.RelatedIDs) == 0 {
		log.Println("error: relatedIDs can't be empty")
		return nil
	}

	userID := user.ID(j.RelatedIDs[0])
	log.Println("creating wyre account")

	pdata, err := m.GetDb().GetAllProfileData(ctx, nil, userID)
	if err != nil {
		log.Println("error while getting profile data: ", err)
		return nil
	}

	_, err = m.GetWyreManager().CreateAccount(ctx, userID, pdata)
	if err != nil {
		log.Println("error while creating wyre account: ", err)
		// retryable once i work out issues?
		return nil
	}

	err = m.GetPusher().Send(userID, &pusher.Message{
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

		err = m.GetJobPublisher().PublishJob(ctx, &job.Job{
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
