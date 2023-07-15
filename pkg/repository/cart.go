package repository

import (
	"errors"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &cartRepository{
		DB: DB,
	}
}

func (cr *cartRepository) QuantityOfProductInCart(userID int, product_id int) (int, error) {

	var cartsQuantity int
	if err := cr.DB.Raw("select quantity from carts where user_id = ? and product_id = ?", userID, product_id).Scan(&cartsQuantity).Error; err != nil {

		return 0, err
	}

	return cartsQuantity, nil

}

func (cr *cartRepository) AddItemToCart(userID int, product_id int, quantity int, productPrice float64) error {

	if err := cr.DB.Exec("insert into carts (user_id,product_id,quantity,total_price) values(?,?,?,?)", userID, product_id, quantity, productPrice).Error; err != nil {
		return err
	}

	return nil

}

func (cr *cartRepository) TotalPriceForProductInCart(userID int, productID int) (float64, error) {

	var totalPrice float64
	if err := cr.DB.Raw("select sum(total_price) as total_price from carts where user_id = ? and product_id = ?", userID, productID).Scan(&totalPrice).Error; err != nil {
		return 0.0, err
	}

	return totalPrice, nil
}

func (cr *cartRepository) UpdateCart(quantity int, price float64, userID int, product_id int) error {

	if err := cr.DB.Exec("update carts set quantity = ?, total_price = ? where user_id = ? and product_id = ?", quantity, price, userID, product_id).Error; err != nil {
		return err
	}

	return nil

}

func (cr *cartRepository) GetTotalPrice(userID int) (models.CartTotal, error) {

	var cartTotal models.CartTotal
	err := cr.DB.Raw("select COALESCE(SUM(total_price), 0) from carts where user_id = ?", userID).Scan(&cartTotal.TotalPrice).Error
	if err != nil {
		return models.CartTotal{}, err
	}

	err = cr.DB.Raw("select name as user_name from users where id = ?", userID).Scan(&cartTotal.UserName).Error
	if err != nil {
		return models.CartTotal{}, err
	}

	var discount_price float64
	discount_price, err = helper.GetCouponDiscountPrice(userID, cartTotal.TotalPrice, cr.DB)
	if err != nil {
		return models.CartTotal{}, err
	}

	cartTotal.FinalPrice = cartTotal.TotalPrice - discount_price
	return cartTotal, nil

}

func (cr *cartRepository) GetQuantityAndTotalPrice(userID int, productID int, cartDetails struct {
	Quantity   int
	TotalPrice float64
}) (struct {
	Quantity   int
	TotalPrice float64
}, error) {

	// select quantity and totalprice = quantity * indiviualproductpriice from carts
	if err := cr.DB.Raw("select quantity,total_price from carts where user_id = ? and product_id = ?", userID, productID).Scan(&cartDetails).Error; err != nil {
		return struct {
			Quantity   int
			TotalPrice float64
		}{}, err
	}

	return cartDetails, nil

}

func (cr *cartRepository) RemoveProductFromCart(userID int, product_id int) error {

	if err := cr.DB.Exec("delete from carts where user_id = ? and product_id = ?", uint(userID), uint(product_id)).Error; err != nil {
		return err
	}

	return nil
}

func (cr *cartRepository) UpdateCartDetails(cartDetails struct {
	Quantity   int
	TotalPrice float64
}, userID int, productID int) error {

	if err := cr.DB.Exec("update carts set quantity = ?,total_price = ? where user_id = ? and product_id = ?", cartDetails.Quantity, cartDetails.TotalPrice, userID, productID).Error; err != nil {
		return err
	}

	return nil

}

func (cr *cartRepository) RemoveFromCart(userID int) ([]models.Cart, error) {

	var cartResponse []models.Cart
	if err := cr.DB.Raw("select carts.product_id,products.movie_name as movie_name,carts.quantity,carts.total_price from carts inner join products on carts.product_id = products.id where carts.user_id = ?", userID).First(&cartResponse).Error; err != nil {
		return []models.Cart{}, err
	}

	return cartResponse, nil

}

func (cr *cartRepository) DisplayCart(userID int) ([]models.Cart, error) {

	var count int
	if err := cr.DB.Raw("select count(*) from carts where user_id = ? ", userID).First(&count).Error; err != nil {
		return []models.Cart{}, err
	}

	if count == 0 {
		return []models.Cart{}, nil
	}

	var cartResponse []models.Cart

	if err := cr.DB.Raw("select carts.user_id,users.name as user_name,carts.product_id,products.movie_name as movie_name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userID).First(&cartResponse).Error; err != nil {
		return []models.Cart{}, err
	}

	return cartResponse, nil

}

func (cr *cartRepository) GetUnUsedCategoryOfferIDS(userID int) ([]int, error) {

	// CATEGORY OFFER RESTORED
	var categoryOfferID []int
	if err := cr.DB.Raw("select category_offer_id from category_offer_useds where user_id = ? and used = false", userID).Scan(&categoryOfferID).Error; err != nil {
		return []int{}, err
	}

	return categoryOfferID, nil

}

func (cr *cartRepository) GetUnUsedProductOfferIDS(userID int) ([]int, error) {

	// PRODUCT OFFER RESTORED
	var productOfferID []int
	if err := cr.DB.Raw("select product_offer_id from product_offer_useds where user_id = ? and used = false", userID).Scan(&productOfferID).Error; err != nil {
		return []int{}, err
	}

	return productOfferID, nil

}

func (cr *cartRepository) UpdateUnUsedCategoryOffer(cOfferID int, userID int) error {

	var offerCount int
	if err := cr.DB.Raw("select offer_count from category_offer_useds where category_offer_id = ?", cOfferID).Scan(&offerCount).Error; err != nil {
		return err
	}

	// code for deleting this record
	if err := cr.DB.Exec("update category_offers set offer_used = offer_used - ? where id = ?", offerCount, cOfferID).Error; err != nil {
		return err
	}

	if err := cr.DB.Exec("delete from category_offer_useds where user_id = ? and used = false", userID).Error; err != nil {
		return err
	}

	return nil

}

func (cr *cartRepository) UpdateUnUsedProductOffer(pOfferID int, userID int) error {

	var offerCount int
	if err := cr.DB.Raw("select offer_count from product_offer_useds where product_offer_id = ?", pOfferID).Scan(&offerCount).Error; err != nil {
		return err
	}

	// code for deleting this record
	if err := cr.DB.Exec("update product_offers set offer_used = offer_used - ? where id = ?", offerCount, pOfferID).Error; err != nil {
		return err
	}

	if err := cr.DB.Exec("delete from product_offer_useds where user_id = ? and used = false", userID).Error; err != nil {
		return err
	}

	return nil

}

func (cr *cartRepository) EmptyCart(userID int) error {

	if err := cr.DB.Exec("delete from carts where user_id = ? ", userID).Error; err != nil {
		return err
	}

	return nil

}

func (cr *cartRepository) GetAllItemsFromCart(userID int) ([]models.Cart, error) {

	var count int

	var cartResponse []models.Cart
	err := cr.DB.Raw("select count(*) from carts where user_id = ?", userID).Scan(&count).Error
	if err != nil {
		return []models.Cart{}, err
	}

	if count == 0 {
		return []models.Cart{}, nil
	}

	err = cr.DB.Raw("select carts.user_id,users.name as user_name,carts.product_id,products.movie_name as movie_name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userID).First(&cartResponse).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if len(cartResponse) == 0 {
				return []models.Cart{}, nil
			}
			return []models.Cart{}, err
		}
		return []models.Cart{}, err
	}

	return cartResponse, nil

}

func (cr *cartRepository) CheckProduct(product_id int) (bool, string, error) {

	var count int
	err := cr.DB.Raw("select count(*) from products where id = ?", product_id).Scan(&count).Error
	if err != nil {
		return false, "", err
	}

	var genre string
	if count > 0 {
		err := cr.DB.Raw("select genres.genre_name from genres inner join products on products.genre_id = genres.id where products.id = ?", product_id).Scan(&genre).Error
		if err != nil {
			return false, "", err
		}
		return true, genre, nil
	}
	return false, "", nil

}

func (cr *cartRepository) ProductExist(product_id int, userID int) (bool, error) {

	var count int
	err := cr.DB.Raw("select count(*) from carts where user_id = ? and product_id = ?", userID, product_id).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil

}

func (cr *cartRepository) GetTotalPriceFromCart(userID int) (float64, error) {

	var totalPrice float64
	err := cr.DB.Raw("select COALESCE(SUM(total_price), 0) from carts where user_id = ?", userID).Scan(&totalPrice).Error
	if err != nil {
		return 0.0, err
	}

	return totalPrice, nil

}

func (cr *cartRepository) UpdateUsedCoupon(coupon string, userID int) (bool, error) {

	var couponID uint
	err := cr.DB.Raw("select id from coupons where coupon = ?", coupon).Scan(&couponID).Error
	if err != nil {
		return false, err
	}

	var count int
	// if a coupon have already been added, replace the order with current coupon and delete the existing coupon
	err = cr.DB.Raw("select count(*) from used_coupons where user_id = ? and used = false", userID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		err = cr.DB.Exec("delete from used_coupons where user_id = ? and used = false", userID).Error
		if err != nil {
			return false, err
		}
	}

	err = cr.DB.Exec("insert into used_coupons (coupon_id,user_id,used) values (?, ?, false)", couponID, userID).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (cr *cartRepository) DoesCartExist(userID int) (bool, error) {

	count := 0
	err := cr.DB.Raw("select count(*) from carts where user_id = ?", userID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count < 1 {
		return false, nil
	}

	return true, nil
}
