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

func (cr *adminRepository) LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {

	var adminCompareDetails domain.Admin
	if err := cr.DB.Raw("select * from admins where email = ? ", adminDetails.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.Admin{}, err
	}

	return adminCompareDetails, nil
}

// check if an admin with specified email already exist
func (cr *adminRepository) CheckAdminAvailability(admin models.AdminSignUp) bool {

	var count int
	if err := cr.DB.Raw("select count(*) from admins where email = ?", admin.Email).Scan(&count).Error; err != nil {
		return false
	}

	return count > 0

}

func (cr *adminRepository) SignUpHandler(admin models.AdminSignUp) (models.AdminDetailsResponse, error) {
	var adminDetails models.AdminDetailsResponse
	if err := cr.DB.Raw("insert into admins (name,email,password) values (?, ?, ?) RETURNING id, name, email", admin.Name, admin.Email, admin.Password).Scan(&adminDetails).Error; err != nil {
		return models.AdminDetailsResponse{}, err
	}

	return adminDetails, nil

}

// Get users details for authenticated admins
func (cr *adminRepository) GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error) {
	// pagination purpose -
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var userDetails []models.UserDetailsAtAdmin

	if err := cr.DB.Raw("select id,name,email,phone,blocked from users limit ? offset ?", count, offset).Scan(&userDetails).Error; err != nil {
		return []models.UserDetailsAtAdmin{}, err
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

func (cr *adminRepository) GetDirectors() ([]domain.Directors, error) {

	var directors []domain.Directors
	if err := cr.DB.Raw("select * from directors").Scan(&directors).Error; err != nil {
		return []domain.Directors{}, err
	}

	return directors, nil

}

func (cr *adminRepository) GetMovieFormat() ([]domain.Movie_Format, error) {
	var formats []domain.Movie_Format
	if err := cr.DB.Raw("select * from movie_formats").Scan(&formats).Error; err != nil {
		return []domain.Movie_Format{}, err
	}

	return formats, nil
}

func (cr *adminRepository) GetMovieLanguages() ([]domain.Movie_Language, error) {

	var languages []domain.Movie_Language
	if err := cr.DB.Raw("select * from movie_languages").Scan(&languages).Error; err != nil {
		return []domain.Movie_Language{}, err
	}

	return languages, nil
}

func (cr *adminRepository) CategoryCount(category models.CategoryUpdate) (models.CategoryUpdateCheck, error) {
	// sub query to check if category name added by the admin already exist (add category where count = 0)
	var categoryCount models.CategoryUpdateCheck
	err := cr.DB.Raw("select (select count(*) from genres where genre_name = ?) as genre_count,(select count(*) from directors where director_name = ?) as director_count,(select count(*) from movie_formats where format = ?) as format_count,(select count(*) from movie_languages where language = ?) as language_count", category.Genre, category.Director, category.Format, category.Language).Scan(&categoryCount).Error
	if err != nil {
		return categoryCount, nil
	}
	return categoryCount, nil

}

// CATEGORY MANAGEMENT
func (cr *adminRepository) AddGenre(genre string) (domain.Genre, error) {

	var gen domain.Genre
	if err := cr.DB.Raw("insert into genres (genre_name) values (?) returning id,genre_name", genre).Scan(&gen).Error; err != nil {
		return domain.Genre{}, err
	}
	return gen, nil

}

func (cr *adminRepository) AddDirector(director string) (domain.Directors, error) {

	var dir domain.Directors
	if err := cr.DB.Raw("insert into directors (director_name) values (?) returning id,director_name", director).Scan(&dir).Error; err != nil {
		return domain.Directors{}, err
	}
	return dir, nil

}

func (cr *adminRepository) AddFormat(format string) (domain.Movie_Format, error) {

	var form domain.Movie_Format
	if err := cr.DB.Raw("insert into movie_formats (format) values  (?) returning id,format", format).Scan(&form).Error; err != nil {
		return domain.Movie_Format{}, err
	}
	return form, nil

}

func (cr *adminRepository) AddLanguage(language string) (domain.Movie_Language, error) {

	var lang domain.Movie_Language
	if err := cr.DB.Raw("insert into movie_languages (language) values (?) returning id,language", language).Scan(&lang).Error; err != nil {
		return domain.Movie_Language{}, nil
	}
	return lang, nil

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

// function which will both block and unblock a user
func (cr *adminRepository) UpdateBlockUserByID(user domain.Users) error {

	err := cr.DB.Exec("update users set blocked = ? where id = ?", user.Blocked, user.ID).Error
	if err != nil {
		fmt.Println("Error updating user:", err)
		return err
	}

	return nil

}
