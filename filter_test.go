package psi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tomyl/psi"
)

func TestFilterReplaceExt(t *testing.T) {
	type testCase struct {
		src string
		dst string
		err error
	}

	ctx := psi.NewContext(context.Background())
	filter := psi.ReplaceExt(".jpg", "_thumb.jpg")

	for _, tc := range []testCase{
		{"foo.jpg", "foo_thumb.jpg", nil},
		{"foo_thumb.jpg", "", psi.ErrSkipFile},
		{"foo.png", "", psi.ErrSkipFile},
	} {
		dst, err := filter(ctx, tc.src)
		require.Equal(t, tc.dst, dst)
		require.ErrorIs(t, err, tc.err)
	}
}

func TestFilterReplaceExtAdd(t *testing.T) {
	type testCase struct {
		src string
		dst string
		err error
	}

	ctx := psi.NewContext(context.Background())
	filter := psi.ReplaceExt("", ".bak")

	for _, tc := range []testCase{
		{"foo", "foo.bak", nil},
		{"foo.bak", "", psi.ErrSkipFile},
	} {
		dst, err := filter(ctx, tc.src)
		require.Equal(t, tc.dst, dst)
		require.ErrorIs(t, err, tc.err)
	}
}
