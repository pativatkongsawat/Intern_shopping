package userController

import (
	"Intern_shopping/database"
	"Intern_shopping/helper"
	"Intern_shopping/models/users"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var now = time.Now()

// SECTION - Create
// NOTE - สร้าง User เดียว
func CreateUser(ctx echo.Context) error {
	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}

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
		return ctx.JSON(500, map[string]interface{}{"message": "Insert Error"})
	}
	return ctx.JSON(200, map[string]interface{}{"message": "Sign up successfully"})
}

// NOTE * สร้างหลาย Users
func CreateUsers(ctx echo.Context) error {
	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}

	// ANCHOR -  - ดึงข้อมูลจาก Body มาใส่ตัวแปร
	data := []users.Users{}
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(400, map[string]interface{}{"message": "Invalid request body"})
	}
	users := []*users.Users{}
	for index, user := range data {
		user.ID = helper.GenerateUUID()
		user.CreatedAt = &now
		user.UpdatedAt = now
		users = append(users, &user)
		if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data[index].Password), bcrypt.DefaultCost); err != nil {
			return ctx.JSON(500, "Failed to hash password")
		} else {
			users[index].Password = string(hashedPassword)
		}
	}
	if errors := userModelHelper.InsertArray(users); errors != nil {
		return ctx.JSON(500, map[string]interface{}{"Error": errors.Error()})
	}

	return ctx.JSON(200, map[string]interface{}{"message": "User created successfully"})
}

// !SECTION - Create

// SECTION - Read
// NOTE - Get user by Id
func GetUserSelf(ctx echo.Context) error {
	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}

	id := ctx.Param("id")
	user, err := userModelHelper.SelectById(id)
	log.Print(id)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{"message": "Select Error"})
	}
	return ctx.JSON(200, user)
}

// NOTE - Get all users
func GetUsers(ctx echo.Context) error {
	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}

	pagination := &helper.Pagination{Row: 5}
	filter := &helper.UserFilter{}

	// *ANCHOR - ดึงค่าจาก QueryParams
	err := echo.QueryParamsBinder(ctx).
		Int("row", &pagination.Row).
		Int("page", &pagination.Page).
		String("sort", &pagination.Sort).
		String("firstname", &filter.Firstname).
		String("lastname", &filter.Lastname).
		String("email", &filter.Email).
		String("add", &filter.Address).
		String("permission_id", &filter.PermissionId).
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

// NOTE - Get all deleted users
func GetDeletedUsers(ctx echo.Context) error {
	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}

	pagination := &helper.Pagination{Row: 5}
	filter := &helper.UserFilter{}

	// *ANCHOR - ดึงค่าจาก QueryParams
	err := echo.QueryParamsBinder(ctx).
		Int("row", &pagination.Row).
		Int("page", &pagination.Page).
		String("sort", &pagination.Sort).
		String("firstname", &filter.Firstname).
		String("lastname", &filter.Lastname).
		String("email", &filter.Email).
		String("add", &filter.Address).
		BindError()
	if err != nil {
		return ctx.JSON(400, map[string]interface{}{"massage": "Error query param"})
	}

	users, err := userModelHelper.SelectDeleted(pagination, filter)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{"Error": err.Error()})
	}
	return ctx.JSON(200, map[string]interface{}{"data": users, "pagination": pagination, "message": "success"})
}

//!SECTION - Read

// SECTION - Update

func UpdateById(ctx echo.Context) error {
	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}

	id := ctx.Param("id")
	userReq := users.UserUpdate{}
	err := ctx.Bind(&userReq)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{"message": "Invalid request body"})
	}
	err = ctx.Validate(userReq)
	if err != nil {
		return ctx.JSON(400, map[string]interface{}{"message": "Invalid Validate User Request " + err.Error()})
	}

	user := users.Users{
		Firstname: userReq.Firstname,
		Lastname:  userReq.Lastname,
		Email:     userReq.Email,
		Address:   userReq.Address,
	}

	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost); err != nil {
		return ctx.JSON(500, "Failed to hash password")
	} else {
		user.Password = string(hashedPassword)
	}

	if result := userModelHelper.UpdateUser(id, user); result != nil {
		return ctx.JSON(500, map[string]interface{}{"message": "Update user error"})
	}

	return ctx.JSON(200, map[string]interface{}{"message": "Update user successfully"})
}

func AdminUpdateUsers(ctx echo.Context) error {
	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}

	data := []*users.Users{}
	err := ctx.Bind(&data)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{"massage": "Invalid request body"})
	}
	if result := userModelHelper.UpdateUserArray(data); result != nil {
		return ctx.JSON(500, map[string]interface{}{"Error": result.Error()})
	}
	return ctx.JSON(200, map[string]interface{}{"massage": "User updated success"})
}

// !SECTION - Update

// SECTION - Delete
func DeleteById(ctx echo.Context) error {
	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}

	id := ctx.Param("id")
	result, Dtime, err := userModelHelper.SoftDelete(id)
	if result != "" {
		return ctx.JSON(500, map[string]interface{}{"message": result, "time": Dtime})
	} else if err != nil {
		return ctx.JSON(500, map[string]interface{}{"message": "Delete user error"})
	}
	return ctx.JSON(200, map[string]interface{}{"message": "Delete user successfully", "time": Dtime})
}

// NOTE - Delete user with array
func DeleteUsers(ctx echo.Context) error {
	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}
	ids := []users.UserDelete{}

	if err := ctx.Bind(&ids); err != nil {
		return ctx.JSON(400, map[string]interface{}{"Error": err.Error()})
	}
	if result := userModelHelper.SoftArrayDelete(ids); result != nil {
		return ctx.JSON(403, map[string]interface{}{"Error": result.Error()})
	}
	return ctx.JSON(200, map[string]interface{}{"message": "Delete user successfully"})

}

// NOTE - Hard Delete User
func RemoveUser(ctx echo.Context) error {
	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}

	id := ctx.Param("id")
	if result := userModelHelper.Remove(id); result != nil {
		return ctx.JSON(500, map[string]interface{}{"Error": result.Error()})
	}
	return ctx.JSON(200, map[string]interface{}{"message": "Removed successfully"})
}

// NOTE - Hard Delete Users
func RemoveUsers(ctx echo.Context) error {
	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}
	ids := []users.UserDelete{}

	if err := ctx.Bind(&ids); err != nil {
		return ctx.JSON(500, map[string]interface{}{"Error": "Invalid request body"})
	}
	if result := userModelHelper.RemoveUsers(ids); result != nil {
		return ctx.JSON(403, map[string]interface{}{"Error": result.Error()})
	}
	return ctx.JSON(200, map[string]interface{}{"message": "Removed successfully"})
}

// !SECTION - Delete
