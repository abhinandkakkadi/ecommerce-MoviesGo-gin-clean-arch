package interfaces


type CartUseCase interface {
	AddToCart(product_id int,userID int)
}