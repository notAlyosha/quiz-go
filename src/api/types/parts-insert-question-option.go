package types

type PartsInsertQuestionOption struct {
	ID int `gorm: "primaryKey;autoIncrement;not null"`

	Question   QuizesQuestionsContainers `gorm: "foreignkey:QuestionID"`
	QuestionID int

	PartText string `gorm: "not null"`

	IsHidden bool

	IsDeleted bool `gorm: "default:false"`
}
