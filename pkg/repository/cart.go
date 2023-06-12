package repository

import (
	"errors"
	"fmt"

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

func (cr *cartRepository) AddToCart(product_id int, userID int, offerDetails models.OfferResponse) ([]models.Cart, error) {

	var cartResponse []models.Cart
	// trancation to achieve all or none property
	tx := cr.DB.Begin()
	var count int
	// to check if product for this particular user exist in the cart. If it does not add a new item else update the quantity
	if err := tx.Raw("select count(*) from carts where user_id = ? and  product_id = ?", userID, product_id).Scan(&count).Error; err != nil {
		tx.Rollback()
		return []models.Cart{}, err
	}

	// to check if the user try to add more quantity than what is already present in stock
	var productQuantity int
	if err := tx.Raw("select quantity from products where id = ?", product_id).Scan(&productQuantity).Error; err != nil {
		tx.Rollback()
		return []models.Cart{}, err
	}

	var cartsQuantity int
	if err := tx.Raw("select quantity from carts where user_id = ? and product_id = ?", userID, product_id).Scan(&cartsQuantity).Error; err != nil {
		tx.Rollback()
		return []models.Cart{}, err
	}

	// // if the product is out of stock
	// if productQuantity == 0 {
	// 	return []models.Cart{}, nil
	// }

	// if the cart is empty and the product we trying to add is out of stock
	var itemsPresentInCart int
	if err := cr.DB.Raw("select count(*) from carts where user_id = ?", userID).Scan(&itemsPresentInCart).Error; err != nil {
		return []models.Cart{}, err
	}

	if itemsPresentInCart == 0 && productQuantity == 0 {
		return []models.Cart{}, errors.New("product out of stock")
	}

	// quantity in carts is equal to quantity in STOCK  -- don't allow to add further products OR product out of stock  -- or if prodctQuanity = 0 - which means the item is out of stock
	if cartsQuantity == productQuantity || productQuantity == 0 {

		if err := tx.Raw("select carts.user_id,users.name as user_name,carts.product_id,products.movie_name as movie_name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userID).First(&cartResponse).Error; err != nil {
			tx.Rollback()
			return []models.Cart{}, err
		}
		return cartResponse, nil
	}

	var totalPrice float64
	var productPrice float64
	// return price of product

	// OFFER DETAILS ARE DONE HERE
	if err := tx.Raw("select price from products where id = ?", product_id).Scan(&productPrice).Error; err != nil {
		tx.Rollback()
		return []models.Cart{}, err
	}

	// if this condition is true that means a offer exist for this product
	if offerDetails.OfferPrice != productPrice {

		var pQuantity int
		if err := tx.Raw("select quantity from carts where product_id = ?", product_id).Scan(&pQuantity).Error; err != nil {
			tx.Rollback()
			return []models.Cart{}, err
		}

		if pQuantity < offerDetails.OfferLimit {
			productPrice = offerDetails.OfferPrice
		}

	}

	fmt.Println(totalPrice)
	// if the product is not already present in the cart - fresh item
	if count == 0 {

		totalPrice = productPrice
		fmt.Println(totalPrice)
		if err := tx.Exec("insert into carts (user_id,product_id,quantity,total_price) values(?,?,?,?)", userID, product_id, 1, totalPrice).Error; err != nil {
			tx.Rollback()
			return []models.Cart{}, err
		}
		fmt.Println("the above thing worked")

	} else {

		// if the product already exist - just iterate the quantity
		if err := tx.Raw("select sum(total_price) as total_price from carts where user_id = ? and product_id = ?", userID, product_id).Scan(&totalPrice).Error; err != nil {
			tx.Rollback()
			return []models.Cart{}, err
		}

		var quantity int
		if err := tx.Raw("select quantity from carts where user_id = ? and product_id = ?", userID, product_id).Scan(&quantity).Error; err != nil {
			tx.Rollback()
			return []models.Cart{}, err

		}

		if err := tx.Exec("update carts set quantity = ?, total_price = ? where user_id = ? and product_id = ?", quantity+1, totalPrice+productPrice, userID, product_id).Error; err != nil {
			tx.Rollback()
			return []models.Cart{}, err

		}
	}
	// list the cart and return
	if err := tx.Raw("select carts.user_id,users.name as user_name,carts.product_id,products.movie_name as movie_name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userID).First(&cartResponse).Error; err != nil {
		tx.Rollback()
		return []models.Cart{}, err
	}
	// commit the transation
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return []models.Cart{}, err
	}

	return cartResponse, nil

}

func (cr *cartRepository) GetTotalPrice(userID int) (models.CartTotal, error) {

	var cartTotal models.CartTotal
	// to return the grand total of all the products in the table
	// return 0 in place of total_price if the record is not present
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
	fmt.Println("this is discount price which is initially 0 ", discount_price)
	return cartTotal, nil
}

func (cr *cartRepository) RemoveFromCart(product_id int, userID int) ([]models.Cart, error) {

	// to check if the product to be removed exist inside the cart
	var count int
	if err := cr.DB.Raw("select count(*) from carts where user_id = ? and product_id = ?", userID, product_id).Scan(&count).Error; err != nil {

		return []models.Cart{}, err
	}

	// have to recheck this
	if count == 0 {
		return []models.Cart{}, errors.New("the product does not exist in the cart")
	}

	var cartDetails struct {
		Quantity   int
		TotalPrice float64
		Price      float64
	}
	// select quantity and totalprice = quantity * indiviualproductpriice from carts
	if err := cr.DB.Raw("select quantity,total_price from carts where user_id = ? and product_id = ?", userID, product_id).Scan(&cartDetails).Error; err != nil {
		return []models.Cart{}, err
	}

	cartDetails.Quantity = cartDetails.Quantity - 1
	// after decrementing one quantity if the quantity = 0. delete that item from the cart
	if cartDetails.Quantity == 0 {
		fmt.Println("it reached here when i intended")
		if err := cr.DB.Exec("delete from carts where user_id = ? and product_id = ?", uint(userID), uint(product_id)).Error; err != nil {

			return []models.Cart{}, err
		}

	}

	if cartDetails.Quantity != 0 {
		fmt.Println("quantity and totalPrice = ", cartDetails)
		// find the price for the product to be removed so that it can be decreased from total price
		if err := cr.DB.Raw("select price from products where id = ?", product_id).Scan(&cartDetails.Price).Error; err != nil {

			return []models.Cart{}, err
		}

		cartDetails.TotalPrice = cartDetails.TotalPrice - cartDetails.Price

		if err := cr.DB.Exec("update carts set quantity = ?,total_price = ? where user_id = ? and product_id = ?", cartDetails.Quantity, cartDetails.TotalPrice, userID, product_id).Error; err != nil {
			return []models.Cart{}, err
		}
	}

	// if err := cr.DB.Exec("delete from carts where user_id = ? and product_id = ?", userID, product_id).Error; err != nil {
	// 	return []models.Cart{}, err
	// }

	// if after decrementing quantity if the carts is empty - return back an empty cart
	if err := cr.DB.Raw("select count(*) from carts where user_id = ?", userID).Scan(&count).Error; err != nil {
		return []models.Cart{}, err
	}

	if count == 0 {

		return []models.Cart{}, nil
	}

	// if all went right sent back the updated cart details
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
	fmt.Println(count)
	if count == 0 {
		return []models.Cart{}, nil
	}

	var cartResponse []models.Cart

	if err := cr.DB.Raw("select carts.user_id,users.name as user_name,carts.product_id,products.movie_name as movie_name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userID).First(&cartResponse).Error; err != nil {
		return []models.Cart{}, err
	}
	fmt.Println(cartResponse)
	return cartResponse, nil
}

func (cr *cartRepository) EmptyCart(userID int) ([]models.Cart, error) {

	var count int
	if err := cr.DB.Raw("select count(*) from carts where user_id = ? ", userID).First(&count).Error; err != nil {
		return []models.Cart{}, err
	}
	var cartResponse []models.Cart
	if count == 0 {
		return []models.Cart{}, nil
	}

	if err := cr.DB.Exec("delete from carts where user_id = ? ", userID).Error; err != nil {
		return []models.Cart{}, err
	}

	// if err := cr.DB.Raw("select carts.user_id,users.name as user_name,carts.product_id,products.movie_name as movie_name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?",userID).First(&cartResponse).Error; err != nil {
	// 	return []models.Cart{},err
	// }

	return cartResponse, nil

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
	fmt.Println("product count", count, product_id)

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
	fmt.Print(count)
	return count > 0, nil
}

func (cr *cartRepository) CouponValidity(coupon string, userID int) (bool, error) {

	// if cart is empty don't let them add coupon

	// check if the coupon exist
	var count int
	err := cr.DB.Raw("select count(*) from coupons where coupon = ?", coupon).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count < 1 {
		return false, errors.New("coupon does not exist")
	}

	// check if the coupon have been revoked or not
	var validity bool
	err = cr.DB.Raw("select validity from coupons where coupon = ?", coupon).Scan(&validity).Error
	if err != nil {
		return false, err
	}

	if !validity {
		return false, errors.New("coupon not valid")
	}

	var MinDiscountPrice float64

	err = cr.DB.Raw("select minimum_price from coupons where coupon = ?", coupon).Scan(&MinDiscountPrice).Error
	if err != nil {
		return false, err
	}

	var totalPrice float64
	err = cr.DB.Raw("select COALESCE(SUM(total_price), 0) from carts where user_id = ?", userID).Scan(&totalPrice).Error
	if err != nil {
		return false, err
	}

	// if the total Price is less than minDiscount price don't allow coupon to be added
	if totalPrice < MinDiscountPrice {
		return false, errors.New("coupon cannot be added as the total amount is less than minimum amount for coupon")
	}

	var couponID uint
	err = cr.DB.Raw("select id from coupons where coupon = ?", coupon).Scan(&couponID).Error
	if err != nil {
		return false, err
	}

	// to check if used have already used this coupon
	err = cr.DB.Raw("select count(*) from used_coupons where coupon_id = ? and user_id = ?", couponID, userID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return false, errors.New("user have already used this coupon")
	}

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
