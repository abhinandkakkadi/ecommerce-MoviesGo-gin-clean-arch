package usecase

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/thnkrn/go-gin-clean-arch/pkg/helper"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) GenerateUser(ctx context.Context,user domain.Users) (domain.Users,error) {
	
	// Check whether the user already exist. If yes, show the error message, since this is signUp
	userExist := c.userRepo.CheckUserAvailability(user)
	
	if userExist {
		return domain.Users{},errors.New("user already exist, sign in")
	}

	// Hash password
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(user.Password),10)
	if err != nil {
		return domain.Users{},errors.New("Internal server error")
	}
	user.Password = string(hashedPassword)



	userData, err := c.userRepo.GenerateUser(user)
	if err != nil {
		return domain.Users{},err
	}

	tokenString,err := helper.GenerateToken(user)
	if err 

	return userData,nil
}

func (c *userUseCase) LoginHandler(ctx context.Context,user domain.Users) (domain.Users,error) {
	user, err := c.userRepo.LoginHandler(ctx,user)
	return user,err
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
