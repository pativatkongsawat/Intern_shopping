package routes

import (
	"Intern_shopping/controller/orderController"
	"Intern_shopping/controller/productController"
	"Intern_shopping/controller/userController"
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func userRoute(e *echo.Echo) {
	userGroup := e.Group("/user")
	product := userGroup.Group("/product")
	order := userGroup.Group("/order")
	userGroup.Use(middleware.JWTAuthMiddleware, middleware.CustomerMiddleware)

	//SECTION - Model User

	// NOTE Get
	userGroup.GET("/profile", userController.GetUserSelf)

	//NOTE - Update
	userGroup.PUT("/profile", userController.UpdateById)

	// NOTE - Delete
	// userGroup.DELETE("/profile/:id", userController.DeleteById)

	// !SECTION - Model User

	//SECTION - Model Order

	order.GET("", orderController.SelfOrderDetail)
	order.POST("", orderController.UserCreateOrder)
	order.DELETE("", orderController.OrderDelete)

	// !SECTION - Model Order

	//SECTION - Model Product

	// NOTE - Get product by product name
	// product.GET("", productController.GetProductBy)
	product.GET("", productController.ProductGetCategory)

	//!SECTION - Model Product

}
