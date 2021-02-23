package main

import (
	"github.com/khoerling/flux/api/lib/wire"
)

const (
	defaultAddress = ":50051"
)

func main() {
	s := wire.InitializeServer()
	s.Serve(defaultAddress)
}
