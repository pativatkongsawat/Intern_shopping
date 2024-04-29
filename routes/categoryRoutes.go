package routes

import (
	"Intern_shopping/controller/categoryController"
	"Intern_shopping/middleware"

	"github.com/labstack/echo/v4"
)

func categoryRoutes(e *echo.Echo) {

	userProductGroup := e.Group("/categorys")

	userProductGroup.Use(middleware.JWTAuthMiddleware, middleware.CustomerMiddleware)

	userProductGroup.GET("", categoryController.GetAllCategory)
	userProductGroup.POST("", categoryController.InsertCategory)
	userProductGroup.PUT("", categoryController.UpdateCategory)
	userProductGroup.DELETE("/:id", categoryController.DeleteCategory)

}
