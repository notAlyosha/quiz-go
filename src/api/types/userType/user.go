package userType

type User struct {
	ID        int     `json:"ID"`
	FrontID   int     `json:"FrontID"`
	Role      string  `json:"Role"`
	IsDeleted bool    `json:"IsDeleted"`
	LogoURL   *string `json:"LogoURL"`
}
