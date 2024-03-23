package entity

import (
	"database/sql"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type QuizesQuestionsContainers struct {
	ID             int
	QuizID         int
	QuestionTypeID int
	QuestionText   string
	PictureURL     sql.NullString
	IsAdditional   bool
	IsDeleted      bool
}

type QuizesQuestionsContainersInput struct {
	QuizID         int
	QuestionTypeID int
	QuestionText   string
	PictureURL     sql.NullString
	IsAdditional   bool
}

func (q QuizesQuestionsContainersInput) Check() error {
	return validation.ValidateStruct(q,
		validation.Field(q.QuizID, validation.Required),
		validation.Field(q.QuestionTypeID, validation.Required),
		validation.Field(q.QuestionText, validation.Required, validation.Length(5, 300)),
		validation.Field(q.PictureURL, is.URL),
		validation.Field(q.IsAdditional, validation.Required),
	)
}
