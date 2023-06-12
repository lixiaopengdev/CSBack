package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

func (Card) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("name").Comment("卡片名称"),
		field.String("description").Optional().Comment("卡片说明"),
		field.String("pic_url").Optional().Comment("卡片图"),
		field.String("thumbnail_url").Optional().Comment("缩略图"),
		field.String("resource_url").Optional().Comment("资源"),
		field.Enum("status").Values("status1", "status2", "status3").Default("status1").Comment("卡片状态"),
		field.Enum("type").Values("type1", "type2", "type3").Default("type1").Comment("卡片类型"),
		field.JSON("script", []string{}).
			Optional(),
		field.Text("script_raw").Optional(),
		field.String("script_url").Optional(),
		field.Uint64("user_id").Optional(),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rule", Rule.Type),
		edge.From("owner", User.Type).Field("user_id").
			Ref("card").Unique(),
	}
}
