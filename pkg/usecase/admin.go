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

func (cr *adminUseCase) LoginHandler(adminDetails domain.Admin) (domain.TokenAdmin, error) {

	adminCompareDetails, err := cr.adminRepository.LoginHandler(adminDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	fmt.Println(err)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	var adminDetailsResponse models.AdminDetailsResponse

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

func (cr *adminUseCase) SignUpHandler(admin models.AdminSignUp) (domain.TokenAdmin, error) {

	if err := validator.New().Struct(admin); err != nil {
		return domain.TokenAdmin{}, err
	}

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

func (cr *adminUseCase) GetUsers(page int) ([]models.UserDetailsResponse, error) {

	userDetails, err := cr.adminRepository.GetUsers(page)
	if err != nil {
		return []models.UserDetailsResponse{}, err
	}

	return userDetails, nil

}

func (cr *adminUseCase) GetGenres() ([]domain.Genre, error) {

	genres, err := cr.adminRepository.GetGenres()
	if err != nil {
		return []domain.Genre{}, err
	}

	return genres, nil

}

func (cr *adminUseCase) AddCategory(category domain.CategoryManagement) (domain.CategoryManagement, error) {

	if category.Genre.ID != 0 {
		err := cr.adminRepository.AddGenre(category.Genre)
		if err != nil {
			return domain.CategoryManagement{}, err
		}
	}

	if category.Director.ID != 0 {
		err := cr.adminRepository.AddDirector(category.Director)
		if err != nil {
			return domain.CategoryManagement{}, err
		}
	}

	if category.Format.ID != 0 {
		err := cr.adminRepository.AddFormat(category.Format)
		if err != nil {
			return domain.CategoryManagement{}, err
		}
	}

	if category.Language.ID != 0 {
		err := cr.adminRepository.AddLanguage(category.Language)
		if err != nil {
			return domain.CategoryManagement{}, err
		}
	}

	return category, nil

}

func (cr *adminUseCase) Delete(genre_id string) error {

	err := cr.adminRepository.Delete(genre_id)
	if err != nil {
		return err
	}
	return nil

}

func (cr *adminUseCase) BlockUser(id string) error {

	user, err := cr.adminRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Blocked {
		user.Blocked = false
	} else {
		user.Blocked = true
	}

	err = cr.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}

	return nil

}
