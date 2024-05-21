package psi

import (
	"path/filepath"
	"strings"
)

type Filter[I any, O any] func(ctx *Context, arg I) (O, error)

func Source() Filter[string, string] {
	return func(ctx *Context, name string) (string, error) {
		return name, nil
	}
}

func ReplaceExt(sourceExt, targetExt string) Filter[string, string] {
	if sourceExt == targetExt {
		panic("source and target extension must be different")
	}
	return func(ctx *Context, name string) (string, error) {
		if strings.HasSuffix(name, targetExt) {
			// Has target extension already. Ignore.
			return "", ErrSkipFile
		}
		base := name
		if sourceExt != "" {
			ext := filepath.Ext(name)
			if ext != sourceExt {
				// Source doesn't match the source extension. Ignore.
				return "", ErrSkipFile
			}
			base = strings.TrimSuffix(name, ext)
		}
		return base + targetExt, nil
	}
}
