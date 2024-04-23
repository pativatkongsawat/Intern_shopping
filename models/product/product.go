package product

import "time"

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
