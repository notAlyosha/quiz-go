package types

type QuestionTypes struct {
	ID                        int
	Name                      string
	Durability                int
	Points                    int
	IsDeleted                 bool
	QuizesQuestionsContainers []QuizesQuestionsContainers
}
