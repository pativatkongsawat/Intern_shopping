package middleware

import (
	"Intern_shopping/controller/auth"
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

		if tokenString == "" {
			return echo.NewHTTPError(401, "Unauthorized missing JWT token")
		}

		token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if !token.Valid {
			// log.Print(token.Valid)
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
		claims := extractClaims(c)
		if claims == nil {
			return echo.NewHTTPError(401, "Please Login")
		}
		// log.Println("Permission id", claims.PermissionID)
		if claims.PermissionID != 1 {
			return echo.NewHTTPError(403, "Admin permission required")
		}

		return next(c)
	}
}

func CustomerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := extractClaims(c)
		if claims == nil {
			return echo.NewHTTPError(401, "Please Login")
		}

		if claims.PermissionID != 0 {
			return echo.NewHTTPError(403, "Please Sign in")
		}

		return next(c)
	}
}

func extractClaims(c echo.Context) *auth.Claims {
	user := c.Get("user").(*jwt.Token)
	if claims, ok := user.Claims.(*auth.Claims); ok && user.Valid {
		// log.Print("Claims in extract", claims)
		return claims
	}
	return nil
}
