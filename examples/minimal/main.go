package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tomyl/psi"
)

func run(root string) error {
	ctx := psi.NewContext(context.Background())

	// List files recursively and write the paths to stdout.
	if err := psi.WalkFiles(ctx, root, psi.Stdout()); err != nil {
		return err
	}

	// Wait for pipeline to finish. Strictly not needed here because psi.Go()
	// was not used to dispatch goroutines.
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
