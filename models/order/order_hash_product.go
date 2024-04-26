package order

import (
	"time"
)

type OrderHasProduct struct {
	OrderId           int     `json:"order_id" gorm:"order_id"`
	ProductId         int     `json:"product_id" gorm:"product_id"`
	Id                int     `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	OrderProductTotal int     `json:"order_product_total" gorm:"order_product_total"`
	OrderProductPrice float64 `json:"order_product_price" gorm:"order_product_price"`
}

type ResponseOrderHasProduct struct {
	OrderId           int                     `json:"order_id" gorm:"order_id"`
	OrderProductTotal int                     `json:"order_product_total" gorm:"order_product_total"`
	OrderProductPrice float64                 `json:"order_product_price" gorm:"order_product_price"`
	TotalPrice        float64                 `json:"total_price" gorm:"total_price"`
	UserId            string                  `json:"user_id" gorm:"user_id"`
	Products          []ResponseProductsOrder `json:"products " gorm:"products"`
	CreateAt          *time.Time              `json:"create_at" gorm:"create_at"`
	UpdatedAt         *time.Time              `json:"updated_at" gorm:"updated_at"`
}

type ResponseProductsOrder struct {
	Id          int     `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name        string  `json:"name" gorm:"name"`
	Description string  `json:"description" gorm:"description"`
	Price       float64 `json:"price" gorm:"price"`
	Quantity    int     `json:"quantity" gorm:"quantity"`
	Image       string  `json:"image" gorm:"image"`
	Category_id int     `json:"category_id" gorm:"category_id"`
	TotalPrice  float64 `json:"total_price" gorm:"total_price"`
}

// func (ResponseProductsOrder) TableName() string {
// 	return "products"
// }

func (OrderHasProduct) TableName() string {
	return "order_has_products"
}
