package jobs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/khoerling/flux/api/lib/db/models/job"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/wyre/account"
	"github.com/khoerling/flux/api/lib/integrations/pusher"
	"github.com/khoerling/flux/api/lib/jobmanager"
	"github.com/lithammer/shortuuid"
)

// RunSnapJob consumes a Pub/Sub message.
func RunSnapJob(ctx context.Context, jobManager jobmanager.Manager, j *job.Job) error {
	defer func() {
		err := jobManager.MarkJobDone(ctx, j)
		if err != nil {
			log.Printf("could not mark job %s as done\n", j.ID)
			log.Println(err)
		}
	}()

	switch j.Kind {
	case job.KindUpdateWyreAccountForUser:
		return runUpdateWyreAccountForUser(ctx, jobManager, j)
	case job.KindCreateWyrePaymentMethodsForUser:
		return runCreateWyrePaymentMethodsForUser(ctx, jobManager, j)
	default:
		return fmt.Errorf("error: unsupported job kind: %s", j.Kind)
	}
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

func runUpdateWyreAccountForUser(ctx context.Context, m jobmanager.Manager, j *job.Job) error {
	if len(j.RelatedIDs) == 0 {
		return fmt.Errorf("relatedIDs can't be empty")
	}

	user, err := m.Db.GetUserByID(ctx, nil, user.ID(j.RelatedIDs[0]))
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("user id does not exist")
	}

	var existingWyreAccount *account.Account
	{
		accounts, err := m.Db.GetWyreAccounts(ctx, nil, user.ID)
		if err != nil {
			return err
		}
		if len(accounts) > 0 {
			existingWyreAccount = accounts[0]
		}
	}

	pdata, err := m.GetDb().GetAllProfileData(ctx, nil, user.ID)
	if err != nil {
		return err
	}

	if existingWyreAccount == nil {
		if pdata.HasWyreAccountPreconditionsMet() {
			log.Println("creating wyre account")
			_, err = m.GetWyreManager().CreateAccount(ctx, user.ID, pdata)
			if err != nil {
				return err
			}
		} else {
			return nil
		}
	} else {
		log.Println("updating wyre account")
		err = m.GetWyreManager().UpdateAccountProfileData(ctx, user.ID, existingWyreAccount, pdata)
		if err != nil {
			return err
		}
	}

	err = m.GetPusher().Send(user.ID, &pusher.Message{
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
			RelatedIDs: []string{string(user.ID)},
			CreatedAt:  now.Unix(),
			UpdatedAt:  now.Unix(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
