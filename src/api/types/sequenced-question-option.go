package types

type SequencedQuestionOption struct {
	ID int

	Question   QuizesQuestionsContainers
	QuestionID int

	OptionText string

	IsDeleted bool
}
