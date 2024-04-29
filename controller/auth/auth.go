package auth

import (
	"Intern_shopping/database"
	"Intern_shopping/helper"
	"Intern_shopping/models/auth"
	"Intern_shopping/models/users"
	"Intern_shopping/models/utils"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(viper.GetString("jwt_secret"))

func GenerateToken(userID *string, PermissionID *int) (string, error) {
	claims := auth.Claims{
		UserID:       *userID,
		PermissionID: *PermissionID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func Login(ctx echo.Context) error {
	// Bind data from request body
	var loginUser users.Users
	if err := ctx.Bind(&loginUser); err != nil {
		return echo.NewHTTPError(400, "Invalid request body")
	}
	user, err := loginHandler(loginUser)
	if err != nil {
		return err
	}

	// Generate JWT token
	if token, err := GenerateToken(&user.ID, &user.PermissionID); err != nil {
		return echo.NewHTTPError(500, "Failed to generate token")
	} else {

		return ctx.JSON(200, map[string]string{
			"status": "Login successful",
			"token":  token,
		})
	}
}
func BackOfficeLogin(ctx echo.Context) error {

	// Bind data from request body
	var loginUser users.Users
	if err := ctx.Bind(&loginUser); err != nil {
		return ctx.JSON(400, utils.ResponseMessage{
			Status:  400,
			Message: "Error binding data",
			Result:  err.Error(),
		})
	}
	user, err := loginHandler(loginUser)
	if err != nil {
		return err
	} else if user.PermissionID != 2 {
		return ctx.JSON(400, utils.ResponseMessage{
			Status:  400,
			Message: "No permission",
		})
	}

	// Generate JWT token
	if token, err := GenerateToken(&user.ID, &user.PermissionID); err != nil {
		return ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Failed to generate token",
			Result:  err.Error(),
		})
	} else {

		return ctx.JSON(200, utils.ResponseMessage{
			Status:  200,
			Message: "Login token generated",
			Result:  token,
		})
	}
}

func loginHandler(userReq users.Users) (user users.Users, err error) {
	if err := database.DBMYSQL.Debug().Where("email = ?", userReq.Email).First(&user).Error; err != nil {
		return user, echo.NewHTTPError(401, "Invalid email or unknown password")
	}
	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password)); err != nil {
		log.Println("Password mismatch: ", err.Error(), userReq.Password)
		return user, echo.NewHTTPError(401, "Invalid email or password")
	}
	return user, nil
}

func Signup(ctx echo.Context) error {
	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}
	now := time.Now()

	// ANCHOR -  - ดึงข้อมูลจาก Body มาใส่ตัวแปร
	var newUser users.CreateUser
	if err := ctx.Bind(&newUser); err != nil {
		return ctx.JSON(400, map[string]interface{}{"message": "Invalid request body"})
	} else if newUser.Email == "" {
		return ctx.JSON(400, map[string]interface{}{"message": "Please enter an email address"})
	}

	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost); err != nil {
		return ctx.JSON(500, "Failed to hash password")
	} else {
		newUser.Password = string(hashedPassword)
	}
	user := users.Users{
		ID:        helper.GenerateUUID(),
		Firstname: newUser.Firstname,
		Lastname:  newUser.Lastname,
		Address:   newUser.Address,
		Email:     newUser.Email,
		Password:  newUser.Password,
		CreatedAt: &now,
		UpdatedAt: now,
		DeletedAt: nil,
	}
	if err := userModelHelper.Insert(&user); err != nil {
		return ctx.JSON(400, utils.ResponseMessage{
			Status:  500,
			Message: "Cannot sign up",
			Result:  err.Error(),
		})
	}
	return ctx.JSON(200, map[string]interface{}{"message": "Sign up successfully"})
}
