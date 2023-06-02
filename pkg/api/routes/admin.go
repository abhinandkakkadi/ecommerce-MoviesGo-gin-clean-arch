package routes

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/handler"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler, orderHandler *handler.OrderHandler, userHandler *handler.UserHandler, couponHandler *handler.CouponHandler) {

	router.POST("/adminsignup", adminHandler.SignUpHandler)
	router.POST("/adminlogin", adminHandler.LoginHandler)
	// api := router.Group("/admin_panel", middleware.AuthorizationMiddleware)
	// api.GET("users", adminHandler.GetUsers)

	router.Use(middleware.AuthorizationMiddleware)
	{
		genres := router.Group("/genres")
		{
			genres.GET("", adminHandler.GetGenres) // change this to get category
			genres.POST("/add_genre", adminHandler.AddCategory)
			genres.GET("/delete_genre/:id", adminHandler.DeleteGenre)
		}

		product := router.Group("/products")
		{
			product.GET("", productHandler.SeeAllProductToAdmin)
			product.GET("/:page", productHandler.SeeAllProductToAdmin)
			product.POST("/add-product", productHandler.AddProduct)
			product.DELETE("/delete-product/:id", productHandler.DeleteProduct)
			product.POST("/update-product", productHandler.UpdateProduct)

		}

		userDetails := router.Group("/users")
		{
			userDetails.GET("", adminHandler.GetUsers)
			userDetails.GET("/:page", adminHandler.GetUsers)
			userDetails.POST("/add-users", userHandler.AddNewUsers)
			userDetails.GET("/block-users/:id", adminHandler.BlockUser)
			userDetails.GET("/unblock-users/:id", adminHandler.UnBlockUser)
		}

		orders := router.Group("/orders")
		{
			orders.GET("", orderHandler.GetAllOrderDetailsForAdmin)
			orders.GET("/:page", orderHandler.GetAllOrderDetailsForAdmin)

			orders.GET("/approve-order/:order_id", orderHandler.ApproveOrder)
			orders.GET("/cancel-order/:order_id", orderHandler.CancelOrderFromAdminSide)
		}

		coupon := router.Group("/coupon")
		{
			coupon.POST("/addcoupon", couponHandler.AddCoupon)
			coupon.GET("", couponHandler.GetCoupon)
			coupon.PATCH("/expire", couponHandler.ExpireCoupon)
		}

	}

}
