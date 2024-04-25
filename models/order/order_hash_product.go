package order

type OrderHasProduct struct {
	OrderId           int     `json:"order_id" gorm:"order_id"`
	ProductId         int     `json:"product_id" gorm:"product_id"`
	Id                int     `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	OrderProductTotal int     `json:"order_product_total" gorm:"order_product_total"`
	OrderProductPrice float64 `json:"order_product_price" gorm:"order_product_price"`
}

func (OrderHasProduct) TableName() string {
	return "order_has_products"
}
