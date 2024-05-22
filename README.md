# psi Ψ

![CI](https://github.com/tomyl/psi/actions/workflows/ci.yml/badge.svg?branch=main&event=push)
[![Go Reference](https://pkg.go.dev/badge/github.com/tomyl/psi.svg)](https://pkg.go.dev/github.com/tomyl/psi)

A Go library for writing script-like programs.

**Pre-alpha software**. Expect API breakage, crashes, data loss, silent data corruption etc.

## Rationale

Shell scripts are great when you want to automate the execution of a series of
commands. However, shell scripts tend to be fragile (portability and dependency
issues, typically poor error handling), inefficient and are hard to maintain
as they grow.

Go, on the other hand, excel at these points. However, for trivial tasks
there's more friction to get started with Go. There are various Go libraries
like [pipe](https://github.com/go-pipe/pipe/tree/v2) and
[script](https://github.com/bitfield/script) that try to bridge the gap.

This library tries to distinguish itself by being a bit more idiomatic by
promoting composition and not hiding contexts and error handling too much.

## Installation

```bash
go get github.com/tomyl/psi@latest
```

## Usage

Trivial example that list files recursively and writes the paths to `stdout`:

```go
import "github.com/tomyl/psi"

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
```

See [examples](./examples) for more advanced examples.

## Concepts

* pipeline: A `psi` pipeline is built from taps and sinks.
* tap: Taps are functions that generate a stream of data and passes it element-wise to a sink.
* sink: Sinks are functions that consume an element of a stream. Some sinks actually pass the data to a different sink.
* predicate: Predicates are functions that return `true` or `false` or an error for a single stream element.
* context: All taps accept `*psi.Context` which is passed to all downstream sinks and predicates. If any sink or predicate return an error, the context is cancelled.

## TODO

* Add more taps and support byte streams.
* Write more documentation.
