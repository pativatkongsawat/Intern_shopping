package ordercontroller

import (
	"Intern_shopping/database"
	"Intern_shopping/models/order"
	"Intern_shopping/models/utils"

	"github.com/labstack/echo"
)

// TODO //Read //Create //Update //Delete func of order | order_has_products | product

func GetOrderAll(ctx echo.Context) error {

	// reqorder := new(order.Requestorder)
	reqorder := []order.Requestorder{}

	if err := ctx.Bind(&reqorder); err != nil {
		return ctx.JSON(400, utils.ResponseMessage{
			Status:  400,
			Message: "Error Bind request order",
			Result:  err.Error(),
		})
	}

	neworder := []order.Order{}

	for _, i := range reqorder {

		orderdata := order.Order{
			Id:        i.Id,
			Create_at: i.Created_at,
			User_id:   i.User_id,
		}

		neworder = append(neworder, orderdata)

	}

	orderModelHelper := order.OrderModelHelper{DB: database.DBMYSQL}

	order, err := orderModelHelper.Insertorder(neworder)

	if err != nil {
		return ctx.JSON(500, utils.ResponseMessage{
			Status: 500,
		})
	}

	return ctx.JSON(200, map[string]interface{}{
		"order":   order,
		"Message": "Order Insert Success",
	})
}
