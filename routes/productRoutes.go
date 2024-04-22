package routes

import (
	"Intern_shopping/controller/productController"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo) {

	e.GET("/product/get", productController.GetProductBy)

}
