package usecase

import (
	"fmt"

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

func (cr *couponUseCase) AddCoupon(coupon models.Coupon) (string, error) {

	// if coupon already exist and if it is expired revalidate it. else give back an error message saying the coupon already exist
	couponExist, err := cr.couponRepository.CouponExist(coupon.Coupon)
	if err != nil {
		return "", err
	}
	fmt.Println("coupon exist :", couponExist)
	if couponExist {
		alreadyValid, err := cr.couponRepository.CouponRevalidateIfExpired(coupon.Coupon)
		if err != nil {
			return "", nil
		}

		if alreadyValid {
			return "The coupon which is valid already exists", nil
		}

		return "Made the coupon valid", nil

	}

	err = cr.couponRepository.AddCoupon(coupon)
	if err != nil {
		return "", err
	}

	return "successfully added the coupon", nil
}

func (cr *couponUseCase) GetCoupon() ([]models.Coupon, error) {

	return cr.couponRepository.GetCoupon()
}
