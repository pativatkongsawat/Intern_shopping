package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var jwtSecret = []byte(viper.GetString("jwt_secret"))

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(401, "Unauthorized missing JWT token")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || token.Valid {
			return echo.NewHTTPError(401, "invalid JWT token")
		}
		c.Set("user", token)
		return next(c)
	}
}
