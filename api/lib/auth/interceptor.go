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
func (verifier JwtVerifier) AuthenticationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (response interface{}, err error) {
	log.Printf("authentication --> %s", info.FullMethod)

	if RPCPublicWhitelist[info.FullMethod] {
		log.Printf("success ✅ (public route)")
		return handler(ctx, req)
	}

	log.Printf("secure route")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, codes.Unauthenticated.String())
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, codes.Unauthenticated.String())
	}
	authorization := values[0]

	expectedPrefix := "Bearer "
	if len(authorization) <= len(expectedPrefix) {
		return nil, status.Errorf(codes.Unauthenticated, codes.Unauthenticated.String())
	}
	if authorization[:len(expectedPrefix)] != expectedPrefix {
		return nil, status.Errorf(codes.Unauthenticated, codes.Unauthenticated.String())
	}
	accessToken := authorization[len(expectedPrefix):]

	claims, err := verifier.ParseAndVerify(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, codes.Unauthenticated.String())
	}
	log.Printf("claims --> %+v", claims)

	md.Set("user-id", claims.Subject)

	log.Printf("authentication success ✅")

	// Last but super important, execute the handler so that the actualy gRPC request is also performed
	return handler(ctx, req)
}
