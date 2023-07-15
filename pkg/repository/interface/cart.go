package interfaces

import "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"

type CartRepository interface {
	AddItemToCart(userID int, product_id int, quantity int, productPrice float64) error
	TotalPriceForProductInCart(userID int, productID int) (float64, error)
	UpdateCart(quantity int, price float64, userID int, product_id int) error
	QuantityOfProductInCart(userID int, product_id int) (int, error)
	RemoveFromCart(userID int) ([]models.Cart, error)
	DisplayCart(userID int) ([]models.Cart, error)
	EmptyCart(userID int) error
	GetTotalPriceFromCart(userID int) (float64, error)
	GetUnUsedCategoryOfferIDS(userID int) ([]int, error)
	UpdateUnUsedCategoryOffer(cOfferID int, userID int) error
	UpdateUnUsedProductOffer(pOfferID int, userID int) error
	GetUnUsedProductOfferIDS(userID int) ([]int, error)
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
	UpdateUsedCoupon(coupon string, userID int) (bool, error)
	DoesCartExist(userID int) (bool, error)
}
