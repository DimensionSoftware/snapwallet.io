package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	f "github.com/khoerling/flux/api/lib/db/simple/firestore"
)

type Rec struct {
}

func (*Rec) ID() string {
	return "foo"
}

func main() {
	ctx := context.Background()

	fs, err := firestore.NewClient(ctx, "ds-snap-sandbox")
	if err != nil {
		panic(err)
	}

	c := f.Collection(fs, []string{
		"pancakes",
	})

	var record Rec
	err = c.Fetch(ctx, "foobar", &record)
	if err != nil {
		panic(err)
	}

	var records []Rec
	err = c.Scan(ctx, &records)
	if err != nil {
		panic(err)
	}

	log.Printf("%#v\n", record)
}
