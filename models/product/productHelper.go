package product

import (
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ProductModelHelper struct {
	DB *gorm.DB
}

func (u *ProductModelHelper) Getproduct(pname string, limit, page int) ([]*Product, int64, error) {

	product := []*Product{}

	var count int64

	offset := (page - 1) * limit

	if err := u.DB.Debug().Where("name LIKE ?", "%"+pname+"%").Limit(limit).Offset(offset).Find(&product).Error; err != nil {

		return nil, 0, err
	}

	if err := u.DB.Debug().Model(product).Count(&count).Error; err != nil {

		return nil, 0, err
	}

	return product, count, nil

}

func (u *ProductModelHelper) InsertProduct(products []*Product) error {

	tx := u.DB.Begin()

	validate := validator.New()

	for _, product := range products {
		if err := validate.Struct(product); err != nil {

			tx.Rollback()

			if _, ok := err.(*validator.InvalidValidationError); ok {
				fmt.Println("Invalid validation error:", err)
				return err
			}

			return fmt.Errorf("validation error: %v", err)
		}
	}

	if err := tx.Debug().Create(&products).Error; err != nil {

		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (u *ProductModelHelper) DeleteProduct(id int) ([]*Product, error) {
	product := []*Product{}
	tx := u.DB.Begin()

	if err := tx.Debug().Where("id = ?", id).First(&product).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Debug().Delete(&product).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return product, nil
}

func (u *ProductModelHelper) UpdateProduct(Productdata []*Product) ([]*Product, error) {

	tx := u.DB.Begin()

	newProductdata := []*Product{}

	for _, product := range Productdata {
		newProduct := map[string]interface{}{
			"Name":        product.Name,
			"Description": product.Description,
			"Price":       product.Price,
			"Quantity":    product.Quantity,
			"Image":       product.Image,
			"Update_at":   product.Update_at,
			"Category_id": product.Category_id,
		}

		if err := tx.Debug().Model(&Product{}).Where("id = ?", product.Id).Updates(&newProduct).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		newProductdata = append(newProductdata, product)
	}

	tx.Commit()
	return newProductdata, nil
}

func (u *ProductModelHelper) SoftDelete(id int) ([]*Product, error) {
	tx := u.DB.Begin()

	product := []*Product{}

	now := time.Now()
	if err := tx.Debug().Model(&Product{}).Where("id = ?", id).Update("deleted_at", &now).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (u *ProductModelHelper) ProductGet(pname, cname, sort string, limit, page int) ([]ProductCategory, int64, error) {

	product := []ProductCategory{}

	// result := "SELECT products.name , products.price  , category.name AS category_name FROM products LEFT JOIN category ON products.category_id = category.id ;"

	offset := (page - 1) * limit

	query := u.DB.Debug().
		Model(&Product{}).
		Select("products.name, products.price, category.name AS category_name").
		Joins("LEFT JOIN category ON products.category_id = category.id").
		Where("products.name LIKE ? AND category.name LIKE ?", "%"+pname+"%", "%"+cname+"%")

	if sort != "" {
		query = query.Order("products.price " + sort)
	}

	query = query.Limit(limit).Offset(offset)

	err := query.Find(&product).Error
	if err != nil {

		return nil, 0, err
	}
	var count int64

	if err := u.DB.Debug().Model(&Product{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return product, count, nil
}
