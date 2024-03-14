package types

type Subject struct {
	ID        int    `gorm: "primaryKey;autoIncrement;not null"`
	Name      string `gorm: "not null;unique"`
	IsDeleted bool   `gorm: "default:false"`
}
