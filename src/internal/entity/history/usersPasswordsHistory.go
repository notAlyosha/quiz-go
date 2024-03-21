package entity

import (
	"time"
)

type Profile struct {
	Id             int
	UserId         int
	Value          string
	AddingDatetime time.Time
	IsDeleted      bool
}
