package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JwtVerifier_authenticateMethod(t *testing.T) {
	a := assert.New(t)
	v := JwtVerifier{}

	ctx := context.Background()
	var err error

	ctx, err = v.authenticateMethod(ctx, "/Flux/PricingData")
	a.NoError(err)
}
