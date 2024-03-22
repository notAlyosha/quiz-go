package entity

import (
	"database/sql"
	"time"
)

type SessionStudentResult struct {
	ID               int
	SessionID        int
	StudentID        int
	SummaryMark      sql.NullInt32
	AdditionalPoints int
	DateTimeEnd      time.Time
	Is_bought        bool
	IsDeleted        bool
}
