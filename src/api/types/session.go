package types

import "time"

type Session struct {
	ID              int
	Quizes          []Quiz
	SessionStatus   SessionStatus
	SessionStudents []SessionStudent
	User            []User
	DateTimeStart   time.Time
	ReserveTime     int
	SessionTime     int
	IsReexam        bool
	IsDeleted       bool
}
