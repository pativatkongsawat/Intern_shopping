package routes

import (
	ordercontroller "Intern_shopping/controller/orderController"
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func OrderRoutes(e *echo.Echo) {

	userOrderGroup := e.Group("/user/order")

	userOrderGroup.Use(middleware.JWTAuthMiddleware, middleware.CustomerMiddleware)

	userOrderGroup.GET("/get", ordercontroller.GetOrderAll)
	userOrderGroup.POST("/create", ordercontroller.InsertOrderAll)

}
