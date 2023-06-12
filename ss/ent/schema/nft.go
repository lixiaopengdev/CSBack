package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// NFT holds the schema definition for the NFT entity.
type NFT struct {
	ent.Schema
}

// Annotations of the User.
func (NFT) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "NFTs"},
	}
}

// https://neocloud.feishu.cn/wiki/wikcnQaLPSLdDSzAzoS5F54G5Nf
// - 系统内部编号
// - 名称
// - 描述
// - 所有者
// - 创造时间
// - 创造者
// - 媒体文件信息（ 类型、大小、分辨率 ... ）
// - 媒体文件链接
// - TokenID
// - 链上处理状态
// - Json 描述文件链接（ 中心化存储与去中心化存储两个链接 ）
// - 媒体文件链接（ 中心化存储与去中心化存储两个链接 ）
// - 文件的 Pin 状态（ 包括时间和当前状态 ），待定，也可能存在 NFT 服务器端数据库
// - 链上交易 Hash 和区块浏览器链接（ 不保存，用拼接生成 ）
// - 智能合约地址
// - 消耗的 Mint Card 数量
// - 分享链接（ 不保存，用拼接生成 ）

// Fields of the NFT.
func (NFT) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("name").Default("nft name").Comment("名称"),
		field.String("desc").Default("default desc").Comment("描述"),
		field.Uint64("user_id").Optional().Comment("所有者"),
		field.Uint64("creator_id").Optional().Comment("创造者"),
		field.JSON("media_info_json", []string{}).Comment("媒体文件信息（ 类型、大小、分辨率 ... ）"),
		field.String("media_url").Optional().Comment("媒体文件链接-中心化存储"),
		field.String("media_durl").Optional().Comment("媒体文件链接-去中心化存储"),
		field.String("token_id").Optional().Comment("TokenID"),
		field.Enum("status").Values("status1", "status2", "status3", "status4").Default("status1").Comment("链上处理状态"),
		field.String("desc_json_url").Optional().Comment("Json 描述文件链接-中心化存储"),
		field.String("desc_json_durl").Optional().Comment("Json 描述文件链接-去中心化存储"),
		field.Enum("pin_status").Values("status1", "status2", "status3", "status4").Default("status1").Comment("文件的 Pin 状态"),
		field.String("contract_address").Default("").Comment("智能合约地址"),
		field.Uint64("mint_card_num").Default(0).Comment("消耗的 Mint Card 数量"),
	}
}

// Edges of the NFT.
func (NFT) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Field("user_id").
			Ref("nft").Unique(),
	}
}

func (NFT) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}
