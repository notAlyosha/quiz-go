package types

type QuizesQuestionsContainers struct {
	ID            int
	Quizes        []Quiz
	QuestionTypes QuestionTypes
	QuestionText  string
	PictureURL    *string
	IsAdditional  bool
	IsDeleted     bool

	ClosedQuestionOptions      []ClosedQuestionOptions
	SequencedQuestionOptions   []SequencedQuestionOption
	PairedAnswerQuestionOption []PairedAnswerQuestionOption
	PartsInsertQuestionOption  []PartsInsertQuestionOption
	GeoWithCheckingOption      []GeoWithCheckingOption
	OpenedQuestion             []OpenedQuestion
}
