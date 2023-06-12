package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Contact holds the schema definition for the Contact entity.
type Contact struct {
	ent.Schema
}

// Fields of the Contact.
func (Contact) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("name").Optional().Comment("名称"),
		field.String("mobile_no").Optional(),
		field.String("email").Optional(),
		field.Uint64("user_id"),
	}
}

// Edges of the Contact.
func (Contact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Field("user_id").
			Ref("contact").Unique().Required(),
	}
}

func (Contact) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
