package handler

import (
	"net/http"

	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	models "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type OtpHandler struct {
	otpUseCase services.OtpUseCase
}

func NewOtpHandler(useCase services.OtpUseCase) *OtpHandler {
	return &OtpHandler{
		otpUseCase: useCase,
	}
}

// @Summary Send OTP
// @Description Send OTP to Authorize user
// @Tags User
// @Accept json
// @Produce json
// @Param phone body models.OTPData true "phone number details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /send-otp [post]
func (ot *OtpHandler) SendOTP(c *gin.Context) {

	var phone models.OTPData
	if err := c.BindJSON(&phone); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest,"fields provided are in wrong format",nil,err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
	}

	err := ot.otpUseCase.SendOTP(phone.PhoneNumber)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError,"Could not send OTP",nil,err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusNoContent,"OTP sent successfully",nil,nil)
	c.JSON(http.StatusNoContent, successRes)

}

// @Summary Verify OTP
// @Description Verify OTP by passing the OTP
// @Tags User
// @Accept json
// @Produce json
// @Param phone body models.VerifyData true "Verify OTP Details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /verify-otp [post]
func (ot *OtpHandler) VerifyOTP(c *gin.Context) {

	var code models.VerifyData
	if err := c.BindJSON(&code); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest,"fields provided are in wrong format",nil,err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	users, err := ot.otpUseCase.VerifyOTP(code)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError,"Could not verify OTP",nil,err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}	

	successRes := response.ClientResponse(http.StatusOK,"Successfully verified OTP",users,nil)
	c.JSON(http.StatusOK, successRes)

}
