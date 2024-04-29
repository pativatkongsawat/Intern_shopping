package routes

import (
	guestcontroller "Intern_shopping/controller/guestController"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) (string, error) {
	e.GET("", guestcontroller.Index)

	authRoute(e)
	userRoute(e)
	adminRoute(e)
	productRoutes(e)
	categoryRoutes(e)
	orderRoutes(e)
	superAdminRoute(e)

	return "Route works as expected", nil
}

// TODO - SuperAdmin for backOffice
