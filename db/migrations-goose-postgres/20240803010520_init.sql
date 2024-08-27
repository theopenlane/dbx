-- +goose Up
-- create "groups" table
CREATE TABLE "groups" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "name" character varying NOT NULL, "description" character varying NULL, "primary_location" character varying NOT NULL, "locations" jsonb NULL, "token" character varying NULL, "region" character varying NOT NULL DEFAULT 'AMER', PRIMARY KEY ("id"));
-- create index "group_name" to table: "groups"
CREATE UNIQUE INDEX "group_name" ON "groups" ("name") WHERE (deleted_at IS NULL);
-- create index "groups_mapping_id_key" to table: "groups"
CREATE UNIQUE INDEX "groups_mapping_id_key" ON "groups" ("mapping_id");
-- create "databases" table
CREATE TABLE "databases" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "mapping_id" character varying NOT NULL, "organization_id" character varying NOT NULL, "name" character varying NOT NULL, "geo" character varying NULL, "dsn" character varying NOT NULL, "token" character varying NULL, "status" character varying NOT NULL DEFAULT 'CREATING', "provider" character varying NOT NULL DEFAULT 'LOCAL', "group_id" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "databases_groups_databases" FOREIGN KEY ("group_id") REFERENCES "groups" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "database_name" to table: "databases"
CREATE UNIQUE INDEX "database_name" ON "databases" ("name") WHERE (deleted_at IS NULL);
-- create index "database_organization_id" to table: "databases"
CREATE UNIQUE INDEX "database_organization_id" ON "databases" ("organization_id") WHERE (deleted_at IS NULL);
-- create index "databases_mapping_id_key" to table: "databases"
CREATE UNIQUE INDEX "databases_mapping_id_key" ON "databases" ("mapping_id");

-- +goose Down
-- reverse: create index "databases_mapping_id_key" to table: "databases"
DROP INDEX "databases_mapping_id_key";
-- reverse: create index "database_organization_id" to table: "databases"
DROP INDEX "database_organization_id";
-- reverse: create index "database_name" to table: "databases"
DROP INDEX "database_name";
-- reverse: create "databases" table
DROP TABLE "databases";
-- reverse: create index "groups_mapping_id_key" to table: "groups"
DROP INDEX "groups_mapping_id_key";
-- reverse: create index "group_name" to table: "groups"
DROP INDEX "group_name";
-- reverse: create "groups" table
DROP TABLE "groups";
