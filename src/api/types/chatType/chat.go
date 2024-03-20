package chatType

type Chat struct {
	ID              int
	AdministratorID int
	Description     string
	InviteLink      string
	ThemeName       string
	IsGroupChat     bool
	IsDeleted       bool
}

type ChatInput struct {
	Description *string
	InviteLink  *string `validator:"uri"`
	ThemeName   *string
}
