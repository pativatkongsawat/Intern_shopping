package routes

import (
	"Intern_shopping/controller/productController"
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo) {
	userProductGroup := e.Group("/user/product")

	userProductGroup.Use(middleware.JWTAuthMiddleware, middleware.CustomerMiddleware)
	userProductGroup.GET("/get", productController.GetProductBy)
	userProductGroup.GET("", productController.ProductGetAll)
	userProductGroup.POST("/create", productController.InsertproductBy)
	userProductGroup.PUT("/update", productController.UpdateProduct)
	userProductGroup.DELETE("/delete", productController.DeleteProductSoft)
	userProductGroup.DELETE("/remove", productController.DeleteProductBy)

}
