package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tomyl/psi"
)

const (
	sourceExt = ".jpg"
	targetExt = "_small.jpg"
)

func run(root string) error {
	// For every jpg that doesn't hasn't been resized already, log its name to
	// stdout and resize it (in parallel, up to GOMAXPROCS goroutines).
	ctx := psi.NewContext(context.Background())
	cmd := psi.Command("convert", "-resize", "50%", psi.Source(), psi.ReplaceExt(sourceExt, targetExt))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	resize := psi.Go(psi.Exec(cmd))
	pipeline := psi.If(psi.HasTarget(sourceExt, targetExt), nil, psi.Seq(psi.Stdout(), resize))

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
