package entity

import validation "github.com/go-ozzo/ozzo-validation"

type OpenedQuestion struct {
	ID         int
	QuestionID int
	AnswerText string
	IsDeleted  bool
}

type OpenedQuestionInput struct {
	QuestionID int
	AnswerText string
}

func (q *OpenedQuestionInput) Check() error {
	return validation.ValidateStruct(q,
		validation.Field(q.QuestionID, validation.Required),
		validation.Field(q.AnswerText, validation.Required, validation.Length(5, 200)),
	)
}

func (q *OpenedQuestionInput) Save() {

}
