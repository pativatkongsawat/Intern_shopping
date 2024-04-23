package product

import (
	"gorm.io/gorm"
)

type ProductModelHelper struct {
	DB *gorm.DB
}

func (u *ProductModelHelper) Getproduct(pname string, limit, page int) ([]Product, int64, error) {

	product := []Product{}

	// tx.Debug().Where("firstname LIKE ? AND lastname LIKE ?", "%"+fname+"%", "%"+lname+"%").Limit(limit).Offset(offset).Find(&user).Error

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

func (u *ProductModelHelper) GetproductAll() ([]Product, error) {

	product := []Product{}

	if err := u.DB.Find(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (u *ProductModelHelper) Insertproduct(products []Product) error {
	tx := u.DB.Begin()

	if err := tx.Debug().Create(&products).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (u *ProductModelHelper) Deleteproduct(id int) ([]*Product, error) {
	product := []*Product{}
	tx := u.DB.Begin()

	if err := tx.Debug().Where("id = ?", id).Delete(product).Error; err != nil {
		return nil, err
	}
	return product, nil

}

func (u *ProductModelHelper) UpdateProduct(Productdata []Product) ([]Product, error) {

	tx := u.DB.Begin()

	newProductdata := []Product{}

	for _, product := range Productdata {
		newProduct := map[string]interface{}{
			"Name":        product.Name,
			"Description": product.Description,
			"Price":       product.Price,
			"Quantity":    product.Quantity,
			"Image":       product.Image,
			"Category_id": product.Category_id,
		}

		if err := tx.Debug().Model(&Product{}).Where("id = ?", product.Id).Updates(newProduct).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		newProductdata = append(newProductdata, product)
	}

	tx.Commit()
	return newProductdata, nil
}
