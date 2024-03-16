package types

import "time"

type ChatMessages struct {
	ID              int
	ChatID          int
	UserID          int
	MessageDateTime time.Time
	MessageText     string
	FileURL         string
	IsDeletedByUser bool
	IsDeleted       bool
}
