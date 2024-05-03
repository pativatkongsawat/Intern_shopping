package guestcontroller

import (
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.JSON(200, "Home Page")
}

func ImageStreamTesting(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	image, err := file.Open()
	if err != nil {
		return err
	}
	return c.Stream(200, "Image/png", image)
}

func ImageBlobTesting(c echo.Context) error {
	file, err := c.FormFile("img")
	if err != nil {
		return err
	}

	image, err := file.Open()
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(image)
	if err != nil {
		return err
	}
	return c.Blob(200, "Image/png", data)
}
