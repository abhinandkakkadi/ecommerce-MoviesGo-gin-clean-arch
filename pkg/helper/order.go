package helper

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/jinzhu/copier"
)


func CombinedOrderDetails(orderDetails models.OrderDetails,userDetails models.UsersProfileDetails,userAddress models.AddressInfo) (models.CombinedOrderDetails,error) {

		var orderCombinedDetails models.CombinedOrderDetails

	
		err := copier.Copy(&orderCombinedDetails, &orderDetails)
		if err != nil {
			return models.CombinedOrderDetails{},err
		}

		err = copier.Copy(&orderCombinedDetails, &userDetails)
		if err != nil {
			return models.CombinedOrderDetails{},err
		}

		err = copier.Copy(&orderCombinedDetails, &userAddress)
		if err != nil {
			return models.CombinedOrderDetails{},err
		}
		
		return orderCombinedDetails,nil
}