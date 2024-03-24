package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

// represents record in database
type Subject struct {
	ID        int
	FrontID   uuid.UUID
	Name      string
	IsDeleted bool
}

type SubjectInput struct {
	FrontID uuid.UUID
	Name    string
}

func (s *SubjectInput) Check() error {
	return validation.ValidateStruct(s,
		validation.Field(s.Name, validation.Required, validation.Length(5, 300)),
	)
}
