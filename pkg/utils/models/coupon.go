package models

import "time"

type Coupon struct {
	ID                 uint    `json:"id"`
	Coupon             string  `json:"coupon"`
	DiscountPercentage int     `json:"discount_percentage"`
	MinimumPrice       float64 `json:"minimum_price"`
	Validity           bool    `json:"validity"`
}

type AddCoupon struct {
	Coupon             string  `json:"coupon" binding:"required"`
	DiscountPercentage int     `json:"discount_percentage" binding:"required"`
	MinimumPrice       float64 `json:"minimum_price" binding:"required"`
	Validity           bool    `json:"validity" binding:"required"`
}

type CouponAddUser struct {
	CouponName string `json:"coupon_name" binding:"required"`
}

type ProductOfferReceiver struct {
	ProductID          uint   `json:"product_id" binding:"required"`
	OfferName          string `json:"offer_name" binding:"required"`
	DiscountPercentage int    `json:"discount_percentage" binding:"required"`
	OfferLimit         int    `json:"offer_limit" binding:"required"`
}

type CategoryOfferReceiver struct {
	GenreID            uint   `json:"genre_id" binding:"required"`
	OfferName          string `json:"offer_name" binding:"required"`
	DiscountPercentage int    `json:"discount_percentage" binding:"required"`
	OfferLimit         int    `json:"offer_limit" binding:"required"`
}

type OfferResponse struct {
	OfferID         uint    `json:"offer_id"`
	OfferName       string  `json:"offer_name"`
	OfferPercentage int     `json:"offer_percentage"`
	OfferPrice      float64 `json:"offer_price"`
	OfferType       string  `json:"offer_type"`
	OfferLimit      int     `json:"offer_limit"`
}

type ProductOfferBriefResponse struct {
	ProductsBrief ProductsBrief
	OfferResponse OfferResponse
}

type ProductOfferLongResponse struct {
	ProductsResponse ProductResponse
	OfferResponse    OfferResponse
}

type ReferralAmount struct {
	ReferralAmount float64 `json:"referral_amount"`
}

type Offer struct {
	ID                 uint
	OfferName          string
	DiscountPercentage int
	StartDate          time.Time
	EndDate            time.Time
	OfferLimit         int
	OfferUsed          int
}

type CombinedOffer struct {
	ProductOffer  Offer
	CategoryOffer Offer
	FinalOffer    OfferResponse
	OriginalPrice float64
}
