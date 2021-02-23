package server

import (
	"net"

	proto "github.com/khoerling/flux/api/lib/protocol"
	"github.com/sendgrid/sendgrid-go"
	"google.golang.org/grpc"
)

// Server represents the grpc server and all its handlers attached
type Server struct {
	proto.UnimplementedAPIServer
	GrpcServer     *grpc.Server
	SendgridClient *sendgrid.Client
}

const sendgridKeyEnvVarName = "SENDGRID_API_KEY"

// ProvideServer instantiates a new grpc server
func ProvideServer(sendgridClient *sendgrid.Client) Server {
	server := Server{
		GrpcServer:     grpc.NewServer(),
		SendgridClient: sendgridClient,
	}
	proto.RegisterAPIServer(server.GrpcServer, &server)
	return server
}

// Serve starts up the grpc server and listens on a tcp port
func (s *Server) Serve(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	err = s.GrpcServer.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}
