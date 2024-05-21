package psi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tomyl/psi"
)

func TestPredicateHasTargetWithExt(t *testing.T) {
	type testCase struct {
		name string
		ok   bool
		err  error
	}

	ctx := psi.NewContext(context.Background())
	pred := psi.HasTarget(".txt", ".txt.gz")

	for _, tc := range []testCase{
		{"testdata/stuff/foo.txt", false, nil},
		{"testdata/stuff/bar.txt", true, nil},
		{"testdata/stuff/baz.txt", true, nil},
		{"testdata/stuff/foo.html", false, psi.ErrSkipFile},
		{"testdata/stuff/foo.txt.gz", false, psi.ErrSkipFile},
		{"testdata/stuff/noext1", false, psi.ErrSkipFile},
		{"testdata/stuff/noext2", false, psi.ErrSkipFile},
		{"testdata/stuff/noext3", false, psi.ErrSkipFile},
	} {
		ok, err := pred(ctx, tc.name)
		require.Equal(t, tc.ok, ok)
		require.ErrorIs(t, err, tc.err)
	}
}

func TestPredicateHasTargetWithoutExt(t *testing.T) {
	type testCase struct {
		name string
		ok   bool
		err  error
	}

	ctx := psi.NewContext(context.Background())
	pred := psi.HasTarget("", ".gz")

	for i, tc := range []testCase{
		{"testdata/stuff/foo.txt", false, nil},
		{"testdata/stuff/bar.txt", true, nil},
		{"testdata/stuff/baz.txt", true, nil},
		{"testdata/stuff/foo.html", false, nil},
		{"testdata/stuff/foo.txt.gz", false, psi.ErrSkipFile},
		{"testdata/stuff/noext1", false, nil},
		{"testdata/stuff/noext2", true, nil},
		{"testdata/stuff/noext3", true, nil},
	} {
		ok, err := pred(ctx, tc.name)
		require.Equal(t, tc.ok, ok, "case %d", i)
		require.ErrorIs(t, err, tc.err, "case %d", i)
	}
}
