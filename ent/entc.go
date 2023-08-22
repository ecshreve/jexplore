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
	jexplore "github.com/ecshreve/jexplore/jex"
	"github.com/entkit/entkit/v2"
)

func main() {
	gqlEx, err := entkit.NewEntgqlExtension(
		entgql.WithConfigPath("./gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./jex/ent.graphql"),
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
		entkit.WithGenerator("expl", entkit.DefaultRefineAdapter),
		entkit.WithGenerator("custom-refine", jexplore.DefaultCustomTemplate, entkit.TargetPath("./expl")),
		entkit.WithGenerator("jsrv", entkit.DefaultServerAdapter),
		entkit.WithGenerator("custom-server", jexplore.DefaultCustomServerTemplate, entkit.TargetPath("./jsrv")),
		entkit.WithGraphqlURL(graphqlUri),
		entkit.WithPrefix("Jepp"),
		entkit.IgnoreUncommittedChanges(),
	)
	if err != nil {
		log.Fatalf("creating entkit extension: %v", err)
	}

	err = entc.Generate("./ent/schema", &gen.Config{
		Package:  "github.com/ecshreve/jexplore/ent",
		Features: []gen.Feature{gen.FeatureVersionedMigration},
	}, entc.Extensions(
		gqlEx,
		entRefine,
	))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
