package types

type Group struct {
	ID        int
	Name      string
	IsDeleted bool
	Users     []User
	Subjects  []Subject
}
