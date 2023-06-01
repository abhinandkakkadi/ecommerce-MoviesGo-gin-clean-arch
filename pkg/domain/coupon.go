package domain

type Coupons struct {
	ID                 uint   `json:"id" gorm:"uniquekey; not null"`
	Coupon             string `json:"coupon" gorm:"coupon"`
	DiscountPercentage int    `json:"discount_percentage"`
	Validity           bool   `json:"validity"`
}

type UsedCoupon struct {
	ID       uint    `json:"id" gorm:"uniquekey not null"`
	CouponID uint    `json:"coupon_id"`
	Coupons  Coupons `json:"-" gorm:"foreignkey:CouponID"`
	UserID   uint    `json:"user_id"`
	Users    Users   `json:"-" gorm:"foreignkey:UserID"`
}
