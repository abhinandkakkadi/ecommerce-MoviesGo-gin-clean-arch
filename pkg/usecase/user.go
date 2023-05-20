package usecase

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	"github.com/thnkrn/go-gin-clean-arch/pkg/helper"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
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

	tokenString,err := helper.GenerateToken(user)

	if err != nil {
		return domain.TokenUsers{},errors.New("could not create token")
	}
	


	return domain.TokenUsers{
		Users: userData,
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

	err = bcrypt.CompareHashAndPassword([]byte(user_details.Password),[]byte(user.Password))
	if err != nil {
		return domain.TokenUsers{},errors.New("Password incorrect")
	}

	tokenString,err := helper.GenerateToken(user)

	if err != nil {
		return domain.TokenUsers{},errors.New("could not create token")
	}
	
	return domain.TokenUsers{
		Users: user_details,
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
