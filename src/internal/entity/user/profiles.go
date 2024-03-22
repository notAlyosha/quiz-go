package entity

import (
	"database/sql"
	"time"
)

type Profile struct {
	Id             int
	UserId         int
	Login          string
	Salt           string
	AddingDatetime time.Time
	Name           sql.NullString
	Email          string
	Phone          string
	IsDeleted      bool
}
