package main

import (
	"Intern_shopping/config"
	"Intern_shopping/database"
	"Intern_shopping/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	config.Init()
	database.Init()
	routes.ProductRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))

}
