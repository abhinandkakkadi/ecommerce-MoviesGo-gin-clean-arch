package helper

import (
	"fmt"
	"time"

	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

func OfferHelper(combinedOfferDetails models.CombinedOffer) models.OfferResponse {

	// if product offer exist check whether it is still active or if it have been expired and also if the user limit for that offer is not exceeded
	currentTime := time.Now()
	if combinedOfferDetails.ProductOffer.OfferName != "" {
		if currentTime.After(combinedOfferDetails.ProductOffer.StartDate) && currentTime.Before(combinedOfferDetails.ProductOffer.EndDate) && combinedOfferDetails.ProductOffer.OfferUsed < combinedOfferDetails.ProductOffer.OfferLimit {
		} else {
			combinedOfferDetails.ProductOffer.OfferName = ""
			combinedOfferDetails.ProductOffer.DiscountPercentage = 0
		}
	}

	// if category offer exist check whether it is still active or if it have been expired and also the user limit for that offer is not exceeded
	if combinedOfferDetails.CategoryOffer.OfferName != "" {
		if currentTime.After(combinedOfferDetails.CategoryOffer.StartDate) && currentTime.Before(combinedOfferDetails.CategoryOffer.EndDate) && combinedOfferDetails.CategoryOffer.OfferUsed < combinedOfferDetails.CategoryOffer.OfferLimit {
		} else {
			combinedOfferDetails.CategoryOffer.OfferName = ""
			combinedOfferDetails.CategoryOffer.DiscountPercentage = 0
		}
	}

	// whichever offer provides greater percentage
	if combinedOfferDetails.ProductOffer.DiscountPercentage > combinedOfferDetails.CategoryOffer.DiscountPercentage {
		combinedOfferDetails.FinalOffer.OfferID = combinedOfferDetails.ProductOffer.ID
		combinedOfferDetails.FinalOffer.OfferName = combinedOfferDetails.ProductOffer.OfferName
		combinedOfferDetails.FinalOffer.OfferPercentage = combinedOfferDetails.ProductOffer.DiscountPercentage
		combinedOfferDetails.FinalOffer.OfferType = "product"
		combinedOfferDetails.FinalOffer.OfferLimit = combinedOfferDetails.ProductOffer.OfferLimit
	} else if combinedOfferDetails.CategoryOffer.DiscountPercentage > combinedOfferDetails.ProductOffer.DiscountPercentage {
		combinedOfferDetails.FinalOffer.OfferID = combinedOfferDetails.CategoryOffer.ID
		combinedOfferDetails.FinalOffer.OfferName = combinedOfferDetails.CategoryOffer.OfferName
		combinedOfferDetails.FinalOffer.OfferPercentage = combinedOfferDetails.CategoryOffer.DiscountPercentage
		combinedOfferDetails.FinalOffer.OfferType = "category"
		combinedOfferDetails.FinalOffer.OfferLimit = combinedOfferDetails.CategoryOffer.OfferLimit
	} else {
		combinedOfferDetails.FinalOffer.OfferName = "sorry no offer at this time"
		combinedOfferDetails.FinalOffer.OfferPercentage = 0
		combinedOfferDetails.FinalOffer.OfferType = "no offer"
		return combinedOfferDetails.FinalOffer
	}
	// select price from Price table and add it to the mix and
	combinedOfferDetails.FinalOffer.OfferPrice = combinedOfferDetails.OriginalPrice - ((float64(combinedOfferDetails.FinalOffer.OfferPercentage) * combinedOfferDetails.OriginalPrice) / 100)

	return combinedOfferDetails.FinalOffer

}

var OfferContainers = map[string][]models.ProductOfferBriefResponse{}

// implement some concurrent stuff and store it in redis
func LatestOfferAlert(productRepo interfaces.ProductRepository) {

	fmt.Println("update this in the future")

}
