package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Stream holds the schema definition for the Stream entity.
type Stream struct {
	ent.Schema
}

// Fields of the Stream.
func (Stream) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("name").Comment("名称"),
		field.Enum("type").Values("video", "audio", "custom").Comment("流类型"),
		field.String("stream_url").Optional().Comment("流地址"),
		field.Uint64("user_id"),
	}
}

// Edges of the Stream.
func (Stream) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Field("user_id").
			Ref("stream").Unique().Required(),
	}
}

func (Stream) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}
