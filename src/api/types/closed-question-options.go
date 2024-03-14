package types

type ClosedQuestionOptions struct {
	ID int `gorm: "primaryKey;autoIncrement;not null"`

	Question   QuizesQuestionsContainers `gorm: "foreignkey:QuestionID"`
	QuestionID int

	IsCorrectOption bool

	IsDeleted bool `gorm: "default:false"`
}
