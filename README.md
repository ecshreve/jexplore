# jexplore

generate gqlserver and data explorer ui from ent schema

## [wip]

![jexplore](https://github.com/ecshreve/jexplore/assets/1425775/8e8be54b-0dc1-4387-aee1-88b977fab542)

## where i found stuff

- https://github.com/entkit/entkit
- https://entkit.com/docs/get-started/introduction
- https://github.com/entkit/entkit-demo
- https://github.com/ent/ent
- https://github.com/ent/contrib
- https://atlasgo.io/getting-started/
- https://gqlgen.com/

## how to do stuff

This repo uses [`run`](https://github.com/amonks/run) as a general task executor, tasks are defined in `./tasks.toml`.

To see available tasks run `run -list` from the root of the repo, and/or read the `./tasks.toml` files.

example:
`run dev` starts the gqlserver and webapp in dev mode, and watches for changes to the app source,
rebuilding as needed.

---

## where / what stuff is

`./file.db`
- sqlite3 database

`./gen.go`
- `go generate` script to generate ent schema / gqlschema / webapp

`./ent/`
- https://entgo.io/docs/getting-started
- schema definition and generated ent schema

`./ent/schema/`
- where schema definitions actually live

`./ent/entc.go`
- ent codegen and plugin configuration

`./ent/migrate/`
- https://entgo.io/docs/versioned-migrations/
- migration files and data seeding

`./jsrv/` + `./expl/`
- https://github.com/entkit/entkit
- gqlserver + react ui, generated from ent schema
