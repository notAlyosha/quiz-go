package types

import "time"

type Session struct {
	ID            int
	SessionStatus SessionStatus
	DateTimeStart time.Time
	ReserveTime   int
	SessionTime   int
	IsReexam      bool
	IsDeleted     bool
}
