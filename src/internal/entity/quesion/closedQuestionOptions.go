package entity

type ClosedQuestionOptions struct {
	ID              int
	QuestionID      int
	OptionText      string
	IsCorrectOption bool
	IsDeleted       bool
}
