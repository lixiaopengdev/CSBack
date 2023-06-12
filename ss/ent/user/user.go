// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldOpenID holds the string denoting the open_id field in the database.
	FieldOpenID = "open_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNickName holds the string denoting the nick_name field in the database.
	FieldNickName = "nick_name"
	// FieldSystemName holds the string denoting the system_name field in the database.
	FieldSystemName = "system_name"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldThumbnailURL holds the string denoting the thumbnail_url field in the database.
	FieldThumbnailURL = "thumbnail_url"
	// FieldSex holds the string denoting the sex field in the database.
	FieldSex = "sex"
	// FieldMobileNo holds the string denoting the mobile_no field in the database.
	FieldMobileNo = "mobile_no"
	// FieldRegionCode holds the string denoting the region_code field in the database.
	FieldRegionCode = "region_code"
	// FieldEmailAddress holds the string denoting the email_address field in the database.
	FieldEmailAddress = "email_address"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// FieldSchoolName holds the string denoting the school_name field in the database.
	FieldSchoolName = "school_name"
	// FieldBio holds the string denoting the bio field in the database.
	FieldBio = "bio"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldIsOnline holds the string denoting the is_online field in the database.
	FieldIsOnline = "is_online"
	// FieldIsShowCollections holds the string denoting the is_show_collections field in the database.
	FieldIsShowCollections = "is_show_collections"
	// FieldIsInvited holds the string denoting the is_invited field in the database.
	FieldIsInvited = "is_invited"
	// FieldNeedPrivacyConfirm holds the string denoting the need_privacy_confirm field in the database.
	FieldNeedPrivacyConfirm = "need_privacy_confirm"
	// FieldCurrentCsFieldID holds the string denoting the current_cs_field_id field in the database.
	FieldCurrentCsFieldID = "current_cs_field_id"
	// FieldCurrentCsFieldName holds the string denoting the current_cs_field_name field in the database.
	FieldCurrentCsFieldName = "current_cs_field_name"
	// FieldPrivateCsFieldID holds the string denoting the private_cs_field_id field in the database.
	FieldPrivateCsFieldID = "private_cs_field_id"
	// FieldPrivateCsFieldName holds the string denoting the private_cs_field_name field in the database.
	FieldPrivateCsFieldName = "private_cs_field_name"
	// FieldRegisterIP holds the string denoting the register_ip field in the database.
	FieldRegisterIP = "register_ip"
	// FieldConstellation holds the string denoting the constellation field in the database.
	FieldConstellation = "constellation"
	// FieldTotalConnections holds the string denoting the total_connections field in the database.
	FieldTotalConnections = "total_connections"
	// EdgeUserAuth holds the string denoting the user_auth edge name in mutations.
	EdgeUserAuth = "user_auth"
	// EdgeCard holds the string denoting the card edge name in mutations.
	EdgeCard = "card"
	// EdgeMessage holds the string denoting the message edge name in mutations.
	EdgeMessage = "message"
	// EdgeDevice holds the string denoting the device edge name in mutations.
	EdgeDevice = "device"
	// EdgeJoinedCsfield holds the string denoting the joined_csfield edge name in mutations.
	EdgeJoinedCsfield = "joined_csfield"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// EdgeHidden holds the string denoting the hidden edge name in mutations.
	EdgeHidden = "hidden"
	// EdgeHistory holds the string denoting the history edge name in mutations.
	EdgeHistory = "history"
	// EdgeToken holds the string denoting the token edge name in mutations.
	EdgeToken = "token"
	// EdgeCreation holds the string denoting the creation edge name in mutations.
	EdgeCreation = "creation"
	// EdgeContact holds the string denoting the contact edge name in mutations.
	EdgeContact = "contact"
	// EdgeSetting holds the string denoting the setting edge name in mutations.
	EdgeSetting = "setting"
	// EdgeNft holds the string denoting the nft edge name in mutations.
	EdgeNft = "nft"
	// EdgeStream holds the string denoting the stream edge name in mutations.
	EdgeStream = "stream"
	// EdgeMask holds the string denoting the mask edge name in mutations.
	EdgeMask = "mask"
	// EdgeTimedew holds the string denoting the timedew edge name in mutations.
	EdgeTimedew = "timedew"
	// EdgeCollection holds the string denoting the collection edge name in mutations.
	EdgeCollection = "collection"
	// EdgeInviteCode holds the string denoting the invite_code edge name in mutations.
	EdgeInviteCode = "invite_code"
	// EdgeFeedback holds the string denoting the feedback edge name in mutations.
	EdgeFeedback = "feedback"
	// EdgeReactionTimedew holds the string denoting the reaction_timedew edge name in mutations.
	EdgeReactionTimedew = "reaction_timedew"
	// EdgeJoins holds the string denoting the joins edge name in mutations.
	EdgeJoins = "joins"
	// EdgeFriendships holds the string denoting the friendships edge name in mutations.
	EdgeFriendships = "friendships"
	// EdgeReactions holds the string denoting the reactions edge name in mutations.
	EdgeReactions = "reactions"
	// Table holds the table name of the user in the database.
	Table = "users"
	// UserAuthTable is the table that holds the user_auth relation/edge.
	UserAuthTable = "user_auths"
	// UserAuthInverseTable is the table name for the User_auth entity.
	// It exists in this package in order to avoid circular dependency with the "user_auth" package.
	UserAuthInverseTable = "user_auths"
	// UserAuthColumn is the table column denoting the user_auth relation/edge.
	UserAuthColumn = "user_id"
	// CardTable is the table that holds the card relation/edge.
	CardTable = "cards"
	// CardInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardInverseTable = "cards"
	// CardColumn is the table column denoting the card relation/edge.
	CardColumn = "user_id"
	// MessageTable is the table that holds the message relation/edge.
	MessageTable = "messages"
	// MessageInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessageInverseTable = "messages"
	// MessageColumn is the table column denoting the message relation/edge.
	MessageColumn = "user_id"
	// DeviceTable is the table that holds the device relation/edge.
	DeviceTable = "devices"
	// DeviceInverseTable is the table name for the Device entity.
	// It exists in this package in order to avoid circular dependency with the "device" package.
	DeviceInverseTable = "devices"
	// DeviceColumn is the table column denoting the device relation/edge.
	DeviceColumn = "user_id"
	// JoinedCsfieldTable is the table that holds the joined_csfield relation/edge. The primary key declared below.
	JoinedCsfieldTable = "joins"
	// JoinedCsfieldInverseTable is the table name for the CSField entity.
	// It exists in this package in order to avoid circular dependency with the "csfield" package.
	JoinedCsfieldInverseTable = "cs_fields"
	// FriendsTable is the table that holds the friends relation/edge. The primary key declared below.
	FriendsTable = "friendships"
	// HiddenTable is the table that holds the hidden relation/edge.
	HiddenTable = "hiddens"
	// HiddenInverseTable is the table name for the Hidden entity.
	// It exists in this package in order to avoid circular dependency with the "hidden" package.
	HiddenInverseTable = "hiddens"
	// HiddenColumn is the table column denoting the hidden relation/edge.
	HiddenColumn = "user_id"
	// HistoryTable is the table that holds the history relation/edge.
	HistoryTable = "user_histories"
	// HistoryInverseTable is the table name for the User_history entity.
	// It exists in this package in order to avoid circular dependency with the "user_history" package.
	HistoryInverseTable = "user_histories"
	// HistoryColumn is the table column denoting the history relation/edge.
	HistoryColumn = "user_id"
	// TokenTable is the table that holds the token relation/edge.
	TokenTable = "agora_tokens"
	// TokenInverseTable is the table name for the Agora_token entity.
	// It exists in this package in order to avoid circular dependency with the "agora_token" package.
	TokenInverseTable = "agora_tokens"
	// TokenColumn is the table column denoting the token relation/edge.
	TokenColumn = "user_id"
	// CreationTable is the table that holds the creation relation/edge.
	CreationTable = "creations"
	// CreationInverseTable is the table name for the Creation entity.
	// It exists in this package in order to avoid circular dependency with the "creation" package.
	CreationInverseTable = "creations"
	// CreationColumn is the table column denoting the creation relation/edge.
	CreationColumn = "user_id"
	// ContactTable is the table that holds the contact relation/edge.
	ContactTable = "contacts"
	// ContactInverseTable is the table name for the Contact entity.
	// It exists in this package in order to avoid circular dependency with the "contact" package.
	ContactInverseTable = "contacts"
	// ContactColumn is the table column denoting the contact relation/edge.
	ContactColumn = "user_id"
	// SettingTable is the table that holds the setting relation/edge.
	SettingTable = "settings"
	// SettingInverseTable is the table name for the Setting entity.
	// It exists in this package in order to avoid circular dependency with the "setting" package.
	SettingInverseTable = "settings"
	// SettingColumn is the table column denoting the setting relation/edge.
	SettingColumn = "user_id"
	// NftTable is the table that holds the nft relation/edge.
	NftTable = "NFTs"
	// NftInverseTable is the table name for the NFT entity.
	// It exists in this package in order to avoid circular dependency with the "nft" package.
	NftInverseTable = "NFTs"
	// NftColumn is the table column denoting the nft relation/edge.
	NftColumn = "user_id"
	// StreamTable is the table that holds the stream relation/edge.
	StreamTable = "streams"
	// StreamInverseTable is the table name for the Stream entity.
	// It exists in this package in order to avoid circular dependency with the "stream" package.
	StreamInverseTable = "streams"
	// StreamColumn is the table column denoting the stream relation/edge.
	StreamColumn = "user_id"
	// MaskTable is the table that holds the mask relation/edge.
	MaskTable = "masks"
	// MaskInverseTable is the table name for the Mask entity.
	// It exists in this package in order to avoid circular dependency with the "mask" package.
	MaskInverseTable = "masks"
	// MaskColumn is the table column denoting the mask relation/edge.
	MaskColumn = "user_id"
	// TimedewTable is the table that holds the timedew relation/edge.
	TimedewTable = "time_dews"
	// TimedewInverseTable is the table name for the TimeDew entity.
	// It exists in this package in order to avoid circular dependency with the "timedew" package.
	TimedewInverseTable = "time_dews"
	// TimedewColumn is the table column denoting the timedew relation/edge.
	TimedewColumn = "user_id"
	// CollectionTable is the table that holds the collection relation/edge.
	CollectionTable = "collections"
	// CollectionInverseTable is the table name for the Collection entity.
	// It exists in this package in order to avoid circular dependency with the "collection" package.
	CollectionInverseTable = "collections"
	// CollectionColumn is the table column denoting the collection relation/edge.
	CollectionColumn = "user_id"
	// InviteCodeTable is the table that holds the invite_code relation/edge.
	InviteCodeTable = "invite_codes"
	// InviteCodeInverseTable is the table name for the Invite_Code entity.
	// It exists in this package in order to avoid circular dependency with the "invite_code" package.
	InviteCodeInverseTable = "invite_codes"
	// InviteCodeColumn is the table column denoting the invite_code relation/edge.
	InviteCodeColumn = "user_id"
	// FeedbackTable is the table that holds the feedback relation/edge.
	FeedbackTable = "feedbacks"
	// FeedbackInverseTable is the table name for the Feedback entity.
	// It exists in this package in order to avoid circular dependency with the "feedback" package.
	FeedbackInverseTable = "feedbacks"
	// FeedbackColumn is the table column denoting the feedback relation/edge.
	FeedbackColumn = "user_id"
	// ReactionTimedewTable is the table that holds the reaction_timedew relation/edge. The primary key declared below.
	ReactionTimedewTable = "reactions"
	// ReactionTimedewInverseTable is the table name for the TimeDew entity.
	// It exists in this package in order to avoid circular dependency with the "timedew" package.
	ReactionTimedewInverseTable = "time_dews"
	// JoinsTable is the table that holds the joins relation/edge.
	JoinsTable = "joins"
	// JoinsInverseTable is the table name for the Join entity.
	// It exists in this package in order to avoid circular dependency with the "join" package.
	JoinsInverseTable = "joins"
	// JoinsColumn is the table column denoting the joins relation/edge.
	JoinsColumn = "user_id"
	// FriendshipsTable is the table that holds the friendships relation/edge.
	FriendshipsTable = "friendships"
	// FriendshipsInverseTable is the table name for the Friendship entity.
	// It exists in this package in order to avoid circular dependency with the "friendship" package.
	FriendshipsInverseTable = "friendships"
	// FriendshipsColumn is the table column denoting the friendships relation/edge.
	FriendshipsColumn = "user_id"
	// ReactionsTable is the table that holds the reactions relation/edge.
	ReactionsTable = "reactions"
	// ReactionsInverseTable is the table name for the Reaction entity.
	// It exists in this package in order to avoid circular dependency with the "reaction" package.
	ReactionsInverseTable = "reactions"
	// ReactionsColumn is the table column denoting the reactions relation/edge.
	ReactionsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldOpenID,
	FieldName,
	FieldNickName,
	FieldSystemName,
	FieldAvatar,
	FieldThumbnailURL,
	FieldSex,
	FieldMobileNo,
	FieldRegionCode,
	FieldEmailAddress,
	FieldBirthday,
	FieldSchoolName,
	FieldBio,
	FieldStatus,
	FieldRole,
	FieldIsOnline,
	FieldIsShowCollections,
	FieldIsInvited,
	FieldNeedPrivacyConfirm,
	FieldCurrentCsFieldID,
	FieldCurrentCsFieldName,
	FieldPrivateCsFieldID,
	FieldPrivateCsFieldName,
	FieldRegisterIP,
	FieldConstellation,
	FieldTotalConnections,
}

var (
	// JoinedCsfieldPrimaryKey and JoinedCsfieldColumn2 are the table columns denoting the
	// primary key for the joined_csfield relation (M2M).
	JoinedCsfieldPrimaryKey = []string{"user_id", "cs_field_id"}
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"user_id", "friend_id"}
	// ReactionTimedewPrimaryKey and ReactionTimedewColumn2 are the table columns denoting the
	// primary key for the reaction_timedew relation (M2M).
	ReactionTimedewPrimaryKey = []string{"time_dew_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultThumbnailURL holds the default value on creation for the "thumbnail_url" field.
	DefaultThumbnailURL string
	// DefaultIsShowCollections holds the default value on creation for the "is_show_collections" field.
	DefaultIsShowCollections bool
	// DefaultIsInvited holds the default value on creation for the "is_invited" field.
	DefaultIsInvited bool
	// DefaultNeedPrivacyConfirm holds the default value on creation for the "need_privacy_confirm" field.
	DefaultNeedPrivacyConfirm bool
)

// Status defines the type for the "status" enum field.
type Status string

// StatusStandalone is the default value of the Status enum.
const DefaultStatus = StatusStandalone

// Status values.
const (
	StatusInfield    Status = "infield"
	StatusStandalone Status = "standalone"
	StatusForbidden  Status = "forbidden"
	StatusMultiMode  Status = "multi_mode"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusInfield, StatusStandalone, StatusForbidden, StatusMultiMode:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for status field: %q", s)
	}
}

// Role defines the type for the "role" enum field.
type Role string

// RoleClient is the default value of the Role enum.
const DefaultRole = RoleClient

// Role values.
const (
	RoleHost   Role = "host"
	RoleAdmin  Role = "admin"
	RoleClient Role = "client"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleHost, RoleAdmin, RoleClient:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}
