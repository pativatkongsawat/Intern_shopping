package productController

import (
	"Intern_shopping/database"
	"Intern_shopping/models/product"
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

	productModelHelper := product.ProductModelHelper{DB: database.DBMYSQL}

	product, count, err := productModelHelper.Getproduct(pname, limit, page)

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
		"Product": product,
	})

}

func InsertproductBy(ctx echo.Context) error {
	productdata := []product.ProductInsert{}
	productModelHelper := product.ProductModelHelper{DB: database.DBMYSQL}
	now := time.Now()

	if err := ctx.Bind(&productdata); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"Message": "Error request Insert Product",
		})
	}

	products := []product.Product{}

	for _, p := range productdata {
		product := product.Product{
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Quantity:    p.Quantity,
			Image:       p.Image,
			Created_at:  &now,
			Update_at:   &now,
			Deleted_at:  nil,
			Category_id: p.Category_id,
		}
		products = append(products, product)
	}

	err := productModelHelper.Insertproduct(products)
	if err != nil {
		log.Println("Error inserting product:", err)
		return ctx.JSON(500, map[string]interface{}{
			"Message": "Error inserting product",
		})
	}

	return ctx.JSON(200, map[string]interface{}{
		"Product": products,
		"Message": "Product insert successfully",
	})
}

func DeleteproductBy(ctx echo.Context) error {

	getid := ctx.QueryParam("id")
	id, err := strconv.Atoi(getid)

	if err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"message": "Error request",
		})
	}
	productModelHelper := product.ProductModelHelper{DB: database.DBMYSQL}

	product, err := productModelHelper.Deleteproduct(id)

	if err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"Message": "Error deleting product",
		})
	}

	return ctx.JSON(200, map[string]interface{}{
		"Product": product,
		"Message": "Product deleted successfully",
	})
}

func ProductGetAll(ctx echo.Context) error {

	productModelHelper := product.ProductModelHelper{DB: database.DBMYSQL}

	product, err := productModelHelper.GetproductAll()

	if err != nil {
		log.Println("Error getting all products")
	}

	return ctx.JSON(200, map[string]interface{}{
		"Product": product,
		"Message": "Successfully retrieved all products",
	})
}

func UpdateProduct(ctx echo.Context) error {

	productdata := []product.ProductUpdate{}

	if err := ctx.Bind(&productdata); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"Message": "Bind Error updating product",
		})
	}
	now := time.Now()
	newproduct := []product.Product{}

	for _, i := range productdata {

		newproductsdata := product.Product{
			Name:        i.Name,
			Description: i.Description,
			Price:       i.Price,
			Quantity:    i.Quantity,
			Image:       i.Image,
			Update_at:   &now,
			Category_id: i.Category_id,
		}

		newproduct = append(newproduct, product.Product(newproductsdata))
	}

	productModelHelper := product.ProductModelHelper{DB: database.DBMYSQL}

	product, err := productModelHelper.UpdateProduct(newproduct)

	if err != nil {
		log.Println("Error updating product")
	}

	return ctx.JSON(200, map[string]interface{}{
		"product": product,
		"Message": "Updated product successfully",
	})
}
