package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type UserUseCase interface {
	UserSignUp(user models.UserDetails) (models.TokenUsers, error)

	LoginHandler(user models.UserDetails) (models.TokenUsers, error)
	AddAddress(address models.AddressInfo,userID int) ([]models.AddressInfoResponse,error)
	UpdateAddress(address models.AddressInfo,addressID int) (models.AddressInfoResponse,error)
}
