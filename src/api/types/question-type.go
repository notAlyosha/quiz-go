package types

type QuestionTypes struct {
	ID int `gorm: "primaryKey;autoIncrement;not null"`

	Name string `gorm: "unique;not null"`

	Durability int `gorm: "not null"`

	Points int `gorm: "not null"`

	IsDeleted bool `gorm: "default:false"`
}
