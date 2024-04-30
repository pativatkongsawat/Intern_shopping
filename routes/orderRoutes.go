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

=======
>>>>>>> 51dc93ff1dfb2b9f83cc01d3b64a513a7044a969
	userOrderGroup.GET("", orderController.SelfOrderDetail)

}
