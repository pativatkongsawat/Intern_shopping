package main

import (
	"Intern_shopping/config"
	"Intern_shopping/database"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	config.Init()
	database.Init()

	e.Logger.Fatal(e.Start(":1323"))

}
