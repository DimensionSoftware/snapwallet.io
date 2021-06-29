package interfaces

import "context"

// Collection is a control structure for collection
type Collection interface {
	Fetch(context.Context, string, *Record) error
	FetchInTx(context.Context, string, *Record) error
	Scan(context.Context, *[]Record) error
	Save(context.Context, *[]Record) error
}

// Record is a record in the database
type Record interface {
	ID() string
}

type Tx interface {
	Tx() interface{}
}
