package repository

import (

"gorm.io/gorm"
interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"

)

type otpRepository struct {
	DB *gorm.DB
}

func NewOtpRepository(DB *gorm.DB) interfaces.OtpRepository  {
	return &otpRepository{
		DB: DB,
	}
}



func (cr *otpRepository) FindUserByMobileNumber(phone string) bool {
	var count int
	if err := cr.DB.Raw("select count(*) from users where phone = ?",phone).Scan(&count).Error; err != nil {
		return false
	}

	return count > 0
}