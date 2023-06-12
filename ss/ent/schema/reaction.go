package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Reaction holds the schema definition for the Reaction entity.
type Reaction struct {
	ent.Schema
}

// Fields of the Reaction.
func (Reaction) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("isLOL").Default(false),
		field.Bool("isOMG").Default(false),
		field.Bool("isCool").Default(false),
		field.Bool("isNooo").Default(false),
		field.Bool("isDAMN").Default(false),

		field.Uint64("time_dew_id"),
		field.Uint64("user_id"),
	}
}

// Edges of the Reaction.
func (Reaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("timedew", TimeDew.Type).
			Required().
			Unique().
			Field("time_dew_id"),
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
	}
}

func (Reaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}
