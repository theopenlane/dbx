-- +goose Up
-- create "groups" table
CREATE TABLE "groups" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "name" character varying NOT NULL, "description" character varying NULL, "primary_location" character varying NOT NULL, "locations" jsonb NULL, "token" character varying NULL, "region" character varying NOT NULL DEFAULT 'AMER', PRIMARY KEY ("id"));
-- create index "group_id" to table: "groups"
CREATE UNIQUE INDEX "group_id" ON "groups" ("id");
-- create index "group_name" to table: "groups"
CREATE UNIQUE INDEX "group_name" ON "groups" ("name") WHERE (deleted_at IS NULL);
-- create "databases" table
CREATE TABLE "databases" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "organization_id" character varying NOT NULL, "name" character varying NOT NULL, "geo" character varying NULL, "dsn" character varying NOT NULL, "token" character varying NULL, "status" character varying NOT NULL DEFAULT 'CREATING', "provider" character varying NOT NULL DEFAULT 'LOCAL', "group_id" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "databases_groups_databases" FOREIGN KEY ("group_id") REFERENCES "groups" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "database_id" to table: "databases"
CREATE UNIQUE INDEX "database_id" ON "databases" ("id");
-- create index "database_name" to table: "databases"
CREATE UNIQUE INDEX "database_name" ON "databases" ("name") WHERE (deleted_at IS NULL);
-- create index "database_organization_id" to table: "databases"
CREATE UNIQUE INDEX "database_organization_id" ON "databases" ("organization_id") WHERE (deleted_at IS NULL);

-- +goose Down
-- reverse: create index "database_organization_id" to table: "databases"
DROP INDEX "database_organization_id";
-- reverse: create index "database_name" to table: "databases"
DROP INDEX "database_name";
-- reverse: create index "database_id" to table: "databases"
DROP INDEX "database_id";
-- reverse: create "databases" table
DROP TABLE "databases";
-- reverse: create index "group_name" to table: "groups"
DROP INDEX "group_name";
-- reverse: create index "group_id" to table: "groups"
DROP INDEX "group_id";
-- reverse: create "groups" table
DROP TABLE "groups";
