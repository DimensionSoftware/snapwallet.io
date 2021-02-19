package server

import (
	"net"

	proto "github.com/khoerling/flux/api/lib/protocol"
	"google.golang.org/grpc"
)

// Server represents the grpc server and all its handlers attached
type Server struct {
	proto.UnimplementedAPIServer
	grpcServer *grpc.Server
}

// NewServer instantiates a new grpc server
func NewServer() *Server {
	server := &Server{
		grpcServer: grpc.NewServer(),
	}
	proto.RegisterAPIServer(server.grpcServer, server)
	return server
}

// Serve starts up the grpc server and listens on a tcp port
func (s *Server) Serve(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	err = s.grpcServer.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}
