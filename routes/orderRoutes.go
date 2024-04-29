package routes

import (
	"Intern_shopping/controller/orderController"
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func orderRoutes(e *echo.Echo) {

	userOrderGroup := e.Group("/orders")

	userOrderGroup.Use(middleware.JWTAuthMiddleware, middleware.CustomerMiddleware)

	userOrderGroup.POST("", orderController.InsertOrder)

	userOrderGroup.DELETE("", orderController.OrderDelete)

	userOrderGroup.GET("", orderController.OrderDetailByUserID)

	userOrderGroup.GET("", orderController.OrdersDetail)

}
