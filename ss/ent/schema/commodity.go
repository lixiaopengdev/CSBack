package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Commodity holds the schema definition for the Commodity entity.
type Commodity struct {
	ent.Schema
}

// Fields of the Commodity.
func (Commodity) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("name").Comment("名称"),
	}
}

// Edges of the Commodity.
func (Commodity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("card", Card.Type),
		edge.To("NFT", NFT.Type),
	}
}

func (Commodity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}
