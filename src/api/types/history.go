package types

type History struct {
	ID           int
	UserID       int
	TableName    string
	ColumnName   string
	RowID        string
	OldValue     string
	NewValue     string
	EditDateTime string
	IsDeleted    bool
}
