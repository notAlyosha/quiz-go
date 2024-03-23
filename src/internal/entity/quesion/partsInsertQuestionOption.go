package entity

import validation "github.com/go-ozzo/ozzo-validation"

type PartsInsertQuestionOption struct {
	ID         int
	QuestionID int
	PartText   string
	IsHidden   bool
	IsDeleted  bool
}

type PartsInsertQuestionOptionInput struct {
	QuestionID int
	PartText   string
	IsHidden   bool
}

func (q *PartsInsertQuestionOptionInput) Check() error {
	return validation.ValidateStruct(q,
		validation.Field(q.QuestionID, validation.Required),
		validation.Field(q.PartText, validation.Required, validation.Length(5, 300)),
		validation.Field(q.IsHidden, validation.Required),
	)
}

func (q *PartsInsertQuestionOptionInput) Save() {

}
