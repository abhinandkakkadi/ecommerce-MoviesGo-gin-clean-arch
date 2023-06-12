package domain

import "gorm.io/gorm"

type ProductOfferUsed struct {
	gorm.Model
	UserID         uint         `json:"user_id"`
	Users          Users        `json:"-" gorm:"foreignkey:UserID"`
	ProductOfferID uint         `json:"product_offer_id"`
	ProductOffer   ProductOffer `json:"-" gorm:"constraint:OnDelete:CASCADE;foreignKey:ProductOfferID"`
	OfferAmount    float64      `json:"offer_amount"`
	OfferCount     int          `json:"offer_count"`
	Used           bool         `json:"used"`
}

type CategoryOfferUsed struct {
	gorm.Model
	UserID          uint          `json:"user_id"`
	Users           Users         `json:"-" gorm:"foreignkey:UserID"`
	CategoryOfferID uint          `json:"product_offer_id"`
	CategoryOffer   CategoryOffer `json:"-" gorm:"constraint:OnDelete:CASCADE;foreignKey:CategoryOfferID"`
	OfferAmount     float64       `json:"offer_amount"`
	OfferCount      int           `json:"offer_count"`
	Used            bool          `json:"used"`
}
