package orderController

import (
	"Intern_shopping/database"
	"Intern_shopping/models/order"
	"Intern_shopping/models/utils"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

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
