package routes

import (
	"Intern_shopping/controller/userController"
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func userRoute(e *echo.Echo) {
	userGroup := e.Group("/user")

	userGroup.Use(middleware.JWTAuthMiddleware, middleware.CustomerMiddleware)

	//SECTION - GET

	//NOTE - Select
	userGroup.GET("/profile/:id", userController.GetUserSelf)

	// !SECTION - GET

	//SECTION - POST

	// !SECTION - POST

	//SECTION - PUT

	//NOTE - Update
	userGroup.PUT("/:id", userController.UpdateById)

	//!SECTION - PUT

}
