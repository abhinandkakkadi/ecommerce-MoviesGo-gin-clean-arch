package interfaces

import "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"

type CartUseCase interface {
	AddToCart(product_id int, userID int) (models.CartResponse, error)
	RemoveFromCart(product_id int, userID int) (models.CartResponse, error)
	DisplayCart(userID int) (models.CartResponse, error)
	EmptyCart(userID int) (models.CartResponse, error)
	ApplyCoupon(coupon string, userID int) error
}
