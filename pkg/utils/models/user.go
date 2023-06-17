package models

type UserDetails struct {
	Name            string `json:"name"`
	Email           string `json:"email" validate:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
	ReferralCode    string `json:"referral_code"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
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
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	ReferralCode string `json:"referral_code"`
}

// user details along with embedded token which can be used by the user to access protected routes
type TokenUsers struct {
	Users UserDetailsResponse
	Token string
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
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
	Street    string `json:"street"`
	City      string `json:"city"`
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
	OldPassword        string `json:"old-password"`
	NewPassword        string `json:"newpassword"`
	ConfirmNewPassword string `json:"confirm-newpassword"`
}

type ResetPassword struct {
	Password  string `json:"password" validate:"required"`
	CPassword string `json:"cpassword" validate:"required"`
}
