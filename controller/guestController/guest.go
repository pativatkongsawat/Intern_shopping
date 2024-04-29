package guestcontroller

import (
	"github.com/labstack/echo/v4"
)

func Index(e echo.Context) error {
	return e.JSON(200, "Home Page")
}
