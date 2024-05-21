package psi

import (
	"errors"
	"io/fs"
	"path/filepath"
)

var ErrSkipFile = errors.New("skip file")

func WalkFiles(ctx *Context, root string, sink Sink[string]) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err := ctx.Err(); err != nil {
			return err
		}
		if err != nil {
			ctx.Cancel(err)
			return err
		}
		if d.IsDir() {
			return nil
		}
		if err := sink(ctx, path); err != nil && !errors.Is(err, ErrSkipFile) {
			ctx.Cancel(err)
			return err
		}
		return nil
	})
}
