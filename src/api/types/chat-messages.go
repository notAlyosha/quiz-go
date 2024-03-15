package types

import "time"

type ChatMessages struct {
	ID int

	Chat   Chat
	ChatID int

	User   User
	UserID int

	MessageDateTime time.Time

	MessageText string

	FileURL         *string
	IsDeletedByUser bool
	IsDeleted       bool
}
