package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("open_id").Optional().Comment("用户openId"),
		field.String("name").Optional().Comment("用户名称"),
		field.String("nick_name").Optional().Comment("用户昵称"),
		field.String("system_name").Optional().Unique().Comment("用户系统名称唯一"),
		field.String("avatar").Optional().Comment("用户虚拟形象"),
		field.String("thumbnail_url").Default("http://thumbnail_placeholder.neoworld.com").Comment("用户虚拟形象"),
		field.String("sex").Optional().Comment("用户性别 w 女 m 男 x 自定义"),
		field.String("mobile_no").Optional().Comment("用户手机号"),
		field.String("region_code").Optional().Comment("用户手机号国家代码"),
		field.String("email_address").Optional().Comment("用户邮箱地址"),
		field.String("birthday").Optional().Comment("生日"),
		field.String("school_name").Optional().Comment("学校名称"),
		field.String("bio").Optional().Comment("个性签名"),
		field.Enum("status").Values("infield", "standalone", "forbidden", "multi_mode").Default("standalone").Comment("用户状态"),
		field.Enum("role").Values("host", "admin", "client").Default("client").Comment("用户身份"),
		field.Bool("is_online").Optional().Comment("当前是否在线"),
		field.Bool("is_show_collections").Default(false).Comment("其他用户能否在主页看到自己收藏的生活流"),
		field.Bool("is_invited").Default(false).Comment("是否使用了邀请码"),
		field.Bool("need_privacy_confirm").Default(true).Comment("是否同意了隐私协议"),
		field.Uint64("current_cs_field_id").Optional().Comment("用户当前所在场ID"),
		field.String("current_cs_field_name").Optional().Comment("用户当前所在场名字"),
		field.Uint64("private_cs_field_id").Optional().Comment("用户自己专属场"),
		field.String("private_cs_field_name").Optional().Comment("用户自己专属场名字"),
		field.String("register_ip").Optional().Comment("用户注册时的IP"),
		field.String("constellation").Optional().Comment("用户星座"),
		field.Int("total_connections").Optional().Comment("用户总建立连接(朋友关系)数"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_auth", User_auth.Type),
		edge.To("card", Card.Type),
		edge.To("message", Message.Type),
		edge.To("device", Device.Type),
		edge.To("joined_csfield", CSField.Type).Through("joins", Join.Type),
		edge.To("friends", User.Type).Through("friendships", Friendship.Type),
		edge.To("hidden", Hidden.Type),
		edge.To("history", User_history.Type),
		edge.To("token", Agora_token.Type),
		edge.To("creation", Creation.Type),
		edge.To("contact", Contact.Type),
		edge.To("setting", Setting.Type),
		edge.To("nft", NFT.Type),
		edge.To("stream", Stream.Type),
		edge.To("mask", Mask.Type),
		edge.To("timedew", TimeDew.Type),
		edge.To("collection", Collection.Type),
		edge.To("invite_code", Invite_Code.Type),
		edge.To("feedback", Feedback.Type),
		edge.From("reaction_timedew", TimeDew.Type).
			Ref("reaction_user").
			Through("reactions", Reaction.Type),
	}
}
