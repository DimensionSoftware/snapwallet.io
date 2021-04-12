package auth

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthenticationInterceptor handles jwt authentication
func (verifier JwtVerifier) AuthenticationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var err error

	ctx, err = verifier.AuthenticateMethod(ctx, info.FullMethod)
	if err != nil {
		return nil, err
	}

	// Last but super important, execute the handler so that the actualy gRPC request is also performed
	// send updated md to context
	return handler(ctx, req)
}

func (verifier JwtVerifier) AuthenticateMethod(ctx context.Context, fullMethod string) (context.Context, error) {
	log.Printf("authentication --> %s", fullMethod)

	if RPCPublicWhitelist[fullMethod] {
		log.Printf("success ✅ (public route)")
		return ctx, nil
	}

	log.Printf("secure route")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "authentication not set")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "authentication not set")
	}
	authorization := values[0]

	expectedPrefix := "Bearer "
	if len(authorization) <= len(expectedPrefix) {
		return ctx, status.Errorf(codes.Unauthenticated, "authentication not set")
	}
	if authorization[:len(expectedPrefix)] != expectedPrefix {
		return ctx, status.Errorf(codes.Unauthenticated, "authentication not set")
	}
	accessToken := authorization[len(expectedPrefix):]

	claims, err := verifier.ParseAndVerify(ctx, TokenKindAccess, accessToken)
	if err != nil {
		log.Printf("%#v\n", err)
		return ctx, status.Errorf(codes.Unauthenticated, "token is invalid or expired")
	}
	log.Printf("claims --> %+v", claims)

	md.Set("user-id", claims.Subject)

	log.Printf("authentication success ✅")

	ctx = metadata.NewIncomingContext(ctx, md)
	return ctx, nil
}
