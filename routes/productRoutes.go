package routes

import (
	"Intern_shopping/controller/productController"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo) {

	e.GET("/product/get", productController.GetProductBy)
	e.GET("/product", productController.ProductGetAll)
	e.POST("/product/create", productController.InsertproductBy)
	e.PUT("/product/update", productController.UpdateProduct)
	e.PUT("/product/delete/soft", productController.DeleteProductSoft)
	e.DELETE("/product/delete/", productController.DeleteProductBy)

}
