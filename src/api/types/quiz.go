package types

import "time"

type Quiz struct {
	ID                        int
	Name                      string
	Description               *string
	SummaryTime               time.Time
	MaxPoints                 string
	MaxAddPoints              string
	IsDeleted                 bool
	Subject                   Subject
	Sessions                  []Session
	QuizesQuestionsContainers []QuizesQuestionsContainers
}
