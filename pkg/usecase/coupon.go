package usecase

import (
	"errors"

	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type couponUseCase struct {
	couponRepository interfaces.CouponRepository
}

func NewCouponUseCase(couponRepo interfaces.CouponRepository) services.CouponUseCase {
	return &couponUseCase{
		couponRepository: couponRepo,
	}
}

func (co *couponUseCase) AddCoupon(coupon models.AddCoupon) (string, error) {

	// if coupon already exist and if it is expired revalidate it. else give back an error message saying the coupon already exist

	couponExist, err := co.couponRepository.CouponExist(coupon.Coupon)
	if err != nil {
		return "", err
	}

	if couponExist {

		alreadyValid, err := co.couponRepository.CouponRevalidateIfExpired(coupon.Coupon)

		if err != nil {
			return "", err
		}

		if alreadyValid {
			return "The coupon which is valid already exists", nil
		}

		return "Made the coupon valid", nil

	}

	err = co.couponRepository.AddCoupon(coupon)
	if err != nil {
		return "", err
	}

	return "successfully added the coupon", nil
}

func (co *couponUseCase) GetCoupon() ([]models.Coupon, error) {

	return co.couponRepository.GetCoupon()

}

func (co *couponUseCase) ExpireCoupon(couponID int) error {

	// check whether coupon exist
	couponExist, err := co.couponRepository.ExistCoupon(couponID)
	if err != nil {
		return err
	}

	// if it exists expire it, if already expired send back relevant message
	if couponExist {
		err = co.couponRepository.CouponAlreadyExpired(couponID)
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("coupon does not exist")

}

func (co *couponUseCase) AddProductOffer(productOffer models.ProductOfferReceiver) error {

	return co.couponRepository.AddProductOffer(productOffer)

}

func (co *couponUseCase) AddCategoryOffer(categoryOffer models.CategoryOfferReceiver) error {

	return co.couponRepository.AddCategoryOffer(categoryOffer)

}
