package routes

import (
	guestcontroller "Intern_shopping/controller/guestController"
	"Intern_shopping/controller/orderController"
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func orderRoutes(e *echo.Echo) {

	userOrderGroup := e.Group("/orders")

	userOrderGroup.Use(middleware.JWTAuthMiddleware, middleware.CustomerMiddleware)

	userOrderGroup.POST("/create", orderController.InsertOrder)

	userOrderGroup.DELETE("/remove", orderController.OrderDelete)

	userOrderGroup.GET("/order", guestcontroller.OrderDetailByUserID)

	userOrderGroup.GET("/orders/detail", guestcontroller.OrdersDetail)

}
