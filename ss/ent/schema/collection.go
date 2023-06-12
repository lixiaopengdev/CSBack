package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Collection holds the schema definition for the Collection entity.
type Collection struct {
	ent.Schema
}

// Fields of the Collection.
func (Collection) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.Enum("type").Values("timedew").Default("timedew").Comment("收藏类型"),
		field.Uint64("item_id").Comment("收藏品对应的id"),
		field.Uint64("user_id"),
	}
}

func (Collection) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "item_id").
			Unique(),
	}
}

// Edges of the Collection.
func (Collection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Field("user_id").
			Ref("collection").Unique().Required(),
	}
}

func (Collection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}
