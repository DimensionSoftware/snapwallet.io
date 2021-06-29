package firestore

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/firestore"
	i "github.com/khoerling/flux/api/lib/db/simple/interfaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Collection create new collection bound to client
func Collection(f *firestore.Client, path []string) collection {
	return collection{
		path:      path,
		firestore: f,
	}
}

// firestore simple firestore collection implementation
type collection struct {
	path      []string
	firestore *firestore.Client
}

func (c collection) Fetch(ctx context.Context, id string, out i.Record) error {
	return c.FetchInTx(ctx, nil, id, out)
}

func (c collection) FetchInTx(ctx context.Context, tx *firestore.Transaction, id string, out i.Record) error {
	if id == "" {
		return fmt.Errorf("Fetch: id was blank")
	}

	fullpath := append(c.path, id)
	ref := c.firestore.Doc(strings.Join(fullpath, "/"))

	snap, err := fetchRef(ctx, ref, tx)
	if err != nil {
		return err
	}
	if snap == nil {
		return fmt.Errorf("Fetch (Not Found): %s", ref.Path)
	}

	snap.DataTo(out)

	return nil
}

func fetchRef(ctx context.Context, ref *firestore.DocumentRef, tx *firestore.Transaction) (*firestore.DocumentSnapshot, error) {
	var (
		snap *firestore.DocumentSnapshot
		err  error
	)

	if tx == nil {
		snap, err = ref.Get(ctx)
	} else {
		snap, err = tx.Get(ref)
	}

	if status.Code(err) == codes.NotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return snap, nil
}

func (c collection) Scan(context.Context, []i.Record) error {
	return nil
}
func (c collection) Save(ctx context.Context, records []i.Record) error {
	return c.SaveWithTx(ctx, nil, records)
}

func (c collection) SaveWithTx(ctx context.Context, tx *firestore.Transaction, records []i.Record) error {

	if len(records) == 0 {
		return nil
	}

	if len(records) == 1 {
		rec := records[0]
		fullpath := append(c.path, rec.ID())
		ref := c.firestore.Doc(strings.Join(fullpath, "/"))

		var (
			err error
		)
		if tx == nil {
			_, err = ref.Set(ctx, &rec)
		} else {
			err = tx.Set(ref, &rec)
		}
		if err != nil {
			return err
		}
	}

	// > 1
	if tx == nil {
		batch := c.firestore.Batch()

		for _, rec := range records {
			fullpath := append(c.path, rec.ID())
			ref := c.firestore.Doc(strings.Join(fullpath, "/"))

			batch.Set(ref, &rec)
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	} else {
		for _, rec := range records {
			fullpath := append(c.path, rec.ID())
			ref := c.firestore.Doc(strings.Join(fullpath, "/"))

			err := tx.Set(ref, &rec)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
