package main

import (
	"log"

	"github.com/khoerling/flux/api/lib/wire"
)

const (
	defaultAddress = ":50051"
)

func main() {
	s, err := wire.InitializeServer()
	if err != nil {
		panic(err)
	}

	log.Println("> listening on port " + defaultAddress[1:])
	s.Serve(defaultAddress)
}
