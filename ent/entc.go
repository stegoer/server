//go:build ignore

package ent

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereFilters(true),
		entgql.WithConfigPath("../gqlgen.yml"),
		entgql.WithSchemaPath("./graph/ent.graphqls"),
	)
	if err != nil {
		log.Fatalf("failed creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(ex),
		entc.TemplateDir("./template"),
	}

	if err := entc.Generate("../ent/schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("failed running ent codegen: %v", err)
	}
}
