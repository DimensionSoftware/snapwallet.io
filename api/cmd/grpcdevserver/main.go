package main

import (
	"context"
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

	u, err := s.Db.GetUserByWyreAccountID(context.Background(), "AC_HCF8AY9NJCV")
	if err != nil {
		panic(err)
	}
	log.Printf("%#v\n", u)

	log.Println("> listening on port " + defaultAddress[1:])
	s.Serve(defaultAddress)
}
