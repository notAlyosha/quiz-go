package entity

import (
	"database/sql"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

// represents record in database
type Session struct {
	ID            int
	FrontID       uuid.UUID
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
	TeacherID     int
	QuizID        int
	StatusID      int
	ReserveTime   int
	DateTimeStart time.Time
	DateTimeEnd   sql.NullTime
	IsSummaryExam bool
}

func (s *SessionInput) Check() error {
	return validation.ValidateStruct(s,
		validation.Field(s.TeacherID, validation.Required),
		validation.Field(s.QuizID, validation.Required),
		validation.Field(s.StatusID, validation.Required),
		validation.Field(s.ReserveTime, validation.Required),
		validation.Field(s.DateTimeStart, validation.Required, validation.Date(time.Stamp)),
		validation.Field(s.DateTimeEnd, validation.Required, validation.Date(time.Stamp)),
		validation.Field(s.IsSummaryExam, validation.Required),
	)
}
