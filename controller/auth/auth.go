package auth

import (
	"Intern_shopping/controller/userController"
	"Intern_shopping/database"
	"Intern_shopping/models/users"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(viper.GetString("jwt_secret"))

type Claims struct {
	UserID       string `json:"user_id"`
	PermissionID uint   `json:"permission_id"`
	jwt.StandardClaims
}

func GenerateToken(userID string, PermissionID uint) (string, error) {
	claims := Claims{
		userID,
		PermissionID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ExtractClaims(c echo.Context) *Claims {
	user := c.Get("user").(*jwt.Token)
	if claims, ok := user.Claims.(*Claims); ok && user.Valid {
		log.Print(claims)
		return claims
	}
	return nil
}

func Login(ctx echo.Context) error {
	// Bind data from request body
	var loginUser users.Users
	if err := ctx.Bind(&loginUser); err != nil {
		return echo.NewHTTPError(400, "Invalid request body")
	}

	// Find user by email
	var user users.Users
	if err := database.DBMYSQL.Where("email = ?", loginUser.Email).First(&user).Error; err != nil {
		return echo.NewHTTPError(401, "Invalid email or unknown password")
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)); err != nil {
		return echo.NewHTTPError(401, "Invalid email or password")
	}

	// Generate JWT token
	token, err := GenerateToken(user.ID, uint(user.PermissionID))
	if err != nil {
		return echo.NewHTTPError(500, "Failed to generate token")
	}

	return ctx.JSON(200, map[string]string{
		"status": "Login successful",
		"token":  token,
	})
}

func Signup(ctx echo.Context) error {
	userController.CreateUser(ctx)
	return nil
}
