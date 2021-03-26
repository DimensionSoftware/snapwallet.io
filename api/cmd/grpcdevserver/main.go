package main

import (
	"log"

	"github.com/khoerling/flux/api/cmd/grpcserver/wire"
)

const (
	defaultAddress = ":50051"
)

func main() {
	s, err := wire.InitializeDevServer()
	if err != nil {
		panic(err)
	}

	log.Println("> listening on port " + defaultAddress[1:])
	s.Serve(defaultAddress)
}
