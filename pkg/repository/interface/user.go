package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type UserRepository interface {
	UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error)
	FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error)
	CheckUserAvailability(email string) bool
	UserBlockStatus(email string) (bool, error)
	LoginHandler(user models.UserDetails) (models.UserDetailsResponse, error)
	AddAddress(address models.AddressInfo, userID int) error
	UpdateAddress(address models.AddressInfo, addressID int, userID int) (models.AddressInfoResponse, error)
	GetAllAddresses(userID int) ([]models.AddressInfoResponse, error)
	GetWalletDetails(userID int) (models.Wallet, error)
	GetAllPaymentOption() ([]models.PaymentDetails, error)
	UserDetails(userID int) (models.UsersProfileDetails, error)
	UpdateUserEmail(email string, userID int) error
	UpdateUserName(name string, userID int) error
	UpdateUserPhone(phone string, userID int) error
	UpdateUserPassword(password string, userID int) error
	UserPassword(userID int) (string, error)
	FindUserByOrderID(orderID string) (models.UsersProfileDetails, error)
	FindUserAddressByOrderID(orderID string) (models.AddressInfo, error)
	AddToWishList(userID int, productID int) error
	GetWishList(userID int) ([]models.WishListResponse, error)
	ProductExistInWishList(productID int, userId int) (bool, error)
	RemoveFromWishList(userID int, productID int) error
	CreateReferralEntry(users models.UserDetailsResponse, userReferral string) error
	GetReferralAndTotalAmount(userID int) (float64, float64, error)
	ResetPassword(userID int, password string) error
	GetUserIdFromReferrals(ReferralCode string) (int, error)
	UpdateReferralAmount(referralAmount float64, referredUserId int, currentUserID int) error
	UpdateSomethingBasedOnUserID(tableName string, columnName string, updateValue float64, userID int) error
}
