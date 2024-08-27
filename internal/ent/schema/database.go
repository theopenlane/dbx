package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	emixin "github.com/theopenlane/entx/mixin"

	"github.com/theopenlane/dbx/internal/ent/hooks"
	"github.com/theopenlane/dbx/internal/ent/mixin"
	"github.com/theopenlane/dbx/pkg/enums"
)

// Database holds the example schema definition for the Database entity
type Database struct {
	ent.Schema
}

// Fields of the Database
func (Database) Fields() []ent.Field {
	return []ent.Field{
		field.String("organization_id").
			Comment("the ID of the organization").
			NotEmpty(),
		field.String("name").
			Comment("the name to the database").
			NotEmpty(),
		field.String("geo").
			Comment("the geo location of the database").
			Optional(),
		field.String("dsn").
			Comment("the DSN to the database").
			NotEmpty(),
		field.String("group_id").
			Comment("the ID of the group"),
		field.String("token").
			Sensitive().
			Comment("the auth token used to connect to the database").
			Optional(), // optional because the token is created after the database is created
		field.Enum("status").
			GoType(enums.DatabaseStatus("")).
			Comment("status of the database").
			Default(string(enums.Creating)),
		field.Enum("provider").
			GoType(enums.DatabaseProvider("")).
			Comment("provider of the database").
			Default(string(enums.Local)),
	}
}

// Indexes of the Database
func (Database) Indexes() []ent.Index {
	return []ent.Index{
		// organization_id should be unique, this will also create a unique name
		index.Fields("organization_id").
			Unique().Annotations(entsql.IndexWhere("deleted_at is NULL")),
		index.Fields("name").
			Unique().Annotations(entsql.IndexWhere("deleted_at is NULL")),
	}
}

// Mixin of the Database
func (Database) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		mixin.SoftDeleteMixin{},
		emixin.IDMixin{},
	}
}

// Edges of the Database
func (Database) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).
			Field("group_id").
			Required().
			Unique().
			Ref("databases"),
	}
}

// Annotations of the Database
func (Database) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// Hooks of the Database
func (Database) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookCreateDatabase(),
		hooks.HookDatabaseDelete(),
	}
}
