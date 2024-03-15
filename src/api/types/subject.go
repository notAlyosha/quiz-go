package types

type Subject struct {
	ID        int
	Name      string
	Groups    []Group
	Quizes    []Quiz
	IsDeleted bool
}
