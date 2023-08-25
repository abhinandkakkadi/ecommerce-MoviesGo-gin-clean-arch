package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/mock"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserSignup(t *testing.T) {

	// ctx := context.Background()

	testCase := map[string]struct {
		input         models.UserDetails
		buildStub     func(useCaseMock *mock.MockUserUseCase, signupData models.UserDetails)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"Valid Signup": {
			input: models.UserDetails{
				Name:            "Abhinand",
				Phone:           "9961088604",
				Email:           "nanduttanvsabhi@gmail.com",
				Password:        "132457689",
				ConfirmPassword: "132457689",
			},
			buildStub: func(useCaseMock *mock.MockUserUseCase, signupData models.UserDetails) {
				// copying signupData to domain.user for pass to Mock usecase
				err := validator.New().Struct(signupData)
				if err != nil {
					fmt.Println("validation failed")
				}

				useCaseMock.EXPECT().UserSignUp(signupData).Times(1).Return(models.TokenUsers{
					Users: models.UserDetailsResponse{
						Id:    1,
						Name:  "Abhinand",
						Email: "nanduttanvsabhi@gmail.com",
						Phone: "9961088604",
					},
					AccessToken: "aksjgnalfiugliagbldfgbldfgbladfjnb",
				}, nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, responseRecorder.Code)

			},
		},
	}

	for testName, test := range testCase {
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mock.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase, test.input)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/signup", userHandler.UserSignUp)

			jsonData, err := json.Marshal(test.input)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockRequest, err := http.NewRequest(http.MethodPost, "/signup", body)
			assert.NoError(t, err)
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}
}
