package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tomyl/psi"
)

func run(root string) error {
	// Define a pipeline that runs "md5sum" (in parallel, up to GOMAXPROCS
	// goroutines).
	cmd := psi.Command("md5sum", psi.Source())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	pipeline := psi.Go(psi.Exec(cmd))

	// List files recursively and run the pipeline for each file.
	ctx := psi.NewContext(context.Background())
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
