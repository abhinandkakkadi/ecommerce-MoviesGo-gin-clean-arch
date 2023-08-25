package usecase

import (
	"errors"

	config "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/config"
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

func (ot *otpUseCase) SendOTP(phone string) error {

	ok := ot.otpRepository.FindUserByMobileNumber(phone)
	if !ok {
		return errors.New("the user does not exist")
	}

	helper.TwilioSetup(ot.cfg.ACCOUNTSID, ot.cfg.AUTHTOKEN)
	_, err := helper.TwilioSendOTP(phone, ot.cfg.SERVICESSID)
	if err != nil {
		return errors.New("error ocurred while generating OTP")
	}

	return nil

}

func (ot *otpUseCase) VerifyOTP(code models.VerifyData) (models.TokenUsers, error) {

	helper.TwilioSetup(ot.cfg.ACCOUNTSID, ot.cfg.AUTHTOKEN)
	err := helper.TwilioVerifyOTP(ot.cfg.SERVICESSID, code.Code, code.User.PhoneNumber)
	if err != nil {
		return models.TokenUsers{}, errors.New("error while verifying")
	}

	// if user is authenticated using OTP send back user details
	userDetails, err := ot.otpRepository.UserDetailsUsingPhone(code.User.PhoneNumber)
	if err != nil {
		return models.TokenUsers{}, err
	}

	accessToken, err := helper.GenerateAccessToken(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	refreshToken, err := helper.GenerateRefreshToke(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	var user models.UserDetailsResponse
	err = copier.Copy(&user, &userDetails)
	if err != nil {
		return models.TokenUsers{}, err
	}

	return models.TokenUsers{
		Users: user,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil

}

func (ot *otpUseCase) SendOTPtoReset(email string) (string, error) {

	// check whether the user exist
	ok, err := ot.otpRepository.FindUserByEmail(email)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("the user does not exist")
	}

	phone, err := ot.otpRepository.GetUserPhoneByEmail(email)
	if err != nil {
		return "", err
	}

	helper.TwilioSetup(ot.cfg.ACCOUNTSID, ot.cfg.AUTHTOKEN)
	_, err = helper.TwilioSendOTP(phone, ot.cfg.SERVICESSID)
	if err != nil {
		return "", errors.New("error ocurred while generating OTP")
	}

	return phone, nil

}

func (ot *otpUseCase) VerifyOTPtoReset(code models.VerifyData) (string, error) {

	helper.TwilioSetup(ot.cfg.ACCOUNTSID, ot.cfg.AUTHTOKEN)
	err := helper.TwilioVerifyOTP(ot.cfg.SERVICESSID, code.Code, code.User.PhoneNumber)
	if err != nil {
		return "", errors.New("error while verifying")
	}

	// if user is authenticated using OTP send back user details
	userDetails, err := ot.otpRepository.UserDetailsUsingPhone(code.User.PhoneNumber)
	if err != nil {
		return "", err
	}

	tokenString, err := helper.GenerateTokenToResetPassword(userDetails)
	if err != nil {
		return "", err
	}

	return tokenString, err

}
