package main

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ecshreve/jexplore/ent"
	jexplore "github.com/ecshreve/jexplore/jex"
	"github.com/urfave/cli/v2"
)

// This file will not be regenerated automatically.

func InitClient(ctx *cli.Context) (*ent.Client, error) {
	client, err := ent.Open(
		ctx.String("DatabaseDriver"),
		ctx.String("DatabaseSourceName"),
	)
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(
		context.Background(),
		//migrate.WithGlobalUniqueID(true),
	); err != nil {
		return nil, err
	}
	return client, nil
}

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return jexplore.NewSchema(client)
}
