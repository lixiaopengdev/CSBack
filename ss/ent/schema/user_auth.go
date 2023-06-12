package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type User_auth struct {
	ent.Schema
}

// Fields of the User_auth.
func (User_auth) Fields() []ent.Field {
	return []ent.Field{

		// Edges of the User.
		field.Enum("type").Values("oauth2", "email", "mobile", "jwt").Comment("验证方式"),
		field.String("oauth_source").Optional().Comment("oauth来源"),
		field.String("email").Optional().Comment("邮箱").Unique(),
		field.String("mobile_no").Optional().Comment("手机号").Unique(),
		field.String("password").Optional().Sensitive().Comment("密码"),
		field.String("access_token").Optional().Comment("oauth access_token"),
		field.String("oauth_token_type").Optional().Comment("oauth access_token"),
		field.String("oauth_refresh_token").Optional().Comment("oauth refresh_token"),
		field.String("oauth_id").Optional().Comment("oauth 唯一标识"),
		field.Bool("is_finished").Optional().Comment("是否注册完成，填完生日，name"),
		field.Time("oauth_expiry").Optional().Comment("oauth过期时间"),

		field.Uint64("user_id"),
	}
}

// Edges of the User_auth.
func (User_auth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Field("user_id").
			Ref("user_auth").
			Unique().Required(),
	}
}

func (User_auth) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
