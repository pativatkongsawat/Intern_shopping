package routes

import (
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) (string, error) {

	authRoute(e)
	userRoute(e)
	adminRoute(e)
	superAdminRoute(e)

	return "Route works as expected", nil
}

// TODO - SuperAdmin for backOffice
