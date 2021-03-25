package main

import (
	"context"
	"log"

	"github.com/khoerling/flux/api/cmd/grpcserver/wire"
	"github.com/khoerling/flux/api/lib/db/models/job"
)

const (
	defaultAddress = ":50051"
)

func main() {
	s, err := wire.InitializeDevServer()
	if err != nil {
		panic(err)
	}

	err = s.JobPublisher.PublishJob(context.Background(), &job.Job{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("> listening on port " + defaultAddress[1:])
	s.Serve(defaultAddress)
}
