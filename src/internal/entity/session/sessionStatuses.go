package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type SessionStatus struct {
	ID        int
	Name      string
	IsDeleted bool
}

type SessionStatusInput struct {
	Name string
}

func (s *SessionStatusInput) Check() error {
	return validation.ValidateStruct(s,
		validation.Field(s.Name, validation.Required, validation.Length(5, 100)),
	)
}
