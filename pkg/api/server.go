package http

import (
	handler "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/handler"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, productHandler *handler.ProductHandler, otpHandler *handler.OtpHandler, adminHandler *handler.AdminHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, couponHandler *handler.CouponHandler, paymentHandler *handler.PaymentHandler) *ServerHTTP {
	router := gin.New()

	router.LoadHTMLGlob("templates/*.html")

	// Use logger from Gin
	router.Use(gin.Logger())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.UserRoutes(router.Group("/"), userHandler, otpHandler, productHandler, cartHandler, orderHandler, paymentHandler)
	routes.AdminRoutes(router.Group("/admin"), adminHandler, productHandler, orderHandler, userHandler, couponHandler)

	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
