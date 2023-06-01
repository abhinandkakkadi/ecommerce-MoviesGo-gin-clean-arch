package domain

type WishList struct {
	ID        uint     `json:"id" gorm:"uniquekey; not null"`
	UserID    uint     `json:"user_id"`
	Users     Users    `json:"-" gorm:"foreignkey:UserID"`
	ProductID uint     `json:"product_id"`
	Products  Products `json:"-" gorm:"foreignkey:ProductID"`
}
