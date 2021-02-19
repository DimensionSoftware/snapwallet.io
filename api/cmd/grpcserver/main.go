package main

import "github.com/khoerling/flux/api/lib/server"

const (
	defaultAddress = ":50051"
)

func main() {
	s := server.NewServer()
	s.Serve(defaultAddress)
}
