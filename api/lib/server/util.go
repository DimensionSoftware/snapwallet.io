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

	var userID user.ID
	if len(vals) > 0 {
		userID = user.ID(vals[0])
	} else {
		return "", status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
	}

	if userID == "" {
		return "", status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
	}

	return user.ID(userID), nil
}
