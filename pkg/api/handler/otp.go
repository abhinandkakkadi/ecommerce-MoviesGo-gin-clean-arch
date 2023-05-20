package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
	models "github.com/thnkrn/go-gin-clean-arch/pkg/utils/models"
)

type OtpHandler struct {
	otpUseCase  services.OtpUseCase
}

func NewOtpHandler(useCase services.OtpUseCase) *OtpHandler {
	return &OtpHandler{
		otpUseCase: useCase,
	}
}

// send OTP
func (cr *OtpHandler) SendOTP(c *gin.Context) {

	var phone models.OTPData
	if err := c.BindJSON(&phone); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Not a valid phone number",
		})
	}

	// err := cr.otpUseCase.VerifyMobileNumberAlreadyPresent(c.Request.Context(),phone.PhoneNumber)

	err := cr.otpUseCase.SendOTP(c.Request.Context(),phone.PhoneNumber)

	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,"OTP Sent Succesfully")
}

// verify OTP
func (cr *OtpHandler) VerifyOTP(c *gin.Context) {

	var code models.VerifyData

	if err := c.BindJSON(&code); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"bad request",
		})
		return
	}

	err := cr.otpUseCase.VerifyOTP(c.Request.Context(),code)

	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,"OTP verified succesfuly")
}