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

// @Summary  OTP login
// @Description Send OTP to Authenticate user
// @Tags User OTP Login
// @Accept json
// @Produce json
// @Param phone body models.OTPData true "phone number details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /send-otp [post]
func (ot *OtpHandler) SendOTP(c *gin.Context) {

	var phone models.OTPData

	if err := c.ShouldBindJSON(&phone); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
	}

	err := ot.otpUseCase.SendOTP(phone.PhoneNumber)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not send OTP", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "OTP sent successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Verify OTP
// @Description Verify OTP by passing the OTP in order to authenticate user
// @Tags User OTP Login
// @Accept json
// @Produce json
// @Param phone body models.VerifyData true "Verify OTP Details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /verify-otp [post]
func (ot *OtpHandler) VerifyOTP(c *gin.Context) {

	var code models.VerifyData

	if err := c.ShouldBindJSON(&code); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	users, err := ot.otpUseCase.VerifyOTP(code)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not verify OTP", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully verified OTP", users, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Send OTP to Reset Password
// @Description Send OTP to number corresponding to the given username
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param email body models.Email true "send OTP Details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /forgot-password [post]
func (ot *OtpHandler) SendOTPtoReset(c *gin.Context) {

	var email models.Email

	if err := c.ShouldBindJSON(&email); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	phone, err := ot.otpUseCase.SendOTPtoReset(email.Email)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not send OTP", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "send OTP to your number ending in ********"+phone[len(phone)-2:], nil, nil)
	c.JSON(http.StatusInternalServerError, successRes)

}

// @Summary Verify OTP To Reset Password
// @Description Verify OTP to get a JWT token which can be used to change password
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param phone body models.VerifyData true "Verify OTP Details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /forgot-password/verify-otp [post]
func (ot *OtpHandler) VerifyOTPToReset(c *gin.Context) {

	var code models.VerifyData

	if err := c.ShouldBindJSON(&code); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	signedToken, err := ot.otpUseCase.VerifyOTPtoReset(code)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not verify OTP", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully verified OTP", signedToken, nil)
	c.JSON(http.StatusOK, successRes)

}
