package types

type User struct {
	ID              int
	FrontID         int
	Role            string
	IsDeleted       bool
	Chats           []Chat
	ChatMessages    []ChatMessages
	Groups          []Group
	Quizes          []Quiz
	SessionStudents []SessionStudent
	Sessions        []Session
	Histories       []History
	LogoURL         *string
}
