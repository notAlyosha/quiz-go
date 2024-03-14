package types

type PairedAnswerQuestionOption struct {
	ID int `gorm: "primaryKey;autoIncrement;not null"`

	Question   QuizesQuestionsContainers `gorm: "foreignkey:QuestionID"`
	QuestionID int

	OptionText string `gorm: "not null"`

	IsDeleted bool `gorm: "default:false"`
}
