package types

type SessionStatus struct {
	ID        int    `gorm: "primaryKey;autoIncrement;not null"`
	Name      string `gorm: "not null"`
	IsDeleted bool   `gorm: "default:false"`
}
