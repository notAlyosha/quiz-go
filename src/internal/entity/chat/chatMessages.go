package entity

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

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

type ChatMessagesInput struct {
	ChatID          int
	UserID          int
	MessageDateTime time.Time
	MessageText     string
	FileURL         string
}

func (c *ChatMessagesInput) Check() error {
	return validation.ValidateStruct(c,
		validation.Field(c.ChatID, validation.Required),
		validation.Field(c.UserID, validation.Required),
		validation.Field(c.MessageDateTime, validation.Required),
		validation.Field(c.MessageText, validation.Required, validation.Length(5, 300)),
		validation.Field(c.FileURL, is.URL),
	)
}

func (c *ChatMessagesInput) Save() {

}
