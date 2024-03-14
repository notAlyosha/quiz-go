package types

import "time"

type Session struct {
	ID       int `gorm: "primaryKey;autoIncrement;not null"`
	Quizes   []Quiz
	Statuses []SessionStatus `gorm: "not null"`

	User      User `gorm: "foreignkey:TeacherID`
	TeacherID int

	DateTimeStart time.Time `gorm: "not null"`
	ReserveTime   int

	SessionTime int

	IsReexam  bool
	IsDeleted bool `gorm: "default:false"`
}
