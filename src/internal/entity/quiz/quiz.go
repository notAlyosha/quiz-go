package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// represents record in database
type Quiz struct {
	ID           int
	FrontID      uuid.UUID
	Name         string
	Description  sql.NullString
	SummaryTime  time.Time
	MaxPoints    string
	MaxAddPoints string
	IsDeleted    bool
}

type QuizInput struct {
	Name         string
	Description  string
	MaxPoints    int
	MaxAddPoints int
}
