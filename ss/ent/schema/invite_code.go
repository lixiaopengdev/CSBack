package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Invite_Code holds the schema definition for the Invite_Code entity.
type Invite_Code struct {
	ent.Schema
}

// Fields of the Invite_Code.
func (Invite_Code) Fields() []ent.Field {
	return []ent.Field{

		// Edges of the User.
		field.Enum("type").Values("register").Default("register").Comment("邀请类型"),
		field.Enum("status").Values("generated", "used").Default("generated").Comment("邀请码状态"),
		field.String("code").Unique().Comment("邀请码"),
		field.Uint64("consumer_id").Optional().Comment("使用邀请码的用户id"),
		field.Uint64("user_id"),
	}
}

// Edges of the Invite_Code.
func (Invite_Code) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Field("user_id").
			Ref("invite_code").Unique().Required(),
	}
}

func (Invite_Code) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}
