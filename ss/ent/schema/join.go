package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Join holds the schema definition for the Join entity.
type Join struct {
	ent.Schema
}

func (Join) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

// Fields of the Join.
func (Join) Fields() []ent.Field {
	return []ent.Field{
		field.Time("join_at").
			Default(time.Now),
		field.Time("leave_at").Optional(),
		field.Enum("status").Values("infield", "temp_leaving", "invited", "leave", "host", "admin").Default("host"),
		field.Uint64("user_id"),
		field.Uint64("cs_field_id"),
	}
}

// Edges of the Join.
func (Join) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
		edge.To("cs_field", CSField.Type).
			Required().
			Unique().
			Field("cs_field_id"),
	}
}
