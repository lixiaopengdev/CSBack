// Code generated by ent, DO NOT EDIT.

package commodity

import (
	"time"
)

const (
	// Label holds the string label denoting the commodity type in the database.
	Label = "commodity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeCard holds the string denoting the card edge name in mutations.
	EdgeCard = "card"
	// EdgeNFT holds the string denoting the nft edge name in mutations.
	EdgeNFT = "NFT"
	// Table holds the table name of the commodity in the database.
	Table = "commodities"
	// CardTable is the table that holds the card relation/edge.
	CardTable = "cards"
	// CardInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardInverseTable = "cards"
	// CardColumn is the table column denoting the card relation/edge.
	CardColumn = "commodity_card"
	// NFTTable is the table that holds the NFT relation/edge.
	NFTTable = "NFTs"
	// NFTInverseTable is the table name for the NFT entity.
	// It exists in this package in order to avoid circular dependency with the "nft" package.
	NFTInverseTable = "NFTs"
	// NFTColumn is the table column denoting the NFT relation/edge.
	NFTColumn = "commodity_nft"
)

// Columns holds all SQL columns for commodity fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
}

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
)
