package entity

import (
	"database/sql"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Profile struct {
	Id             int
	UserId         int
	Login          string
	Salt           string
	AddingDatetime time.Time
	Name           sql.NullString
	Email          sql.NullString
	Phone          sql.NullString
	IsDeleted      bool
}

type ProfileInput struct {
	Login string
	Name  sql.NullString
	Email string
	Phone string
}

func (p ProfileInput) Check() error {
	return validation.ValidateStruct(p,
		validation.Field(p.Login, validation.Required),
		validation.Field(p.Name, validation.Required),
	)
}
