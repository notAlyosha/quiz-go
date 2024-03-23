package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type GeoWithCheckingOption struct {
	ID                  int
	QuestionID          int
	Langitude           string
	Latitude            string
	IsPlaceVisited      bool
	CheckQuestionText   string
	CheckQuestionAnswer string
	IsDeleted           bool
}

type GeoWithCheckingOptionInput struct {
	QuestionID          int
	Langitude           string
	Latitude            string
	IsPlaceVisited      bool
	CheckQuestionText   string
	CheckQuestionAnswer string
}

func (q *GeoWithCheckingOptionInput) Check() error {
	return validation.ValidateStruct(q,
		validation.Field(q.Langitude, validation.Required, is.Latitude),
		validation.Field(q.Latitude, validation.Required, is.Latitude),
		validation.Field(q.IsPlaceVisited, validation.Required),
		validation.Field(q.CheckQuestionText, validation.Required),
		validation.Field(q.CheckQuestionAnswer, validation.Required),
	)
}

func (q *GeoWithCheckingOptionInput) Save() {

}
