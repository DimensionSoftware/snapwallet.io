package server

import (
	"net"

	"cloud.google.com/go/firestore"
	"github.com/khoerling/flux/api/lib/auth"
	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/filemanager"
	"github.com/khoerling/flux/api/lib/integrations/twilio"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	proto "github.com/khoerling/flux/api/lib/protocol"
	"github.com/plaid/plaid-go/plaid"
	"github.com/pusher/pusher-http-go"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sfreiberg/gotwilio"
	"google.golang.org/grpc"
)

// Server represents the grpc server and all its handlers attached
type Server struct {
	proto.UnimplementedFluxServer `wire:"-"`
	GrpcServer                    *grpc.Server
	Sendgrid                      *sendgrid.Client
	Twilio                        *gotwilio.Twilio
	TwilioConfig                  *twilio.Config
	Firestore                     *firestore.Client
	FileManager                   *filemanager.Manager
	Db                            *db.Db
	Wyre                          *wyre.Client
	WyreManager                   *wyre.Manager
	Plaid                         *plaid.Client
	JwtSigner                     *auth.JwtSigner
	JwtVerifier                   *auth.JwtVerifier
	AuthManager                   *auth.Manager
	Pusher                        *pusher.Client
}

const sendgridKeyEnvVarName = "SENDGRID_API_KEY"

// Maximum upload of 25 MB
const maxMsgSizeBytes = 1024 * 1024 * 25

func ProvideGrpcServer(jwtVerifier *auth.JwtVerifier) *grpc.Server {
	return grpc.NewServer(grpc.UnaryInterceptor(jwtVerifier.AuthenticationInterceptor), grpc.MaxRecvMsgSize(maxMsgSizeBytes))
}

// Serve starts up the grpc server and listens on a tcp port
func (s *Server) Serve(address string) error {
	proto.RegisterFluxServer(s.GrpcServer, s)

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
