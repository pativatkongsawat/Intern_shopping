package routes

import (
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func adminRoute(e *echo.Echo) {
	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.JWTAuthMiddleware, middleware.AdminMiddleware)

	// SECTION - GET
	// adminGroup.GET("", adminController.GetAdmins)

	// !SECTION - GET

	//SECTION - POST
	// adminGroup.POST("/signup", adminController.CreateAdmin)

	// !SECTION - POST

	//SECTION - PUT
	// adminGroup.PUT("/:id", adminController.UpdateAdmin)

	// !SECTION - PUT

	//SECTION - DELETE
	// adminGroup.DELETE("/:id", adminController.DeleteAdmin)

	// !SECTION - DELETE
}
