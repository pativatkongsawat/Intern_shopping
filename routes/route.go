package routes

import (
	guestcontroller "Intern_shopping/controller/guestController"
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) (string, error) {
	e.GET("", guestcontroller.Index)

	//FIXME - Route ยังไม่จัด และต้องจัดการกับสิทธิการใช้งาน
	test := e.Group("/test")
	test.Use(middleware.JWTAuthMiddleware)
	test.GET("/order", guestcontroller.OrderDetailByUserID)
	e.GET("/orders/detail", guestcontroller.OrdersDetail)

	authRoute(e)
	userRoute(e)
	adminRoute(e)
	productRoutes(e)
	categoryRoutes(e)
	orderRoutes(e)

	return "Route works as expected", nil
}

// TODO - SuperAdmin for backOffice
