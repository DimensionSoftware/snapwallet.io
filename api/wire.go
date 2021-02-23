package main

import (
	"github.com/google/wire"
	"github.com/khoerling/flux/api/lib/server"
	"github.com/sendgrid/sendgrid-go"
)

// wire.go

// InitializeServer creates the main server container
func InitializeServer() server.Server {
	wire.Build(sendgrid.NewSendClient)
	return server.Server{}
}
