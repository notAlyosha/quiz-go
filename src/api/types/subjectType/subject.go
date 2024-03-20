package subjectType

type Subject struct {
	ID        int
	Name      string
	IsDeleted bool
}

type SubjectInput struct {
	Name *string `validator:"required"`
}
