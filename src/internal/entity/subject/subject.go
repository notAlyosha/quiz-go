package entity

import validation "github.com/go-ozzo/ozzo-validation"

// represents record in database
type Subject struct {
	ID        int
	Name      string
	IsDeleted bool
}

type SubjectInput struct {
	Name string
}

func (s *SubjectInput) Check() error {
	return validation.ValidateStruct(s,
		validation.Field(s, validation.Required, validation.Length(5, 300)),
	)
}
