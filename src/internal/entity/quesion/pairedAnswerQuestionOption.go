package entity

type PairedAnswerQuestionOption struct {
	ID         int
	QuestionID int
	OptionText string
	IsDeleted  bool
}
