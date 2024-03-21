package entity

type Subject struct {
	ID        int
	Name      string
	IsDeleted bool
}

type SubjectInput struct {
	Name *string
}
