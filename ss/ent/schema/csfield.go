package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// CSField holds the schema definition for the CSField entity.
type CSField struct {
	ent.Schema
}

// Fields of the CSField.
func (CSField) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("name").Optional().Comment("场名称"),
		field.Enum("status").Values("creating", "opening", "end").Comment("场状态"),
		field.Enum("type").Values("empty", "video", "audio", "text", "mixed").Comment("场类型"),
		field.Enum("mode").Values("single", "multi").Comment("场模式"),
		field.Enum("private_level").Values("public", "ghost", "private").Default("public").Comment("场隐私等级"),
		field.Uint64("user_id"),
		field.Uint64("master_id").Optional(),
	}
}

// Edges of the CSField.
func (CSField) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("joined_user", User.Type).
			Ref("joined_csfield").
			Through("joins", Join.Type),
	}
}

func (CSField) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}
