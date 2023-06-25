package domain

import "time"

type Products struct {
	ID                  uint        `json:"id" gorm:"unique;not null"`
	MovieName           string      `json:"movie_name"`
	SKU                 string      `json:"sku"`
	GenreID             uint        `json:"genre_id"`
	Genre               Genre       `json:"-" gorm:"foreignkey:GenreID;constraint:OnDelete:CASCADE"`
	Language            string      `json:"language"`
	Director            string      `json:"director"`
	ReleaseYear         string      `json:"release_year"`
	Format              string      `json:"format"`
	ProductsDescription string      `json:"products_discription"`
	RunTime             float64     `json:"run_time"`
	StudioID            uint        `json:"studio_id"`
	MovieStudio         MovieStudio `json:"-" gorm:"foreignkey:StudioID;constraint:OnDelete:CASCADE"`
	Quantity            int         `json:"quantity"`
	Price               float64     `json:"price"`
	Delete              bool        `json:"delete" gorm:"default:false"`
}

type Genre struct {
	ID        uint   `json:"id" gorm:"unique; not null"`
	GenreName string `json:"genre_name"`
}

type Rating struct {
	ID        uint     `json:"id" gorm:"unique; not null"`
	ProductID uint     `json:"product_id"`
	Products  Products `json:"-" gorm:"foreignkey:ProductID"`
	Rating    int      `json:"rating"`
}

type MovieStudio struct {
	ID     uint   `json:"id" gorm:"unique; not null"`
	Studio string `json:"studio"`
}

type ProductOffer struct {
	ID                 uint      `json:"id" gorm:"unique; not null"`
	ProductID          uint      `json:"product_id"`
	Products           Products  `json:"-" gorm:"foreignkey:ProductID"`
	OfferName          string    `json:"offer_name"`
	DiscountPercentage int       `json:"discount_percentage"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	OfferLimit         int       `json:"offer_limit"`
	OfferUsed          int       `json:"offer_used"`
}

type CategoryOffer struct {
	ID                 uint      `json:"id" gorm:"unique; not null"`
	GenreID            uint      `json:"genre_id"`
	Genre              Genre     `json:"-" gorm:"foreignkey:GenreID"`
	OfferName          string    `json:"offer_name"`
	DiscountPercentage int       `json:"discount_percentage"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	OfferLimit         int       `json:"offer_limit"`
	OfferUsed          int       `json:"offer_used"`
}
