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

func (co *couponRepository) DidUserAlreadyUsedThisCoupon(coupon string, userID int) (bool, error) {

	var count int
	err := co.DB.Raw("select count(*) from used_coupons where coupon_id = (select id from coupons where coupon = ?) and user_id = ?", coupon, userID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil

}

func (co *couponRepository) CouponExist(couponName string) (bool, error) {

	var count int
	err := co.DB.Raw("select count(*) from coupons where coupon = ?", couponName).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil

}

func (co *couponRepository) GetCouponMinimumAmount(coupon string) (float64, error) {

	var MinDiscountPrice float64
	err := co.DB.Raw("select minimum_price from coupons where coupon = ?", coupon).Scan(&MinDiscountPrice).Error
	if err != nil {
		return 0.0, err
	}
	return MinDiscountPrice, nil
}

func (co *couponRepository) CouponValidity(couponName string) (bool, error) {

	var validity bool
	err := co.DB.Raw("select validity from coupons where coupon = ?", couponName).Scan(&validity).Error
	if err != nil {
		return false, err
	}

	return validity, nil

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

	var valid bool
	err := co.DB.Raw("select validity from coupons where id = ?", couponID).Scan(&valid).Error
	if err != nil {
		return err
	}

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
		err = co.DB.Exec("delete from product_offers where product_id = ?", productOffer.ProductID).Error
		if err != nil {
			return err
		}
	}

	startDate := time.Now()
	endDate := time.Now().Add(time.Hour * 24 * 5)
	err = co.DB.Exec("INSERT INTO product_offers (product_id, offer_name, discount_percentage, start_date, end_date, offer_limit,offer_used) VALUES (?, ?, ?, ?, ?, ?, ?)", productOffer.ProductID, productOffer.OfferName, productOffer.DiscountPercentage, startDate, endDate, productOffer.OfferLimit, 0).Error
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
	err = co.DB.Exec("INSERT INTO category_offers (genre_id, offer_name, discount_percentage, start_date, end_date, offer_limit,offer_used) VALUES (?, ?, ?, ?, ?, ?, ?)", categoryOffer.GenreID, categoryOffer.OfferName, categoryOffer.DiscountPercentage, startDate, endDate, categoryOffer.OfferLimit, 0).Error
	if err != nil {
		return err
	}

	return nil

}

func (co *couponRepository) OfferDetails(productID int, genre string) (models.CombinedOffer, error) {

	var offer models.OfferResponse

	// first we will check whether what all offer's exists and we will choose the one which provides maximum discount
	var pOff models.Offer
	var cOff models.Offer

	// get details of product offer
	err := co.DB.Raw("select id,offer_name,discount_percentage,start_date,end_date,offer_limit,offer_used from product_offers where product_id = ?", productID).Scan(&pOff).Error
	if err != nil {
		return models.CombinedOffer{}, err
	}

	var price float64
	err = co.DB.Raw("select price from products where id = ?", productID).Scan(&price).Error
	if err != nil {
		return models.CombinedOffer{}, err
	}

	// get details of category offer
	err = co.DB.Raw("select id,offer_name,discount_percentage,start_date,end_date,offer_limit,offer_used from category_offers where genre_id = (select id from genres where genre_name  = ?)", genre).Scan(&cOff).Error
	if err != nil {
		return models.CombinedOffer{}, err
	}

	return models.CombinedOffer{
		ProductOffer:  pOff,
		CategoryOffer: cOff,
		FinalOffer:    offer,
		OriginalPrice: price,
	}, nil

}

func (co *couponRepository) CheckIfProductOfferAlreadyUsed(offerDetails models.OfferResponse, product_id int, userID int) (models.OfferResponse, error) {

	var used bool
	err := co.DB.Raw("select used from product_offer_useds where user_id = ? and product_offer_id = ?", userID, offerDetails.OfferID).Scan(&used).Error
	if err != nil {
		return models.OfferResponse{}, err
	}

	if used {
		err := co.DB.Raw("select price from products where id = ? ", product_id).Scan(&offerDetails.OfferPrice).Error
		if err != nil {
			return models.OfferResponse{}, err
		}
		return offerDetails, nil
	}

	return offerDetails, nil
}

func (co *couponRepository) CheckIfCategoryOfferAlreadyUsed(offerDetails models.OfferResponse, product_id int, userID int) (models.OfferResponse, error) {

	var used bool
	err := co.DB.Raw("select used from category_offer_useds where user_id = ? and category_offer_id = ?", userID, offerDetails.OfferID).Scan(&used).Error
	if err != nil {
		return models.OfferResponse{}, err
	}

	if used {
		err := co.DB.Raw("select price from products where id = ? ", product_id).Scan(&offerDetails.OfferPrice).Error
		if err != nil {
			return models.OfferResponse{}, err
		}
		return offerDetails, nil
	}

	return offerDetails, nil

}

func (co *couponRepository) OfferUpdateProduct(offerDetails models.OfferResponse, userID int) error {

	var count int
	err := co.DB.Raw("select count(*) from product_offer_useds where product_offer_id = ? and user_id = ? ", offerDetails.OfferID, userID).Scan(&count).Error
	if err != nil {
		return err
	}

	// find genre of the movie
	var genreID int
	err = co.DB.Raw("select genre_id from category_offers inner join category_offer_useds on category_offers.id = category_offer_useds.category_offer_id where user_id = ? and used = false", userID).Scan(&genreID).Error
	if err != nil {
		return err
	}
	// the user haven't used product offer. also the user did't use category offer yet. -  if both condition come true - allow it
	if genreID == 0 {
		if count == 0 {
			co.DB.Exec("insert into product_offer_useds (user_id,product_offer_id,offer_amount,offer_count,used) values (?,?,?,?,?)", userID, offerDetails.OfferID, offerDetails.OfferPrice, 1, false).Scan(&count)
		} else {
			err = co.DB.Exec("update product_offer_useds set offer_count = offer_count + 1 where product_offer_id = ? and user_id = ?", offerDetails.OfferID, userID).Error
			if err != nil {
				return err
			}
		}

		err = co.DB.Exec("update product_offers set offer_used = offer_used + 1 where id = ?", offerDetails.OfferID).Error
		if err != nil {
			return err
		}
	}

	return nil

}

func (co *couponRepository) OfferUpdateCategory(offerDetails models.OfferResponse, userID int) error {

	var count int
	err := co.DB.Raw("select count(*) from category_offer_useds where category_offer_id = ? and user_id = ? ", offerDetails.OfferID, userID).Scan(&count).Error
	if err != nil {
		return err
	}

	var productID int
	err = co.DB.Raw("select product_id from product_offers inner join product_offer_useds on product_offers.id = product_offer_useds.product_offer_id where user_id = ? and used = false", userID).Scan(&productID).Error
	if err != nil {
		return err
	}

	// The user have not yet used this offer. create one
	if productID == 0 {
		if count == 0 {
			co.DB.Exec("insert into category_offer_useds (user_id,category_offer_id,offer_amount,offer_count,used) values (?,?,?,?,?)", userID, offerDetails.OfferID, offerDetails.OfferPrice, 1, false).Scan(&count)
		} else {
			err = co.DB.Exec("update category_offer_useds set offer_count = offer_count + 1 where category_offer_id = ? and user_id = ?", offerDetails.OfferID, userID).Error
			if err != nil {
				return err
			}
		}

		err = co.DB.Exec("update category_offers set offer_used = offer_used + 1 where id = ?", offerDetails.OfferID).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// this is the most complicated function in this program
func (co *couponRepository) GetPriceBasedOnOffer(product_id int, userID int) (float64, error) {

	var quantity int
	err := co.DB.Raw("select quantity from carts where product_id = ?", product_id).Scan(&quantity).Error
	if err != nil {
		return 0.0, err
	}

	var originalPrice float64
	err = co.DB.Raw("select price from products where id = ?", product_id).Scan(&originalPrice).Error
	if err != nil {
		return 0.0, err
	}

	var productOfferPrice float64
	var categoryOfferPrice float64
	var OriginalProductPrice = originalPrice

	// check if the product have a offer using offer
	// for that first find the offer id for that particular product, and now check in the product_offers_used table to check whether
	// it is used or not i.e (if is is used that means the order have been already done)
	var pOfferID int
	err = co.DB.Raw("select id from product_offers where product_id = ?", product_id).Scan(&pOfferID).Error
	if err != nil {
		return 0.0, err
	}

	var pOfferCount int
	err = co.DB.Raw("select count(*) from product_offer_useds where product_offer_id = ? and user_id = ? and used = false", pOfferID, userID).Scan(&pOfferCount).Error
	if err != nil {
		return 0.0, err
	}

	var genreID int
	err = co.DB.Raw("select genre_id from products where id = ? ", product_id).Scan(&genreID).Error
	if err != nil {
		return 0.0, err
	}

	var cOfferID int
	err = co.DB.Raw("select id from category_offers where genre_id = ?", genreID).Scan(&cOfferID).Error
	if err != nil {
		return 0.0, err
	}

	var cOfferCount int
	err = co.DB.Raw("select count(*) from category_offer_useds where category_offer_id = ? and user_id = ? and used = false", cOfferID, userID).Scan(&cOfferCount).Error
	if err != nil {
		return 0.0, err
	}

	if pOfferCount > 0 {

		var offerCount int
		err = co.DB.Raw("select offer_count from product_offer_useds where product_offer_id = ? and user_id = ?", pOfferID, userID).Scan(&offerCount).Error
		if err != nil {
			return 0.0, err
		}

		var offerAmount float64
		err = co.DB.Raw("select offer_amount from product_offer_useds where product_offer_id = ? and user_id = ?", pOfferID, userID).Scan(&offerAmount).Error
		if err != nil {
			return 0.0, err
		}

		if quantity <= offerCount {
			productOfferPrice = offerAmount
		}

	}

	if cOfferCount > 0 {
		var offerCount int
		err = co.DB.Raw("select offer_count from category_offer_useds where category_offer_id = ? and user_id = ?", cOfferID, userID).Scan(&offerCount).Error
		if err != nil {
			return 0.0, err
		}

		var offerAmount float64
		err = co.DB.Raw("select offer_amount from category_offer_useds where category_offer_id = ? and user_id = ?", cOfferID, userID).Scan(&offerAmount).Error
		if err != nil {
			return 0.0, err
		}

		if quantity <= offerCount {
			fmt.Println("if quantity is less than 5 it have to sent back 400")
			categoryOfferPrice = offerAmount
		}

	}

	if categoryOfferPrice == 0 && productOfferPrice == 0 {
		return OriginalProductPrice, nil
	}

	if productOfferPrice != 0 {

		err = co.DB.Exec("update product_offer_useds set offer_count = offer_count - 1 where product_offer_id = ? and user_id = ?", pOfferID, userID).Error
		if err != nil {
			return 0.0, err
		}

		err = co.DB.Exec("update product_offers set offer_used = offer_used - 1 where id = ?", pOfferID).Error
		if err != nil {
			return 0.0, err
		}

	}

	if categoryOfferPrice != 0 {

		err = co.DB.Exec("update category_offer_useds set offer_count = offer_count - 1 where category_offer_id = ? and user_id = ?", cOfferID, userID).Error
		if err != nil {
			return 0.0, err
		}

		err = co.DB.Exec("update category_offers set offer_used = offer_used - 1 where id = ?", cOfferID).Error
		if err != nil {
			return 0.0, err
		}
	}

	if categoryOfferPrice > productOfferPrice {
		return categoryOfferPrice, nil
	} else if productOfferPrice > categoryOfferPrice {
		return productOfferPrice, nil
	}
	return OriginalProductPrice, nil

}

func (co *couponRepository) GetReferralAmount(userID int) (models.ReferralAmount, error) {

	// get referral amount associated with the user
	var referralAmount models.ReferralAmount
	err := co.DB.Raw("select referral_amount from referrals where user_id = ?", userID).Scan(&referralAmount).Error
	if err != nil {
		return models.ReferralAmount{}, err
	}
	return referralAmount, nil

}

func (co *couponRepository) DiscountReason(userID int, tableName string, discountLabel string, discountApplied *[]string) error {

	var count int
	err := co.DB.Raw("select count(*) from "+tableName+" where used = false and user_id = ?", userID).Scan(&count).Error
	if err != nil {
		return err
	}

	if count != 0 {
		*discountApplied = append(*discountApplied, discountLabel)
		count = 0
	}

	return nil

}
