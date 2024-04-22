package productRequest

import (
	"Intern_shopping/models/product/productResponse"

	"gorm.io/gorm"
)

type ProductModelHelper struct {
	DB *gorm.DB
}

func (u *ProductModelHelper) Getproduct(pname string, limit, page int) (*productResponse.Product, int64, error) {

	product := []productResponse.Product{}

	// tx.Debug().Where("firstname LIKE ? AND lastname LIKE ?", "%"+fname+"%", "%"+lname+"%").Limit(limit).Offset(offset).Find(&user).Error

	var count int64

	offset := (page - 1) * limit

	if err := u.DB.Debug().Where("name LIKE ?", "%"+pname+"%").Limit(limit).Offset(offset).Find(&product).Error; err != nil {

		return nil, 0, err
	}

	if err := u.DB.Debug().Model(product).Count(&count).Error; err != nil {

		return nil, 0, err
	}

	return nil, count, nil

}

func (u *ProductModelHelper) Insertproduct(*productResponse.Product) error {

	product := []productResponse.Product{}
	tx := u.DB.Begin()

	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
