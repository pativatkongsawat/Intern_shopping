package orderController

import (
	"Intern_shopping/database"
	"Intern_shopping/helper"
	"Intern_shopping/models/auth"
	"Intern_shopping/models/order"
	"Intern_shopping/models/utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// @Tags Order
// @Summary Super Admin Delete Order
// @Description Super Admin Delete Order
// @Accept json
// @Produce json
// @Param Request body order.OrderCreateRequest true "Update Product"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /back-office/order [post]
func UserCreateOrder(ctx echo.Context) error {
	now := time.Now()
	reqOrder := order.OrderCreateRequest{}
	if err := ctx.Bind(&reqOrder); err != nil {
		return ctx.JSON(400, utils.ResponseMessage{
			Status:  400,
			Message: "Error binding request order",
			Result:  err.Error(),
		})
	}
	claim := ctx.Get("user").(*jwt.Token)
	userClaim := claim.Claims.(*auth.Claims)
	creator := userClaim.UserID

	// NOTE - Calculate
	var totalPrice float64
	for _, v := range reqOrder.Products {
		totalPrice += v.Price * float64(v.Quantity)
	}

	log.Print(totalPrice)
	newOrder := order.Order{
		UserId:     userClaim.UserID,
		CreateAt:   &now,
		TotalPrice: float64(totalPrice),
		CreatedBy:  creator,
	}

	orderModelHelper := order.OrderModelHelper{DB: database.DBMYSQL}

	result, err := orderModelHelper.InsertOrder(&newOrder)

	if err != nil {
		return ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Error creating order",
			Result:  err.Error(),
		})
	}

	createdhasstock, err := orderModelHelper.InsertOrderHasProduct(result.Id, &reqOrder.Products)

	if err != nil {
		return ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Error inserting order",
			Result:  err.Error(),
		})
	}

	return utils.ResponseData(ctx, 200, map[string]interface{}{
		"Order": createdhasstock,
	}, "Success")
}

// @Tags Order
// @Summary Super Admin Delete Order
// @Description Super Admin Delete Order
// @Accept json
// @Produce json
// @Param id query int true "id"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /back-office/order [delete]
func OrderDelete(ctx echo.Context) error {

	getid := ctx.QueryParam("id")

	id, err := strconv.Atoi(getid)

	if err != nil {
		return ctx.JSON(400, utils.ResponseMessage{
			Status:  400,
			Message: "Error invalid id",
			Result:  err.Error(),
		})
	}

	orderModelHelper := order.OrderModelHelper{DB: database.DBMYSQL}

	order, orderhas, err := orderModelHelper.DeleteOrder(int64(id))

	if err != nil {

		return ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Error deleting order",
			Result:  err.Error(),
		})
	}

	return ctx.JSON(200, map[string]interface{}{
		"Order":    order,
		"Orderhas": orderhas,
		"Message":  "Order deleted successfully",
	})
}

// @Tags Order
// @Summary SelfOrderDetail
// @Description SelfOrderDetail
// @Accept json
// @Produce json
// @Param user_id query string true "user_id"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /orders [get]
func SelfOrderDetail(e echo.Context) error {
	orderModelHelper := order.OrderModelHelper{DB: database.DBMYSQL}
	claim := e.Get("user").(*jwt.Token)
	userClaim := claim.Claims.(*auth.Claims)
	updaterId := userClaim.UserID

	orders, err := orderModelHelper.GetOrderByUserID(updaterId)
	if err != nil {
		log.Error(err.Error())
		return e.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Can not Get Orders",
		})
	}
	return e.JSON(200, map[string]interface{}{
		"Orders":  orders,
		"Message": "Successfully orders",
	})

}

// @Tags Order
// @Summary Super Admin Get SelfOrderDetail
// @Description Super Admin Get SelfOrderDetail
// @Accept json
// @Produce json
// @Param id query int true "id"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /back-office/order/detail [get]
func SuperAdminOrderDetailByUserID(e echo.Context) error {
	orderModelHelper := order.OrderModelHelper{DB: database.DBMYSQL}
	id := e.QueryParam("id")

	orders, err := orderModelHelper.GetOrderByUserID(id)
	if err != nil {
		log.Error(err.Error())
		return e.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Can not Get Orders",
		})
	}
	if orders == nil {
		return e.JSON(404, utils.ResponseMessage{
			Status:  404,
			Message: "Order not found",
		})
	}
	return e.JSON(200, map[string]interface{}{
		"Orders":  orders,
		"Message": "Successfully orders",
	})

}

// @Tags Order
// @Summary SuperAdmin Get Order
// @Description SuperAdmin Get Order
// @Accept json
// @Produce json
// @Param row query integer false "Number of rows per page (default 5)"
// @Param page query integer false "Page number (default 1)"
// @Param sort query string false "Sorting order (asc or desc)"
// @Param id query integer false "Order ID"
// @Param user query string false "User ID"
// @Param price query number false "Total price"
// @Param operator query string false "Price operator (eq, gt, lt)"
// @Param create query string false "Created at date (format: yyyy-mm-dd)"
// @Param update query string false "Updated at date (format: yyyy-mm-dd)"
// @Param status query string false "Order status"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /back-office//order/all [get]
func SuperAdminAllOrdersDetail(e echo.Context) error {
	orderModelHelper := order.OrderModelHelper{DB: database.DBMYSQL}
	var createAt time.Time
	var updateAt time.Time
	pagination := &helper.Pagination{Row: 5,
		Page: 1}
	filter := &helper.OrderFilter{
		TotalPrice: 0,
	}
	if err := echo.QueryParamsBinder(e).
		Int("row", &pagination.Row).
		Int("page", &pagination.Page).
		String("sort", &pagination.Sort).
		Int("id", &filter.Id).
		String("user", &filter.UserId).
		Float64("price", &filter.TotalPrice).
		String("operator", &filter.Operator).
		Time("create", &createAt, "2006-01-02").
		Time("update", &updateAt, "2006-01-02").
		String("status", &filter.Status).
		BindError(); err != nil {
		return e.JSON(400, utils.ResponseMessage{
			Status:  400,
			Message: "Error query param",
		})
	}
	log.Print("create", &createAt, "update", &updateAt)
	if createAt.Format("2006-01-02") != "0001-01-01" {
		filter.CreateAt = &createAt
	}
	if updateAt.Format("2006-01-02") != "0001-01-01" {
		filter.UpdatedAt = &updateAt
	}
	orders, err := orderModelHelper.GetOrdersDetail(pagination, filter)
	if err != nil {
		log.Error(err.Error())
		return e.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Can not Get Orders",
		})
	}
	if orders == nil {
		return e.JSON(404, utils.ResponseMessage{
			Status:  404,
			Message: "Order not found",
		})
	}
	return e.JSON(200, map[string]interface{}{
		"Orders":  orders,
		"Page":    pagination,
		"Message": "Successfully retrieved all orders",
	})
}
