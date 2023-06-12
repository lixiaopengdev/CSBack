package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Friendship struct {
	ent.Schema
}

func (Friendship) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

// Fields of the Friendship.
func (Friendship) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").Values("invite", "invited", "established", "rejected", "forbidden").Comment("关系状态"),
		field.Enum("request_type").Values("normal", "good", "close", "custom", "none").Default("none").Comment("请求新建关系类型"),
		field.Enum("curr_type").Values("normal", "good", "close", "custom", "none").Default("none").Comment("当前关系类型"),
		field.Uint64("user_id"),
		field.Uint64("friend_id"),
	}
}

// Edges of the Friendship.
func (Friendship) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
		edge.To("friend", User.Type).
			Required().
			Unique().
			Field("friend_id"),
	}
}
