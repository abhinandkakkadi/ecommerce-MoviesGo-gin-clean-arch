package interfaces

import "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"

type CouponUseCase interface {
	AddCoupon(coupon models.Coupon) (string, error)
	GetCoupon() ([]models.Coupon, error)
	ExpireCoupon(couponID int) error
}
