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

// create new collection bound to client
func CollectionInTx(path []string, f *firestore.Client, tx *firestore.Transaction) collection {
	return collection{
		path:      path,
		firestore: f,
		tx:        tx,
	}
}

func Collection(path []string, f *firestore.Client) collection {
	return collection{
		path:      path,
		firestore: f,
	}
}

// firestore simple firestore collection implementation
type collection struct {
	path      []string
	firestore *firestore.Client
	tx        *firestore.Transaction
}

func (c collection) Fetch(ctx context.Context, out *i.Record, path []string) error {
	if len(path) == 0 {
		return fmt.Errorf("Fetch: path was empty")
	}

	fullpath := append(c.path, path...)
	ref := c.firestore.Doc(strings.Join(fullpath, "/"))

	var (
		snap *firestore.DocumentSnapshot
		err  error
	)

	if c.tx == nil {
		snap, err = ref.Get(ctx)
	} else {
		snap, err = c.tx.Get(ref)
	}

	if status.Code(err) == codes.NotFound {
		//return fmt.Errorf("Fetch: doc not found")
		return err
	}
	if err != nil {
		return err
	}

	snap.DataTo(out)

	return nil
}

func (c collection) Scan(*[]i.Record) error {

}
func (c collection) Save(*[]i.Record) {

}
