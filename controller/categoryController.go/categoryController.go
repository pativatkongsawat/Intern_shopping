package categoryController

import (
	"Intern_shopping/database"
	"Intern_shopping/models/category"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InsertCategory(ctx echo.Context) error {
	categorydata := []category.Category{}

	categoryModelHelper := category.CategoryModelHelper{DB: database.DBMYSQL}

	if err := ctx.Bind(&categorydata); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"Message": "Error req creating category",
		})
	}

	categorys := []category.Category{}

	for _, i := range categorydata {
		category := category.Category{
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
	categoryModelHelper := category.CategoryModelHelper{DB: database.DBMYSQL}

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

	categoryModelHelper := category.CategoryModelHelper{DB: database.DBMYSQL}

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

	categorydata := []category.CategoryUpdate{}

	if err := ctx.Bind(&categorydata); err != nil {
		log.Println("Error Bind ")
		return err
	}

	categoryModelHelper := category.CategoryModelHelper{DB: database.DBMYSQL}
	categorys := []category.Category{}

	for _, i := range categorydata {

		newCategory := category.CategoryUpdate{
			Id:   i.Id,
			Name: i.Name,
		}
		categorys = append(categorys, category.Category(newCategory))
	}

	newCategory, err := categoryModelHelper.UpdateCategory(categorys)
	if err != nil {
		log.Println("Error Update category")
	}

	return ctx.JSON(200, map[string]interface{}{
		"category": newCategory,
		"Message":  "Updated category Successfully",
	})
}
