package repository

import (
	"errors"
	"fmt"

	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

// check whether the user is already present in the database . If there recommend to login
func (c *userDatabase) CheckUserAvailability(user models.UserDetails) bool {

	var count int
	query := fmt.Sprintf("select count(*) from users where email='%s'", user.Email)
	if err := c.DB.Raw(query).Scan(&count).Error; err != nil {
		return false
	}
	// if count is greater than 0 that means the user already exist
	return count > 0

}

// retrieve the user details form the database
func (c *userDatabase) FindUserByEmail(user models.UserDetails) (models.UserSignInResponse, error) {

	var user_details models.UserSignInResponse

	err := c.DB.Raw(`
		SELECT *
		FROM users where email = ? and blocked = false
		`, user.Email).Scan(&user_details).Error

	if err != nil {
		return models.UserSignInResponse{}, errors.New("error checking user details")
	}

	return user_details, nil

}

func (c *userDatabase) UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error) {

	var userDetails models.UserDetailsResponse
	err := c.DB.Raw("INSERT INTO users (name, email, password, phone) VALUES (?, ?, ?, ?) RETURNING id, name, email, phone", user.Name, user.Email, user.Password, user.Phone).Scan(&userDetails).Error

	if err != nil {
		return models.UserDetailsResponse{}, err
	}

	return userDetails, nil
}

func (c *userDatabase) LoginHandler(user models.UserDetails) (models.UserDetailsResponse, error) {
	var userResponse models.UserDetailsResponse
	err := c.DB.Save(&userResponse).Error
	return userResponse, err
}

func (cr *userDatabase) AddAddress(address models.AddressInfo, userID int) ([]models.AddressInfoResponse, error) {

	address.UserID = uint(userID)
	fmt.Println(address)
	err := cr.DB.Exec("insert into addresses (user_id,house_name,state,pin,street,city) values (?, ?, ?, ?, ?, ?)", address.UserID, address.HouseName, address.State, address.Pin, address.Street, address.City).Error
	if err != nil {
		return []models.AddressInfoResponse{}, err
	}

	var addressResponse []models.AddressInfoResponse
	err = cr.DB.Raw("select * from addresses where user_id = ?", address.UserID).Scan(&addressResponse).Error
	if err != nil {
		return []models.AddressInfoResponse{}, err
	}

	return addressResponse, nil

}

func (cr *userDatabase) UpdateAddress(address models.AddressInfo, addressID int) (models.AddressInfoResponse, error) {

	fmt.Println(address)
	err := cr.DB.Exec("update addresses set house_name = ?, state = ?, pin = ?, street = ?, city = ? where id = ? and user_id = ?", address.HouseName, address.State, address.Pin, address.Street, address.City, addressID, address.UserID).Error
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

func (cr *userDatabase) GetAllAddresses(userID int) ([]models.AddressInfoResponse, error) {

	var addressResponse []models.AddressInfoResponse
	err := cr.DB.Raw("select * from addresses where user_id = ?", userID).Scan(&addressResponse).Error
	if err != nil {
		return []models.AddressInfoResponse{}, err
	}

	return addressResponse, nil

}

func (cr *userDatabase) GetAllPaymentOption() ([]models.PaymentDetails, error) {

	var paymentMethods []models.PaymentDetails
	err := cr.DB.Raw("select * from payment_methods").Scan(&paymentMethods).Error
	if err != nil {
		return []models.PaymentDetails{}, err
	}

	return paymentMethods, nil

}
