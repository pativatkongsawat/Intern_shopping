package main

import (
	"Intern_shopping/config"
	"Intern_shopping/database"
	"Intern_shopping/routes"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	config.Init()
	database.Init()
	msg, err := routes.Init(e)
	if err != nil {
		panic(err)
	}
	fmt.Print(msg)

	e.Logger.Fatal(e.Start(":1323"))

}
