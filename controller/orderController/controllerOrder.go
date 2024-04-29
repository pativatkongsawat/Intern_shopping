package orderController

import (
	"Intern_shopping/database"
	"Intern_shopping/models/auth"
	"Intern_shopping/models/order"
	"Intern_shopping/models/utils"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// @Tags Order
// @Summary Insert Order
// @Description Insert Order from the database
// @Accept json
// @Produce json
// @Param Request body []order.Requestorder true "Sturct Order to insert"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /orders [post]
func InsertOrder(ctx echo.Context) error {

	now := time.Now()
	reqOrder := order.Requestorder{}
	if err := ctx.Bind(&reqOrder); err != nil {
		return ctx.JSON(400, utils.ResponseMessage{
			Status:  400,
			Message: "Error binding request order",
			Result:  err.Error(),
		})
	}

	newOrder := order.Order{

		UserId:     reqOrder.UserId,
		CreateAt:   &now,
		TotalPrice: reqOrder.TotalPrice,
	}

	orderModelHelper := order.OrderModelHelper{DB: database.DBMYSQL}
	createdOrder, err := orderModelHelper.InsertOrder(&newOrder)

	if err != nil {
		return ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Error creating order",
			Result:  err.Error(),
		})
	}

	createdhasstock, err := orderModelHelper.InsertOrderHasProduct(createdOrder.Id, reqOrder.Products)

	if err != nil {
		return ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Error inserting order",
			Result:  err.Error(),
		})
	}

	return ctx.JSON(200, map[string]interface{}{

		"Order Has Product": createdhasstock,
		"Order":             createdOrder,
		"Request Order":     reqOrder,
		"Message":           "Success",
	})
}

// @Tags Order
// @Summary Delete Order
// @Description Delete Order from the database
// @Accept json
// @Produce json
// @Param id query int true "id"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /orders [delete]
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

func OrderDetailByUserID(e echo.Context) error {
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
// @Summary Get all Order
// @Description Get all Order from the database
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /orders [get]
func OrdersDetail(e echo.Context) error {
	orderModelHelper := order.OrderModelHelper{DB: database.DBMYSQL}

	orders, err := orderModelHelper.GetOrdersDetail()
	if err != nil {
		log.Error(err.Error())
		return e.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Can not Get Orders",
		})
	}
	return e.JSON(200, map[string]interface{}{
		"Orders":  orders,
		"Message": "Successfully retrieved all orders",
	})
}
