package types

import "time"

type Quiz struct {
	ID           int `gorm: "primaryKey;autoIncrement;not null"`
	Subject      Subject
	Name         string `gorm: "not null"`
	Description  *string
	SummaryTime  time.Time
	MaxPoints    string `gorm: "not null;check: MaxPoints > 0`
	MaxAddPoints string `gorm: "not null;check: MaxPoints > 0`
	IsDeleted    bool   `gorm: "default:false"`
}
