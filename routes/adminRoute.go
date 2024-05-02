package routes

import (
	"Intern_shopping/controller/categoryController"
	"Intern_shopping/controller/orderController"
	"Intern_shopping/controller/productController"
	"Intern_shopping/controller/userController"
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func adminRoute(e *echo.Echo) {
	adminGroup := e.Group("/admin")
	userGroup := adminGroup.Group("/user")
	productGroup := adminGroup.Group("/product")
	categoryGroup := adminGroup.Group("/category")
	orderGroup := adminGroup.Group("/order")

	adminGroup.Use(middleware.JWTAuthMiddleware, middleware.AdminMiddleware)

	// SECTION - USER
	// NOTE - Get all users
	userGroup.GET("", userController.GetUsers)
	// NOTE - GET all deleted users
	userGroup.GET("/deleted", userController.GetDeletedUsers)

	// NOTE - Create Multiple Users
	userGroup.POST("", userController.CreateUsers)

	// NOTE - UPDATE
	// userGroup.PUT("/:id", adminController.UpdateAdmin)

	userGroup.PUT("/:id", userController.UpdateById)
	userGroup.PUT("", userController.AdminUpdateUsers)

	// NOTE - DELETE
	// userGroup.DELETE("/:id", adminController.DeleteAdmin)
	userGroup.DELETE("/delete/:id", userController.DeleteById)
	// userGroup.DELETE("/delete", userController.DeleteUsers)

	// TODO - For only Request Needed
	// userGroup.DELETE("/remove/:id", userController.RemoveUser)
	// userGroup.DELETE("/removes", userController.RemoveUsers)
	// !SECTION - USER

	// SECTION - ORDER
	orderGroup.GET("", orderController.SelfOrderDetail)

	orderGroup.POST("", orderController.UserCreateOrder)

	orderGroup.DELETE("", orderController.OrderDelete)
	// !SECTION - ORDER

	// SECTION - PRODUCT
	//NOTE - Get product by name
	productGroup.GET("/by", productController.GetProductBy)

	//NOTE - Get products
	productGroup.GET("", productController.ProductGetAll)

	productGroup.GET("/category", productController.ProductGetCategory)

	productGroup.POST("", productController.InsertproductBy)

	productGroup.PUT("", productController.UpdateProduct)

	productGroup.DELETE("/hide/:id", productController.DeleteProductSoft)

	productGroup.DELETE("/:id", productController.DeleteProductBy)
	// !SECTION - Products

	//SECTION - CATEGORY - PRODUCT
	categoryGroup.GET("", categoryController.GetAllCategory)

	categoryGroup.POST("", categoryController.InsertCategory)

	categoryGroup.PUT("", categoryController.UpdateCategory)

	categoryGroup.DELETE("/:id", categoryController.DeleteCategory)
	//!SECTION - CATEGORY - PRODUCT
}

func superAdminRoute(e *echo.Echo) {
	superAdmin := e.Group("/back-office")
	superAdmin.Use(middleware.JWTAuthMiddleware, middleware.SuperAdminMiddleware)
	userGroup := superAdmin.Group("/user")
	productGroup := superAdmin.Group("/product")
	categoryGroup := superAdmin.Group("/category")
	orderGroup := superAdmin.Group("/order")

	// SECTION - USER
	// NOTE - Get all users
	userGroup.GET("", userController.GetUsers)
	// NOTE - GET all deleted users
	userGroup.GET("/deleted", userController.GetDeletedUsers)

	// NOTE - Create Multiple Users
	userGroup.POST("", userController.CreateUsers)

	// NOTE - UPDATE
	// userGroup.PUT("/:id", adminController.UpdateAdmin)

	userGroup.PUT("/:id", userController.UpdateById)
	userGroup.PUT("", userController.AdminUpdateUsers)

	// NOTE - DELETE
	// userGroup.DELETE("/:id", adminController.DeleteAdmin)
	userGroup.DELETE("/delete/:id", userController.DeleteById)
	// userGroup.DELETE("/delete", userController.DeleteUsers)

	// TODO - For only Request Needed
	// userGroup.DELETE("/remove/:id", userController.RemoveUser)
	// userGroup.DELETE("/removes", userController.RemoveUsers)
	// !SECTION - USER

	// SECTION - ORDER

	orderGroup.POST("", orderController.UserCreateOrder)

	orderGroup.DELETE("", orderController.OrderDelete)

	// SECTION - ORDER HAS PRODUCTS

	//Get order-products detail
<<<<<<< HEAD
	superAdmin.GET("/order/detail/", orderController.SuperAdminOrderDetailByUserID)
	superAdmin.GET("/orders/detail", orderController.SuperAdminAllOrdersDetail)
	superAdmin.GET("/orders/fil", orderController.SuperAdminAllOrdersFil)
=======
	orderGroup.GET("/detail/", orderController.SuperAdminOrderDetailByUserID)
	orderGroup.GET("/all", orderController.SuperAdminAllOrdersDetail)
>>>>>>> 80f7caf032fe931d0e2acf71a39bb5fbb31bcc3a

	//!SECTION - ORDER HAS PRODUCTS

	// !SECTION - ORDER

	// SECTION - PRODUCT
	//NOTE - Get product by name
	productGroup.GET("/by", productController.GetProductBy)

	//NOTE - Get products
	productGroup.GET("", productController.ProductGetAll)

	productGroup.GET("/category", productController.ProductGetCategory)

	productGroup.POST("", productController.InsertproductBy)

	productGroup.PUT("", productController.UpdateProduct)

	productGroup.DELETE("/hide/:id", productController.DeleteProductSoft)

	productGroup.DELETE("/:id", productController.DeleteProductBy)
	// !SECTION - Products

	//SECTION - CATEGORY - PRODUCT
	categoryGroup.GET("", categoryController.GetAllCategory)

	categoryGroup.POST("", categoryController.InsertCategory)

	categoryGroup.PUT("", categoryController.UpdateCategory)

	categoryGroup.DELETE("/:id", categoryController.DeleteCategory)
	//!SECTION - CATEGORY - PRODUCT
}
