package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/thnkrn/go-gin-clean-arch/cmd/api/docs"
	handler "github.com/thnkrn/go-gin-clean-arch/pkg/api/handler"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, productHandler *handler.ProductHandler,otpHandler *handler.OtpHandler) *ServerHTTP {
	router := gin.New()

	// Use logger from Gin
	router.Use(gin.Logger())

	// Swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Request JWT

// USER SIDE
	//sign up functionality for user
	router.POST("/signup", userHandler.GenerateUser)
	//login functionality for user
	router.POST("/login", userHandler.LoginHandler)
	//send and verify otp along with user autherisation
	router.POST("/send-otp",otpHandler.SendOTP)
	router.POST("/verify-otp",otpHandler.VerifyOTP)

	product := router.Group("/products")
{
	product.GET("",productHandler.ShowAllProducts)
	product.GET("/:id",productHandler.ShowIndividualProducts)
}

// ADMIN SIDE

	router.POST("/adminsignup",adminHandler.LoginHandler)



	// Auth middleware
	// api := router.Group("/moviesgo", middleware.AuthorizationMiddleware)

	// api.GET("users", userHandler.FindAll)
	// api.GET("users/:id", userHandler.FindByID)
	// api.POST("users", userHandler.Save)
	// api.DELETE("users/:id", userHandler.Delete)

	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
