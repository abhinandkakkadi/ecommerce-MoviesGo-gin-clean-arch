package interfaces

import (
	"context"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type UserUseCase interface {
	UserSignUp(user models.UserDetails) (models.TokenUsers, error)
	LoginHandler(user models.UserLogin) (models.TokenUsers, error)
	AddAddress(address models.AddressInfo, userID int) error
	UpdateAddress(address models.AddressInfo, addressID int, userID int) (models.AddressInfoResponse, error)
	Checkout(userID int) (models.CheckoutDetails, error)
	UserDetails(userID int) (models.UsersProfileDetails, error)
	GetAllAddress(userID int) ([]models.AddressInfoResponse, error)
	UpdateUserDetails(body models.UsersProfileDetails, ctx context.Context) (models.UsersProfileDetails, error)
	UpdatePassword(ctx context.Context, body models.UpdatePassword) error
	AddToWishList(productID int, userID int) error
	GetWishList(userID int) ([]models.WishListResponse, error)
	RemoveFromWishList(productID int, userID int) error
	ApplyReferral(userID int) (string, error)
	ResetPassword(userID int, pass models.ResetPassword) error
}
