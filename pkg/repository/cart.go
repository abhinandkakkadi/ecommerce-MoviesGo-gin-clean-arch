package repository

import (
	"errors"
	"fmt"

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

func (cr *cartRepository) AddToCart(product_id int, userID int) ([]models.Cart, error) {

	var cartResponse []models.Cart

	tx := cr.DB.Begin()
	var count int
	if err := tx.Raw("select count(*) from carts where user_id = ? and  product_id = ?", userID, product_id).Scan(&count).Error; err != nil {
		tx.Rollback()
		return []models.Cart{}, err
	}

	// if the user try to add more quantity than what is already present in stock
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

	if productQuantity == 0 {
		return []models.Cart{}, nil
	}
	// quantity in carts is equal to quantity in STOCK  -- don't allow to add further products
	if cartsQuantity == productQuantity {

		if err := tx.Raw("select carts.user_id,users.name as user_name,carts.product_id,products.movie_name as movie_name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userID).First(&cartResponse).Error; err != nil {
			tx.Rollback()
			return []models.Cart{}, err
		}
		return cartResponse, nil
	}

	var totalPrice float64
	var productPrice float64
	if err := tx.Raw("select price from products where id = ?", product_id).Scan(&productPrice).Error; err != nil {
		tx.Rollback()
		return []models.Cart{}, err
	}

	fmt.Println(totalPrice)

	if count == 0 {
		totalPrice = productPrice
		fmt.Println(totalPrice)
		if err := tx.Exec("insert into carts (user_id,product_id,quantity,total_price) values(?,?,?,?)", userID, product_id, 1, totalPrice).Error; err != nil {
			tx.Rollback()
			return []models.Cart{}, err
		}
		fmt.Println("the above thing worked")
	} else {

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

	if err := tx.Raw("select carts.user_id,users.name as user_name,carts.product_id,products.movie_name as movie_name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userID).First(&cartResponse).Error; err != nil {
		tx.Rollback()
		return []models.Cart{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return []models.Cart{}, err
	}

	return cartResponse, nil

}

func (cr *cartRepository) GetTotalPrice(userID int) (models.CartTotal, error) {

	var cartTotal models.CartTotal
	// return 0 id the record is not present
	err := cr.DB.Raw("select COALESCE(SUM(total_price), 0) from carts where user_id = ?", userID).Scan(&cartTotal.TotalPrice).Error
	if err != nil {
		return models.CartTotal{}, err
	}

	err = cr.DB.Raw("select name as user_name from users where id = ?", userID).Scan(&cartTotal.UserName).Error
	if err != nil {
		return models.CartTotal{}, err
	}

	return cartTotal, nil
}

func (cr *cartRepository) RemoveFromCart(product_id int, userID int) ([]models.Cart, error) {

	tx := cr.DB.Begin()

	var count int
	if err := tx.Raw("select count(*) from carts where user_id = ? and product_id = ?", userID, product_id).Scan(&count).Error; err != nil {
		tx.Rollback()
		return []models.Cart{}, err
	}

	if count == 0 {
		return []models.Cart{}, nil
	}

	var cartDetails struct {
		Quantity   int
		TotalPrice float64
		Price      float64
	}

	if err := tx.Raw("select quantity,total_price from carts where user_id = ? and product_id = ?", userID, product_id).Scan(&cartDetails).Error; err != nil {
		tx.Rollback()
		return []models.Cart{}, err
	}

	cartDetails.Quantity = cartDetails.Quantity - 1
	if cartDetails.Quantity == 0 {
		if err := tx.Exec("delete from carts where user_id = ? and product_id = ?", userID, product_id).Error; err != nil {
			tx.Rollback()
			return []models.Cart{}, err
		}
	}

	fmt.Println("quantity and totalPrice = ", cartDetails)
	if err := tx.Raw("select price from products where id = ?", product_id).Scan(&cartDetails.Price).Error; err != nil {
		tx.Rollback()
		return []models.Cart{}, err
	}

	cartDetails.TotalPrice = cartDetails.TotalPrice - cartDetails.Price

	if err := tx.Exec("update carts set quantity = ?,total_price = ? where user_id = ? and product_id = ?", cartDetails.Quantity, cartDetails.TotalPrice, userID, product_id).Error; err != nil {
		tx.Rollback()
		return []models.Cart{}, err
	}

	// if err := cr.DB.Exec("delete from carts where user_id = ? and product_id = ?", userID, product_id).Error; err != nil {
	// 	return []models.Cart{}, err
	// }

	if err := tx.Raw("select count(*) from carts where user_id = ? ", userID).Scan(&count).Error; err != nil {
		tx.Rollback()
		return []models.Cart{}, err
	}

	if count == 0 {
		return []models.Cart{}, nil
	}

	var cartResponse []models.Cart

	if err := tx.Raw("select carts.user_id,users.name as user_name,carts.product_id,products.movie_name as movie_name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userID).First(&cartResponse).Error; err != nil {
		tx.Rollback()
		return []models.Cart{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
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

func (cr *cartRepository) CheckProduct(product_id int) (bool, error) {

	var count int
	err := cr.DB.Raw("select count(*) from products where id = ?", product_id).Scan(&count).Error
	if err != nil {
		return false, err
	}
	fmt.Println("product count", count, product_id)
	return count > 0, nil

}
