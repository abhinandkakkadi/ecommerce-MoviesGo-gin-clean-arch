package usecase

import (
	"errors"
	"fmt"

	domain "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	helper "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
}

func NewAdminUseCase(repo interfaces.AdminRepository) services.AdminUseCase {
	return &adminUseCase{
		adminRepository: repo,
	}
}

func (cr *adminUseCase) LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error) {

	// getting details of the admin based on the email provided
	adminCompareDetails, err := cr.adminRepository.LoginHandler(adminDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	// compare password from database and that provided from admins
	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	fmt.Println(err)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	var adminDetailsResponse models.AdminDetailsResponse

	//  copy all details except password and sent it back to the front end
	err = copier.Copy(&adminDetailsResponse, &adminCompareDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	tokenString, err := helper.GenerateTokenAdmin(adminDetailsResponse)

	if err != nil {
		return domain.TokenAdmin{}, err
	}

	// var admin models.AdminDetails

	// err = copier.Copy(&admin, &adminCompareDetails)
	// if err != nil {
	// 	return domain.TokenAdmin{}, err
	// }

	return domain.TokenAdmin{
		Admin: adminDetailsResponse,
		Token: tokenString,
	}, nil

}

// signup handler for the admin
func (cr *adminUseCase) SignUpHandler(admin models.AdminSignUp) (domain.TokenAdmin, error) {

	// validator package to check the constraints specified in the struct which is used to retrieve these details
	if err := validator.New().Struct(admin); err != nil {
		return domain.TokenAdmin{}, err
	}

	// check whether the admin already exist in the database -
	userExist := cr.adminRepository.CheckAdminAvailability(admin)
	if userExist {
		return domain.TokenAdmin{}, errors.New("admin already exist, sign in")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return domain.TokenAdmin{}, errors.New("Internal server error")
	}
	admin.Password = string(hashedPassword)

	adminDetails, err := cr.adminRepository.SignUpHandler(admin)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	tokenString, err := helper.GenerateTokenAdmin(adminDetails)

	if err != nil {
		return domain.TokenAdmin{}, err
	}

	// var adminDetails models.AdminDetails

	// err = copier.Copy(&adminDetails, &adminToken)
	// if err != nil {
	// 	return domain.TokenAdmin{}, err
	// }

	return domain.TokenAdmin{
		Admin: adminDetails,
		Token: tokenString,
	}, nil

}

func (cr *adminUseCase) GetUsers(page int) ([]models.UserDetailsAtAdmin, error) {

	userDetails, err := cr.adminRepository.GetUsers(page)
	if err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}

// business logic to get all the category
func (cr *adminUseCase) GetFullCategory() (domain.CategoryResponse, error) {

	genres, err := cr.adminRepository.GetGenres()
	if err != nil {
		return domain.CategoryResponse{}, err
	}

	directors, err := cr.adminRepository.GetDirectors()
	if err != nil {
		return domain.CategoryResponse{}, err
	}

	formats, err := cr.adminRepository.GetMovieFormat()
	if err != nil {
		return domain.CategoryResponse{}, err
	}

	languages, err := cr.adminRepository.GetMovieLanguages()
	if err != nil {
		return domain.CategoryResponse{}, err
	}

	return domain.CategoryResponse{
		Genre:          genres,
		Directors:      directors,
		Movie_Format:   formats,
		Movie_Language: languages,
	}, nil

}

// add new category
func (cr *adminUseCase) AddCategory(category models.CategoryUpdate) (domain.CategoryManagement, error) {

	var (
		genre    domain.Genre
		director domain.Directors
		format   domain.Movie_Format
		language domain.Movie_Language
		err      error
	)
	var count int
	// to check if a category with same name exists in the database
	categoryCount, err := cr.adminRepository.CategoryCount(category)
	if err != nil {
		return domain.CategoryManagement{}, nil
	}

	if category.Genre != "" && categoryCount.GenreCount == 0 {
		count++
		genre, err = cr.adminRepository.AddGenre(category.Genre)
		if err != nil {
			return domain.CategoryManagement{}, err
		}
	}

	if category.Director != "" && categoryCount.DirectorCount == 0 {
		count++
		director, err = cr.adminRepository.AddDirector(category.Director)
		if err != nil {
			return domain.CategoryManagement{}, err
		}
	}

	if category.Format != "" && categoryCount.FormatCount == 0 {
		count++
		format, err = cr.adminRepository.AddFormat(category.Format)
		if err != nil {
			return domain.CategoryManagement{}, err
		}
	}

	if category.Language != "" && categoryCount.LanguageCount == 0 {
		count++
		language, err = cr.adminRepository.AddLanguage(category.Language)
		if err != nil {
			return domain.CategoryManagement{}, err
		}
	}

	// if count = 0 that means no category was added
	if count == 0 {
		return domain.CategoryManagement{}, errors.New("no new category added")
	}

	return domain.CategoryManagement{
		Genre:    genre,
		Director: director,
		Format:   format,
		Language: language,
	}, nil

}

func (cr *adminUseCase) Delete(genre_id string) error {

	err := cr.adminRepository.Delete(genre_id)
	if err != nil {
		return err
	}
	return nil

}

// block user
func (cr *adminUseCase) BlockUser(id string) error {

	user, err := cr.adminRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Blocked {
		return errors.New("already blocked")
	} else {
		user.Blocked = true
	}

	err = cr.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}

	return nil

}

// unblock user
func (cr *adminUseCase) UnBlockUser(id string) error {

	user, err := cr.adminRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Blocked {
		user.Blocked = false
	} else {
		return errors.New("already unblocked")
	}

	err = cr.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}

	return nil

}
