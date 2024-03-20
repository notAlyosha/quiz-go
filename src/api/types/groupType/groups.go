package groupType

type Group struct {
	ID        int
	Name      string
	IsDeleted bool
}

type GroupInput struct {
	Name *string
}
