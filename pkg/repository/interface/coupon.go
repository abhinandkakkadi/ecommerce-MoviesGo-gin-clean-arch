package interfaces

import "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"

type CouponRepository interface {
	CouponExist(couponName string) (bool, error)
	CouponRevalidateIfExpired(couponName string) (bool, error)
	AddCoupon(coupon models.AddCoupon) error
	GetCoupon() ([]models.Coupon, error)
	ExistCoupon(couponID int) (bool, error)
	CouponAlreadyExpired(couponID int) error
	AddProductOffer(productOffer models.ProductOfferReceiver) error
	AddCategoryOffer(categoryOffer models.CategoryOfferReceiver) error
	OfferDetails(productID int, genre string) (models.OfferResponse, error)
	CheckIfOfferAlreadyUsed(offerDetails models.OfferResponse, product_id int, userID int) (models.OfferResponse, error)
	OfferUpdate(offerDetails models.OfferResponse, userID int) error
	GetReferralAmount(userID int) (models.ReferralAmount, error)
	GetPriceBasedOnOffer(product_id int, userID int) (float64, error)

	DiscountReason(userID int) ([]string, error)
}
