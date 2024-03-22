package entity

import (
	"database/sql"
	"time"
)

// represents record in database
type Session struct {
	ID            int
	FrontID       int
	TeacherID     int
	QuizID        int
	StatusID      int
	ReserveTime   int
	DateTimeStart time.Time
	DateTimeEnd   sql.NullTime
	IsSummaryExam bool
	IsDeleted     bool
}

type SessionInput struct {
	DateTimeStart *time.Time
	ReserveTime   *int
	SessionTime   *int
}
