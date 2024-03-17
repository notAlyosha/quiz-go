package orderType

import "time"

type ReexamOrder struct {
	Id            int
	MainSessionId int
	TeacherId     int
	StudentId     int
	OrderDatetime time.Time
	IsDeleted     bool
}
