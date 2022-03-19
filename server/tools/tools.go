//go:build tools
// +build tools

package tools

import (
	_ "entgo.io/ent/cmd/ent"
	_ "github.com/99designs/gqlgen"
	_ "github.com/cosmtrek/air@latest"
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate@latest"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/incu6us/goimports-reviser"
	_ "golang.org/x/tools/cmd/godoc"
	_ "mvdan.cc/gofumpt"
)
