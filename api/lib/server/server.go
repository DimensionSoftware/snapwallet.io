package server

import (
	"net"

	"cloud.google.com/go/firestore"
	"github.com/khoerling/flux/api/lib/auth"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	proto "github.com/khoerling/flux/api/lib/protocol"
	"github.com/plaid/plaid-go/plaid"
	"github.com/sendgrid/sendgrid-go"
	"google.golang.org/grpc"
)

// Server represents the grpc server and all its handlers attached
type Server struct {
	proto.UnimplementedFluxServer
	GrpcServer     *grpc.Server
	SendgridClient *sendgrid.Client
	Firestore      *firestore.Client
	Wyre           *wyre.Client
	Plaid          *plaid.Client
	JwtSigner      *auth.JwtSigner
}

const sendgridKeyEnvVarName = "SENDGRID_API_KEY"

// ProvideServer instantiates a new grpc server
func ProvideServer(
	sendgridClient *sendgrid.Client,
	firestore *firestore.Client,
	wyre wyre.Client,
	plaid *plaid.Client,
	jwtSigner auth.JwtSigner,
) Server {
	server := Server{
		GrpcServer:     grpc.NewServer(),
		SendgridClient: sendgridClient,
		Firestore:      firestore,
		Wyre:           &wyre,
		Plaid:          plaid,
		JwtSigner:      &jwtSigner,
	}
	proto.RegisterFluxServer(server.GrpcServer, &server)
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
