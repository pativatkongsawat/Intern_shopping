package routes

import (
	"Intern_shopping/controller/orderController"
	"Intern_shopping/controller/userController"
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func adminRoute(e *echo.Echo) {
	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.JWTAuthMiddleware, middleware.AdminMiddleware)

	// SECTION - READ
	// NOTE - Get all users
	adminGroup.GET("", userController.GetUsers)
	// NOTE - GET all deleted users
	adminGroup.GET("/deleted", userController.GetDeletedUsers)

	// NOTE - CREATE
	// adminGroup.POST("/signup", adminController.CreateAdmin)
	// NOTE - Create Multiple Users
	adminGroup.POST("", userController.CreateUsers)

	// NOTE - UPDATE
	// adminGroup.PUT("/:id", adminController.UpdateAdmin)

	adminGroup.PUT("/:id", userController.UpdateById)
	adminGroup.PUT("", userController.AdminUpdateUsers)

	// NOTE - DELETE
	// adminGroup.DELETE("/:id", adminController.DeleteAdmin)
	adminGroup.DELETE("/user/delete/:id", userController.DeleteById)
	// adminGroup.DELETE("/users/delete", userController.DeleteUsers)

	// TODO - For only Request Needed
	// adminGroup.DELETE("/user/remove/:id", userController.RemoveUser)
	// adminGroup.DELETE("/users/remove", userController.RemoveUsers)

	// !SECTION - USER

	// SECTION - ORDER

	// !SECTION - ORDER
}

func superAdminRoute(e *echo.Echo) {
	superAdmin := e.Group("/back-office/admin")
	superAdmin.Use(middleware.JWTAuthMiddleware, middleware.SuperAdminMiddleware)

	//SECTION - ORDER

	// NOTE - Get order detail information
	// superAdmin.GET("/order/detail", orderController.OrderDetailByUserID)

	// !SECTION - ORDER

	// SECTION - ORDER HAS PRODUCTS

	//Get order-products detail
	superAdmin.GET("/order/detail/", orderController.SuperAdminOrderDetailByUserID)
	superAdmin.GET("/orders/detail", orderController.SuperAdminAllOrdersDetail)
	superAdmin.GET("/orders/fil", orderController.SuperAdminAllOrdersFil)

	//!SECTION - ORDER HAS PRODUCTS
}
