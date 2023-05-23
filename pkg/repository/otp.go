package repository

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type otpRepository struct {
	DB *gorm.DB
}

func NewOtpRepository(DB *gorm.DB) interfaces.OtpRepository {
	return &otpRepository{
		DB: DB,
	}
}

func (cr *otpRepository) FindUserByMobileNumber(phone string) bool {

	var count int
	if err := cr.DB.Raw("select count(*) from users where phone = ?", phone).Scan(&count).Error; err != nil {
		return false
	}

	return count > 0

}

func (cr *otpRepository) UserDetailsUsingPhone(phone string) (domain.Users, error) {

	var usersDetails domain.Users
	if err := cr.DB.Raw("select * from users where phone = ?", phone).Scan(&usersDetails).Error; err != nil {
		return domain.Users{}, err
	}

	return usersDetails, nil

}
