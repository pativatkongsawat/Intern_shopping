package routes

import (
	"Intern_shopping/controller/userController"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	userGroup := e.Group("/user")

	//SECTION - GET

	//NOTE - Select
	userGroup.GET("", userController.GetUsers)

	// !SECTION - GET

	//SECTION - POST

	//NOTE - Create
	userGroup.POST("/signup", userController.CreateUser)

	// !SECTION - POST

	//SECTION - PUT

	//NOTE - Update
	userGroup.PUT("/:id", userController.UpdateById)

	//!SECTION - PUT
}
