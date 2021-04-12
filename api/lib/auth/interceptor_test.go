package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_JwtVerifier_authenticateMethod(t *testing.T) {
	tt := []struct {
		desc           string
		fullMethod     string
		expectedStatus *status.Status
	}{
		{
			desc:           "tc 1",
			fullMethod:     "/Flux/PricingData",
			expectedStatus: nil,
		},
		{
			desc:           "tc 2",
			fullMethod:     "/Flux/ViewerData",
			expectedStatus: status.New(codes.Unauthenticated, "authentication not set"),
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			a := assert.New(t)
			v := JwtVerifier{}

			ctx := context.Background()
			var err error

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
