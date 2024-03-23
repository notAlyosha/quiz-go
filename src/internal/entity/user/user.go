package entity

import (
	"database/sql"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	uuid "github.com/satori/go.uuid"
)

// represents record in database
type User struct {
	ID        int            `json:"ID"`
	FrontID   string         `json:"front_id"`
	Role      string         `json:"role"`
	LogoURL   sql.NullString `json:"LogoURL"`
	IsDeleted bool           `json:"IsDeleted"`
}

// struct for authentification
type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// data that going to client
type UserResponse struct {
	FrontID uuid.UUID `json:"front_id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Email   string    `json:"email,omitempty"`
	Role    string    `json:"role,omitempty"`
}

type UserCreate struct {
	Role     string `json:"role,omitempty"`
	Login    string
	Password string         `json:"password"`
	Name     sql.NullString `json:"name,omitempty"`
	Email    sql.NullString `json:"email,omitempty"`
	Phone    sql.NullString `json:"phone"`
	LogoURL  sql.NullString `json:"LogoURL"`
}

func (uc UserCreate) Check() error {
	return validation.ValidateStruct(&uc,
		validation.Field(&uc.Email, validation.Required, validation.Length(15, 255), is.Email),
		validation.Field(&uc.Login, validation.Required, validation.Length(15, 255)),
		validation.Field(&uc.LogoURL, validation.Required),
		validation.Field(&uc.Name, validation.Required),
		validation.Field(&uc.Role, validation.Required),
	)
}
