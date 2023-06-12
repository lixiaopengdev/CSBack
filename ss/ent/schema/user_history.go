package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User_history holds the schema definition for the User_history entity.
type User_history struct {
	ent.Schema
}

// Fields of the User_history.
func (User_history) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.Enum("type").Values("csfield", "action", "unknown").Comment("用户历史类型"),
		field.String("name").Comment("设备特征码"),
		field.String("resource_url").Optional().Comment("资源url"),
		field.Uint64("user_id"),
	}
}

// Edges of the User_history.
func (User_history) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Field("user_id").
			Ref("history").Unique().Required(),
	}
}

func (User_history) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}
