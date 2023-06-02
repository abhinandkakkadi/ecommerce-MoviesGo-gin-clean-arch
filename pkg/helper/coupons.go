package helper

import (
	"fmt"

	"gorm.io/gorm"
)

func GetCouponDiscountPrice(userID int, TotalPrice float64, DB *gorm.DB) (float64, error) {
	fmt.Println("code reached coupon discount price")
	var count int
	err := DB.Raw("select count(*) from used_coupons where user_id = ? and used = false", userID).Scan(&count).Error
	if err != nil {
		return 0.0, err
	}

	if count < 0 {
		return 0.0, nil
	}

	var couponID int
	err = DB.Raw("select coupon_id from used_coupons where user_id = ? and used = false", userID).Scan(&couponID).Error
	if err != nil {
		return 0.0, err
	}

	type CouponDetails struct {
		DiscountPercentage int
		MinimumPrice       float64
	}

	var coupD CouponDetails
	err = DB.Raw("select discount_percentage,minimum_price from coupons where id = ?", couponID).Scan(&coupD).Error
	if err != nil {
		return 0.0, err
	}

	var totalPrice float64
	err = DB.Raw("select COALESCE(SUM(total_price), 0) from carts where user_id = ?", userID).Scan(&totalPrice).Error
	if err != nil {
		return 0.0, err
	}

	if totalPrice < coupD.MinimumPrice {
		return 0.0, nil
	}

	return ((float64(coupD.DiscountPercentage) * totalPrice) / 100), nil

}
