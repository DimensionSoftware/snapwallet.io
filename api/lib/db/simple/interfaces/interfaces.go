package interfaces

import "context"

// Collection is a control structure for collection
type Collection interface {
	Fetch(context.Context, string, interface{}) error
	FetchInTx(context.Context, Tx, string, interface{}) error
	Scan(context.Context, interface{}) error
	Save(context.Context, interface{}) error
	SaveInTx(context.Context, Tx, interface{}) error
}

// Record is a record in the database
type Record interface {
	ID() string
}

type Tx interface {
	Tx() interface{}
}
