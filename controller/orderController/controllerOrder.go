package ordercontroller

import (
	"Intern_shopping/database"
	"Intern_shopping/models/order"
	"log"

	"github.com/labstack/echo/v4"
)

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

	Order := []order.Order{}

	orderModelHelper := order.OrderModelHelper{DB: database.DBMYSQL}

	if err := ctx.Bind(&Order); err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"Message": err.Error,
		})
	}

	if err, _ := orderModelHelper.Insertorder(Order); err != nil {
		log.Println("Error inserting order")
	}

	return ctx.JSON(200, map[string]interface{}{
		"order":   Order,
		"Message": "Success",
	})
}
