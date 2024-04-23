package middleware

import (
	"Intern_shopping/controller/auth"
	"log"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var jwtSecret = []byte(viper.GetString("jwt_secret"))

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		str := c.Request().Header.Get("Authorization")
		tokenString := strings.TrimPrefix(str, "Bearer ")

		log.Print(tokenString)
		if tokenString == "" {
			return echo.NewHTTPError(401, "Unauthorized missing JWT token")
		}

		token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		log.Println(token.Claims)
		if !token.Valid {
			log.Print(token.Valid)
			return echo.NewHTTPError(401, "invalid JWT token")
		}
		if err != nil {
			// log.Println(err.Error())
			return echo.NewHTTPError(500, "Error")
		}
		c.Set("user", token)
		return next(c)
	}
}

// NOTE - Admin only
func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := auth.ExtractClaims(c)
		if claims == nil {
			return echo.NewHTTPError(401, "Unauthorized missing JWT token")
		}

		if claims.PermissionID != 1 {
			return echo.NewHTTPError(403, "Admin permission required")
		}

		return next(c)
	}
}

func CustomerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		log.Println("user", user)
		claims := user.Claims.(*auth.Claims)

		log.Println("claim", claims)
		if claims == nil {
			return echo.NewHTTPError(401, "Unauthorized missing JWT token")
		}

		if claims.PermissionID != 0 {
			return echo.NewHTTPError(403, "Please Sign in")
		}

		return next(c)
	}
}
