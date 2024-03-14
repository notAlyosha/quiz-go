package types

import "time"

type ChatMessages struct {
	ID int `gorm: "primaryKey;autoIncrement;not null"`

	Chat Chat
	User User

	MessageDateTime time.Time

	MessageText string

	FileURL         *string
	IsDeletedByUser bool
	IsDeleted       bool `gorm: "default:false"`
}
