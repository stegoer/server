//go:build ignore

package ent

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

const (
	schemaPath = "./schema"

	gqlgenConfigPath     = "./gqlgen.yml"
	entGraphqlSchemaPath = "./graph/ent.graphqls"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereFilters(true),
		entgql.WithConfigPath(gqlgenConfigPath),
		entgql.WithSchemaPath(entGraphqlSchemaPath),
	)
	if err != nil {
		log.Fatalf("failed creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(ex),
	}

	if err := entc.Generate(schemaPath, &gen.Config{}, opts...); err != nil {
		log.Fatalf("failed running ent codegen: %v", err)
	}
}
