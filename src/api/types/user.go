package types

type User struct {
	ID        int    `gorm: "primaryKey;autoIncrement;not null"`
	FrontID   int    `gorm: "not null; unique"`
	Role      string `gorm: "not null"`
	LogoURL   *string
	IsDeleted bool `gorm: "default:false"`
}
