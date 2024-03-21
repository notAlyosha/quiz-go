package state

type QuizState struct {
	CurrnetQuestion int
	Questions       []Question
	Points          int
}
