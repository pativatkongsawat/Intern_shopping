package routes

import (
	"Intern_shopping/controller/productController"
	"Intern_shopping/controller/userController"
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func userRoute(e *echo.Echo) {
	userGroup := e.Group("/user")

	userGroup.Use(middleware.JWTAuthMiddleware, middleware.CustomerMiddleware)

	//SECTION - Model User

	// NOTE Get
	userGroup.GET("/profile", userController.GetUserSelf)

	//NOTE - Update
	userGroup.PUT("/profile", userController.UpdateById)

	// NOTE - Delete
	userGroup.DELETE("/profile/", userController.DeleteById)

	// !SECTION - Model User

	//SECTION - Model Order

	// !SECTION - Model Order

	//SECTION - Model Product

	// NOTE - Get product by product name
	userGroup.GET("/by", productController.GetProductBy)
	userGroup.GET("", productController.ProductGetAll)

	//!SECTION - Model Product

}
