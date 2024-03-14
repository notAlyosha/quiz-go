package types

type OpenedQuestion struct {
	ID int `gorm: "primaryKey;autoIncrement;not null"`

	Question   QuizesQuestionsContainers `gorm: "foreignkey:QuestionID"`
	QuestionID int

	AnswerText string `gorm: "not null"`

	IsDeleted bool `gorm: "default:false"`
}
