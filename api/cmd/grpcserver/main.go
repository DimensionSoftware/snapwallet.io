package main

import (
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

	println("> listening on port " + defaultAddress[1:])
	s.Serve(defaultAddress)
}
