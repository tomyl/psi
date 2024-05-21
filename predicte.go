package psi

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type Predicate[T any] func(ctx *Context, arg T) (ok bool, err error)

func HasTarget(sourceExt, targetExt string) Predicate[string] {
	if sourceExt == targetExt {
		panic("source and target extension must be different")
	}
	return func(ctx *Context, name string) (bool, error) {
		if strings.HasSuffix(name, targetExt) {
			// This is the target. Ignore it.
			return false, ErrSkipFile
		}
		base := name
		if sourceExt != "" {
			ext := filepath.Ext(name)
			if ext != sourceExt {
				// Source doesn't match the source extension. Ignore.
				return false, ErrSkipFile
			}
			base = strings.TrimSuffix(name, ext)
		}
		targetName := base + targetExt
		if _, err := os.Stat(targetName); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				// Has no target
				return false, nil
			}
			// Unexpected error
			return false, err
		}
		// Has target
		return true, nil
	}
}
