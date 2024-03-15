package types

type ClosedQuestionOptions struct {
	ID int

	Question   QuizesQuestionsContainers
	QuestionID int

	IsCorrectOption bool

	IsDeleted bool
}
