package productResponse

import "time"

type Product struct {
	Id          int        `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name        string     `json:"name" gorm:"name"`
	Description string     `json:"description" gorm:"description"`
	Price       float64    `json:"price" gorm:"price"`
	Quantity    int        `json:"quantity" gorm:"quantity"`
	Imageurl    string     `json:"imageurl" gorm:"imageURl"`
	Created_at  *time.Time `json:"created_at" gorm:"created_at"`
	Update_at   *time.Time `json:"update_at" gorm:"update_at"`
	Delete_at   *time.Time `json:"delete_at" gorm:"delete_at"`
	Category_id int        `json:"category_id" gorm:"category_id"`
}

func (Product) Tablename() string {
	return "products"
}

type InsertProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"price"`
	Quantity    int     `json:"quantity" gorm:"quantity"`
	Imageurl    string  `json:"imageurl" gorm:"imageURl"`
}
