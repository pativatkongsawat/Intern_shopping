package orderController

import (
	"Intern_shopping/database"
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
// @Summary Delete Order
// @Description Delete Order from the database
// @Accept json
// @Produce json
// @Param Request body order.OrderCreateRequest true "Update Product"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /orders [post]
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

// @Tags Order
// @Summary SelfOrderDetail
// @Description SelfOrderDetail
// @Accept json
// @Produce json
// @Param id query int true "id"
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
	return e.JSON(200, map[string]interface{}{
		"Orders":  orders,
		"Message": "Successfully orders",
	})

}

func SuperAdminAllOrdersDetail(e echo.Context) error {
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
