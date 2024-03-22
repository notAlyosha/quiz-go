package entity

import "database/sql"

type QuizesQuestionsContainers struct {
	ID             int
	QuizID         int
	QuestionTypeID int
	QuestionText   string
	PictureURL     sql.NullString
	IsAdditional   bool
	IsDeleted      bool
}
