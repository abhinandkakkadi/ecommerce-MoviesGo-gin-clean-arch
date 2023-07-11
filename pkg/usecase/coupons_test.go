package usecase

// import (
// 	"errors"
// 	"testing"

// 	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/mock"
// 	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// )

// func Test_coupon_AddCoupon(t *testing.T) {

// 	ctrl := gomock.NewController(t)
// 	couponRepo := mock.NewMockCouponRepository(ctrl)
// 	couponUseCase := NewCouponUseCase(couponRepo)

// 	testData := []struct {
// 		name           string
// 		input          models.AddCoupon
// 		buildStub      func(adminRepo *mock.MockCouponRepository)
// 		expectedOutput string
// 		expectedError  error
// 	}{
// 		{
// 			name: "databse error while adding coupon",
// 			input: models.AddCoupon{
// 				Coupon:             "ONAMOFFER",
// 				DiscountPercentage: 20,
// 				MinimumPrice:       1000,
// 				Validity:           true,
// 			},

// 			buildStub: func(couponRepo *mock.MockCouponRepository) {
// 				couponRepo.EXPECT().
// 					CouponExist("ONAMOFFER").
// 					Times(1).
// 					Return(false, errors.New("some error in database"))
// 			},
// 			expectedOutput: "",
// 			expectedError:  errors.New("some error in database"),
// 		},
// 		{
// 			name: "coupon already exists - error while revalidating",
// 			input: models.AddCoupon{
// 				Coupon:             "ONAMOFFER",
// 				DiscountPercentage: 20,
// 				MinimumPrice:       1000,
// 				Validity:           true,
// 			},
// 			buildStub: func(couponRepo *mock.MockCouponRepository) {
// 				couponRepo.EXPECT().
// 					CouponExist("ONAMOFFER").
// 					Times(1).Return(true, nil)
// 				couponRepo.EXPECT().
// 					CouponRevalidateIfExpired("ONAMOFFER").
// 					Times(1).
// 					Return(false, errors.New("error in database while validating"))
// 			},
// 			expectedOutput: "",
// 			expectedError:  errors.New("error in database while validating"),
// 		},
// 		{
// 			name: "The coupon which is valid already exists",
// 			input: models.AddCoupon{
// 				Coupon:             "ONAMOFFER",
// 				DiscountPercentage: 20,
// 				MinimumPrice:       1000,
// 				Validity:           true,
// 			},
// 			buildStub: func(couponRepo *mock.MockCouponRepository) {
// 				couponRepo.EXPECT().
// 					CouponExist("ONAMOFFER").
// 					Times(1).Return(true, nil)
// 				couponRepo.EXPECT().
// 					CouponRevalidateIfExpired("ONAMOFFER").
// 					Times(1).
// 					Return(true, nil)
// 			},
// 			expectedOutput: "The coupon which is valid already exists",
// 			expectedError:  nil,
// 		},
// 		{
// 			name: "Made the coupon valid",
// 			input: models.AddCoupon{
// 				Coupon:             "ONAMOFFER",
// 				DiscountPercentage: 20,
// 				MinimumPrice:       1000,
// 				Validity:           true,
// 			},
// 			buildStub: func(couponRepo *mock.MockCouponRepository) {
// 				couponRepo.EXPECT().
// 					CouponExist("ONAMOFFER").
// 					Times(1).Return(true, nil)
// 				couponRepo.EXPECT().
// 					CouponRevalidateIfExpired("ONAMOFFER").
// 					Times(1).
// 					Return(false, nil)
// 			},
// 			expectedOutput: "Made the coupon valid",
// 			expectedError:  nil,
// 		},
// 		{
// 			name: "The coupon which is valid already exists",
// 			input: models.AddCoupon{
// 				Coupon:             "ONAMOFFER",
// 				DiscountPercentage: 20,
// 				MinimumPrice:       1000,
// 				Validity:           true,
// 			},
// 			buildStub: func(couponRepo *mock.MockCouponRepository) {
// 				couponRepo.EXPECT().
// 					CouponExist("ONAMOFFER").
// 					Times(1).Return(true, nil)
// 				couponRepo.EXPECT().
// 					CouponRevalidateIfExpired("ONAMOFFER").
// 					Times(1).
// 					Return(true, nil)
// 			},
// 			expectedOutput: "The coupon which is valid already exists",
// 			expectedError:  nil,
// 		},
// 		{
// 			name: "successfully added the coupon",
// 			input: models.AddCoupon{
// 				Coupon:             "ONAMOFFER",
// 				DiscountPercentage: 20,
// 				MinimumPrice:       1000,
// 				Validity:           true,
// 			},
// 			buildStub: func(couponRepo *mock.MockCouponRepository) {
// 				couponRepo.EXPECT().
// 					CouponExist("ONAMOFFER").
// 					Times(1).Return(false, nil)
// 				couponRepo.EXPECT().
// 					AddCoupon(models.AddCoupon{
// 						Coupon:             "ONAMOFFER",
// 						DiscountPercentage: 20,
// 						MinimumPrice:       1000,
// 						Validity:           true,
// 					}).
// 					Times(1).
// 					Return(nil)
// 			},
// 			expectedOutput: "successfully added the coupon",
// 			expectedError:  nil,
// 		},
// 	}

// 	for _, tt := range testData {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.buildStub(couponRepo)

// 			resp, err := couponUseCase.AddCoupon(tt.input)

// 			assert.Equal(t, tt.expectedError, err)

// 			assert.Equal(t, tt.expectedOutput, resp)
// 		})
// 	}

// }
