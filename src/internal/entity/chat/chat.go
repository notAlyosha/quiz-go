package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type Chat struct {
	ID              int
	FrontID         uuid.UUID
	AdministratorID int
	Description     string
	InviteLink      string
	ThemeName       string
	IsGroupChat     bool
	IsDeleted       bool
}

type ChatInput struct {
	FrontID         uuid.UUID
	AdministratorID int
	Description     string
	InviteLink      string
	ThemeName       string
	IsGroupChat     bool
}

func (c *ChatInput) Check() error {
	return validation.ValidateStruct(c,
		validation.Field(c.AdministratorID),
		validation.Field(c.Description, validation.Length(5, 300)),
		validation.Field(c.IsGroupChat, validation.Required),
		validation.Field(c.ThemeName, validation.Required),
	)
}
