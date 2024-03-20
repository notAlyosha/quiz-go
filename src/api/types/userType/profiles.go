package userType

import "time"

type Profile struct {
	Id             int
	UserId         int
	Login          string
	Salt           string
	AddingDatetime time.Time
	Name           string
	Email          string
	Phone          string
	IsDeleted      bool
}
