package entity

import validation "github.com/go-ozzo/ozzo-validation"

type ClosedQuestionOptions struct {
	ID              int
	QuestionID      int
	OptionText      string
	IsCorrectOption bool
	IsDeleted       bool
}

type ClosedQuestionOptionsInput struct {
	OptionText      string
	IsCorrectOption bool
}

func (q *ClosedQuestionOptionsInput) Check() error {
	return validation.ValidateStruct(q,
		validation.Field(q.OptionText, validation.Required, validation.Length(5, 200)),
		validation.Field(q.IsCorrectOption, validation.Required),
	)
}

func (q *ClosedQuestionOptionsInput) Save() {

}
