//go:build ignore

package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	// supported ent database drivers
	_ "github.com/lib/pq"           // postgres driver
	_ "github.com/theopenlane/entx" // overlay for sqlite
	"github.com/theopenlane/utils/testutils"
	_ "github.com/tursodatabase/libsql-client-go/libsql" // libsql driver

	"modernc.org/sqlite" // sqlite driver (non-cgo)

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"

	atlas "ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/sqltool"
	"github.com/theopenlane/dbx/internal/ent/generated/migrate"
)

func init() {
	sql.Register("sqlite3", &sqlite.Driver{})
}

func main() {
	ctx := context.Background()

	// Create a local migration directory able to understand Atlas migration file format for replay.
	atlasDir, err := atlas.NewLocalDir("migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

	gooseDirSqlite, err := sqltool.NewGooseDir("migrations-goose-sqlite")
	if err != nil {
		log.Fatalf("failed creating goose migration directory: %v", err)
	}

	gooseDirPG, err := sqltool.NewGooseDir("migrations-goose-postgres")
	if err != nil {
		log.Fatalf("failed creating goose migration directory: %v", err)
	}

	// Migrate diff options.
	baseOpts := []schema.MigrateOption{
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	}

	sqliteOpts := append(baseOpts, schema.WithDialect(dialect.SQLite))
	postgresOpts := append(baseOpts, schema.WithDialect(dialect.Postgres))

	if len(os.Args) != 2 {
		log.Fatalln("migration name is required. Use: 'go run -mod=mod create_migration.go <name>'")
	}

	sqliteDBURI, ok := os.LookupEnv("ATLAS_SQLITE_DB_URI")
	if !ok {
		log.Fatalln("failed to load the ATLAS_SQLITE_DB_URI env var")
	}

	pgDBURI, ok := os.LookupEnv("ATLAS_POSTGRES_DB_URI")
	if !ok {
		log.Fatalln("failed to load the ATLAS_POSTGRES_DB_URI env var")
	}

	maxConnections := 10

	tf, err := testutils.GetPostgresDockerTest(pgDBURI, 5*time.Minute, maxConnections)
	if err != nil {
		log.Fatalf("failed creating postgres test container: %v", err)
	}

	defer testutils.TeardownFixture(tf)

	// Generate migrations using Atlas support for sqlite (note the Ent dialect option passed above).
	atlasOpts := append(baseOpts,
		schema.WithDialect(dialect.Postgres),
		schema.WithDir(atlasDir),
		schema.WithFormatter(atlas.DefaultFormatter),
	)

	if err := migrate.NamedDiff(ctx, tf.URI, os.Args[1], atlasOpts...); err != nil {
		log.Fatalf("failed generating atlas migration file: %v", err)
	}

	// Generate migrations using Goose support for sqlite
	gooseOptsSQLite := append(sqliteOpts, schema.WithDir(gooseDirSqlite))

	if err = migrate.NamedDiff(ctx, sqliteDBURI, os.Args[1], gooseOptsSQLite...); err != nil {
		log.Fatalf("failed generating goose migration file for sqlite: %v", err)
	}

	// Generate migrations using Goose support for postgres
	gooseOptsPG := append(postgresOpts, schema.WithDir(gooseDirPG))

	if err = migrate.NamedDiff(ctx, tf.URI, os.Args[1], gooseOptsPG...); err != nil {
		log.Fatalf("failed generating goose migration file for postgres: %v", err)
	}
}
