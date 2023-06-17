package models

type Cart struct {
	ProductID  uint    `json:"product_id"`
	MovieName  string  `json:"movie_name"`
	Quantity   float64 `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

type CartResponse struct {
	UserName   string
	TotalPrice float64
	Cart       []Cart
}

type CartTotal struct {
	UserName       string  `json:"user_name"`
	TotalPrice     float64 `json:"total_price"`
	FinalPrice     float64 `json:"final_price"`
	DiscountReason string
}
