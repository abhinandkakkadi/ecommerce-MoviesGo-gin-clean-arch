package repository

import (
	"errors"
	"fmt"

	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"gorm.io/gorm"
)

type couponRepository struct {
	DB *gorm.DB
}

func NewCouponRepository(DB *gorm.DB) interfaces.CouponRepository {
	return &couponRepository{
		DB: DB,
	}
}

func (co *couponRepository) CouponExist(couponName string) (bool, error) {

	var count int
	err := co.DB.Raw("select count(*) from coupons where coupon = ?", couponName).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil

}

func (co *couponRepository) CouponRevalidateIfExpired(couponName string) (bool, error) {

	var isValid bool
	err := co.DB.Raw("select validity from coupons where coupon = ?", couponName).Scan(&isValid).Error
	if err != nil {
		return false, err
	}

	if isValid {
		return true, nil
	}

	err = co.DB.Exec("update coupons set validity = true where coupon = ?", couponName).Error
	if err != nil {
		return false, err
	}

	return false, nil

}

func (co *couponRepository) AddCoupon(coupon models.AddCoupon) error {
	fmt.Println("from add coupon repository: ", coupon)
	err := co.DB.Exec("insert into coupons (coupon,discount_percentage,minimum_price,validity) values (?, ?, ?, ?)", coupon.Coupon, coupon.DiscountPercentage, coupon.MinimumPrice, true).Error
	if err != nil {
		return nil
	}

	return nil
}

func (co *couponRepository) GetCoupon() ([]models.Coupon, error) {

	var coupons []models.Coupon
	err := co.DB.Raw("select id,coupon,discount_percentage,minimum_price,Validity from coupons").Scan(&coupons).Error
	if err != nil {
		return []models.Coupon{}, err
	}

	return coupons, nil
}

func (co *couponRepository) ExistCoupon(couponID int) (bool, error) {

	var count int
	err := co.DB.Raw("select count(*) from coupons where id = ?", couponID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (co *couponRepository) CouponAlreadyExpired(couponID int) error {
	fmt.Println("the code reached here")
	var valid bool
	err := co.DB.Raw("select validity from coupons where id = ?", couponID).Scan(&valid).Error
	if err != nil {
		return err
	}
	fmt.Println("the validity = ", valid)
	if valid {
		err := co.DB.Exec("update coupons set validity = false where id = ?", couponID).Error
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("already expired")
}
