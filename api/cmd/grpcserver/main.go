package main

import (
	"context"
	"log"
	"net"

	proto "github.com/khoerling/flux/api/lib/protocol"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	proto.UnimplementedAPIServer
}

func (s *server) UserData(ctx context.Context, in *proto.UserDataRequest) (*proto.UserDataResponse, error) {
	log.Printf("Received: %v", in)

	resp := &proto.UserDataResponse{
		User: &proto.User{
			Id:    -1,
			Email: "bill@microsoft.com",
		},
	}

	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterAPIServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
