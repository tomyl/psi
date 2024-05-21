package psi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tomyl/psi"
)

func TestContextWait(t *testing.T) {
	ctx := psi.NewContext(context.Background())
	require.NoError(t, ctx.Wait())
}
