package interfaces

import "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"

type CartRepository interface {
	AddToCart(product_id int, userID int, offerDetails models.OfferResponse) ([]models.Cart, error)
	RemoveFromCart(userID int) ([]models.Cart, error)
	DisplayCart(userID int) ([]models.Cart, error)
	EmptyCart(userID int) ([]models.Cart, error)
	GetQuantityAndTotalPrice(userID int, product_id int, cartDetails struct {
		Quantity   int
		TotalPrice float64
	}) (struct {
		Quantity   int
		TotalPrice float64
	}, error)
	UpdateCartDetails(cartDetails struct {
		Quantity   int
		TotalPrice float64
	}, userID int, productID int) error
	RemoveProductFromCart(userID int, product_id int) error
	GetTotalPrice(userID int) (models.CartTotal, error)
	GetAllItemsFromCart(userID int) ([]models.Cart, error)
	CheckProduct(product_id int) (bool, string, error)
	ProductExist(product_id int, userID int) (bool, error)
	CouponValidity(coupon string, userID int) (bool, error)
	DoesCartExist(userID int) (bool, error)
}
