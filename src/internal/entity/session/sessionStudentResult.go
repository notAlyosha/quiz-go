package entity

import (
	"database/sql"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type SessionStudentResult struct {
	ID               int
	FrontID          uuid.UUID
	SessionID        int
	StudentID        int
	SummaryMark      sql.NullInt32
	AdditionalPoints int
	DateTimeEnd      time.Time
	Is_bought        bool
	IsDeleted        bool
}

type SessionStudentResultInput struct {
	FrontID          uuid.UUID
	SessionID        int
	StudentID        int
	SummaryMark      sql.NullInt32
	AdditionalPoints int
	DateTimeEnd      time.Time
}

func (s SessionStudentResultInput) Check() error {
	return validation.ValidateStruct(s,
		validation.Field(s.SessionID, validation.Required),
		validation.Field(s.StudentID, validation.Required),
		validation.Field(s.SummaryMark, validation.Required),
	)
}
