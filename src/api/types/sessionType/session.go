package types

import "time"

type Session struct {
	ID            int
	DateTimeStart time.Time
	ReserveTime   int
	SessionTime   int
	IsReexam      bool
	IsDeleted     bool
}

type SessionInput struct {
	DateTimeStart *time.Time
	ReserveTime   *int
	SessionTime   *int
}
