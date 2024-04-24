package routes

import (
	"Intern_shopping/controller/userController"
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func adminRoute(e *echo.Echo) {
	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.JWTAuthMiddleware, middleware.AdminMiddleware)

	// SECTION - READ
	// NOTE - Get all users
	adminGroup.GET("/users", userController.GetUsers)
	// NOTE - GET all deleted users
	adminGroup.GET("/deleted-users", userController.GetDeletedUsers)

	// !SECTION - READ

	//SECTION - CREATE
	// adminGroup.POST("/signup", adminController.CreateAdmin)
	// NOTE - Create Multiple Users
	adminGroup.POST("/users/create", userController.CreateUsers)

	// !SECTION - CREATE

	//SECTION - UPDATE
	// adminGroup.PUT("/:id", adminController.UpdateAdmin)

	adminGroup.PUT("/user/edit/:id", userController.UpdateById)
	adminGroup.PUT("/users/edit", userController.AdminUpdateUsers)

	// !SECTION - UPDATE

	//SECTION - DELETE
	// adminGroup.DELETE("/:id", adminController.DeleteAdmin)
	adminGroup.DELETE("/user/delete/:id", userController.DeleteById)
	adminGroup.DELETE("/users/delete", userController.DeleteUsers)
	adminGroup.DELETE("/user/remove/:id", userController.RemoveUser)
	adminGroup.DELETE("/users/remove", userController.RemoveUsers)

	// !SECTION - DELETE
}
