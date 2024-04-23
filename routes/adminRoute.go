package routes

import (
	"Intern_shopping/controller/userController"
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func adminRoute(e *echo.Echo) {
	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.JWTAuthMiddleware, middleware.AdminMiddleware)

	// SECTION - GET
	// NOTE - Get all users
	adminGroup.GET("/users", userController.GetUsers)
	// NOTE - GET all deleted users
	adminGroup.GET("/deleted-users", userController.GetDeletedUsers)

	// !SECTION - GET

	//SECTION - POST
	// adminGroup.POST("/signup", adminController.CreateAdmin)

	// !SECTION - POST

	//SECTION - PUT
	// adminGroup.PUT("/:id", adminController.UpdateAdmin)

	// !SECTION - PUT

	//SECTION - DELETE
	// adminGroup.DELETE("/:id", adminController.DeleteAdmin)
	adminGroup.DELETE("/user/:id", userController.DeleteById)
	adminGroup.DELETE("/user/remove/:id", userController.RemoveUser)

	// !SECTION - DELETE
}
