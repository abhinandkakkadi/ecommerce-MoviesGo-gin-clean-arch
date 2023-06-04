package interfaces

import "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"

type CouponRepository interface {
	CouponExist(couponName string) (bool, error)
	CouponRevalidateIfExpired(couponName string) (bool, error)
	AddCoupon(coupon models.AddCoupon) error
	GetCoupon() ([]models.Coupon, error)
	ExistCoupon(couponID int) (bool, error)
	CouponAlreadyExpired(couponID int) error
}
