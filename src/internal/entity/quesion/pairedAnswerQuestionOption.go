package entity

import validation "github.com/go-ozzo/ozzo-validation"

type PairedAnswerQuestionOption struct {
	ID         int
	QuestionID int
	OptionText string
	IsDeleted  bool
}

type PairedAnswerQuestionOptionInput struct {
	QuestionID int
	OptionText string
}

func (q *PairedAnswerQuestionOptionInput) Check() error {
	return validation.ValidateStruct(q,
		validation.Field(q.QuestionID, validation.Required),
		validation.Field(q.OptionText, validation.Required, validation.Length(5, 100)),
	)
}

func (q *PairedAnswerQuestionOptionInput) Save() {

}
