package repository

import (
	"errors"
	"fmt"
	"time"

	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"gorm.io/gorm"
)

type couponRepository struct {
	DB *gorm.DB
}

func NewCouponRepository(DB *gorm.DB) interfaces.CouponRepository {
	return &couponRepository{
		DB: DB,
	}
}

func (co *couponRepository) CouponExist(couponName string) (bool, error) {

	var count int
	err := co.DB.Raw("select count(*) from coupons where coupon = ?", couponName).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil

}

func (co *couponRepository) CouponRevalidateIfExpired(couponName string) (bool, error) {

	var isValid bool
	err := co.DB.Raw("select validity from coupons where coupon = ?", couponName).Scan(&isValid).Error
	if err != nil {
		return false, err
	}

	if isValid {
		return true, nil
	}

	err = co.DB.Exec("update coupons set validity = true where coupon = ?", couponName).Error
	if err != nil {
		return false, err
	}

	return false, nil

}

func (co *couponRepository) AddCoupon(coupon models.AddCoupon) error {
	fmt.Println("from add coupon repository: ", coupon)
	err := co.DB.Exec("insert into coupons (coupon,discount_percentage,minimum_price,validity) values (?, ?, ?, ?)", coupon.Coupon, coupon.DiscountPercentage, coupon.MinimumPrice, true).Error
	if err != nil {
		return nil
	}

	return nil
}

func (co *couponRepository) GetCoupon() ([]models.Coupon, error) {

	var coupons []models.Coupon
	err := co.DB.Raw("select id,coupon,discount_percentage,minimum_price,Validity from coupons").Scan(&coupons).Error
	if err != nil {
		return []models.Coupon{}, err
	}

	return coupons, nil
}

func (co *couponRepository) ExistCoupon(couponID int) (bool, error) {

	var count int
	err := co.DB.Raw("select count(*) from coupons where id = ?", couponID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (co *couponRepository) CouponAlreadyExpired(couponID int) error {
	fmt.Println("the code reached here")
	var valid bool
	err := co.DB.Raw("select validity from coupons where id = ?", couponID).Scan(&valid).Error
	if err != nil {
		return err
	}
	fmt.Println("the validity = ", valid)
	if valid {
		err := co.DB.Exec("update coupons set validity = false where id = ?", couponID).Error
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("already expired")
}

func (co *couponRepository) AddProductOffer(productOffer models.ProductOfferReceiver) error {

	// check if the offer with the offer name already exist in the database
	var count int
	err := co.DB.Raw("select count(*) from product_offers where offer_name = ? and product_id = ?", productOffer.OfferName, productOffer.ProductID).Scan(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("the offer already exists")
	}

	// if there is any other offer for this product delete that before adding this one
	count = 0
	err = co.DB.Raw("select count(*) from product_offers where product_id = ?", productOffer.ProductID).Scan(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		fmt.Println("reached here since offer already exist and is ready to delete")
		err = co.DB.Exec("delete from product_offers where product_id = ?", productOffer.ProductID).Error
		if err != nil {
			return err
		}
	}

	startDate := time.Now()
	endDate := time.Now().Add(time.Hour * 24 * 5)
	fmt.Println(productOffer)
	err = co.DB.Exec("INSERT INTO product_offers (product_id, offer_name, discount_percentage, start_date, end_date) VALUES (?, ?, ?, ?, ?)", productOffer.ProductID, productOffer.OfferName, productOffer.DiscountPercentage, startDate, endDate).Error
	if err != nil {
		return err
	}

	return nil

}

func (co *couponRepository) AddCategoryOffer(categoryOffer models.CategoryOfferReceiver) error {

	// check if the offer with the offer name already exist in the database
	var count int
	err := co.DB.Raw("select count(*) from category_offers where offer_name = ?", categoryOffer.OfferName).Scan(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("the offer already exists")
	}

	// if there is any other offer for this category delete that before adding this one
	count = 0
	err = co.DB.Raw("select count(*) from category_offers where genre_id = ?", categoryOffer.GenreID).Scan(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {

		err = co.DB.Exec("delete from category_offers where genre_id = ?", categoryOffer.GenreID).Error
		if err != nil {
			return err
		}
	}

	startDate := time.Now()
	endDate := time.Now().Add(time.Hour * 24 * 5)
	fmt.Println(categoryOffer)
	err = co.DB.Exec("INSERT INTO category_offers (genre_id, offer_name, discount_percentage, start_date, end_date) VALUES (?, ?, ?, ?, ?)", categoryOffer.GenreID, categoryOffer.OfferName, categoryOffer.DiscountPercentage, startDate, endDate).Error
	if err != nil {
		return err
	}

	return nil

}

func (co *couponRepository) OfferDetails(productID int, genre string) (models.OfferResponse, error) {

	var offer models.OfferResponse

	// first we will check whether what all offer's exists and we will choose the one which provides maximum discount

	type Offer struct {
		OfferName          string
		DiscountPercentage int
		StartDate          time.Time
		EndDate            time.Time
	}
	var pOff Offer
	var cOff Offer
	err := co.DB.Raw("select offer_name,discount_percentage,start_date,end_date from product_offers where product_id = ?", productID).Scan(&pOff).Error
	if err != nil {
		return models.OfferResponse{}, err
	}

	var genreID int
	err = co.DB.Raw("select id from genres where genre_name  = ?", genre).Scan(&genreID).Error
	if err != nil {
		return models.OfferResponse{}, err
	}

	var price float64
	err = co.DB.Raw("select price from products where id = ?", productID).Scan(&price).Error
	if err != nil {
		return models.OfferResponse{}, err
	}
	fmt.Println("price of the product is ", price)

	err = co.DB.Raw("select offer_name,discount_percentage,start_date,end_date from category_offers where genre_id = ?", genreID).Scan(&cOff).Error
	if err != nil {
		return models.OfferResponse{}, err
	}
	fmt.Println("product offer details : ", pOff)
	fmt.Println("category offer details : ", cOff)

	// if product offer exist check whether it is still active or if it have been expired
	currentTime := time.Now()
	if pOff.OfferName != "" {
		fmt.Println("reached here and checking is done in product offer")
		if currentTime.After(pOff.StartDate) && currentTime.Before(pOff.EndDate) {
			fmt.Println("Offer is currently active!")
		} else {
			pOff.OfferName = ""
			pOff.DiscountPercentage = 0
		}
	}

	// if category offer exist check whether it is still active or if it have been expired
	if cOff.OfferName != "" {
		fmt.Println("reached here and checking is done in category offer")
		if currentTime.After(cOff.StartDate) && currentTime.Before(cOff.EndDate) {
			fmt.Println("Offer is currently active!")
		} else {
			cOff.OfferName = ""
			cOff.DiscountPercentage = 0
		}
	}

	if pOff.DiscountPercentage > cOff.DiscountPercentage {
		offer.OfferName = pOff.OfferName
		offer.OfferPercentage = pOff.DiscountPercentage
	} else if cOff.DiscountPercentage > pOff.DiscountPercentage {
		offer.OfferName = cOff.OfferName
		offer.OfferPercentage = cOff.DiscountPercentage
	} else {
		offer.OfferName = "sorry no offer at this time"
		offer.OfferPercentage = 0
		return offer, err
	}
	// select price from Price table and add it to the mix and
	offer.OfferPrice = price - ((float64(offer.OfferPercentage) * price) / 100)
	fmt.Println("discounted price : ", offer.OfferPrice)
	return offer, err
}
