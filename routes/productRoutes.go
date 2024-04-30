package routes

import (
	"Intern_shopping/controller/productController"
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func productRoutes(e *echo.Echo) {
	userProductGroup := e.Group("/products")

	userProductGroup.Use(middleware.JWTAuthMiddleware, middleware.CustomerMiddleware)

	userProductGroup.GET("/by", productController.GetProductBy)
	userProductGroup.GET("", productController.ProductGetAll)
	userProductGroup.GET("/category", productController.ProductGetCategory)
	userProductGroup.POST("", productController.InsertproductBy)
	userProductGroup.PUT("", productController.UpdateProduct)
	userProductGroup.DELETE("/hide/:id", productController.DeleteProductSoft)
	userProductGroup.DELETE("/:id", productController.DeleteProductBy)

}
