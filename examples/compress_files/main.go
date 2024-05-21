package main

import (
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/tomyl/psi"
)

func compressor(w io.Writer) io.WriteCloser {
	return gzip.NewWriter(w)
}

func run(root string) error {
	// For every file that isn't compressed already, log its name to stdout and
	// compress it (in parallel, up to GOMAXPROCS goroutines).
	ctx := psi.NewContext(context.Background())
	pipeline := psi.If(psi.HasTarget("", ".gz"), nil, psi.Seq(psi.Stdout(), psi.Go(psi.Transform(compressor))))

	if err := psi.WalkFiles(ctx, root, pipeline); err != nil {
		return err
	}

	return ctx.Wait()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "expected directory as argument")
		os.Exit(1)
	}

	root := os.Args[1]

	if err := run(root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
