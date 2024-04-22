package productController

import (
	"Intern_shopping/database"
	"Intern_shopping/models/product/productRequest"
	"Intern_shopping/models/product/productResponse"
	"log"
	"strconv"
	"time"

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

func InsertproductBy(ctx echo.Context) error {

	productdata := []productResponse.InsertProduct{}

	productModelHelper := productRequest.ProductModelHelper{DB: database.DBMYSQL}
	now := time.Now()

	if err := ctx.Bind(&productdata); err != nil {
		log.Println("Error Bind Product")
	}

	products := []productResponse.Product{}

	for _, p := range productdata {
		product := productResponse.Product{

			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Quantity:    p.Quantity,
			Imageurl:    p.Imageurl,
			Created_at:  &now,
			Update_at:   &now,
			Delete_at:   nil,
		}
		products = append(products, product)
	}

	if err := productModelHelper.Insertproduct(products); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"message": "Error inserting product",
		})
	}

	return ctx.JSON(200, map[string]interface{}{
		"Product": products,
		"Message": "success",
	})

}
