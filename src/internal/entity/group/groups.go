package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	uuid "github.com/satori/go.uuid"
)

// represents record in database
type Group struct {
	ID        int
	FrontID   uuid.UUID
	Name      string
	IsDeleted bool
}

type GroupInput struct {
	FrontID uuid.UUID
	Name    string
}

func (g *GroupInput) Check() error {
	return validation.ValidateStruct(g,
		validation.Field(g.FrontID, validation.Required, is.UUID),
		validation.Field(g.Name, validation.Length(5, 30)),
	)
}
