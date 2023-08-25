package models

type UserDetails struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required" validate:"email"`
	Phone           string `json:"phone" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmpassword" binding:"required"`
	ReferralCode    string `json:"referral_code"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required"`
}

// user details shown after logging in
type UserDetailsResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserDetailsAtAdmin struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	BlockStatus bool   `json:"block_status"`
}

// show in users profile / also used to update user details
type UsersProfileDetails struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	ReferralCode string `json:"referral_code" binding:"required"`
}

// user details along with embedded token which can be used by the user to access protected routes
type TokenUsers struct {
	Users        UserDetailsResponse
	AccessToken  string
	RefreshToken string
}

type UserSignInResponse struct {
	Id       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type AddressInfo struct {
	Name      string `json:"name" binding:"required" validate:"required"`
	HouseName string `json:"house_name" binding:"required" validate:"required"`
	State     string `json:"state" binding:"required" validate:"required"`
	Pin       string `json:"pin" binding:"required" validate:"required"`
	Street    string `json:"street" binding:"required"`
	City      string `json:"city" binding:"required"`
}

type AddressInfoResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
	Street    string `json:"street"`
	City      string `json:"city"`
}

type PaymentDetails struct {
	ID           uint   `json:"id"`
	Payment_Name string `json:"payment_name"`
}

type CheckoutDetails struct {
	AddressInfoResponse []AddressInfoResponse
	Payment_Method      []PaymentDetails
	Wallet              Wallet
	ReferralAmount      ReferralAmount
	Cart                []Cart
	Grand_Total         float64
	Total_Price         float64
	DiscountReason      []string
}

type UpdatePassword struct {
	OldPassword        string `json:"old-password" binding:"required"`
	NewPassword        string `json:"newpassword" binding:"required"`
	ConfirmNewPassword string `json:"confirm-newpassword" binding:"required "`
}

type ResetPassword struct {
	Password  string `json:"password" binding:"required" validate:"required"`
	CPassword string `json:"cpassword" binding:"required" validate:"required"`
}
