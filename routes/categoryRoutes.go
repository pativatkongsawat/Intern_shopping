package routes

import (
	"Intern_shopping/controller/categoryController"

	"github.com/labstack/echo/v4"
)

func CategoryRoutes(e *echo.Echo) {
	e.GET("/category", categoryController.GetAllCategory)
	e.POST("/category/create", categoryController.InsertCategory)
	e.PUT("/category/update/", categoryController.UpdateCategory)
	e.DELETE("/category/del/:id", categoryController.DeleteCategory)

}
