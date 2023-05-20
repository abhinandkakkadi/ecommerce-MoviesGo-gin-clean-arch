// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/thnkrn/go-gin-clean-arch/pkg/api"
	"github.com/thnkrn/go-gin-clean-arch/pkg/api/handler"
	"github.com/thnkrn/go-gin-clean-arch/pkg/config"
	"github.com/thnkrn/go-gin-clean-arch/pkg/db"
	"github.com/thnkrn/go-gin-clean-arch/pkg/repository"
	"github.com/thnkrn/go-gin-clean-arch/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	productRepository := repository.NewProductRepository(gormDB)
	otpRepository := repository.NewOtpRepository(gormDB)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productHandler := handler.NewProductHandler(productUseCase)
	otpUseCase := usecase.NewOtpUseCase(cfg,otpRepository)
	otpHandler := handler.NewOtpHandler(otpUseCase)
	serverHTTP := http.NewServerHTTP(userHandler, productHandler, otpHandler)
	return serverHTTP, nil
}
