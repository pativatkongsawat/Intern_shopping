package order

type Order_hash_product struct {
	Order_id            int     `json:"order_id" gorm:"order_id"`
	Product_id          int     `json:"product_id" gorm:"product_id"`
	Id                  int     `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Order_product_total float64 `json:"order_product_total" gorm:"order_product_total"`
}

func (Order_hash_product) TableName() string {
	return "order_has_products"
}
