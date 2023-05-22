package usecase

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	domain "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	helper "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
}


func NewAdminUseCase(repo interfaces.AdminRepository) services.AdminUseCase  {
	return &adminUseCase{
		adminRepository: repo,
	}
}



func (cr *adminUseCase) LoginHandler(c context.Context,adminDetails domain.Admin) (domain.TokenAdmin,error) {

	adminCompareDetails,err := cr.adminRepository.LoginHandler(c,adminDetails)
	if err != nil {
		return domain.TokenAdmin{},err
	}

	
	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password),[]byte(adminDetails.Password))
	fmt.Println(err)
	if err != nil {
		return domain.TokenAdmin{},err
	}

	tokenString,err := helper.GenerateTokenAdmin(adminCompareDetails)

	if err != nil {
		return domain.TokenAdmin{},err
	}

	var admin models.AdminDetails

	err = copier.Copy(&admin,&adminCompareDetails)
	if err != nil {
		return domain.TokenAdmin{},err
	}

	return domain.TokenAdmin{
		Admin: admin,
		Token: tokenString,
	},nil
	

}

func (cr *adminUseCase) SignupHandler(c context.Context,admin domain.Admin) (domain.TokenAdmin,error) {

	if err := validator.New().Struct(admin); err != nil {
		return domain.TokenAdmin{},err
	}

	userExist := cr.adminRepository.CheckAdminAvailability(admin)
	
	if userExist {
		return domain.TokenAdmin{},errors.New("admin already exist, sign in")
	}

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(admin.Password),10)
	if err != nil {
		return domain.TokenAdmin{},errors.New("Internal server error")
	}
	admin.Password = string(hashedPassword)

	admin,err = cr.adminRepository.SignupHandler(c,admin)
	if err != nil {
		return domain.TokenAdmin{},err
	}

	tokenString,err := helper.GenerateTokenAdmin(admin)

	if err != nil {
		return domain.TokenAdmin{},err
	}

	var adminDetails models.AdminDetails

	err = copier.Copy(&adminDetails,&admin)
	if err != nil {
		return domain.TokenAdmin{},err
	}

	return domain.TokenAdmin{
		Admin: adminDetails,
		Token: tokenString,
	},nil


}


func (cr *adminUseCase) GetUsers(c context.Context) ([]models.UserDetails,error) {

	userDetails,err := cr.adminRepository.GetUsers(c)
	if err != nil {
		return []models.UserDetails{},err
	}
	return userDetails,nil
	
}

func (cr *adminUseCase) GetGenres(c context.Context) ([]domain.Genre,error) {
	fmt.Println("the code reached here")
	genres,err := cr.adminRepository.GetGenres(c)
	if err != nil {
		return []domain.Genre{},err
	}
	fmt.Println("the code reached here")
	return genres,nil
	
}

func (cr *adminUseCase) AddGenre(c context.Context,genre domain.Genre) (domain.Genre,error) {

	genre_added,err := cr.adminRepository.AddGenre(c,genre)
	if err != nil {
		return domain.Genre{},err
	}

	return genre_added,nil
}

func (cr *adminUseCase) Delete(c context.Context,genre_id string) error {

	err := cr.adminRepository.Delete(c,genre_id)
	if err != nil {
		return err
	}
	return nil
}


func (cr *adminUseCase) BlockUser(c context.Context,id string) error {

	user,err := cr.adminRepository.GetUserByID(c,id)
	if err != nil {
		return err
	}
	fmt.Println(user)

	if user.Blocked {
		user.Blocked = false
	} else {
		user.Blocked = true
	}

	err = cr.adminRepository.UpdateUserByID(c,user)

	if err != nil {
		return nil
	}

	return nil

}