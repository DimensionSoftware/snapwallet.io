package main

import (
	"github.com/joho/godotenv"
	"github.com/khoerling/flux/api/lib/wire"
)

const (
	defaultAddress = ":50051"
)

func main() {
	godotenv.Load()
	s, err := wire.InitializeServer()
	if err != nil {
		panic(err)
	}
	s.Serve(defaultAddress)
}
