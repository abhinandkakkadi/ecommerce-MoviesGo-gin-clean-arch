package routes

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/handler"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler, orderHandler *handler.OrderHandler, userHandler *handler.UserHandler, couponHandler *handler.CouponHandler) {

	router.POST("/adminlogin", adminHandler.LoginHandler)

	router.Use(middleware.AuthorizationMiddleware)
	{
		router.GET("/dashboard", adminHandler.DashBoard)
		router.GET("/sales-report/:period", adminHandler.FilteredSalesReport)
		router.POST("/createadmin", adminHandler.CreateAdmin)

		genres := router.Group("/genres")
		{
			genres.GET("", adminHandler.GetGenres) // change this to get category
			genres.POST("/add_genre", adminHandler.AddGenres)
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
			orders.PUT("/refund-order/:order_id", orderHandler.RefundUser)
		}

		offer := router.Group("/offer")
		{
			coupon := offer.Group("/coupons")
			{
				coupon.POST("/addcoupon", couponHandler.AddCoupon)
				coupon.GET("", couponHandler.GetCoupon)
				coupon.PATCH("/expire/:id", couponHandler.ExpireCoupon)
			}

			offer.POST("/product-offer", couponHandler.AddProdcutOffer)
			offer.POST("/category-offer", couponHandler.AddCategoryOffer)

		}

	}

}
