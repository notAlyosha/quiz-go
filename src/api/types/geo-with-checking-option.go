package types

type GeoWithCheckingOption struct {
	ID int

	Question   QuizesQuestionsContainers
	QuestionID int

	Langitude string
	Lantitude string

	IsPlaceVisited bool

	CheckQuestionText   string
	CheckQuestionAnswer string

	IsDeleted bool
}
