# Database Support

## Dependencies

1. [ent](https://entgo.io/) - insane entity mapping tool, definitely not an ORM but kind of an ORM
1. [atlas](https://atlasgo.io/) - Schema generation and migration
1. [entx](https://github.com/theopenlane/entx) - Wrapper to interact with the ent

## Supported Drivers

1. [libsql](https://github.com/tursodatabase/libsql)
1. [sqlite](https://gitlab.com/cznic/sqlite)
1. [postgres](https://github.com/lib/pq)

## Local Development

### Config Examples

#### Libsql

1. This will write to a local file `dbx.db`, already included in `.gitignore`

```yaml
db:
  debug: true
  driver_name: "libsql"
  primary_db_source: "file:dbx.db"
  run_migrations: true
```

#### Sqlite

1. This will write to a local file `dbx.db`, already included in `.gitignore`

```yaml
db:
  debug: true
  driverName: sqlite3
  primaryDbSource: "dbx.db"
  runMigrations: true
```

#### Postgres

1. Postgres is included in `docker/docker-compose-pg.yml` and the same instance can be used for development. The following connection string should work when using `task docker:all:up`

```yaml
db:
  debug: true
  driverName: postgres
  primaryDbSource: "postgres://postgres:password@postgres:5432?sslmode=disable"
  runMigrations: true
```

#### Turso

1. Replace the url with your turso database url and token

```yaml
db:
  debug: true
  driverName: libsql
  primaryDbSource: "https://openlane-golanglemonade.turso.io?authToken=$TURSO_TOKEN"  # set TURSO_TOKEN to value
  runMigrations: false
```
