package routes

import (
	"Intern_shopping/controller/auth"

	"github.com/labstack/echo/v4"
)

func authRoute(e *echo.Echo) {
	authRoute := e.Group("/auth")

	authRoute.POST("/signup", auth.Signup)
	authRoute.POST("/login", auth.Login)

	authAdminBackOffice(e)
}

func authAdminBackOffice(e *echo.Echo) {
	authRoute := e.Group("/api/intern-shop/admin")

	authRoute.POST("/login", auth.BackOfficeLogin)
}
