package main

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	f "github.com/khoerling/flux/api/lib/db/simple/firestore"
)

type Rec struct {
	ID string `firestore:"id"`
}

func (r Rec) GetID() string {
	return r.ID
}

func (r Rec) GetData() map[string]interface{} {
	return map[string]interface{}{
		"id":        r.ID,
		"createdAt": time.Now(),
	}
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

	err = c.Save(ctx, Rec{ID: "foobar"})
	if err != nil {
		log.Println(err)
	}

	var record Rec
	err = c.Fetch(ctx, "foobar", &record)
	if err != nil {
		log.Println(err)
	}

	log.Printf("record: %#v\n", record)

	var records []Rec
	err = c.Scan(ctx, &records)
	if err != nil {
		log.Println(err)
	}

	log.Printf("records: %#v\n", records)
}
