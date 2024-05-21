//go:build tools
// +build tools

package tools

import (
	_ "entgo.io/ent/cmd/ent"
	_ "golang.org/x/vuln/cmd/govulncheck"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
