// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// 用户openId
	OpenID string `json:"open_id,omitempty"`
	// 用户名称
	Name string `json:"name,omitempty"`
	// 用户昵称
	NickName string `json:"nick_name,omitempty"`
	// 用户系统名称唯一
	SystemName string `json:"system_name,omitempty"`
	// 用户虚拟形象
	Avatar string `json:"avatar,omitempty"`
	// 用户虚拟形象
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	// 用户性别 w 女 m 男 x 自定义
	Sex string `json:"sex,omitempty"`
	// 用户手机号
	MobileNo string `json:"mobile_no,omitempty"`
	// 用户手机号国家代码
	RegionCode string `json:"region_code,omitempty"`
	// 用户邮箱地址
	EmailAddress string `json:"email_address,omitempty"`
	// 生日
	Birthday string `json:"birthday,omitempty"`
	// 学校名称
	SchoolName string `json:"school_name,omitempty"`
	// 个性签名
	Bio string `json:"bio,omitempty"`
	// 用户状态
	Status user.Status `json:"status,omitempty"`
	// 用户身份
	Role user.Role `json:"role,omitempty"`
	// 当前是否在线
	IsOnline bool `json:"is_online,omitempty"`
	// 其他用户能否在主页看到自己收藏的生活流
	IsShowCollections bool `json:"is_show_collections,omitempty"`
	// 是否使用了邀请码
	IsInvited bool `json:"is_invited,omitempty"`
	// 是否同意了隐私协议
	NeedPrivacyConfirm bool `json:"need_privacy_confirm,omitempty"`
	// 用户当前所在场ID
	CurrentCsFieldID uint64 `json:"current_cs_field_id,omitempty"`
	// 用户当前所在场名字
	CurrentCsFieldName string `json:"current_cs_field_name,omitempty"`
	// 用户自己专属场
	PrivateCsFieldID uint64 `json:"private_cs_field_id,omitempty"`
	// 用户自己专属场名字
	PrivateCsFieldName string `json:"private_cs_field_name,omitempty"`
	// 用户注册时的IP
	RegisterIP string `json:"register_ip,omitempty"`
	// 用户星座
	Constellation string `json:"constellation,omitempty"`
	// 用户总建立连接(朋友关系)数
	TotalConnections int `json:"total_connections,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// UserAuth holds the value of the user_auth edge.
	UserAuth []*User_auth `json:"user_auth,omitempty"`
	// Card holds the value of the card edge.
	Card []*Card `json:"card,omitempty"`
	// Message holds the value of the message edge.
	Message []*Message `json:"message,omitempty"`
	// Device holds the value of the device edge.
	Device []*Device `json:"device,omitempty"`
	// JoinedCsfield holds the value of the joined_csfield edge.
	JoinedCsfield []*CSField `json:"joined_csfield,omitempty"`
	// Friends holds the value of the friends edge.
	Friends []*User `json:"friends,omitempty"`
	// Hidden holds the value of the hidden edge.
	Hidden []*Hidden `json:"hidden,omitempty"`
	// History holds the value of the history edge.
	History []*User_history `json:"history,omitempty"`
	// Token holds the value of the token edge.
	Token []*Agora_token `json:"token,omitempty"`
	// Creation holds the value of the creation edge.
	Creation []*Creation `json:"creation,omitempty"`
	// Contact holds the value of the contact edge.
	Contact []*Contact `json:"contact,omitempty"`
	// Setting holds the value of the setting edge.
	Setting []*Setting `json:"setting,omitempty"`
	// Nft holds the value of the nft edge.
	Nft []*NFT `json:"nft,omitempty"`
	// Stream holds the value of the stream edge.
	Stream []*Stream `json:"stream,omitempty"`
	// Mask holds the value of the mask edge.
	Mask []*Mask `json:"mask,omitempty"`
	// Timedew holds the value of the timedew edge.
	Timedew []*TimeDew `json:"timedew,omitempty"`
	// Collection holds the value of the collection edge.
	Collection []*Collection `json:"collection,omitempty"`
	// InviteCode holds the value of the invite_code edge.
	InviteCode []*Invite_Code `json:"invite_code,omitempty"`
	// Feedback holds the value of the feedback edge.
	Feedback []*Feedback `json:"feedback,omitempty"`
	// ReactionTimedew holds the value of the reaction_timedew edge.
	ReactionTimedew []*TimeDew `json:"reaction_timedew,omitempty"`
	// Joins holds the value of the joins edge.
	Joins []*Join `json:"joins,omitempty"`
	// Friendships holds the value of the friendships edge.
	Friendships []*Friendship `json:"friendships,omitempty"`
	// Reactions holds the value of the reactions edge.
	Reactions []*Reaction `json:"reactions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [23]bool
}

// UserAuthOrErr returns the UserAuth value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserAuthOrErr() ([]*User_auth, error) {
	if e.loadedTypes[0] {
		return e.UserAuth, nil
	}
	return nil, &NotLoadedError{edge: "user_auth"}
}

// CardOrErr returns the Card value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CardOrErr() ([]*Card, error) {
	if e.loadedTypes[1] {
		return e.Card, nil
	}
	return nil, &NotLoadedError{edge: "card"}
}

// MessageOrErr returns the Message value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MessageOrErr() ([]*Message, error) {
	if e.loadedTypes[2] {
		return e.Message, nil
	}
	return nil, &NotLoadedError{edge: "message"}
}

// DeviceOrErr returns the Device value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) DeviceOrErr() ([]*Device, error) {
	if e.loadedTypes[3] {
		return e.Device, nil
	}
	return nil, &NotLoadedError{edge: "device"}
}

// JoinedCsfieldOrErr returns the JoinedCsfield value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) JoinedCsfieldOrErr() ([]*CSField, error) {
	if e.loadedTypes[4] {
		return e.JoinedCsfield, nil
	}
	return nil, &NotLoadedError{edge: "joined_csfield"}
}

// FriendsOrErr returns the Friends value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FriendsOrErr() ([]*User, error) {
	if e.loadedTypes[5] {
		return e.Friends, nil
	}
	return nil, &NotLoadedError{edge: "friends"}
}

// HiddenOrErr returns the Hidden value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) HiddenOrErr() ([]*Hidden, error) {
	if e.loadedTypes[6] {
		return e.Hidden, nil
	}
	return nil, &NotLoadedError{edge: "hidden"}
}

// HistoryOrErr returns the History value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) HistoryOrErr() ([]*User_history, error) {
	if e.loadedTypes[7] {
		return e.History, nil
	}
	return nil, &NotLoadedError{edge: "history"}
}

// TokenOrErr returns the Token value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TokenOrErr() ([]*Agora_token, error) {
	if e.loadedTypes[8] {
		return e.Token, nil
	}
	return nil, &NotLoadedError{edge: "token"}
}

// CreationOrErr returns the Creation value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CreationOrErr() ([]*Creation, error) {
	if e.loadedTypes[9] {
		return e.Creation, nil
	}
	return nil, &NotLoadedError{edge: "creation"}
}

// ContactOrErr returns the Contact value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ContactOrErr() ([]*Contact, error) {
	if e.loadedTypes[10] {
		return e.Contact, nil
	}
	return nil, &NotLoadedError{edge: "contact"}
}

// SettingOrErr returns the Setting value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SettingOrErr() ([]*Setting, error) {
	if e.loadedTypes[11] {
		return e.Setting, nil
	}
	return nil, &NotLoadedError{edge: "setting"}
}

// NftOrErr returns the Nft value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) NftOrErr() ([]*NFT, error) {
	if e.loadedTypes[12] {
		return e.Nft, nil
	}
	return nil, &NotLoadedError{edge: "nft"}
}

// StreamOrErr returns the Stream value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) StreamOrErr() ([]*Stream, error) {
	if e.loadedTypes[13] {
		return e.Stream, nil
	}
	return nil, &NotLoadedError{edge: "stream"}
}

// MaskOrErr returns the Mask value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MaskOrErr() ([]*Mask, error) {
	if e.loadedTypes[14] {
		return e.Mask, nil
	}
	return nil, &NotLoadedError{edge: "mask"}
}

// TimedewOrErr returns the Timedew value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TimedewOrErr() ([]*TimeDew, error) {
	if e.loadedTypes[15] {
		return e.Timedew, nil
	}
	return nil, &NotLoadedError{edge: "timedew"}
}

// CollectionOrErr returns the Collection value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CollectionOrErr() ([]*Collection, error) {
	if e.loadedTypes[16] {
		return e.Collection, nil
	}
	return nil, &NotLoadedError{edge: "collection"}
}

// InviteCodeOrErr returns the InviteCode value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) InviteCodeOrErr() ([]*Invite_Code, error) {
	if e.loadedTypes[17] {
		return e.InviteCode, nil
	}
	return nil, &NotLoadedError{edge: "invite_code"}
}

// FeedbackOrErr returns the Feedback value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FeedbackOrErr() ([]*Feedback, error) {
	if e.loadedTypes[18] {
		return e.Feedback, nil
	}
	return nil, &NotLoadedError{edge: "feedback"}
}

// ReactionTimedewOrErr returns the ReactionTimedew value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReactionTimedewOrErr() ([]*TimeDew, error) {
	if e.loadedTypes[19] {
		return e.ReactionTimedew, nil
	}
	return nil, &NotLoadedError{edge: "reaction_timedew"}
}

// JoinsOrErr returns the Joins value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) JoinsOrErr() ([]*Join, error) {
	if e.loadedTypes[20] {
		return e.Joins, nil
	}
	return nil, &NotLoadedError{edge: "joins"}
}

// FriendshipsOrErr returns the Friendships value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FriendshipsOrErr() ([]*Friendship, error) {
	if e.loadedTypes[21] {
		return e.Friendships, nil
	}
	return nil, &NotLoadedError{edge: "friendships"}
}

// ReactionsOrErr returns the Reactions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReactionsOrErr() ([]*Reaction, error) {
	if e.loadedTypes[22] {
		return e.Reactions, nil
	}
	return nil, &NotLoadedError{edge: "reactions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldIsOnline, user.FieldIsShowCollections, user.FieldIsInvited, user.FieldNeedPrivacyConfirm:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldCurrentCsFieldID, user.FieldPrivateCsFieldID, user.FieldTotalConnections:
			values[i] = new(sql.NullInt64)
		case user.FieldOpenID, user.FieldName, user.FieldNickName, user.FieldSystemName, user.FieldAvatar, user.FieldThumbnailURL, user.FieldSex, user.FieldMobileNo, user.FieldRegionCode, user.FieldEmailAddress, user.FieldBirthday, user.FieldSchoolName, user.FieldBio, user.FieldStatus, user.FieldRole, user.FieldCurrentCsFieldName, user.FieldPrivateCsFieldName, user.FieldRegisterIP, user.FieldConstellation:
			values[i] = new(sql.NullString)
		case user.FieldCreateTime, user.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = uint64(value.Int64)
		case user.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				u.CreateTime = value.Time
			}
		case user.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				u.UpdateTime = value.Time
			}
		case user.FieldOpenID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field open_id", values[i])
			} else if value.Valid {
				u.OpenID = value.String
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldNickName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nick_name", values[i])
			} else if value.Valid {
				u.NickName = value.String
			}
		case user.FieldSystemName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field system_name", values[i])
			} else if value.Valid {
				u.SystemName = value.String
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = value.String
			}
		case user.FieldThumbnailURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field thumbnail_url", values[i])
			} else if value.Valid {
				u.ThumbnailURL = value.String
			}
		case user.FieldSex:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sex", values[i])
			} else if value.Valid {
				u.Sex = value.String
			}
		case user.FieldMobileNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile_no", values[i])
			} else if value.Valid {
				u.MobileNo = value.String
			}
		case user.FieldRegionCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field region_code", values[i])
			} else if value.Valid {
				u.RegionCode = value.String
			}
		case user.FieldEmailAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email_address", values[i])
			} else if value.Valid {
				u.EmailAddress = value.String
			}
		case user.FieldBirthday:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field birthday", values[i])
			} else if value.Valid {
				u.Birthday = value.String
			}
		case user.FieldSchoolName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field school_name", values[i])
			} else if value.Valid {
				u.SchoolName = value.String
			}
		case user.FieldBio:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bio", values[i])
			} else if value.Valid {
				u.Bio = value.String
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = user.Status(value.String)
			}
		case user.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				u.Role = user.Role(value.String)
			}
		case user.FieldIsOnline:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_online", values[i])
			} else if value.Valid {
				u.IsOnline = value.Bool
			}
		case user.FieldIsShowCollections:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_show_collections", values[i])
			} else if value.Valid {
				u.IsShowCollections = value.Bool
			}
		case user.FieldIsInvited:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_invited", values[i])
			} else if value.Valid {
				u.IsInvited = value.Bool
			}
		case user.FieldNeedPrivacyConfirm:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field need_privacy_confirm", values[i])
			} else if value.Valid {
				u.NeedPrivacyConfirm = value.Bool
			}
		case user.FieldCurrentCsFieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field current_cs_field_id", values[i])
			} else if value.Valid {
				u.CurrentCsFieldID = uint64(value.Int64)
			}
		case user.FieldCurrentCsFieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field current_cs_field_name", values[i])
			} else if value.Valid {
				u.CurrentCsFieldName = value.String
			}
		case user.FieldPrivateCsFieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field private_cs_field_id", values[i])
			} else if value.Valid {
				u.PrivateCsFieldID = uint64(value.Int64)
			}
		case user.FieldPrivateCsFieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field private_cs_field_name", values[i])
			} else if value.Valid {
				u.PrivateCsFieldName = value.String
			}
		case user.FieldRegisterIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field register_ip", values[i])
			} else if value.Valid {
				u.RegisterIP = value.String
			}
		case user.FieldConstellation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field constellation", values[i])
			} else if value.Valid {
				u.Constellation = value.String
			}
		case user.FieldTotalConnections:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_connections", values[i])
			} else if value.Valid {
				u.TotalConnections = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUserAuth queries the "user_auth" edge of the User entity.
func (u *User) QueryUserAuth() *UserAuthQuery {
	return (&UserClient{config: u.config}).QueryUserAuth(u)
}

// QueryCard queries the "card" edge of the User entity.
func (u *User) QueryCard() *CardQuery {
	return (&UserClient{config: u.config}).QueryCard(u)
}

// QueryMessage queries the "message" edge of the User entity.
func (u *User) QueryMessage() *MessageQuery {
	return (&UserClient{config: u.config}).QueryMessage(u)
}

// QueryDevice queries the "device" edge of the User entity.
func (u *User) QueryDevice() *DeviceQuery {
	return (&UserClient{config: u.config}).QueryDevice(u)
}

// QueryJoinedCsfield queries the "joined_csfield" edge of the User entity.
func (u *User) QueryJoinedCsfield() *CSFieldQuery {
	return (&UserClient{config: u.config}).QueryJoinedCsfield(u)
}

// QueryFriends queries the "friends" edge of the User entity.
func (u *User) QueryFriends() *UserQuery {
	return (&UserClient{config: u.config}).QueryFriends(u)
}

// QueryHidden queries the "hidden" edge of the User entity.
func (u *User) QueryHidden() *HiddenQuery {
	return (&UserClient{config: u.config}).QueryHidden(u)
}

// QueryHistory queries the "history" edge of the User entity.
func (u *User) QueryHistory() *UserHistoryQuery {
	return (&UserClient{config: u.config}).QueryHistory(u)
}

// QueryToken queries the "token" edge of the User entity.
func (u *User) QueryToken() *AgoraTokenQuery {
	return (&UserClient{config: u.config}).QueryToken(u)
}

// QueryCreation queries the "creation" edge of the User entity.
func (u *User) QueryCreation() *CreationQuery {
	return (&UserClient{config: u.config}).QueryCreation(u)
}

// QueryContact queries the "contact" edge of the User entity.
func (u *User) QueryContact() *ContactQuery {
	return (&UserClient{config: u.config}).QueryContact(u)
}

// QuerySetting queries the "setting" edge of the User entity.
func (u *User) QuerySetting() *SettingQuery {
	return (&UserClient{config: u.config}).QuerySetting(u)
}

// QueryNft queries the "nft" edge of the User entity.
func (u *User) QueryNft() *NFTQuery {
	return (&UserClient{config: u.config}).QueryNft(u)
}

// QueryStream queries the "stream" edge of the User entity.
func (u *User) QueryStream() *StreamQuery {
	return (&UserClient{config: u.config}).QueryStream(u)
}

// QueryMask queries the "mask" edge of the User entity.
func (u *User) QueryMask() *MaskQuery {
	return (&UserClient{config: u.config}).QueryMask(u)
}

// QueryTimedew queries the "timedew" edge of the User entity.
func (u *User) QueryTimedew() *TimeDewQuery {
	return (&UserClient{config: u.config}).QueryTimedew(u)
}

// QueryCollection queries the "collection" edge of the User entity.
func (u *User) QueryCollection() *CollectionQuery {
	return (&UserClient{config: u.config}).QueryCollection(u)
}

// QueryInviteCode queries the "invite_code" edge of the User entity.
func (u *User) QueryInviteCode() *InviteCodeQuery {
	return (&UserClient{config: u.config}).QueryInviteCode(u)
}

// QueryFeedback queries the "feedback" edge of the User entity.
func (u *User) QueryFeedback() *FeedbackQuery {
	return (&UserClient{config: u.config}).QueryFeedback(u)
}

// QueryReactionTimedew queries the "reaction_timedew" edge of the User entity.
func (u *User) QueryReactionTimedew() *TimeDewQuery {
	return (&UserClient{config: u.config}).QueryReactionTimedew(u)
}

// QueryJoins queries the "joins" edge of the User entity.
func (u *User) QueryJoins() *JoinQuery {
	return (&UserClient{config: u.config}).QueryJoins(u)
}

// QueryFriendships queries the "friendships" edge of the User entity.
func (u *User) QueryFriendships() *FriendshipQuery {
	return (&UserClient{config: u.config}).QueryFriendships(u)
}

// QueryReactions queries the "reactions" edge of the User entity.
func (u *User) QueryReactions() *ReactionQuery {
	return (&UserClient{config: u.config}).QueryReactions(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("create_time=")
	builder.WriteString(u.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(u.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("open_id=")
	builder.WriteString(u.OpenID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("nick_name=")
	builder.WriteString(u.NickName)
	builder.WriteString(", ")
	builder.WriteString("system_name=")
	builder.WriteString(u.SystemName)
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(u.Avatar)
	builder.WriteString(", ")
	builder.WriteString("thumbnail_url=")
	builder.WriteString(u.ThumbnailURL)
	builder.WriteString(", ")
	builder.WriteString("sex=")
	builder.WriteString(u.Sex)
	builder.WriteString(", ")
	builder.WriteString("mobile_no=")
	builder.WriteString(u.MobileNo)
	builder.WriteString(", ")
	builder.WriteString("region_code=")
	builder.WriteString(u.RegionCode)
	builder.WriteString(", ")
	builder.WriteString("email_address=")
	builder.WriteString(u.EmailAddress)
	builder.WriteString(", ")
	builder.WriteString("birthday=")
	builder.WriteString(u.Birthday)
	builder.WriteString(", ")
	builder.WriteString("school_name=")
	builder.WriteString(u.SchoolName)
	builder.WriteString(", ")
	builder.WriteString("bio=")
	builder.WriteString(u.Bio)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", u.Role))
	builder.WriteString(", ")
	builder.WriteString("is_online=")
	builder.WriteString(fmt.Sprintf("%v", u.IsOnline))
	builder.WriteString(", ")
	builder.WriteString("is_show_collections=")
	builder.WriteString(fmt.Sprintf("%v", u.IsShowCollections))
	builder.WriteString(", ")
	builder.WriteString("is_invited=")
	builder.WriteString(fmt.Sprintf("%v", u.IsInvited))
	builder.WriteString(", ")
	builder.WriteString("need_privacy_confirm=")
	builder.WriteString(fmt.Sprintf("%v", u.NeedPrivacyConfirm))
	builder.WriteString(", ")
	builder.WriteString("current_cs_field_id=")
	builder.WriteString(fmt.Sprintf("%v", u.CurrentCsFieldID))
	builder.WriteString(", ")
	builder.WriteString("current_cs_field_name=")
	builder.WriteString(u.CurrentCsFieldName)
	builder.WriteString(", ")
	builder.WriteString("private_cs_field_id=")
	builder.WriteString(fmt.Sprintf("%v", u.PrivateCsFieldID))
	builder.WriteString(", ")
	builder.WriteString("private_cs_field_name=")
	builder.WriteString(u.PrivateCsFieldName)
	builder.WriteString(", ")
	builder.WriteString("register_ip=")
	builder.WriteString(u.RegisterIP)
	builder.WriteString(", ")
	builder.WriteString("constellation=")
	builder.WriteString(u.Constellation)
	builder.WriteString(", ")
	builder.WriteString("total_connections=")
	builder.WriteString(fmt.Sprintf("%v", u.TotalConnections))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}