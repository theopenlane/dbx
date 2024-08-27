package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/theopenlane/entx"
	emixin "github.com/theopenlane/entx/mixin"

	"github.com/theopenlane/dbx/internal/ent/hooks"
	"github.com/theopenlane/dbx/internal/ent/mixin"
	"github.com/theopenlane/dbx/pkg/enums"
)

// Group holds the schema definition for the Group entity
type Group struct {
	ent.Schema
}

// Fields of the Group
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("the name of the group in turso").
			NotEmpty(),
		field.String("description").
			Comment("the description of the group").
			Optional(),
		field.String("primary_location").
			Comment("the primary of the group").
			NotEmpty(),
		field.Strings("locations").
			Comment("the replica locations of the group").
			Optional(),
		field.String("token").
			Sensitive().
			Comment("the auth token used to connect to the group").
			Optional(), // optional because the token is created after the group is created
		field.Enum("region").
			GoType(enums.Region("")).
			Comment("region the group").
			Default(string(enums.Amer)),
	}
}

// Mixin of the Group
func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		emixin.IDMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Hooks of the Group
func (Group) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookGroupCreate(),
		hooks.HookGroupUpdate(),
		hooks.HookGroupDelete(),
	}
}

// Edges of the Database
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("databases", Database.Type).
			Annotations(entx.CascadeAnnotationField("Group")),
	}
}

func (Group) Indexes() []ent.Index {
	return []ent.Index{
		// names should be unique, but ignore deleted names
		index.Fields("name").
			Unique().Annotations(entsql.IndexWhere("deleted_at is NULL")),
	}
}

// Annotations of the Group
func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
