package interfaces

import "context"

// Collection is a control structure for collection
type Collection interface {
	Fetch(context.Context, *Record, string) error
	Scan(context.Context, *[]Record) error
	Save(context.Context, *[]Record) error
}

// Record is a record in the database
type Record interface {
	ID() string
}
