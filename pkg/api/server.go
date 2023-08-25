package http

import (
	"log"

	_ "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/handler"
	handler "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/handler"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/routes"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
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

	router.Use(gin.Logger())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.UserRoutes(router.Group("/"), userHandler, otpHandler, productHandler, cartHandler, orderHandler, paymentHandler)
	routes.AdminRoutes(router.Group("/admin"), adminHandler, productHandler, orderHandler, userHandler, couponHandler)

	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start(productRepo interfaces.ProductRepository, infoLog *log.Logger, errorLog *log.Logger) {
	go helper.LatestOfferAlert(productRepo)
	infoLog.Printf("starting server on :3000")
	err := sh.engine.Run(":3000")
	errorLog.Fatal(err)
}
