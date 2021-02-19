package server

import (
	"context"
	"log"

	proto "github.com/khoerling/flux/api/lib/protocol"
)

// UserData is an rpc handler
func (s *Server) UserData(ctx context.Context, in *proto.UserDataRequest) (*proto.UserDataResponse, error) {
	log.Printf("Received: %v", in)

	resp := &proto.UserDataResponse{
		User: &proto.User{
			Id:    -1,
			Email: "bill@microsoft.com",
		},
	}

	return resp, nil
}
