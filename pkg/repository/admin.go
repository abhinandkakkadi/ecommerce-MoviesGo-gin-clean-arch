package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"gorm.io/gorm"
)



type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB)  interfaces.AdminRepository {
	return &adminRepository{
		DB: DB,
	}
}

func (cr *adminRepository) LoginHandler(c context.Context,adminDetails domain.Admin) (domain.Admin,error) {
	
	var adminCompareDetails domain.Admin
	if err := cr.DB.Raw("select * from admins where email = ? ",adminDetails.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.Admin{},err
	}
	
	return adminCompareDetails,nil
}

func (cr *adminRepository) CheckAdminAvailability(admin domain.Admin) bool {

	var count int
	if err := cr.DB.Raw("select count(*) from admins where email = ?",admin.Email).Scan(&count).Error; err != nil {
		return false
	}

	return count > 0
}



func (cr *adminRepository) SignupHandler(c context.Context,admin domain.Admin) (domain.Admin,error) {

	query := fmt.Sprintf("insert into admins (id,name,email,password) values ('%d','%s','%s','%s')",admin.ID,admin.Name,admin.Email,admin.Password)
	
	if err := cr.DB.Exec(query).Error; err != nil {
		return domain.Admin{},err
	}

	return admin,nil
}

func (cr *adminRepository) GetUsers(c context.Context) ([]models.UserDetails,error) {

	var userDetails []models.UserDetails
	
	if err := cr.DB.Raw("select id,name,email,phone from users").Scan(&userDetails).Error; err != nil {
		return []models.UserDetails{},err
	}

	return userDetails,nil
}

func (cr *adminRepository) GetGenres(c context.Context) ([]domain.Genre,error) {

	var genres []domain.Genre
	if err := cr.DB.Raw("select * from genres").Scan(&genres).Error; err != nil {
		return []domain.Genre{},err
	}

	return genres,nil
}

func (cr *adminRepository) AddGenre(c context.Context,genre domain.Genre) (domain.Genre,error) {

	if err := cr.DB.Exec("insert into genres (id,genre_name) values (?,?)",genre.ID,genre.Genre_Name).Error; err != nil {
		return domain.Genre{},err
	}
	return genre,nil
}


func (cr *adminRepository) Delete(c context.Context,genre_id string) error {
	
	id,_ := strconv.Atoi(genre_id)

	query := fmt.Sprintf("delete from genres where id = '%d'",id)
	
	if err := cr.DB.Exec(query).Error; err != nil {
		return err
	}

	return nil
}




func (cr *adminRepository) GetUserByID(c context.Context,id string) (domain.Users,error) {

	user_id,_ := strconv.Atoi(id)

	query := fmt.Sprintf("select * from users where id = '%d'",user_id)
	var userDetails domain.Users

	if err := cr.DB.Raw(query).Scan(&userDetails).Error; err != nil {
		return domain.Users{},err
	}

	return userDetails,nil
}

func (cr *adminRepository) UpdateUserByID(c context.Context,user domain.Users) error {
	fmt.Println(user)
	err := cr.DB.Exec("update users set blocked = ? where id = ?", user.Blocked, user.ID).Error
	if err != nil {
		fmt.Println("Error updating user:", err)
		return err
	}

	return nil
}