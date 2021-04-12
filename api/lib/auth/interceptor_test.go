package auth

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func Test_JwtVerifier_authenticateMethod(t *testing.T) {
	tt := []struct {
		desc                string
		fullMethod          string
		mdAuthorizationDesc string
		mdAuthorization     *[]string
		expectedStatus      *status.Status
	}{
		// Public Routes
		{
			fullMethod:          "/Flux/PricingData",
			mdAuthorizationDesc: "unset",
			mdAuthorization:     nil,
			expectedStatus:      nil,
		},
		{
			fullMethod:          "/Flux/OneTimePasscode",
			mdAuthorizationDesc: "unset",
			mdAuthorization:     nil,
			expectedStatus:      nil,
		},
		{
			fullMethod:          "/Flux/OneTimePasscodeVerify",
			mdAuthorizationDesc: "unset",
			mdAuthorization:     nil,
			expectedStatus:      nil,
		},

		{
			fullMethod:          "/Flux/WyreWebhook",
			mdAuthorizationDesc: "unset",
			mdAuthorization:     nil,
			expectedStatus:      nil,
		},
		{
			fullMethod:          "/Flux/TokenExchange",
			mdAuthorizationDesc: "unset",
			mdAuthorization:     nil,
			expectedStatus:      nil,
		},
		// Test token header edge cases
		{
			fullMethod:          "/Flux/ViewerData",
			mdAuthorizationDesc: "empty array",
			mdAuthorization:     &[]string{},
			expectedStatus:      status.New(codes.Unauthenticated, "authentication not set"),
		},
		{
			fullMethod:          "/Flux/ViewerData",
			mdAuthorizationDesc: "too short #1",
			mdAuthorization: &[]string{
				"",
			},
			expectedStatus: status.New(codes.Unauthenticated, "authentication not set"),
		},
		{
			fullMethod:          "/Flux/ViewerData",
			mdAuthorizationDesc: "too short #2",
			mdAuthorization: &[]string{
				"foo",
			},
			expectedStatus: status.New(codes.Unauthenticated, "authentication not set"),
		},
		{
			fullMethod:          "/Flux/ViewerData",
			mdAuthorizationDesc: "too short #3",
			mdAuthorization: &[]string{
				"Bearer",
			},
			expectedStatus: status.New(codes.Unauthenticated, "authentication not set"),
		},
		{
			fullMethod:          "/Flux/ViewerData",
			mdAuthorizationDesc: "too short #4",
			mdAuthorization: &[]string{
				"Bearer ",
			},
			expectedStatus: status.New(codes.Unauthenticated, "authentication not set"),
		},
		{
			fullMethod:          "/Flux/ViewerData",
			mdAuthorizationDesc: "lowercase bearer",
			mdAuthorization: &[]string{
				"bearer foobar",
			},
			expectedStatus: status.New(codes.Unauthenticated, "authentication not set"),
		},
		{
			fullMethod:          "/Flux/ViewerData",
			mdAuthorizationDesc: "valid Bearer header but invalid token",
			mdAuthorization: &[]string{
				"Bearer foobar",
			},
			expectedStatus: status.New(codes.Unauthenticated, "token is invalid or expired"),
		},
		// Secure Routes
		{
			fullMethod:          "/Flux/ViewerData",
			mdAuthorizationDesc: "unset",
			mdAuthorization:     nil,
			expectedStatus:      status.New(codes.Unauthenticated, "authentication not set"),
		},
		{
			fullMethod:          "/Flux/ViewerProfileData",
			mdAuthorizationDesc: "unset",
			mdAuthorization:     nil,
			expectedStatus:      status.New(codes.Unauthenticated, "authentication not set"),
		},
	}

	ctx := context.Background()

	for _, tc := range tt {
		tc := tc
		t.Run(fmt.Sprintf("%s given %s auth should have status %s and message %s", tc.fullMethod, tc.mdAuthorizationDesc, tc.expectedStatus.Code().String(), tc.expectedStatus.Message()), func(t *testing.T) {
			t.Parallel()

			a := assert.New(t)
			v := JwtVerifier{}

			var err error

			if tc.mdAuthorization != nil {
				ctx = metadata.NewIncomingContext(ctx, metadata.MD{
					"authorization": *tc.mdAuthorization,
				})
			} else {
				ctx = metadata.NewIncomingContext(ctx, metadata.MD{})
			}

			ctx, err = v.authenticateMethod(ctx, tc.fullMethod)
			if tc.expectedStatus == nil {
				a.NoError(err)
			} else {
				s, ok := status.FromError(err)
				a.True(ok)
				a.Equal(tc.expectedStatus, s)
			}
		})
	}
}
