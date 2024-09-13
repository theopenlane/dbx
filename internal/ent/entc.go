//go:build ignore

// See Upstream docs for more details: https://entgo.io/docs/code-gen/#use-entc-as-a-package

package main

import (
	"log"
	"net/http"
	"os"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/theopenlane/entx"
	"github.com/theopenlane/entx/genhooks"
	"github.com/theopenlane/go-turso"
	"github.com/theopenlane/iam/entfga"
	"github.com/theopenlane/iam/fgax"
	"gocloud.dev/secrets"
)

var (
	graphSchemaDir = "./schema/"
	graphQueryDir  = "./query/"
)

func main() {
	xExt, err := entx.NewExtension(
		entx.WithJSONScalar(),
	)
	if err != nil {
		log.Fatalf("creating entx extension: %v", err)
	}

	// Ensure the schema directory exists before running entc.
	_ = os.Mkdir("schema", 0755)

	gqlExt, err := entgql.NewExtension(
		// Tell Ent to generate a GraphQL schema for
		// the Ent schema in a file named ent.graphql.
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("schema/ent.graphql"),
		entgql.WithConfigPath("gqlgen.yml"),
		entgql.WithWhereInputs(true),
		entgql.WithSchemaHook(xExt.GQLSchemaHooks()...),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	if err := entc.Generate("./internal/ent/schema", &gen.Config{
		Target:    "./internal/ent/generated",
		Templates: entgql.AllTemplates,
		Hooks: []gen.Hook{
			genhooks.GenSchema(graphSchemaDir),
			genhooks.GenQuery(graphQueryDir),
		},
		Package: "github.com/theopenlane/dbx/internal/ent/generated",
		Features: []gen.Feature{
			gen.FeatureVersionedMigration,
			gen.FeaturePrivacy,
			gen.FeatureEntQL,
			gen.FeatureNamedEdges,
			gen.FeatureSchemaConfig,
			gen.FeatureIntercept,
		},
	},
		entc.Dependency(
			entc.DependencyType(&secrets.Keeper{}),
		),
		entc.Dependency(
			entc.DependencyName("Authz"),
			entc.DependencyType(fgax.Client{}),
		),
		entc.Dependency(
			entc.DependencyName("Turso"),
			entc.DependencyType(&turso.Client{}),
		),
		entc.Dependency(
			entc.DependencyType(&http.Client{}),
		),
		entc.TemplateDir("./internal/ent/templates"),
		entc.Extensions(
			gqlExt,
			entfga.New(
				entfga.WithSoftDeletes(),
			),
		)); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
