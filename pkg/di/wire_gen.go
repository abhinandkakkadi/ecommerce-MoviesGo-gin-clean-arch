// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/handler"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/config"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/db"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	cartRepository := repository.NewCartRepository(gormDB)
	userRepository := repository.NewUserRepository(gormDB)
	productRepository := repository.NewProductRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository, cartRepository,productRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	productUseCase := usecase.NewProductUseCase(productRepository,cartRepository)
	productHandler := handler.NewProductHandler(productUseCase)
	otpRepository := repository.NewOtpRepository(gormDB)
	otpUseCase := usecase.NewOtpUseCase(cfg, otpRepository)
	otpHandler := handler.NewOtpHandler(otpUseCase)
	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	adminHandler := handler.NewAdminHandler(adminUseCase)

	cartuseCase := usecase.NewCartUseCase(cartRepository)
	cartHandler := handler.NewCartHandler(cartuseCase)

	orderRepository := repository.NewOrderRepository(gormDB)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, cartRepository, userRepository)
	orderHandler := handler.NewOrderHandler(orderUseCase)

	serverHTTP := http.NewServerHTTP(userHandler, productHandler, otpHandler, adminHandler, cartHandler, orderHandler)
	return serverHTTP, nil
}
