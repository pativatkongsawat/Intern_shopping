package main

import (
	"Intern_shopping/config"
	"Intern_shopping/controller/userController"
	"Intern_shopping/database"
	"Intern_shopping/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	config.Init()
	database.Init()
	routes.UserRoute(e)

	e.POST("/admin/user/create", userController.CreateUsers)

	e.Logger.Fatal(e.Start(":1323"))

}
