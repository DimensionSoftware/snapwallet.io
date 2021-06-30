package interfaces

import "context"

// Collection is a control structure for collection
type Collection interface {
	Fetch(context.Context, string, *Record) error
	FetchInTx(context.Context, Tx, string, *Record) error
	Scan(context.Context, *[]Record) error
	Save(context.Context, *[]Record) error
	SaveInTx(context.Context, Tx, *[]Record) error
}

// Record is a record in the database
type Record interface {
	ID() string
}

type Tx interface {
	Tx() interface{}
}
