package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Setting holds the schema definition for the Setting entity.
type Setting struct {
	ent.Schema
}

// Fields of the Setting.
func (Setting) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.Bool("friends_online").Default(false).Optional(),
		field.Bool("time_dew_from_friends").Default(false).Optional(),
		field.Bool("detailed_notification").Default(false).Optional(),
		field.Bool("receive_field_invitation").Default(false).Optional(),
		field.Bool("see_my_location").Default(false).Optional(),
		field.Bool("camera").Default(false).Optional(),
		field.Bool("microphone").Default(false).Optional(),
		field.Bool("health_data").Default(false).Optional(),
		field.Bool("time_dew_location").Default(false).Optional(),
		field.Bool("time_dew_microphone").Default(false).Optional(),
		field.Bool("time_dew_Lora").Default(false).Optional(),
		field.Bool("public_collection").Default(false).Optional(),
		field.Uint64("user_id"),
	}
}

// Edges of the Setting.
func (Setting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Field("user_id").
			Ref("setting").Unique().Required(),
	}
}

func (Setting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}
