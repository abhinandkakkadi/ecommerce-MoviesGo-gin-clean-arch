package interfaces

import "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"

type PaymentUseCase interface {
	MakePaymentRazorPay(orderID string, userID int) (models.CombinedOrderDetails, string, error)
	SavePaymentDetails(paymentID string, razorID string, orderID string) error
}
