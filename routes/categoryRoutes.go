package routes

import (
	"Intern_shopping/controller/categoryController.go"

	"github.com/labstack/echo/v4"
)

func CategoryRoutes(e *echo.Echo) {

	e.POST("/category/create", categoryController.InsertCategory)
	e.GET("/category", categoryController.GetAllCategory)
	e.DELETE("/category/del/:id", categoryController.DeleteCategory)
	e.PUT("/category/update/", categoryController.UpdateCategory)

}
