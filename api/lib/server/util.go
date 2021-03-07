package server

import (
	"context"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"google.golang.org/grpc/metadata"
)

// GetUserIDFromIncomingContext gets the user id from the incoming context
func GetUserIDFromIncomingContext(ctx context.Context) user.ID {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	vals := md.Get("user-id")

	if len(vals) == 0 {
		return ""
	}

	return user.ID(vals[0])
}
