package routes

import (
	"Intern_shopping/controller/userController"
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func userRoute(e *echo.Echo) {
	userGroup := e.Group("/user")

	userGroup.Use(middleware.JWTAuthMiddleware, middleware.CustomerMiddleware)

	//SECTION - READ

	userGroup.GET("/profile/:id", userController.GetUserSelf)

	// !SECTION - READ

	//SECTION - CREATE

	// !SECTION - CREATE

	//SECTION - UPDATE

	//NOTE - Update
	userGroup.PUT("edit/:id", userController.UpdateById)

	//!SECTION - UPDATE

	// SECTION - DELETE
	userGroup.DELETE("/:id", userController.DeleteById)

	// !SECTION - DELETE

}
