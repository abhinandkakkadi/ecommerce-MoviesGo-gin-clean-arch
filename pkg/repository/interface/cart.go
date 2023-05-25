package interfaces


type CartRepository interface {
	AddToCart(product_id int,userID int)
}