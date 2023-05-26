package usecase

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/jinzhu/copier"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
	cartRepo interfaces.CartRepository
}

func NewUserUseCase(repo interfaces.UserRepository, cartRepositiry interfaces.CartRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
		cartRepo: cartRepositiry,
	}
}

func (c *userUseCase) UserSignUp(user models.UserDetails) (models.TokenUsers, error) {

	// Check whether the user already exist. If yes, show the error message, since this is signUp
	userExist := c.userRepo.CheckUserAvailability(user)

	if userExist {
		return models.TokenUsers{}, errors.New("user already exist, sign in")
	}

	// if user.Password != user.ConfirmPassword {
	// 	return domain.TokenUsers{},errors.New("password does not match")
	// }

	// Hash password since details are validated
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return models.TokenUsers{}, errors.New("internal server error")
	}
	user.Password = string(hashedPassword)

	// add user details to the database
	userData, err := c.userRepo.UserSignUp(user)
	if err != nil {
		return models.TokenUsers{}, err
	}

	// crete a JWT token string for the user
	tokenString, err := helper.GenerateTokenUsers(userData)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	// copies all the details except the password of the user
	var userDetails models.UserDetailsResponse
	err = copier.Copy(&userDetails, &userData)
	if err != nil {
		return models.TokenUsers{}, err
	}

	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil
}

func (c *userUseCase) LoginHandler(user models.UserDetails) (models.TokenUsers, error) {

	// checking if a username exist with this email address
	ok := c.userRepo.CheckUserAvailability(user)
	if !ok {
		return models.TokenUsers{}, errors.New("the user does not exist")
	}

	// Get the user details in order to check the password, in this case ( The same function can be reused in future )
	user_details, err := c.userRepo.FindUserByEmail(user)
	if err != nil {
		return models.TokenUsers{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user_details.Password), []byte(user.Password))
	if err != nil {
		return models.TokenUsers{}, errors.New("password incorrect")
	}

	var userDetails models.UserDetailsResponse
	err = copier.Copy(&userDetails, &user_details)
	if err != nil {
		return models.TokenUsers{}, err
	}

	tokenString, err := helper.GenerateTokenUsers(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token")
	}

	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil

}

func (cr *userUseCase) AddAddress(address models.AddressInfo, userID int) ([]models.AddressInfoResponse, error) {

	addressResponse, err := cr.userRepo.AddAddress(address, userID)
	if err != nil {
		return []models.AddressInfoResponse{}, err
	}

	return addressResponse, nil
}

func (cr *userUseCase) UpdateAddress(address models.AddressInfo, addressID int) (models.AddressInfoResponse, error) {

	return cr.userRepo.UpdateAddress(address, addressID)
}

func (cr *userUseCase) Checkout(userID int) (models.CheckoutDetails, error) {

	allUserAddress, err := cr.userRepo.GetAllAddresses(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	paymentDetails, err := cr.userRepo.GetAllPaymentOption()
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	cartItems, err := cr.cartRepo.GetAllItemsFromCart(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	grandTotal, err := cr.cartRepo.GetTotalPrice(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	return models.CheckoutDetails{
		AddressInfoResponse: allUserAddress,
		Payment_Method:      paymentDetails,
		Cart:                cartItems,
		Grand_Total:         grandTotal.TotalPrice,
	}, nil
}

func (cr *userUseCase) UserDetails(userID int) (models.UsersProfileDetails, error) {

	return cr.userRepo.UserDetails(userID)

}


func (cr *userUseCase) GetAllAddress(userID int) ([]models.AddressInfoResponse,error) {

	userAddress,err := cr.userRepo.GetAllAddresses(userID)

	if err != nil {
		return []models.AddressInfoResponse{},nil
	}	

	return userAddress,nil
	
}