package http

import (
	handler "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/handler"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/routes"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, productHandler *handler.ProductHandler, otpHandler *handler.OtpHandler, adminHandler *handler.AdminHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler) *ServerHTTP {
	router := gin.New()

	// Use logger from Gin
	router.Use(gin.Logger())

	routes.UserRoutes(router.Group("/"),userHandler,otpHandler,productHandler,cartHandler,orderHandler)
	routes.AdminRoutes(router.Group("/admin"),adminHandler,productHandler)

	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
