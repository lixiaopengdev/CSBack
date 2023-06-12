package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// TimeDew holds the schema definition for the TimeDew entity.
type TimeDew struct {
	ent.Schema
}

func (TimeDew) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

// Fields of the TimeDew.
func (TimeDew) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("name").Optional().Comment("timedew名称"),
		field.String("description").Optional().Comment("timedew描述"),
		field.JSON("raw_data", []string{}).Optional().Comment("timedew原始数据"),
		field.Text("speechs").Optional().Comment("timedew语音数据"),
		field.String("place").Optional().Comment("timedew位置数据"),
		field.Text("generated_content").Optional().Comment("timedew生成内容"),
		field.String("prompt_seq").Optional().Comment("timedew生成过程中用到的Prompt方法顺序"),
		field.Text("prompt_seq_full_text").Optional().Comment("timedew生成过程中用到的Prompt全文本"),
		field.String("joined_label").Optional().Comment("timedew生成过程中用到的labels"),
		field.String("pic_url").Optional().Comment("原始图片"),
		field.String("thumbnail_url").Optional().Comment("缩略图"),
		field.String("resource_url").Optional().Comment("资源"),
		field.Enum("status").Values("status1", "status2", "status3").Default("status1").Comment("timedew状态"),
		field.Enum("type").Values("user", "cs_field", "system", "invite").Default("user").Comment("timedew类型"),
		field.Uint64("user_id").Unique(),
		field.Uint64("cs_field_id").Optional(),
		field.Uint64("target_id").Optional(),
		field.Strings("members").Optional(),
	}
}

// Edges of the TimeDew.
func (TimeDew) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Field("user_id").
			Ref("timedew").Unique().Required(),

		edge.To("reaction_user", User.Type).Through("reactions", Reaction.Type),
	}
}
