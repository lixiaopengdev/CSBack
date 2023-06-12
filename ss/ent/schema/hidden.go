package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Hidden holds the schema definition for the Hidden entity.
type Hidden struct {
	ent.Schema
}

// Fields of the Hidden.
func (Hidden) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("user_id"),
		field.Uint64("hidden_id"),
	}
}

func (Hidden) Indexes() []ent.Index {
	return []ent.Index{

		index.Fields("user_id", "hidden_id").
			Unique(),
	}
}

// Edges of the Hidden.
func (Hidden) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Field("user_id").
			Ref("hidden").Unique().Required(),
	}
}

func (Hidden) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}
