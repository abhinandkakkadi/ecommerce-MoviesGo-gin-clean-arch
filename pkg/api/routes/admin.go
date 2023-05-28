package routes

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/handler"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler, orderHandler *handler.OrderHandler) {

	router.POST("/adminsignup", adminHandler.SignUpHandler)
	router.POST("/adminlogin", adminHandler.LoginHandler)
	// api := router.Group("/admin_panel", middleware.AuthorizationMiddleware)
	// api.GET("users", adminHandler.GetUsers)

	router.Use(middleware.AuthorizationMiddleware)
	{
		genres := router.Group("/genres")
		genres.GET("", adminHandler.GetGenres)
		genres.POST("/add_genre", adminHandler.AddCategory)
		genres.GET("/delete_genre/:id", adminHandler.DeleteGenre)

		product := router.Group("/products")
		product.POST("/add-product", productHandler.AddProduct)
		product.DELETE("/delete-product/:id", productHandler.DeleteProduct)

		router.GET("/users/block-users/:id", adminHandler.BlockUser)

		userDetails := router.Group("/users")
		userDetails.GET("", adminHandler.GetUsers)
		userDetails.GET("/:page", adminHandler.GetUsers)

		router.GET("/orders",orderHandler.GetAllOrderDetailsForAdmin)
		// router.GET("/approve-order/:order_id",orderHandler.ApproveOrder)

	}

}
