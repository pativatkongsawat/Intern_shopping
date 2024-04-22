package productController

import (
	"Intern_shopping/database"
	"Intern_shopping/models/product/productRequest"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Fil struct {
	Totalpage     int
	Prevpage      int
	Nextpage      int
	Totalrows     int
	TotalNextpage int
	Totalprevpage int
}

func GetProductBy(ctx echo.Context) error {

	pname := ctx.QueryParam("pname")
	getlimit := ctx.QueryParam("limit")
	getpage := ctx.QueryParam("page")

	limit, _ := strconv.Atoi(getlimit)
	page, _ := strconv.Atoi(getpage)

	// userModelHelper := {DB: database.DBMYSQL}

	productModelHelper := productRequest.ProductModelHelper{DB: database.DBMYSQL}

	user, count, err := productModelHelper.Getproduct(pname, limit, page)

	if err != nil {
		log.Println("Error getting product")
	}

	totalpage := count / int64(limit)

	return ctx.JSON(200, map[string]interface{}{
		"Meta": Fil{
			Totalpage:     int(totalpage),
			Totalrows:     int(count),
			TotalNextpage: int(totalpage) - page,
			Totalprevpage: page - 1,
			Prevpage:      page - 1,
			Nextpage:      page + 1,
		},
		"user": user,
	})

}

// func InsertproductBy(ctx echo.Context) error {

// 	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}

// }
