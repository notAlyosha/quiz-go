package entity

import (
	"time"

	"github.com/google/uuid"
)

type ReexamOrder struct {
	Id            int
	FrontID       uuid.UUID
	MainSessionId int
	TeacherId     int
	StudentId     int
	OrderDatetime time.Time
	IsDeleted     bool
}

type ReexamOrderInput struct {
	OrderDatetime *time.Time
}
