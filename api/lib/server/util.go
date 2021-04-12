package server

import (
	"context"
	"log"

	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

// LookupUserFromIncomingContext attempts to lookup the user from the incoming context
func LookupUserFromIncomingContext(ctx context.Context, db *db.Db) (*user.User, error) {
	userID := GetUserIDFromIncomingContext(ctx)

	u, err := db.GetUserByID(ctx, nil, user.ID(userID))
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unknown, "An unknown error ocurred; please try again.")
	}

	return u, nil
}

// RequireUserFromIncomingContext requires a user from the incoming context and will a throw an error if the lookup fails
func RequireUserFromIncomingContext(ctx context.Context, db db.Db) (*user.User, error) {
	u, err := LookupUserFromIncomingContext(ctx, db)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
	}

	return u, nil
}
