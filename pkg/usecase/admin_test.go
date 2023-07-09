package usecase

import (
	"errors"
	"testing"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_user_AddAddress(t *testing.T) {

	ctrl := gomock.NewController(t)
	adminRepo := mock.NewMockAdminRepository(ctrl)
	adminUseCase := NewAdminUseCase(adminRepo)

	testData := []struct {
		name           string
		input          string
		buildStub      func(adminRepo mock.MockAdminRepository)
		expectedOutput error
	}{
		{
			name:  "success",
			input: "1",
			buildStub: func(adminRepo mock.MockAdminRepository) {
				adminRepo.EXPECT().
					Delete("1").
					Times(1).
					Return(nil)
			},
			expectedOutput: nil,
		},
		{
			name:  "genre not present",
			input: "2",
			buildStub: func(adminRepo mock.MockAdminRepository) {
				adminRepo.EXPECT().
					Delete("2").
					Times(1).
					Return(errors.New("genre not present"))
			},
			expectedOutput: errors.New("genre not present"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*adminRepo)

			err := adminUseCase.Delete(tt.input)

			assert.Equal(t, tt.expectedOutput, err)
		})
	}

}
