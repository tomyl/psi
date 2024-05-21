package psi_test

import (
	"bytes"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tomyl/psi"
)

func TestSinkWrite(t *testing.T) {
	var buf bytes.Buffer
	ctx := psi.NewContext(context.Background())
	sink := psi.Write(&buf)
	require.NoError(t, sink(ctx, "foo"))
	require.Equal(t, "foo\n", buf.String())
}

func TestSinkCommand(t *testing.T) {
	var buf bytes.Buffer
	ctx := psi.NewContext(context.Background())
	cmd := psi.Command("echo", "foo", psi.Source())
	cmd.Stdout = &buf
	sink := psi.Exec(cmd)
	require.NoError(t, sink(ctx, "bar"))
	require.Equal(t, "foo bar\n", buf.String())
}

func TestSinkGo(t *testing.T) {
	err := errors.New("error 1")

	ctx := psi.NewContext(context.Background())
	sink := psi.Go(func(ctx *psi.Context, name string) error {
		time.Sleep(100 * time.Millisecond)
		return err
	})

	require.NoError(t, sink(ctx, "foo"))
	require.ErrorIs(t, ctx.Wait(), err)
}
