package interfaces

import "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"

type CartRepository interface {
	AddToCart(product_id int, userID int, offerDetails models.OfferResponse) ([]models.Cart, error)
	RemoveFromCart(product_id int, userID int, priceDecrement float64) ([]models.Cart, error)
	DisplayCart(userID int) ([]models.Cart, error)
	EmptyCart(userID int) ([]models.Cart, error)
	GetTotalPrice(userID int) (models.CartTotal, error)
	GetAllItemsFromCart(userID int) ([]models.Cart, error)
	CheckProduct(product_id int) (bool, string, error)
	ProductExist(product_id int, userID int) (bool, error)
	CouponValidity(coupon string, userID int) (bool, error)
}
