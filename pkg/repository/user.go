package repository

import (
	"errors"
	"fmt"

	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"gorm.io/gorm"
)

type UserDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &UserDatabase{DB}
}

// check whether the user is already present in the database . If there recommend to login
func (c *UserDatabase) CheckUserAvailability(email string) bool {

	var count int
	query := fmt.Sprintf("select count(*) from users where email='%s'", email)
	if err := c.DB.Raw(query).Scan(&count).Error; err != nil {
		return false
	}
	// if count is greater than 0 that means the user already exist
	return count > 0

}

// retrieve the user details form the database
func (c *UserDatabase) FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error) {

	var userDetails models.UserSignInResponse

	err := c.DB.Raw(`
		SELECT *
		FROM users where email = ? and blocked = false
		`, user.Email).Scan(&userDetails).Error

	if err != nil {
		return models.UserSignInResponse{}, errors.New("error checking user details")
	}

	return userDetails, nil

}

func (c *UserDatabase) UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error) {

	var userDetails models.UserDetailsResponse
	err := c.DB.Raw(`INSERT INTO users (name, email, phone, password) VALUES ($1, $2, $3, $4) RETURNING id, name, email, phone`, user.Name, user.Email, user.Phone, user.Password).Scan(&userDetails).Error

	if err != nil {
		return models.UserDetailsResponse{}, err
	}

	return userDetails, nil
}

func (c *UserDatabase) LoginHandler(user models.UserDetails) (models.UserDetailsResponse, error) {

	var userResponse models.UserDetailsResponse
	err := c.DB.Save(&userResponse).Error
	return userResponse, err

}

func (cr *UserDatabase) AddAddress(address models.AddressInfo, userID int) error {

	fmt.Println(address)
	err := cr.DB.Exec("insert into addresses (user_id,name,house_name,state,pin,street,city) values (?, ?, ?, ?, ?, ?, ?)", userID, address.Name, address.HouseName, address.State, address.Pin, address.Street, address.City).Error
	if err != nil {
		return err
	}

	return nil

}

func (cr *UserDatabase) UpdateAddress(address models.AddressInfo, addressID int, userID int) (models.AddressInfoResponse, error) {

	err := cr.DB.Exec("update addresses set house_name = ?, state = ?, pin = ?, street = ?, city = ? where id = ? and user_id = ?", address.HouseName, address.State, address.Pin, address.Street, address.City, addressID, userID).Error
	if err != nil {
		return models.AddressInfoResponse{}, err
	}

	var addressResponse models.AddressInfoResponse
	err = cr.DB.Raw("select * from addresses where id = ?", addressID).Scan(&addressResponse).Error
	if err != nil {
		return models.AddressInfoResponse{}, err
	}

	return addressResponse, nil

}

func (cr *UserDatabase) GetAllAddresses(userID int) ([]models.AddressInfoResponse, error) {

	var addressResponse []models.AddressInfoResponse
	err := cr.DB.Raw(`select * from addresses where user_id = $1`, userID).Scan(&addressResponse).Error
	if err != nil {
		return []models.AddressInfoResponse{}, err
	}

	return addressResponse, nil

}

func (cr *UserDatabase) GetAllPaymentOption() ([]models.PaymentDetails, error) {

	var paymentMethods []models.PaymentDetails
	err := cr.DB.Raw("select * from payment_methods").Scan(&paymentMethods).Error
	if err != nil {
		return []models.PaymentDetails{}, err
	}

	return paymentMethods, nil

}

func (cr *UserDatabase) GetWalletDetails(userID int) (models.Wallet, error) {

	var walletDetails models.Wallet
	err := cr.DB.Raw("select wallet_amount from wallets where user_id = ?", userID).Scan(&walletDetails).Error
	if err != nil {
		return models.Wallet{}, err
	}

	return walletDetails, nil

}

func (cr *UserDatabase) UserDetails(userID int) (models.UsersProfileDetails, error) {

	var userDetails models.UsersProfileDetails
	err := cr.DB.Raw("select users.name,users.email,users.phone,referrals.referral_code from users inner join referrals on users.id = referrals.user_id where users.id = ?", userID).Row().Scan(&userDetails.Name, &userDetails.Email, &userDetails.Phone, &userDetails.ReferralCode)
	if err != nil {
		return models.UsersProfileDetails{}, err
	}

	return userDetails, nil
}

func (cr *UserDatabase) UpdateUserEmail(email string, userID int) error {

	err := cr.DB.Exec("update users set email = ? where id = ?", email, userID).Error
	if err != nil {
		return err
	}
	return nil

}

func (cr *UserDatabase) UpdateUserPhone(phone string, userID int) error {

	err := cr.DB.Exec("update users set phone = ? where id = ?", phone, userID).Error
	if err != nil {
		return err
	}
	return nil

}

func (cr *UserDatabase) UpdateUserName(name string, userID int) error {

	err := cr.DB.Exec("update users set name = ? where id = ?", name, userID).Error
	if err != nil {
		return err
	}
	return nil

}

func (cr *UserDatabase) UpdateUserPassword(password string, userID int) error {

	err := cr.DB.Exec("update users set password = ? where id = ?", password, userID).Error
	if err != nil {
		return err
	}
	return nil

}

func (cr *UserDatabase) UserPassword(userID int) (string, error) {

	var userPassword string
	err := cr.DB.Raw("select password from users where id = ?", userID).Scan(&userPassword).Error
	if err != nil {
		return "", err
	}
	return userPassword, nil

}

func (cr *UserDatabase) FindUserByOrderID(orderId string) (models.UsersProfileDetails, error) {

	var userDetails models.UsersProfileDetails
	err := cr.DB.Raw("select users.name,users.email,users.phone from users inner join orders on orders.user_id = users.id where order_id = ?", orderId).Scan(&userDetails).Error
	if err != nil {
		return models.UsersProfileDetails{}, err
	}

	return userDetails, nil
}

// get the shipping address of
func (cr *UserDatabase) FindUserAddressByOrderID(orderID string) (models.AddressInfo, error) {

	var shipmentAddress models.AddressInfo
	err := cr.DB.Raw("select addresses.name,addresses.house_name,addresses.street,addresses.city,addresses.state,addresses.pin from addresses inner join orders on orders.address_id = addresses.id where order_id = ?", orderID).Scan(&shipmentAddress).Error
	if err != nil {
		return models.AddressInfo{}, err
	}

	return shipmentAddress, nil
}

func (cr *UserDatabase) UserBlockStatus(email string) (bool, error) {

	var isBlocked bool
	err := cr.DB.Raw("select blocked from users where email = ?", email).Scan(&isBlocked).Error
	if err != nil {
		return false, err
	}

	return isBlocked, nil
}

func (cr *UserDatabase) ProductExistInWishList(productID int, userID int) (bool, error) {

	var count int
	err := cr.DB.Raw("select count(*) from wish_lists where user_id = ? and product_id = ? ", userID, productID).Scan(&count).Error
	if err != nil {
		return false, errors.New("error checking user product already present")
	}

	return count > 0, nil

}

func (cr *UserDatabase) AddToWishList(userID int, productID int) error {

	err := cr.DB.Exec("insert into wish_lists (user_id,product_id) values (?, ?)", userID, productID).Error
	if err != nil {
		return err
	}

	return nil
}

func (cr *UserDatabase) GetWishList(userID int) ([]models.WishListResponse, error) {

	var wishList []models.WishListResponse
	err := cr.DB.Raw("select products.id as product_id, products.movie_name as product_name,products.price as product_price from products inner join wish_lists on products.id = wish_lists.product_id where wish_lists.user_id = ? ", userID).Scan(&wishList).Error
	if err != nil {
		return []models.WishListResponse{}, err
	}

	return wishList, nil

}

func (cr *UserDatabase) RemoveFromWishList(userID int, productID int) error {

	err := cr.DB.Exec("delete from wish_lists where user_id = ? and product_id = ?", userID, productID).Error
	if err != nil {
		return err
	}

	return nil

}

func (cr *UserDatabase) CreateReferralEntry(userDetails models.UserDetailsResponse, userReferral string) error {

	err := cr.DB.Exec("insert into referrals (user_id,referral_code,referral_amount) values (?,?,?)", userDetails.Id, userReferral, 0).Error
	if err != nil {
		return err
	}

	return nil

}

func (cr *UserDatabase) GetUserIdFromReferrals(ReferralCode string) (int, error) {

	var referredUserId int
	err := cr.DB.Raw("select user_id from referrals where referral_code = ?", ReferralCode).Scan(&referredUserId).Error
	if err != nil {
		return 0, nil
	}

	return referredUserId, nil
}

func (cr *UserDatabase) UpdateReferralAmount(referralAmount float64, referredUserId int, currentUserID int) error {

	err := cr.DB.Exec("update referrals set referral_amount = ?,referred_user_id = ? where user_id = ? ", referralAmount, referredUserId, currentUserID).Error
	if err != nil {
		return err
	}

	// find the current amount in referred users referral table and add 100 with that
	err = cr.DB.Exec("update referrals set referral_amount = referral_amount + ? where user_id = ? ", referralAmount, referredUserId).Error
	if err != nil {
		return err
	}

	return nil

}

func (cr *UserDatabase) GetReferralAndTotalAmount(userID int) (float64, float64, error) {

	// first check whether the cart is empty -- do this for coupon too
	var cartDetails struct {
		ReferralAmount  float64
		TotalCartAmount float64
	}

	err := cr.DB.Raw("SELECT (SELECT referral_amount FROM referrals WHERE user_id = ?) AS referral_amount, COALESCE(SUM(total_price), 0) AS total_cart_amount FROM carts WHERE user_id = ?", userID, userID).Scan(&cartDetails).Error
	if err != nil {
		return 0.0, 0.0, err
	}

	return cartDetails.ReferralAmount, cartDetails.TotalCartAmount, nil

}

func (cr *UserDatabase) UpdateSomethingBasedOnUserID(tableName string, columnName string, updateValue float64, userID int) error {

	err := cr.DB.Exec("update "+tableName+" set "+columnName+" = ? where user_id = ?", updateValue, userID).Error
	if err != nil {
		cr.DB.Rollback()
		return err
	}
	return nil

}

func (cr *UserDatabase) ResetPassword(userID int, password string) error {

	err := cr.DB.Exec("update users set password = ? where id = ?", password, userID).Error
	if err != nil {
		return err
	}

	return nil
}
