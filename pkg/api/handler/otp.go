package handler

import (
	"net/http"

	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	models "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/gin-gonic/gin"
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

	err := cr.otpUseCase.SendOTP(phone.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,"OTP Sent successfully")

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

	user,err := cr.otpUseCase.VerifyOTP(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,user)
	
}