package routes

import (
	guestcontroller "Intern_shopping/controller/guestController"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) (string, error) {
	e.GET("", guestcontroller.Index)

	// SECTION - Only for testing purposes
	// e.GET("test/users", guestcontroller.TestGetUsers)
	// e.GET("test/create/users", guestcontroller.TestCreateUsers)
	// !SECTION - Only for testing purposes

	authRoute(e)
	userRoute(e)
	adminRoute(e)
	ProductRoutes(e)
	CategoryRoutes(e)

	return "Route works as expected", nil
}

// TODO - SuperAdmin for backOffice
//! /api/v1/admin/login