package productRequest

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id          int        `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name        string     `json:"name" gorm:"name"`
	Description string     `json:"description" gorm:"description"`
	Price       float64    `json:"price" gorm:"price"`
	Quantity    int        `json:"quantity" gorm:"quantity"`
	Image       string     `json:"image" gorm:"image"`
	Created_at  *time.Time `json:"created_at" gorm:"created_at"`
	Update_at   *time.Time `json:"update_at" gorm:"update_at"`
	Deleted_at  *time.Time `json:"deleted_at" gorm:"deleted_at"`
	Category_id int        `json:"category_id" gorm:"category_id"`
}

type ProductInsert struct {
	Name        string  `json:"name" gorm:"name"`
	Description string  `json:"description" gorm:"description"`
	Price       float64 `json:"price" gorm:"price"`
	Quantity    int     `json:"quantity" gorm:"quantity"`
	Image       string  `json:"image" gorm:"image"`
	Category_id int     `json:"category_id" gorm:"category_id"`
}

func (Product) TableName() string {
	return "products"
}

type ProductModelHelper struct {
	DB *gorm.DB
}

func (u *ProductModelHelper) Getproduct(pname string, limit, page int) (*Product, int64, error) {

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

	return nil, count, nil

}

func (u *ProductModelHelper) Insertproduct(products []Product) error {
	tx := u.DB.Begin()

	if err := tx.Debug().Create(&products).Error; err != nil {

		return err
	}

	return nil
}

// func (u *ProductModelHelper) Updateproduct(product []productResponse.Product, id int) (*productResponse.Product, error) {

// 	productupdate := []productResponse.Product{}
// 	// tx.Debug().Model(&User{}).Where("id = ?", User_id).Updates(&user)

// 	return nil, nil
// }

func (u *ProductModelHelper) Deleteproduct(id int) ([]*Product, error) {
	product := []*Product{}
	tx := u.DB.Begin()

	if err := tx.Debug().Where("id = ?", id).Delete(product).Error; err != nil {
		return nil, err
	}
	return product, nil

}

func (u *ProductModelHelper) SoftDelete(id int) ([]Product, error) {

	return nil, nil
}
