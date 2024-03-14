package types

type History struct {
	ID           int    `gorm: "primaryKey;autoIncrement;not null"`
	UserID       int    `gorm: "not null"`
	TableName    string `gorm: "not null"`
	ColumnName   string `gorm: "not null"`
	RowID        string `gorm: "not null"`
	OldValue     string `gorm: "not null"`
	NewValue     string `gorm: "not null"`
	EditDateTime string `gorm: "not null"`
	IsDeleted    bool   `gorm: "default:false"`
}
