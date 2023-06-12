package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Bundle holds the schema definition for the Bundle entity.
type Bundle struct {
	ent.Schema
}

// Fields of the Bundle.
func (Bundle) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.Uint64("verionID").Optional().Comment("版本ID"),
		field.String("bundle_url").Optional().Comment("bundle资源地址"),
		field.Enum("status").Values("status1", "status2", "status3").Default("status1").Comment("bundle状态"),
		field.Enum("platform").Values("iPhone", "Android").Default("iPhone").Comment("平台"),
		field.Uint64("mask_id").Optional(),
	}
}

// Edges of the Bundle.
func (Bundle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Mask.Type).Field("mask_id").
			Ref("bundle").Unique(),
	}
}

func (Bundle) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}
