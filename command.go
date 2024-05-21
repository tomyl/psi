package psi

import (
	"io"
)

type Cmd struct {
	Path   string
	Args   []string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer

	filters []Filter[string, string]
}

func Command(name string, args ...any) *Cmd {
	cmd := &Cmd{
		Path:    name,
		Args:    make([]string, len(args)),
		filters: make([]Filter[string, string], len(args)),
	}
	for i, arg := range args {
		switch v := arg.(type) {
		case string:
			cmd.Args[i] = v
		case Filter[string, string]:
			cmd.filters[i] = v
		default:
			panic("unsupported type")
		}
	}
	return cmd
}
