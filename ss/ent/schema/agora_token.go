package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Agora_token holds the schema definition for the Agora_token entity.
type Agora_token struct {
	ent.Schema
}

// Fields of the Agora_token.
func (Agora_token) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("token").Comment("声网token"),
		field.Uint64("user_id"),
	}
}

// Edges of the Agora_token.
func (Agora_token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Field("user_id").
			Ref("token").Unique().Required(),
	}
}

func (Agora_token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}
