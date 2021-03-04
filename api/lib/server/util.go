package server

import (
	"context"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// GetUserIDFromIncomingContext gets the user id from the incoming context
func GetUserIDFromIncomingContext(ctx context.Context) (user.ID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
	}

	vals := md.Get("user-id")

	if len(vals) == 0 {
		return "", status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
	}
	userID := user.ID(vals[0])

	if userID == "" {
		return "", status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
	}

	return user.ID(userID), nil
}
