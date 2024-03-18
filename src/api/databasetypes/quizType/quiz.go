package quesionType

import "time"

type Quiz struct {
	ID           int
	Name         string
	Description  *string
	SummaryTime  time.Time
	MaxPoints    string
	MaxAddPoints string
	IsDeleted    bool
}
