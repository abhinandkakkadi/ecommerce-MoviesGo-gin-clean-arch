package helper

import (
	"strconv"
	"time"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

func CombinedOrderDetails(orderDetails models.OrderDetails, userDetails models.UsersProfileDetails, userAddress models.AddressInfo) (models.CombinedOrderDetails, error) {

	var orderCombinedDetails models.CombinedOrderDetails

	err := copier.Copy(&orderCombinedDetails, &orderDetails)
	if err != nil {
		return models.CombinedOrderDetails{}, err
	}

	err = copier.Copy(&orderCombinedDetails, &userDetails)
	if err != nil {
		return models.CombinedOrderDetails{}, err
	}

	err = copier.Copy(&orderCombinedDetails, &userAddress)
	if err != nil {
		return models.CombinedOrderDetails{}, err
	}

	return orderCombinedDetails, nil
}

func GetTimeFromPeriod(timePeriod string) (time.Time, time.Time) {

	endDate := time.Now()

	if timePeriod == "week" {
		startDate := endDate.AddDate(0, 0, -6)
		return startDate, endDate
	}

	if timePeriod == "month" {
		startDate := endDate.AddDate(0, -1, 0)
		return startDate, endDate
	}

	if timePeriod == "year" {
		startDate := endDate.AddDate(0, -1, 0)
		return startDate, endDate
	}

	return endDate.AddDate(0, 0, -6), endDate

}

func CopyOrderDetails(orderDetails domain.Order, orderBody models.OrderIncoming) domain.Order {

	id := uuid.New().ID()
	str := strconv.Itoa(int(id))
	orderDetails.OrderId = str[:8]
	orderDetails.AddressID = orderBody.AddressID
	orderDetails.PaymentMethodID = orderBody.PaymentID
	orderDetails.UserID = int(orderBody.UserID)
	orderDetails.Approval = false
	orderDetails.ShipmentStatus = "processing"
	orderDetails.PaymentStatus = "not paid"

	return orderDetails

}
