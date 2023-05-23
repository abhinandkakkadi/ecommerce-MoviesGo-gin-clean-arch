package repository

import (
	"errors"
	"fmt"

	domain "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

// check whether the user is already present in the database . If there recommend to login
	func (c *userDatabase) CheckUserAvailability(user domain.Users) bool {
		
		var count int
		query := fmt.Sprintf("select count(*) from users where email='%s'",user.Email)
		if err := c.DB.Raw(query).Scan(&count).Error; err != nil {
			return false
		}
		// if count is greater than 0 that means the user already exist
		return count > 0
		
	}

// retrieve the user details form the database
	func (c *userDatabase) FindUserByEmail(user domain.Users) (domain.Users,error) {

		var user_details domain.Users

		err := c.DB.Raw(`
		SELECT *
		FROM users where email = ? and blocked = false
		`,user.Email).Scan(&user_details).Error

		if err != nil {
		return domain.Users{},errors.New("error checking user details")
		}

		return user_details,nil

	}

	func (c *userDatabase) UserSignUp(user domain.Users) (domain.Users,error) {
		
		query := fmt.Sprintf("insert into users (id,name,email,password,phone) values ('%d','%s','%s','%s','%s')",user.ID,user.Name,user.Email,user.Password,user.Phone)
	
		if err := c.DB.Exec(query).Error; err != nil {
			return domain.Users{},err
		}

		return user,nil
}

func (c *userDatabase) LoginHandler(user domain.Users) (domain.Users,error) {
	err := c.DB.Save(&user).Error
	return user,err
}

