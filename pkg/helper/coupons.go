package helper

import (
	"gorm.io/gorm"
)

func GetCouponDiscountPrice(userID int, TotalPrice float64, DB *gorm.DB) (float64, error) {

	// If there is no coupons added for this user, return 0 as discount price
	var count int
	err := DB.Raw("select count(*) from used_coupons where user_id = ? and used = false", userID).Scan(&count).Error
	if err != nil {
		return 0.0, err
	}

	if count < 0 {
		return 0.0, nil
	}

	type CouponDetails struct {
		DiscountPercentage int
		MinimumPrice       float64
	}

	// take the discount percentage and minimum price to check the condition ( !! Actually this is not needed. As all the conditions were checked while adding the coupon !!)
	// just discount percentage would work fine - should refactor this in the future
	var coupD CouponDetails
	err = DB.Raw("select discount_percentage,minimum_price from coupons where id = (select coupon_id from used_coupons where user_id = ? and used = false)", userID).Scan(&coupD).Error
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
