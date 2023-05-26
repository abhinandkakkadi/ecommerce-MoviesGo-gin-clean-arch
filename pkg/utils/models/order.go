package models

type OrderIncoming struct {
	UserID    uint `json:"user_id"`
	PaymentID uint `json:"payment_id"`
	AddressID uint `json:"address_id"`
}
