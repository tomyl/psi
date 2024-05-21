package psi

import (
	"io"
	"os"
	"os/exec"
)

type Sink[T any] func(ctx *Context, arg T) error

func Write(w io.Writer) Sink[string] {
	return func(ctx *Context, line string) error {
		_, err := w.Write([]byte(line + "\n"))
		return err
	}
}

func Stdout() Sink[string] {
	return Write(os.Stdout)
}

func Exec(cmd *Cmd) Sink[string] {
	return func(ctx *Context, name string) (err error) {
		for i, filter := range cmd.filters {
			if filter != nil {
				arg, err := filter(ctx, name)
				if err != nil {
					return err
				}
				cmd.Args[i] = arg
			}
		}
		ecmd := exec.CommandContext(ctx, cmd.Path, cmd.Args...)
		ecmd.Stdin = cmd.Stdin
		ecmd.Stdout = cmd.Stdout
		ecmd.Stderr = cmd.Stderr
		return ecmd.Run()
	}
}

func Transform(transformer func(io.Writer) io.WriteCloser) Sink[string] {
	return func(ctx *Context, name string) (err error) {
		src, err := os.Open(name)
		if err != nil {
			return err
		}

		defer src.Close()

		archiveName := name + ".gz"

		defer func() {
			if err != nil {
				_ = os.Remove(archiveName)
			}
		}()

		dst, err := os.Create(archiveName)
		if err != nil {
			return err
		}

		defer dst.Close()

		w := transformer(dst)

		if _, err := io.Copy(w, src); err != nil {
			return err
		}

		if err := w.Close(); err != nil {
			return err
		}

		return dst.Close()
	}
}

func If[T any](pred Predicate[T], ifSink, elseSink Sink[T]) Sink[T] {
	return func(ctx *Context, arg T) error {
		ok, err := pred(ctx, arg)
		if err != nil {
			return err
		}
		if ok {
			if ifSink != nil {
				return ifSink(ctx, arg)
			}
			return nil
		}
		if elseSink != nil {
			return elseSink(ctx, arg)
		}
		return nil
	}
}

func Go[T any](sink Sink[T]) Sink[T] {
	return func(ctx *Context, arg T) error {
		ctx.Go(func() error {
			return sink(ctx, arg)
		})
		return nil
	}
}

func Seq[T any](sinks ...Sink[T]) Sink[T] {
	return func(ctx *Context, arg T) error {
		for _, sink := range sinks {
			if err := sink(ctx, arg); err != nil {
				return err
			}
		}
		return nil
	}
}
