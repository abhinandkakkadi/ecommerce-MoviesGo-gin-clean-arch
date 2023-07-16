package interfaces

import "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"

type CouponRepository interface {
	CouponExist(couponName string) (bool, error)
	CouponRevalidateIfExpired(couponName string) (bool, error)
	GetCouponMinimumAmount(coupon string) (float64, error)
	CouponValidity(couponName string) (bool, error)
	AddCoupon(coupon models.AddCoupon) error
	GetCoupon() ([]models.Coupon, error)
	ExistCoupon(couponID int) (bool, error)
	CouponAlreadyExpired(couponID int) error
	AddProductOffer(productOffer models.ProductOfferReceiver) error
	AddCategoryOffer(categoryOffer models.CategoryOfferReceiver) error
	OfferDetails(productID int, genreID string) (models.CombinedOffer, error)
	CheckIfProductOfferAlreadyUsed(offerDetails models.OfferResponse, product_id int, userID int) (models.OfferResponse, error)
	CheckIfCategoryOfferAlreadyUsed(offerDetails models.OfferResponse, product_id int, userID int) (models.OfferResponse, error)
	DidUserAlreadyUsedThisCoupon(coupon string, userID int) (bool, error)
	// OfferUpdate(offerDetails models.OfferResponse, userID int) error
	OfferUpdateProduct(offerDetails models.OfferResponse, userID int) error
	OfferUpdateCategory(offerDetails models.OfferResponse, userID int) error

	GetReferralAmount(userID int) (models.ReferralAmount, error)
	GetPriceBasedOnOffer(product_id int, userID int) (float64, error)
	// DiscountReason(userID int) ([]string, error)
	DiscountReason(userID int, tableName string, discountLabel string, discountApplied *[]string) error
}
