package utils

import "github.com/labstack/echo/v4"

func ResponseData(ctx echo.Context, code int, data interface{}, message string) error {
	return ctx.JSON(code, map[string]interface{}{
		"message": message,
		"data":    data,
	},
	)
}
