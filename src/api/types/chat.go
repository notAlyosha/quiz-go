package types

type Chat struct {
	ID              int  `gorm: "primaryKey;autoIncrement;not null"`
	User            User `gorm: "foreignkey:AdministratorID"`
	AdministratorID int
	Description     *string
	InviteLink      *string
	ThemeName       *string

	IsGroupChat bool `gorm: "not null"`

	IsDeleted bool `gorm: "default:false"`
}
