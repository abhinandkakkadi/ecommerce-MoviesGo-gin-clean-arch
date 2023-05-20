package repository

import (
	"context"
	"errors"
	"fmt"

	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

// check whether the user is already present in the database
	func (c *userDatabase) CheckUserAvailability(user domain.Users) bool {
		
		var count int
		query := fmt.Sprintf("select count(*) from users where email='%s'",user.Email)
		if err := c.DB.Raw(query).Scan(&count).Error; err != nil {
			return false
		}
		return count > 0
		
	}

// retreive the user details form the database

	func (c *userDatabase) FindUserByEmail(user domain.Users) (domain.Users,error) {

		var user_details domain.Users

		err := c.DB.Raw(`
		SELECT *
		FROM users where email = ?
		`,user.Email).Scan(&user_details).Error

		if err != nil {
		return domain.Users{},errors.New("Error checking user details")
		}

		return user_details,nil

	}

func (c *userDatabase) GenerateUser(user domain.Users) (domain.Users,error) {
	query := fmt.Sprintf("insert into users (id,name,email,password,phone) values ('%d','%s','%s','%s','%s')",user.ID,user.Name,user.Email,user.Password,user.Phone)
	
	if err := c.DB.Exec(query).Error; err != nil {
		return domain.Users{},err
	}

	return user,nil
}

func (c *userDatabase) LoginHandler(ctx context.Context, user domain.Users) (domain.Users,error) {
	err := c.DB.Save(&user).Error
	return user,err
}

func (c *userDatabase) FindAll(ctx context.Context) ([]domain.Users, error) {
	var users []domain.Users
	err := c.DB.Find(&users).Error

	return users, err
}

func (c *userDatabase) FindByID(ctx context.Context, id uint) (domain.Users, error) {
	var user domain.Users
	err := c.DB.First(&user, id).Error

	return user, err
}

func (c *userDatabase) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	err := c.DB.Save(&user).Error

	return user, err
}

func (c *userDatabase) Delete(ctx context.Context, user domain.Users) error {
	err := c.DB.Delete(&user).Error

	return err
}
