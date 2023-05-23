package usecase

import (
	"errors"

	config "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/config"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	helper "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/jinzhu/copier"
)

type otpUseCase struct {
	cfg           config.Config
	otpRepository interfaces.OtpRepository
}

func NewOtpUseCase(cfg config.Config, repo interfaces.OtpRepository) services.OtpUseCase {
	return &otpUseCase{
		cfg:           cfg,
		otpRepository: repo,
	}
}

func (cr *otpUseCase) SendOTP(phone string) error {

	ok := cr.otpRepository.FindUserByMobileNumber(phone)
	if !ok {
		return errors.New("the user does not exist")
	}

	helper.TwilioSetup(cr.cfg.ACCOUNTSID, cr.cfg.AUTHTOKEN)
	_, err := helper.TwilioSendOTP(phone, cr.cfg.SERVICESSID)
	if err != nil {
		return errors.New("error ocurred while generating OTP")
	}

	return nil

}

func (cr *otpUseCase) VerifyOTP(code models.VerifyData) (domain.TokenUsers, error) {

	helper.TwilioSetup(cr.cfg.ACCOUNTSID, cr.cfg.AUTHTOKEN)
	err := helper.TwilioVerifyOTP(cr.cfg.SERVICESSID, code.Code, code.User.PhoneNumber)
	if err != nil {
		return domain.TokenUsers{}, errors.New("error while verifying")
	}

	// if user is authenticated using OTP send back user details
	userDetails, err := cr.otpRepository.UserDetailsUsingPhone(code.User.PhoneNumber)
	if err != nil {
		return domain.TokenUsers{}, err
	}

	tokenString, err := helper.GenerateTokenUsers(userDetails)

	var user models.UserDetails
	err = copier.Copy(&user, &userDetails)
	if err != nil {
		return domain.TokenUsers{}, err
	}

	return domain.TokenUsers{
		Users: user,
		Token: tokenString,
	}, nil

}
