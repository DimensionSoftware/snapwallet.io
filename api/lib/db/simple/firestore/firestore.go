package firestore

import (
	"context"
	"fmt"
	"log"
	"reflect"
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

func (c collection) Fetch(ctx context.Context, id string, out interface{}) error {
	return c.FetchInTx(ctx, nil, id, out)
}

func (c collection) FetchInTx(ctx context.Context, tx *firestore.Transaction, id string, out interface{}) error {
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

func (c collection) Scan(ctx context.Context, out interface{}) error {
	ref := c.firestore.Collection(strings.Join(c.path, "/"))

	docs, err := ref.Documents(ctx).GetAll()
	if err != nil {
		return err
	}
	log.Printf("inner docs: %#v", docs)

	outType := reflect.TypeOf(out)
	sliceType := outType.Elem()
	records := reflect.Zero(sliceType)

	for _, doc := range docs {
		rec := reflect.Zero(sliceType.Elem())
		err := doc.DataTo(&rec)
		if err != nil {
			return err
		}
		records = reflect.Append(records, rec)
	}

	//out_ := reflect.NewAt(reflect.TypeOf(out))
	//out_ := reflect.PtrTo(outType)
	reflect.Indirect(reflect.ValueOf(out)).Set(records)

	log.Printf("inner records: %#v", records)
	return nil
}

func (c collection) Save(ctx context.Context, in interface{}) error {
	return c.SaveInTx(ctx, nil, in)
}

func (c collection) SaveInTx(ctx context.Context, tx *firestore.Transaction, in interface{}) error {

	var records []i.Record

	switch r := in.(type) {
	case i.Record:
		records = []i.Record{r}
	case []i.Record:
		records = r
	default:
		return fmt.Errorf("SaveInTx: must be []interfaces.Record or interfaces.Record")
	}

	if len(records) == 0 {
		return nil
	}

	if len(records) == 1 {
		rec := records[0]

		id := rec.GetID()
		if id == "" {
			return fmt.Errorf("SaveInTx: id was blank")
		}

		fullpath := append(c.path, id)
		ref := c.firestore.Doc(strings.Join(fullpath, "/"))

		var (
			err error
		)
		if tx == nil {
			_, err = ref.Set(ctx, rec.GetData())
		} else {
			err = tx.Set(ref, rec.GetData())
		}
		if err != nil {
			return err
		}

		return nil
	}

	// > 1
	if tx == nil {
		batch := c.firestore.Batch()

		for _, rec := range records {
			id := rec.GetID()
			if id == "" {
				return fmt.Errorf("SaveInTx: id was blank")
			}

			fullpath := append(c.path, id)
			ref := c.firestore.Doc(strings.Join(fullpath, "/"))

			batch.Set(ref, rec.GetData())
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	} else {
		for _, rec := range records {
			id := rec.GetID()
			if id == "" {
				return fmt.Errorf("SaveInTx: id was blank")
			}

			fullpath := append(c.path, id)
			ref := c.firestore.Doc(strings.Join(fullpath, "/"))

			err := tx.Set(ref, rec.GetData())
			if err != nil {
				return err
			}
		}
	}

	return nil
}
