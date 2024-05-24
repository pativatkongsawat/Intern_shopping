package routes

import (
	guestcontroller "Intern_shopping/controller/guestController"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) (string, error) {

	e.POST("upload/stream", guestcontroller.ImageStreamTesting)
	e.POST("upload/blob", guestcontroller.ImageBlobTesting)

	authRoute(e)
	userRoute(e)
	adminRoute(e)
	superAdminRoute(e)

	return "Route works as expected", nil
}
