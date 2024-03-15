package types

type Chat struct {
	ID              int
	User            []User
	AdministratorID int
	Description     *string
	InviteLink      *string
	ThemeName       *string

	IsGroupChat bool

	IsDeleted bool
}
