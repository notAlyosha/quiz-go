package entity

import (
	"time"
)

// represents record in database
type UserPasswordHistroy struct {
	ID             int
	UserID         int
	Value          string
	AddingDatetime time.Time
	IsDeleted      bool
}
