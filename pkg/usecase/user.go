package usecase

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/copier"
	domain "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) GenerateUser(ctx context.Context,user domain.Users) (domain.TokenUsers,error) {
	
	// Check whether the user already exist. If yes, show the error message, since this is signUp
	userExist := c.userRepo.CheckUserAvailability(user)
	
	if userExist {
		return domain.TokenUsers{},errors.New("user already exist, sign in")
	}

	// Hash password
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(user.Password),10)
	if err != nil {
		return domain.TokenUsers{},errors.New("Internal server error")
	}
	user.Password = string(hashedPassword)



	userData, err := c.userRepo.GenerateUser(user)
	if err != nil {
		return domain.TokenUsers{},err
	}

	tokenString,err := helper.GenerateTokenUsers(user)

	if err != nil {
		return domain.TokenUsers{},errors.New("could not create token")
	}

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

func (c *userUseCase) LoginHandler(ctx context.Context,user domain.Users) (domain.TokenUsers,error) {

	// checking if a username exist with this email address
	ok := c.userRepo.CheckUserAvailability(user)
	if !ok {
		return domain.TokenUsers{},errors.New("The user does not exist")
	}

	// Get the user details in order to check the password, in this case ( The same function can be reused in future )
	user_details,err := c.userRepo.FindUserByEmail(user)
	if err != nil {
		return domain.TokenUsers{},err
	}

	fmt.Println(user)
	fmt.Println(user_details)

	err = bcrypt.CompareHashAndPassword([]byte(user_details.Password),[]byte(user.Password))
	if err != nil {
		return domain.TokenUsers{},errors.New("Password incorrect")
	}

	tokenString,err := helper.GenerateTokenUsers(user)

	if err != nil {
		return domain.TokenUsers{},errors.New("could not create token")
	}

	var userDetails models.UserDetails
	// err = mapstructure.Decode(&userDetails,&user_details)
	// if err != nil {
	// 	return domain.TokenUsers{},errors.New("internal server error")
	// }

	err = copier.Copy(&userDetails,&user_details)
	if err != nil {
		return domain.TokenUsers{},err
	}
	
	return domain.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	},nil

	// user, err = c.userRepo.LoginHandler(ctx,user)
	// return user,err
}

func (c *userUseCase) FindAll(ctx context.Context) ([]domain.Users, error) {
	users, err := c.userRepo.FindAll(ctx)
	return users, err
}

func (c *userUseCase) FindByID(ctx context.Context, id uint) (domain.Users, error) {
	user, err := c.userRepo.FindByID(ctx, id)
	return user, err
}

func (c *userUseCase) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	user, err := c.userRepo.Save(ctx, user)

	return user, err
}

func (c *userUseCase) Delete(ctx context.Context, user domain.Users) error {
	err := c.userRepo.Delete(ctx, user)

	return err
}
