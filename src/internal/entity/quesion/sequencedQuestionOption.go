package entity

import validation "github.com/go-ozzo/ozzo-validation"

type SequencedQuestionOption struct {
	ID         int
	QuestionID int
	OptionText string
	IsDeleted  bool
}

type SequencedQuestionOptionInput struct {
	QuestionID int
	OptionText string
}

func (q *SequencedQuestionOptionInput) Check() error {
	return validation.ValidateStruct(q,
		validation.Field(q.QuestionID, validation.Required),
		validation.Field(q.OptionText, validation.Required, validation.Length(5, 300)),
	)
}

func (q *SequencedQuestionOptionInput) Save() {

}
