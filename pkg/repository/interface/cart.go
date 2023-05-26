package interfaces

import "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"


type CartRepository interface {
	AddToCart(product_id int,userID int)  ([]models.Cart,error)
	RemoveFromCart(product_id int,userID int) ([]models.Cart,error)
	DisplayCart(userID int) ([]models.Cart,error)
	EmptyCart(userID int) ([]models.Cart,error) 
	GetTotalPrice(userID int) (models.CartTotal,error)
}