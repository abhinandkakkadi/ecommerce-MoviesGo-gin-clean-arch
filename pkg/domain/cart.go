package domain

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID     uint     `json:"user_id" gorm:"uniquekey; not null"`
	Users      Users    `json:"-" gorm:"foreignkey:UserID"`
	ProductID  uint     `json:"product_id"`
	Products   Products `json:"-" gorm:"foreignkey:ProductID"`
	Quantity   float64  `json:"quantity"`
	TotalPrice float64  `json:"total_price"`
}
