package routes

import (
	"Intern_shopping/controller/orderController"
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func orderRoutes(e *echo.Echo) {

	userOrderGroup := e.Group("/orders")

	userOrderGroup.Use(middleware.JWTAuthMiddleware, middleware.CustomerMiddleware)

	userOrderGroup.POST("", orderController.UserCreateOrder)

	userOrderGroup.DELETE("", orderController.OrderDelete)

<<<<<<< HEAD
	userOrderGroup.GET("", orderController.OrderDetailByUserID)

	userOrderGroup.GET("", orderController.OrdersDetail)
=======
	userOrderGroup.GET("/order", orderController.SelfOrderDetail)
>>>>>>> 2a5617e68f843c10ad2848e130f5ac5357576912

}
