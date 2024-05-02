package userController

import (
	"Intern_shopping/database"
	"Intern_shopping/helper"
	"Intern_shopping/models/auth"
	"Intern_shopping/models/users"
	"Intern_shopping/models/utils"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var now = time.Now()

// SECTION - Create
// NOTE - สร้าง User เดียว ย้ายไป auth เป็น Sign up แทน

// NOTE * สร้างหลาย Users

// @Tags User
// @Summary Super Admin Create User
// @Description Super Admin Create User
// @Accept json
// @Produce json
// @Param Request body []users.Users true "Sturct User to insert"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /admin/user [post]
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

// @Tags User
// @Summary User Get User by Id
// @Description User Get User by Id
// @Accept json
// @Produce json
// @Param id query string true "Id User"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /user/profile [get]
func GetUserSelf(ctx echo.Context) error {
	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}

	claim := ctx.Get("user").(*jwt.Token)
	userClaim := claim.Claims.(*auth.Claims)

	user, err := userModelHelper.SelectById(userClaim.UserID)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{"message": "Select Error"})
	}
	ctx.SetParamValues(user.Firstname)
	return ctx.JSON(200, user)
}

// NOTE - Get all users

// @Tags User
// @Summary Admin Get User
// @Description Admin Get User
// @Accept json
// @Produce json
// @Param row query int false "row"
// @Param page query int false "page"
// @Param sort query string false "sort"
// @Param firstname query string false "firstname"
// @Param lastname query string false "lastname"
// @Param email query string false "email"
// @Param add query string false "add"
// @Param permission_id query string false "permission_id"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /admin/user [get]
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
	return ctx.JSON(200, utils.ResponseMessage{
		Status:  200,
		Message: "success",
		Result:  map[string]interface{}{"data": users, "pagination": pagination},
	})
}

// NOTE - Get all deleted users

// @Tags User
// @Summary Admin Get User Delete
// @Description Admin Get User Delete
// @Accept json
// @Produce json
// @Param row query int false "row"
// @Param page query int false "page"
// @Param sort query string false "sort"
// @Param firstname query string false "firstname"
// @Param lastname query string false "lastname"
// @Param email query string false "email"
// @Param add query string false "add"
// @Param permission_id query string false "permission_id"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /admin/user/deleted [get]
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

// @Tags User
// @Summary Admin Update User
// @Description Admin Update User
// @Accept json
// @Produce json
// @Param UserID query string false "Id User"
// @Param Request body users.UserUpdate true "Update User"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /admin/user/id [put]
func UpdateById(ctx echo.Context) error {
	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}

	claim := ctx.Get("user").(*jwt.Token)
	userClaim := claim.Claims.(*auth.Claims)
	id := userClaim.UserID

	userReq := users.UserUpdate{}
	if err := ctx.Bind(&userReq); err != nil {
		return ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Invalid request body",
			Result:  err.Error(),
		})
	}
	if reflect.ValueOf(userReq).IsZero() {
		return ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Invalid request body",
			Result:  "No value from request body",
		})
	}

	// err = ctx.Validate(&userReq)
	// if err != nil {
	// 	return ctx.JSON(400, map[string]interface{}{"message": "Invalid Validate User Request " + err.Error()})
	// }

	user := users.Users{
		Firstname: userReq.Firstname,
		Lastname:  userReq.Lastname,
		Email:     userReq.Email,
		Password:  userReq.Password,
		Address:   userReq.Address,
		UpdatedAt: now,
	}

	if result := userModelHelper.UpdateUser(id, id, &user); result != nil {
		return ctx.JSON(500, map[string]interface{}{"message": "Update user error"})
	}

	return ctx.JSON(200, map[string]interface{}{"message": "Update user successfully"})
}

// @Tags  User
// @Summary Admin Update User
// @Description Admin Update User
// @Accept json
// @Produce json
// @Param Request body []users.AdminUserMultiUpdate true "Sturct User to update"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /admin/user [put]
func AdminUpdateUsers(ctx echo.Context) error {
	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}

	data := []*users.AdminUserMultiUpdate{}
	err := ctx.Bind(&data)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{"massage": "Invalid request body"})
	}
	if reflect.ValueOf(len(data)).IsZero() {
		return ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "No value from request body",
			Result:  data,
		})
	}
	if result := userModelHelper.UpdateUserArray(data); result != nil {
		return ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Update user error",
			Result:  result.Error(),
		})
	}
	return ctx.JSON(200, map[string]interface{}{"massage": "User updated success"})
}

// !SECTION - Update

// SECTION - Delete

// @Tags User
// @Summary Admin Soft Delete User
// @Description Admin Soft Delete User
// @Accept json
// @Produce json
// @Param id path string true "Id User"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /admin/user/delete/{id} [delete]
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
