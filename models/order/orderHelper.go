package order

import (
	"database/sql"
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

type OrderModelHelper struct {
	DB *gorm.DB
}

// function for unmarshal response order has product model
func rowUnmarshal(rows *sql.Rows, orders []ResponseOrderHasProduct) (*[]ResponseOrderHasProduct, error) {
	defer rows.Close()

	for rows.Next() {
		var order ResponseOrderHasProduct
		var productsJSON string
		err := rows.Scan(&order.OrderId, &order.UserId, &order.CreateAt, &order.UpdatedAt, &order.CreatedBy, &productsJSON, &order.OrderProductTotal, &order.TotalPrice)
		if err != nil {
			return nil, err
		}

		var products []ResponseProductsOrder
		err = json.Unmarshal([]byte(productsJSON), &products)
		if err != nil {
			return nil, err
		}

		order.Products = products

		orders = append(orders, order)
	}
	return &orders, nil
}

// NOTE - Get the order
func (u *OrderModelHelper) GetOrderByUserID(user_id string) (*[]ResponseOrderHasProduct, error) {
	var data []ResponseOrderHasProduct
	db := u.DB

	rows, err := db.Table("orders").
		Select("orders.id AS order_id, orders.user_id AS user_id, orders.create_at, orders.updated_at, orders.created_by, JSON_ARRAYAGG(JSON_OBJECT('id', products.id, 'name', products.name, 'description', products.description, 'quantity', order_has_products.order_product_total, 'price', products.price, 'image', products.image, 'total_products_price', order_has_products.order_product_total * products.price, 'category_id', products.category_id)) AS products, SUM(order_has_products.order_product_total) AS total_products, SUM(order_has_products.order_product_total * products.price) AS total_price").
		Joins("JOIN order_has_products ON orders.id = order_has_products.order_id").
		Joins("JOIN products ON order_has_products.product_id = products.id").
		Group("orders.id, orders.user_id").Where("orders.deleted_at IS NULL and orders.user_id =?", user_id).
		Rows()
	if err != nil {
		return nil, err
	}
	orders, err := rowUnmarshal(rows, data)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// NOTE - Get order details
// ! Super Admin Only
func (u *OrderModelHelper) GetOrdersDetail() (*[]ResponseOrderHasProduct, error) {
	var data []ResponseOrderHasProduct
	db := u.DB

	rows, err := db.Table("orders").
		Select("orders.id AS order_id, orders.user_id AS user_id, orders.create_at, orders.updated_at, orders.created_by, JSON_ARRAYAGG(JSON_OBJECT('id', products.id, 'name', products.name,'description', products.description, 'quantity', order_has_products.order_product_total, 'price', products.price, 'image', products.image, 'total_products_price', order_has_products.order_product_total * products.price, 'category_id', products.category_id)) AS products, SUM(order_has_products.order_product_total) AS total_products, orders.total_price").
		Joins("JOIN order_has_products ON orders.id = order_has_products.order_id").
		Joins("JOIN products ON order_has_products.product_id = products.id").
		Group("orders.id, orders.user_id").
		Rows()
	if err != nil {
		return nil, err
	}
	orders, err := rowUnmarshal(rows, data)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (u *OrderModelHelper) InsertOrder(orders *Order) (*Order, error) {
	tx := u.DB.Begin()
	if err := tx.Debug().Table("orders").Create(&orders).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return orders, nil
}

func (u *OrderModelHelper) InsertOrderHasProduct(orderId int, products *[]RequestProducts) (*[]OrderHasProduct, error) {

	tx := u.DB.Begin()
	orderDetail := []OrderHasProduct{}
	for _, p := range *products {
		orderhas := OrderHasProduct{
			ProductId:         p.Id,
			OrderId:           orderId,
			OrderProductTotal: p.Quantity,
		}

		if err := tx.Debug().Create(&orderhas).Error; err != nil {
			tx.Rollback()
			log.Println("Error creating order has product:", err)
			return nil, err
		}

		orderDetail = append(orderDetail, orderhas)
	}

	tx.Commit()
	return &orderDetail, nil
}

func (u *OrderModelHelper) DeleteOrder(orderId int64) (*Order, []OrderHasProduct, error) {

	order := Order{}
	orderhas := []OrderHasProduct{}
	tx := u.DB.Begin()

	if err := tx.Debug().Where("id = ?", orderId).Delete(&order).Error; err != nil {
		tx.Rollback()
		log.Println("Error deleting order has products:", err)
		return nil, nil, err

	}

	if err := tx.Debug().Where("order_id = ?", orderId).Delete(&orderhas).Error; err != nil {
		tx.Rollback()
		log.Println("Error deleting order :", err)
		return nil, nil, err
	}

	tx.Commit()

	return &order, orderhas, nil
}
