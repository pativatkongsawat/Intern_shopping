package ordercontroller

import (
	"Intern_shopping/database"
	"Intern_shopping/models/order"
	"time"

	"github.com/labstack/echo/v4"
)

// TODO //Read //Create //Update //Delete func of order | order_has_products | product

func GetOrderAll(ctx echo.Context) error {

	orderModelHelper := order.OrderModelHelper{DB: database.DBMYSQL}

	order, err := orderModelHelper.GetAllorder()

	if err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"Message": err.Error(),
		})
	}

	return ctx.JSON(200, map[string]interface{}{
		"Order":   order,
		"Message": "Success ",
	})

}

func InsertOrderAll(ctx echo.Context) error {

	orderdata := []order.OrderInsert{}

	now := time.Now()

	orderModelHelper := order.OrderModelHelper{DB: database.DBMYSQL}

	if err := ctx.Bind(&orderdata); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"Message": err.Error(),
		})
	}

	neworder := []order.Order{}

	for _, i := range orderdata {
		orderss := order.Order{
			Create_at:  &now,
			Updated_at: &now,
			Deleted_at: nil,
			User_id:    i.User_id,
		}

		neworder = append(neworder, order.Order(orderss))
	}

	order, err := orderModelHelper.Insertorder(neworder)

	if err != nil {
		return err
	}

	return nil

}
