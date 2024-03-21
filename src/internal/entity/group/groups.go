package entity

import uuid "github.com/satori/go.uuid"

type Group struct {
	ID        int
	FrontID   uuid.UUID
	Name      string
	IsDeleted bool
}

type GroupInput struct {
	FrontID uuid.UUID
	Name    *string
}
