//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

const (
	schemaPath = "./schema"

	versionedMigrationsFeatureName = "sql/versioned-migration"

	gqlgenConfigPath  = "../gqlgen.yml"
	graphqlSchemaPath = "../graph/ent.graphqls"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereFilters(true),
		entgql.WithConfigPath(gqlgenConfigPath),
		entgql.WithSchemaPath(graphqlSchemaPath),
	)
	if err != nil {
		log.Panicf("failed creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(ex),
		entc.FeatureNames(versionedMigrationsFeatureName),
	}

	if err := entc.Generate(schemaPath, &gen.Config{}, opts...); err != nil {
		log.Panicf("failed running ent codegen: %v", err)
	}
}
