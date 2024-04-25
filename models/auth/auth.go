package auth

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserID       string `json:"user_id"`
	PermissionID int    `json:"permission_id"`
	jwt.StandardClaims
}
