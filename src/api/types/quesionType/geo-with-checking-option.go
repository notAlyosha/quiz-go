package types

type GeoWithCheckingOption struct {
	ID                  int
	QuestionID          int
	Langitude           string
	Lantitude           string
	IsPlaceVisited      bool
	CheckQuestionText   string
	CheckQuestionAnswer string
	IsDeleted           bool
}
