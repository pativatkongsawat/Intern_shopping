package categoryController

import (
	"Intern_shopping/database"
	"Intern_shopping/models/category"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Tags Category
// @Summary Admin Insert a new Category
// @Description Admin Insert a new Category
// @Accept json
// @Produce json
// @Param Request body []category.Category true "Array Category to Inset"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /admin/category [post]
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

// @Tags Category
// @Summary Admin Get all Category
// @Description Admin Get all Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /admin/category [get]
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

// @Tags Category
// @Summary Admin Delete category
// @Description Admin Delete category
// @Accept json
// @Produce json
// @Param id path int true "Id Category"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /admin/category/{id} [delete]
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

// @Tags Category
// @Summary Admin Update Category
// @Description Admin Update Category from the database
// @Accept json
// @Produce json
// @Param Request body []category.CategoryUpdate true "Update Category"
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /admin/category [put]
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
