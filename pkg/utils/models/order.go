package models

type OrderIncoming struct {
	UserID    uint `json:"user_id"`
	PaymentID uint `json:"payment_id"`
	AddressID uint `json:"address_id"`
}

type OrderSuccessResponse struct {
	OrderID        string `json:"order_id"`
	ShipmentStatus string `json:"order_status"`
}

// type ProductsBrief struct {
// 	ID             int    `json:"id"`
// 	Movie_Name     string `json:"movie_name"`
// 	Genre          string `json:"genre"`
// 	Movie_Language string `json:"movie_language"`
// }

type OrderDetails struct {
	
	OrderId    				string    
	GrandTotal 			  float64   
	ShipmentStatus    string

}

type OrderProductDetails struct {

	ProductID  uint				`json:"product_id"`
	MovieName  string  `json:"movie_name"`
	Quantity   int      `json:"quantity"`
	TotalPrice float64  `json:"total_price"`
	
}

type FullOrderDetails struct {
	OrderDetails  OrderDetails
	OrderProductDetails  []OrderProductDetails
}