package models

type Coupon struct {
	ID                 uint    `json:"id"`
	Coupon             string  `json:"coupon"`
	DiscountPercentage int     `json:"discount_percentage"`
	MinimumPrice       float64 `json:"minimum_price"`
	Validity           bool    `json:"validity"`
}

type AddCoupon struct {
	Coupon             string  `json:"coupon"`
	DiscountPercentage int     `json:"discount_percentage"`
	MinimumPrice       float64 `json:"minimum_price"`
	Validity           bool    `json:"validity"`
}

type CouponAddUser struct {
	CouponName string `json:"coupon_name"`
}

type ProductOfferReceiver struct {
	ProductID          uint   `json:"product_id"`
	OfferName          string `json:"offer_name"`
	OfferDescription   string `json:"offer_description"`
	DiscountPercentage int    `json:"discount_percentage"`
	OfferLimit         int    `json:"offer_limit"`
}

type OfferResponse struct {
	OfferName       string  `json:"offer_name"`
	OfferPercentage int     `json:"offer_percentage"`
	OfferPrice      float64 `json:"offer_price"`
}

type ProductOfferBriefResponse struct {
	ProductsBrief ProductsBrief
	OfferResponse OfferResponse
}
