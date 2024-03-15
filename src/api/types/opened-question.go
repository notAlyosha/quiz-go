package types

type OpenedQuestion struct {
	ID int

	Question   QuizesQuestionsContainers
	QuestionID int

	AnswerText string

	IsDeleted bool
}
