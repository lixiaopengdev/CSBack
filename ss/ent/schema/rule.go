package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Rule holds the schema definition for the Rule entity.
type Rule struct {
	ent.Schema
}

// Fields of the Rule.
func (Rule) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("name").Comment("名称"),
	}
}

// Edges of the Rule.
func (Rule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("used", Card.Type).
			Ref("rule"),
	}
}

func (Rule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
