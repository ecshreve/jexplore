// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build ignore
// +build ignore

package main

import (
	"log"
	"os"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"github.com/entkit/entkit/v2"
)

func main() {
	// The codegen is executed from internal/todo/gen.go.
	// So the path for the config file, ent schema, and the GQL schema
	// starts from internal/todo.
	gqlEx, err := entkit.NewEntgqlExtension(
		entgql.WithConfigPath("./gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./ent.graphql"),
		entgql.WithWhereInputs(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	graphqlUri := os.Getenv("GRAPHQL_URI")
	if graphqlUri == "" {
		graphqlUri = "http://localhost/query"
	}

	entRefine, err := entkit.NewExtension(
		//entkit.WithGenerator(filepath.Join("typescript-project"), entkit.DefaultTypescriptAdapter),
		entkit.WithGenerator("refine-project", entkit.DefaultRefineAdapter),
		//entkit.WithGenerator(
		//	"other-refine-project",
		//	entkit.DefaultRefineAdapter,
		//	entkit.TargetPath(filepath.Join("other-refine-project-root/project")),
		//),
		entkit.WithGenerator("my-server", entkit.DefaultServerAdapter),
		entkit.WithGraphqlURL(graphqlUri),
		entkit.WithPrefix("Jepp"),
		entkit.IgnoreUncommittedChanges(),
	)
	if err != nil {
		log.Fatalf("creating entkit extension: %v", err)
	}

	err = entc.Generate("./ent/schema", &gen.Config{
		Package: "github.com/ecshreve/jexplore/ent",
		IDType: &field.TypeInfo{
			Type: field.TypeString,
		},
		Features: []gen.Feature{gen.FeatureVersionedMigration},
	}, entc.Extensions(
		gqlEx,
		entRefine,
	))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
