package userController

import (
	"Intern_shopping/database"
	"Intern_shopping/helper"
	"Intern_shopping/models/user/userRequest"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var now = time.Now()

// SECTION - Create
// NOTE - สร้าง User เดียว
func CreateUser(ctx echo.Context) error {
	userModelHelper := userRequest.DatabaseRequest{DB: database.DBMYSQL}

	// ANCHOR -  - ดึงข้อมูลจาก Body มาใส่ตัวแปร
	var newUser userRequest.CreateUser
	if err := ctx.Bind(&newUser); err != nil {
		return ctx.JSON(400, map[string]interface{}{"message": "Invalid request body"})
	} else if newUser.Email == "" {
		return ctx.JSON(400, map[string]interface{}{"message": "Please enter an email address"})
	}
	var CheckEmail userRequest.User
	duplicate := userModelHelper.DB.Debug().Where("email = ?", newUser.Email).Find(&CheckEmail)
	if duplicate.RowsAffected != 0 {
		return ctx.JSON(400, map[string]interface{}{"message": "Email already exists"})
	}
	if duplicate.Error != nil {
		return ctx.JSON(500, map[string]interface{}{"message": "Email Select Error"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.JSON(500, "Failed to hash password")
	}
	newUser.Password = string(hashedPassword)

	user := userRequest.User{
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
		return ctx.JSON(500, map[string]interface{}{"message": "Insert Error"})
	}
	return ctx.JSON(200, map[string]interface{}{"message": "Sign up successfully"})
}

// NOTE * สร้างหลาย Users
func CreateUsers(ctx echo.Context) error {
	userModelHelper := userRequest.DatabaseRequest{DB: database.DBMYSQL}

	// ANCHOR -  - ดึงข้อมูลจาก Body มาใส่ตัวแปร
	data := []userRequest.User{}
	err := ctx.Bind(&data)
	if err != nil {
		return ctx.JSON(400, map[string]interface{}{"message": "Invalid request body"})
	}
	users := []*userRequest.User{}
	for _, user := range data {
		user.ID = helper.GenerateUUID()
		user.CreatedAt = &now
		user.UpdatedAt = now
		users = append(users, &user)
	}
	errors := userModelHelper.InsertArray(users)
	if errors != nil {
		return ctx.JSON(400, map[string]interface{}{"message": "Invalid insert body"})
	}

	return ctx.JSON(200, map[string]interface{}{"message": "User created successfully"})
}

// !SECTION - Create

// SECTION - Read

func GetUsers(ctx echo.Context) error {
	userModelHelper := userRequest.DatabaseRequest{DB: database.DBMYSQL}
	pagination := &helper.Pagination{
		Row:  5,
		Page: 1,
	}
	filter := &helper.UserFilter{}

	// *ANCHOR - ดึงค่าจาก QueryParams
	err := echo.QueryParamsBinder(ctx).
		Int("row", &pagination.Row).
		Int("page", &pagination.Page).
		String("sort", &pagination.Sort).
		String("firstname", &filter.Firstname).
		String("lastname", &filter.Lastname).
		BindError()
	if err != nil {
		return ctx.JSON(400, map[string]interface{}{"massage": "Error query param"})
	}

	users, err := userModelHelper.SelectAll(pagination, filter)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{"Error": err.Error()})
	}
	return ctx.JSON(200, map[string]interface{}{"data": users, "pagination": pagination, "message": "success"})
}

//!SECTION - Read

// SECTION - Update

func UpdateById(ctx echo.Context) error {
	userModelHelper := userRequest.DatabaseRequest{DB: database.DBMYSQL}
	id := ctx.Param("id")
	fields := userRequest.UserUpdate{}
	err := ctx.Bind(&fields)
	user := userRequest.User{
		Firstname: fields.Firstname,
		Lastname:  fields.Lastname,
		Email:     fields.Email,
		Address:   fields.Address,
	}
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{"message": "Invalid request body"})
	}
	result := userModelHelper.UpdateUser(id, user)
	if result != nil {
		return ctx.JSON(500, map[string]interface{}{"message": "Update user error"})
	}
	return ctx.JSON(200, map[string]interface{}{"message": "Update user successfully"})
}

// !SECTION - Update
