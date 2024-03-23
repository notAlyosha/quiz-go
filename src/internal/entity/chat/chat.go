package entity

import validation "github.com/go-ozzo/ozzo-validation"

type Chat struct {
	ID              int
	AdministratorID int
	Description     string
	InviteLink      string
	ThemeName       string
	IsGroupChat     bool
	IsDeleted       bool
}

type ChatInput struct {
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
