package productController

import (
	"Intern_shopping/database"
	"Intern_shopping/models/product"
	"Intern_shopping/models/utils"
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
		return ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Error Get Product",
			Result:  err.Error(),
		})
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
		return ctx.JSON(400, utils.ResponseMessage{
			Status:  400,
			Message: "Bind error",
			Result:  err.Error(),
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
		return ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Error Insert Product",
			Result:  err.Error(),
		})
	}

	return ctx.JSON(200, map[string]interface{}{
		"Product": products,
		"Message": "Product insert successfully",
	})
}

func DeleteProductBy(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(400, utils.ResponseMessage{
			Status:  400,
			Message: "Invalid id",
			Result:  err.Error(),
		})
	}

	productModelHelper := product.ProductModelHelper{DB: database.DBMYSQL}
	deletedProduct, err := productModelHelper.DeleteProduct(id)
	if err != nil {
		return ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Delete product failed",
			Result:  err.Error(),
		})
	}

	return ctx.JSON(200, map[string]interface{}{
		"product": deletedProduct,
		"message": "Product deleted successfully",
	})
}

func ProductGetAll(ctx echo.Context) error {

	productModelHelper := product.ProductModelHelper{DB: database.DBMYSQL}

	product, err := productModelHelper.GetproductAll()

	if err != nil {
		return ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Can not Get Product",
		})
	}

	return ctx.JSON(200, map[string]interface{}{
		"Product": product,
		"Message": "Successfully retrieved all products",
	})
}

func UpdateProduct(ctx echo.Context) error {

	productdata := []*product.ProductUpdate{}

	if err := ctx.Bind(&productdata); err != nil {
		return ctx.JSON(400, utils.ResponseMessage{
			Status:  400,
			Message: "Error Bind Productdata",
			Result:  err.Error(),
		})
	}
	now := time.Now()
	newproduct := []*product.Product{}

	for _, i := range productdata {

		newproductsdata := product.Product{
			Id:          i.Id,
			Name:        i.Name,
			Description: i.Description,
			Price:       i.Price,
			Quantity:    i.Quantity,
			Image:       i.Image,
			Update_at:   &now,
			Category_id: i.Category_id,
		}

		newproduct = append(newproduct, &newproductsdata)
	}

	productModelHelper := product.ProductModelHelper{DB: database.DBMYSQL}

	product, err := productModelHelper.UpdateProduct(newproduct)

	if err != nil {
		return ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: " product update failed",
			Result:  err.Error(),
		})
	}

	return ctx.JSON(200, map[string]interface{}{
		"product": product,
		"Message": "Updated product successfully",
	})
}

func DeleteProductSoft(ctx echo.Context) error {

	productModelHelper := product.ProductModelHelper{DB: database.DBMYSQL}

	getid := ctx.Param("id")

	id, err := strconv.Atoi(getid)

	if err != nil {
		return ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Error Delete Soft Product",
			Result:  err.Error(),
		})
	}

	product := productModelHelper.SoftDelete(id)

	return ctx.JSON(200, map[string]interface{}{
		"product": product,
		"Message": "Soft deleted product successfully deleted",
	})

}

func Delelelele(ctx echo.Context) error {
	return nil
}
