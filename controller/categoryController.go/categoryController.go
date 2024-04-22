package categoryController

import (
	"Intern_shopping/database"
	"Intern_shopping/models/category/categoryRequest"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InsertCategory(ctx echo.Context) error {
	categorydata := []categoryRequest.Category{}

	categoryModelHelper := categoryRequest.CategoryModelHelper{DB: database.DBMYSQL}

	if err := ctx.Bind(&categorydata); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"Message": "Error req creating category",
		})
	}

	categorys := []categoryRequest.Category{}

	for _, i := range categorydata {
		category := categoryRequest.Category{
			Id:   i.Id,
			Name: i.Name,
		}
		categorys = append(categorys, category)
	}

	if err := categoryModelHelper.InsertCategory(categorys); err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"Message": "Error inserting category",
		})
	}

	return ctx.JSON(200, map[string]interface{}{
		"Category": categorys,
		"Message":  "Success Insert category successfully",
	})
}

func GetAllCategory(ctx echo.Context) error {

	// category := []categoryRequest.Category{}
	categoryModelHelper := categoryRequest.CategoryModelHelper{DB: database.DBMYSQL}

	category, err := categoryModelHelper.GetAllCategory()

	if err != nil {
		return err
	}

	return ctx.JSON(200, map[string]interface{}{
		"category": category,
		"Message":  "Success",
	})

}

func DeleteCategory(ctx echo.Context) error {

	getid := ctx.Param("id")
	id, err := strconv.Atoi(getid)

	if err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"Message": "Error request Id category",
		})
	}

	categoryModelHelper := categoryRequest.CategoryModelHelper{DB: database.DBMYSQL}

	category, err := categoryModelHelper.DeleleteCategory(id)

	if err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"Message": "Error Delete category",
		})
	}

	return ctx.JSON(200, map[string]interface{}{
		"Category": category,
		"Message":  "Delete category successfully",
	})
}

func UpdateCategory(ctx echo.Context) error {
	return nil
}
