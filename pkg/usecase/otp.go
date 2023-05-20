package usecase

import (
	"errors"

	config "github.com/thnkrn/go-gin-clean-arch/pkg/config"
	helper "github.com/thnkrn/go-gin-clean-arch/pkg/helper"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils/models"
	"golang.org/x/net/context"
)

type otpUseCase struct {
	cfg config.Config
	repo interfaces.OtpRepository
}

func NewOtpUseCase(cfg config.Config,repo interfaces.OtpRepository) services.OtpUseCase {
	return &otpUseCase{
		cfg: cfg,
		repo: repo,
	}
}

// func (cr *otpUseCase) VerifyMobileNumberAlreadyPresent(ctx context.Context,phone string) error {

// 	// ok := VerifyMobileNumberAlreadyPresent(ctx,phone)

// 	// if !ok {
// 	// 	return errors.New("Mobile number not of a valid user")
// 	// }

// 	return nil
// }

func (cr *otpUseCase) SendOTP(ctx context.Context,phone string) error {

	ok := cr.repo.FindUserByMobileNumber(phone)
	if !ok {
		return errors.New("the user does not exist")
	}

	helper.TwilioSetup(cr.cfg.ACCOUNTSID,cr.cfg.AUTHTOKEN)
	_,err := helper.TwilioSendOTP(phone,cr.cfg.SERVICESSID)

	if err != nil {
		return errors.New("error occured while generating OTP")
	}

	return nil
}

func (cr *otpUseCase) VerifyOTP(ctx context.Context, code models.VerifyData) error {

	err := helper.TwilioVerifyOTP(cr.cfg.SERVICESSID,code.Code,code.User.PhoneNumber)

	if err != nil {
		return errors.New("error while verifying")
	}

	return nil
}