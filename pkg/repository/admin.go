package repository

import (
	"errors"
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

func (ad *adminRepository) LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {

	var adminCompareDetails domain.Admin
	if err := ad.DB.Raw("select * from admins where email = ? ", adminDetails.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.Admin{}, err
	}

	return adminCompareDetails, nil
}

// check if an admin with specified email already exist
func (ad *adminRepository) CheckAdminAvailability(admin models.AdminSignUp) bool {

	var count int
	if err := ad.DB.Raw("select count(*) from admins where email = ?", admin.Email).Scan(&count).Error; err != nil {
		return false
	}

	return count > 0

}

func (ad *adminRepository) SignUpHandler(admin models.AdminSignUp) (models.AdminDetailsResponse, error) {
	var adminDetails models.AdminDetailsResponse
	if err := ad.DB.Raw("insert into admins (name,email,password) values (?, ?, ?) RETURNING id, name, email", admin.Name, admin.Email, admin.Password).Scan(&adminDetails).Error; err != nil {
		return models.AdminDetailsResponse{}, err
	}

	return adminDetails, nil

}

// Get users details for authenticated admins
func (ad *adminRepository) GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error) {
	// pagination purpose -
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var userDetails []models.UserDetailsAtAdmin

	if err := ad.DB.Raw("select id,name,email,phone,blocked from users limit ? offset ?", count, offset).Scan(&userDetails).Error; err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}

func (ad *adminRepository) GetGenres() ([]domain.Genre, error) {

	var genres []domain.Genre
	if err := ad.DB.Raw("select * from genres").Scan(&genres).Error; err != nil {
		return []domain.Genre{}, err
	}

	return genres, nil

}

// func (ad *adminRepository) GetDirectors() ([]domain.Directors, error) {

// 	var directors []domain.Directors
// 	if err := ad.DB.Raw("select * from directors").Scan(&directors).Error; err != nil {
// 		return []domain.Directors{}, err
// 	}

// 	return directors, nil

// }

// func (ad *adminRepository) GetMovieFormat() ([]domain.Movie_Format, error) {
// 	var formats []domain.Movie_Format
// 	if err := ad.DB.Raw("select * from movie_formats").Scan(&formats).Error; err != nil {
// 		return []domain.Movie_Format{}, err
// 	}

// 	return formats, nil
// }

// func (ad *adminRepository) GetMovieLanguages() ([]domain.Movie_Language, error) {

// 	var languages []domain.Movie_Language
// 	if err := ad.DB.Raw("select * from movie_languages").Scan(&languages).Error; err != nil {
// 		return []domain.Movie_Language{}, err
// 	}

// 	return languages, nil
// }

func (ad *adminRepository) CategoryCount(category models.CategoryUpdate) (models.CategoryUpdateCheck, error) {
	// sub query to check if category name added by the admin already exist (add category where count = 0)
	var categoryCount models.CategoryUpdateCheck
	err := ad.DB.Raw("select (select count(*) from genres where genre_name = ?) as genre_count,(select count(*) from directors where director_name = ?) as director_count,(select count(*) from movie_formats where format = ?) as format_count,(select count(*) from movie_languages where language = ?) as language_count", category.Genre, category.Director, category.Format, category.Language).Scan(&categoryCount).Error
	if err != nil {
		return categoryCount, nil
	}
	return categoryCount, nil

}

// CATEGORY MANAGEMENT
func (ad *adminRepository) AddGenre(genre string) (domain.Genre, error) {

	var gen domain.Genre
	if err := ad.DB.Raw("insert into genres (genre_name) values (?) returning id,genre_name", genre).Scan(&gen).Error; err != nil {
		return domain.Genre{}, err
	}
	return gen, nil

}

// func (ad *adminRepository) AddDirector(director string) (domain.Directors, error) {

// 	var dir domain.Directors
// 	if err := ad.DB.Raw("insert into directors (director_name) values (?) returning id,director_name", director).Scan(&dir).Error; err != nil {
// 		return domain.Directors{}, err
// 	}
// 	return dir, nil

// }

// func (ad *adminRepository) AddFormat(format string) (domain.Movie_Format, error) {

// 	var form domain.Movie_Format
// 	if err := ad.DB.Raw("insert into movie_formats (format) values  (?) returning id,format", format).Scan(&form).Error; err != nil {
// 		return domain.Movie_Format{}, err
// 	}
// 	return form, nil

// }

// func (ad *adminRepository) AddLanguage(language string) (domain.Movie_Language, error) {

// 	var lang domain.Movie_Language
// 	if err := ad.DB.Raw("insert into movie_languages (language) values (?) returning id,language", language).Scan(&lang).Error; err != nil {
// 		return domain.Movie_Language{}, nil
// 	}
// 	return lang, nil

// }

func (ad *adminRepository) Delete(genre_id string) error {

	id, err := strconv.Atoi(genre_id)
	if err != nil {
		return err
	}
	var count int
	if err := ad.DB.Raw("select count(*) from genres where id = ?").Scan(&count).Error; err != nil {
		return err
	}
	if count < 1 {
		return errors.New("genre for given id does not exist")
	}

	query := fmt.Sprintf("delete from genres where id = '%d'", id)
	if err := ad.DB.Exec(query).Error; err != nil {
		return err
	}

	return nil

}

func (ad *adminRepository) GetUserByID(id string) (domain.Users, error) {

	user_id, err := strconv.Atoi(id)
	if err != nil {
		return domain.Users{}, err
	}

	var count int
	if err := ad.DB.Raw("select count(*) from users where id = ?").Scan(&count).Error; err != nil {
		return domain.Users{}, err
	}
	if count < 1 {
		return domain.Users{}, errors.New("user for the given id does not exist")
	}

	query := fmt.Sprintf("select * from users where id = '%d'", user_id)
	var userDetails domain.Users

	if err := ad.DB.Raw(query).Scan(&userDetails).Error; err != nil {
		return domain.Users{}, err
	}

	return userDetails, nil
}

// function which will both block and unblock a user
func (ad *adminRepository) UpdateBlockUserByID(user domain.Users) error {

	err := ad.DB.Exec("update users set blocked = ? where id = ?", user.Blocked, user.ID).Error
	if err != nil {
		fmt.Println("Error updating user:", err)
		return err
	}

	return nil

}
