package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type UserRepository interface {
	UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error)
	FindUserByEmail(user models.UserDetails) (models.UserSignInResponse, error)
	CheckUserAvailability(user models.UserDetails) bool
	LoginHandler(user models.UserDetails) (models.UserDetailsResponse, error)
	AddAddress(address models.AddressInfo, userID int) ([]models.AddressInfoResponse, error)
	UpdateAddress(address models.AddressInfo, addressID int) (models.AddressInfoResponse, error)
	GetAllAddresses(userID int) ([]models.AddressInfoResponse, error)
	GetAllPaymentOption() ([]models.PaymentDetails, error)
	UserDetails(userID int) (models.UsersProfileDetails,error) 
}
