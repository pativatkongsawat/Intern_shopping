package main

import (
	"Intern_shopping/config"
	"Intern_shopping/database"
	_ "Intern_shopping/docs"
	"Intern_shopping/routes"
	"fmt"

	"github.com/go-playground/validator"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if errs := cv.validator.Struct(i); errs != nil {
		var errMsg []string
		for _, err := range errs.(validator.ValidationErrors) {
			errMsg = append(errMsg, err.Field())
		}
		return echo.NewHTTPError(400, map[string]interface{}{
			"message": "Error Input field is required",
			"field":   errMsg,
		})
	}
	return nil
}

// @title			Intern_shopping
// @version		1.0
// @description	This is a sample server.
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:1323
// @BasePath		/
// @schemes		http
// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func main() {

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	config.Init()
	database.Init()
	msg, err := routes.Init(e)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Print(msg)

	e.Logger.Fatal(e.Start(":1323"))

}
