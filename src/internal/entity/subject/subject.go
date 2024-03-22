package entity

// represents record in database
type Subject struct {
	ID        int
	Name      string
	IsDeleted bool
}

type SubjectInput struct {
	Name *string
}
