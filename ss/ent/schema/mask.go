package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Mask holds the schema definition for the Mask entity.
type Mask struct {
	ent.Schema
}

// Fields of the Mask.
func (Mask) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("name").Comment("面具名称"),
		field.String("desc").Default("description for a mask").Comment("面具描述"),
		field.String("GUID").Comment("GUID"),
		field.String("thumbnail_url").Default("http://192.168.50.193/assets/favicon.ico").Comment("缩略图"),
		field.Enum("status").Values("status1", "status2", "status3").Default("status1").Comment("面具状态"),
		field.Enum("type").Values("type1", "type2", "type3").Default("type1").Comment("面具类型"),
		field.Uint64("user_id").Optional(),
	}
}

// Edges of the Mask.
func (Mask) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bundle", Bundle.Type),
		edge.From("owner", User.Type).Field("user_id").
			Ref("mask").Unique(),
	}
}

func (Mask) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}
