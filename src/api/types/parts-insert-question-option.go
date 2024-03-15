package types

type PartsInsertQuestionOption struct {
	ID int

	Question   QuizesQuestionsContainers
	QuestionID int

	PartText string

	IsHidden bool

	IsDeleted bool
}
