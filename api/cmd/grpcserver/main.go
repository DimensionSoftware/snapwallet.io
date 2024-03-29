package main

import (
	"log"

	"github.com/khoerling/flux/api/cmd/grpcserver/wire"
)

const (
	defaultAddress = ":50051"
)

// sigh.......
func main() {
	s, err := wire.InitializeServer()
	if err != nil {
		panic(err)
	}

	log.Println("> listening on port " + defaultAddress[1:])
	s.Serve(defaultAddress)
}
