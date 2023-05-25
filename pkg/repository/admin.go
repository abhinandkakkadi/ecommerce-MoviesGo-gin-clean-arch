package repository

import (
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

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminRepository{
		DB: DB,
	}
}

func (cr *adminRepository) LoginHandler(adminDetails domain.Admin) (domain.Admin, error) {

	var adminCompareDetails domain.Admin
	if err := cr.DB.Raw("select * from admins where email = ? ", adminDetails.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.Admin{}, err
	}

	return adminCompareDetails, nil
}

func (cr *adminRepository) CheckAdminAvailability(admin domain.Admin) bool {

	var count int
	if err := cr.DB.Raw("select count(*) from admins where email = ?", admin.Email).Scan(&count).Error; err != nil {
		return false
	}

	return count > 0

}

func (cr *adminRepository) SignUpHandler(admin domain.Admin) (domain.Admin, error) {

	query := fmt.Sprintf("insert into admins (id,name,email,password) values ('%d','%s','%s','%s')", admin.ID, admin.Name, admin.Email, admin.Password)
	if err := cr.DB.Exec(query).Error; err != nil {
		return domain.Admin{}, err
	}

	return admin, nil

}

// Get users details for authenticated admins
func (cr *adminRepository) GetUsers() ([]models.UserDetailsResponse, error) {

	var userDetails []models.UserDetailsResponse

	if err := cr.DB.Raw("select id,name,email,phone from users").Scan(&userDetails).Error; err != nil {
		return []models.UserDetailsResponse{}, err
	}

	return userDetails, nil

}

func (cr *adminRepository) GetGenres() ([]domain.Genre, error) {

	var genres []domain.Genre
	if err := cr.DB.Raw("select * from genres").Scan(&genres).Error; err != nil {
		return []domain.Genre{}, err
	}

	return genres, nil

}

// CATEGORY MANAGEMENT
func (cr *adminRepository) AddGenre(genre domain.Genre) error {

	if err := cr.DB.Exec("insert into genres (id,genre_name) values (?,?)", genre.ID, genre.Genre_Name).Error; err != nil {
		return err
	}
	return nil

}

func (cr *adminRepository) AddDirector(director domain.Directors) error {

	if err := cr.DB.Exec("insert into directors (id,director_name) values (?,?)", director.ID, director.Director_Name).Error; err != nil {
		return err
	}
	return nil

}

func (cr *adminRepository) AddFormat(format domain.Movie_Format) error {

	if err := cr.DB.Exec("insert into movie_formats (id,format) values (?,?)", format.ID, format.Format).Error; err != nil {
		return err
	}
	return nil

}

func (cr *adminRepository) AddLanguage(language domain.Movie_Language) error {

	if err := cr.DB.Exec("insert into movie_languages (id,language) values (?,?)", language.ID, language.Language).Error; err != nil {
		return err
	}
	return nil

}

func (cr *adminRepository) Delete(genre_id string) error {

	id, _ := strconv.Atoi(genre_id)
	query := fmt.Sprintf("delete from genres where id = '%d'", id)
	if err := cr.DB.Exec(query).Error; err != nil {
		return err
	}

	return nil

}

func (cr *adminRepository) GetUserByID(id string) (domain.Users, error) {

	user_id, _ := strconv.Atoi(id)

	query := fmt.Sprintf("select * from users where id = '%d'", user_id)
	var userDetails domain.Users

	if err := cr.DB.Raw(query).Scan(&userDetails).Error; err != nil {
		return domain.Users{}, err
	}

	return userDetails, nil
}

func (cr *adminRepository) UpdateBlockUserByID(user domain.Users) error {

	err := cr.DB.Exec("update users set blocked = ? where id = ?", user.Blocked, user.ID).Error
	if err != nil {
		fmt.Println("Error updating user:", err)
		return err
	}

	return nil

}
