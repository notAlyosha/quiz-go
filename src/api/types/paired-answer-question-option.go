package types

type PairedAnswerQuestionOption struct {
	ID int

	Question   QuizesQuestionsContainers
	QuestionID int

	OptionText string

	IsDeleted bool
}
