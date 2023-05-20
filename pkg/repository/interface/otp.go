package interfaces

type OtpRepository interface {
	FindUserByMobileNumber(phone string) bool
	// VerifyMobileNumberAlreadyPresent(ctx context.Context,phone string) bool
}