package types

type GeoWithCheckingOption struct {
	ID int `gorm: "primaryKey;autoIncrement;not null"`

	Question   QuizesQuestionsContainers `gorm: "foreignkey:QuestionID"`
	QuestionID int

	Langitude string `gorm: "not null"`
	Lantitude string `gorm: "not null"`

	IsPlaceVisited bool `gorm: default:false"`

	CheckQuestionText   string `gorm: "not null"`
	CheckQuestionAnswer string `gorm: "not null"`

	IsDeleted bool `gorm: "default:false"`
}
