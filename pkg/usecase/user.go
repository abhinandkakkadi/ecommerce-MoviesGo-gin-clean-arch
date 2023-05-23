package usecase

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	domain "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/jinzhu/copier"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) UserSignUp(user domain.Users) (domain.TokenUsers,error) {
	
	// Check whether the user already exist. If yes, show the error message, since this is signUp
	userExist := c.userRepo.CheckUserAvailability(user)
	
	if userExist {
		return domain.TokenUsers{},errors.New("user already exist, sign in")
	}

	// Hash password since details are validated
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(user.Password),10)
	if err != nil {
		return domain.TokenUsers{},errors.New("internal server error")
	}
	user.Password = string(hashedPassword)

	// add user details to the database
	userData, err := c.userRepo.UserSignUp(user)
	if err != nil {
		return domain.TokenUsers{},err
	}

	// crete a JWT token string for the user
	tokenString,err := helper.GenerateTokenUsers(user)
	if err != nil {
		return domain.TokenUsers{},errors.New("could not create token due to some internal error")
	}

	// copies all the details except the password of the user 
	var userDetails models.UserDetails
	err = copier.Copy(&userDetails,&userData)
	if err != nil {
		return domain.TokenUsers{},err
	}

	return domain.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	},nil
}

func (c *userUseCase) LoginHandler(user domain.Users) (domain.TokenUsers,error) {

	// checking if a username exist with this email address
	ok := c.userRepo.CheckUserAvailability(user)
	if !ok {
		return domain.TokenUsers{},errors.New("the user does not exist")
	}

	// Get the user details in order to check the password, in this case ( The same function can be reused in future )
	user_details,err := c.userRepo.FindUserByEmail(user)
	if err != nil {
		return domain.TokenUsers{},err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user_details.Password),[]byte(user.Password))
	if err != nil {
		return domain.TokenUsers{},errors.New("password incorrect")
	}

	tokenString,err := helper.GenerateTokenUsers(user)
	if err != nil {
		return domain.TokenUsers{},errors.New("could not create token")
	}

	var userDetails models.UserDetails
	err = copier.Copy(&userDetails,&user_details)
	if err != nil {
		return domain.TokenUsers{},err
	}
	
	return domain.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	},nil


}


