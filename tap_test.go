package psi_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tomyl/psi"
)

func TestTapWalkFiles(t *testing.T) {
	// Compress .txt files, then remove the source. If the .gz file already
	// exists, just remove the source.

	var actions []string

	remove := func(ctx *psi.Context, name string) error {
		actions = append(actions, "remove "+name)
		return nil
	}

	compress := func(ctx *psi.Context, name string) error {
		actions = append(actions, "compress "+name)
		return nil
	}

	var buf bytes.Buffer

	ctx := psi.NewContext(context.Background())
	pipe := psi.If(psi.HasTarget(".txt", ".txt.gz"), remove, psi.Seq(compress, remove, psi.Write(&buf)))
	require.NoError(t, psi.WalkFiles(ctx, "testdata/stuff", pipe))
	require.NoError(t, ctx.Wait())
	require.Equal(t, "testdata/stuff/foo.txt\n", buf.String())
	require.Equal(t, []string{
		"remove testdata/stuff/baz.txt",
		"compress testdata/stuff/foo.txt",
		"remove testdata/stuff/foo.txt",
	}, actions)
}
