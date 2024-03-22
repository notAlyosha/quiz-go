package entity

import (
	"database/sql"

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
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	Login   string
	Role    string  `json:"role,omitempty"`
	LogoURL *string `json:"LogoURL"`
}
