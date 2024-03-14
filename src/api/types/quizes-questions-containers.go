package types

type QuizesQuestionsContainers struct {
	ID     int `gorm: "primaryKey;autoIncrement;not null"`
	Quizes []Quiz

	QuestionTypes QuestionTypes
	QuestionText  string `gorm: "not null"`
	PictureURL    *string

	IsAdditional bool
	IsDeleted    bool `gorm: "default:false"`
}
