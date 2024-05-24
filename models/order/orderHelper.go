package order

import (
	"Intern_shopping/helper"
	"database/sql"
	"encoding/json"
	"log"
	"math"
	"strconv"

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
		err := rows.Scan(&order.OrderId, &order.UserId, &order.CreateAt, &order.UpdatedAt, &order.CreatedBy, &productsJSON, &order.OrderProductTotal, &order.TotalPrice, &order.Status)
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
		Select("orders.id AS order_id, orders.user_id AS user_id, orders.create_at, orders.updated_at, orders.created_by, JSON_ARRAYAGG(JSON_OBJECT('id', products.id, 'name', products.name, 'description', products.description, 'quantity', order_has_products.order_product_total, 'price', products.price, 'image', products.image, 'total_products_price', order_has_products.order_product_total * products.price, 'category_id', products.category_id)) AS products, SUM(order_has_products.order_product_total) AS total_products, orders.total_price, orders.status").
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
	if orders == nil || len(*orders) == 0 {
		return nil, nil
	}
	return orders, nil
}

// NOTE - Get order details
// ! Super Admin Only
func (u *OrderModelHelper) GetOrdersDetail(p *helper.Pagination, f *helper.OrderFilter) (*[]ResponseOrderHasProduct, error) {
	var data []ResponseOrderHasProduct
	db := u.DB

	// Building the query
	query := db.Debug().Table("orders").
		Select("orders.id AS order_id, orders.user_id AS user_id, orders.create_at, orders.updated_at, orders.created_by, " +
			"JSON_ARRAYAGG(JSON_OBJECT('id', products.id, 'name', products.name, 'description', products.description, " +
			"'quantity', order_has_products.order_product_total, 'price', products.price, 'image', products.image, " +
			"'total_products_price', order_has_products.order_product_total * products.price, 'category_id', products.category_id)) AS products, " +
			"SUM(order_has_products.order_product_total) AS total_products, orders.total_price, orders.status").
		Joins("JOIN order_has_products ON orders.id = order_has_products.order_id").
		Joins("JOIN products ON order_has_products.product_id = products.id").
		Group("orders.id, orders.user_id").
		Order(p.Sort).Count(&p.TotalRows)

	// Adding WHERE conditions
	if f.Id != 0 {
		query = query.Where("orders.id LIKE ?", "%"+strconv.Itoa(f.Id)+"%").Count(&p.TotalRows)
	}
	if f.UserId != "" {
		query = query.Where("orders.user_id LIKE ?", "%"+f.UserId+"%").Count(&p.TotalRows)
	}
	if f.Status != "" {
		query = query.Where("orders.status = ?", f.Status).Count(&p.TotalRows)
	}
	if f.MinPrice > 0 && f.MaxPrice > f.MinPrice {
		query = query.Where("orders.total_price >= ?", f.MinPrice).Count(&p.TotalRows)
	}
	if f.MaxPrice > 0 && f.MaxPrice > f.MinPrice {
		query = query.Where("orders.total_price <= ?", f.MaxPrice).Count(&p.TotalRows)
	}
	if f.MinPrice > 0 && f.MaxPrice > 0 {
		query = query.Where("orders.total_price >= ? AND orders.total_price <= ?", f.MinPrice, f.MaxPrice).Count(&p.TotalRows)
	}
	// switch f.Operator {
	// case "more":
	// 	query = query.Where("orders.total_price > ?", f.TotalPrice).Count(&p.TotalRows)
	// case "less":
	// 	query = query.Where("orders.total_price < ?", f.TotalPrice).Count(&p.TotalRows)
	// case "equal":
	// 	query = query.Where("orders.total_price <= ?", f.TotalPrice).Count(&p.TotalRows)
	// case "=":
	// 	query = query.Where("orders.total_price = ?", f.TotalPrice).Count(&p.TotalRows)
	// default:
	// 	query = query.Where("orders.total_price >= ?", f.TotalPrice).Count(&p.TotalRows)
	// }
	if f.CreateAt != nil {
		query = query.Where("DATE(orders.create_at) = ?", f.CreateAt.Format("2006-01-02")).Count(&p.TotalRows)
	}
	if f.UpdatedAt != nil {
		query = query.Where("DATE(orders.updated_at) = ?", f.UpdatedAt.Format("2006-01-02")).Count(&p.TotalRows)
	}

	query.Limit(p.Row).Offset((p.Page - 1) * p.Row)

	// Execute the query
	rows, err := query.Rows()
	if err != nil {
		return nil, err
	}

	// Calculate total pages
	p.TotalPages = math.Ceil(float64(p.TotalRows) / float64(p.Row))

	// Unmarshal rows into data slice
	orders, err := rowUnmarshal(rows, data)
	if err != nil {
		return nil, err
	}

	if orders == nil || len(*orders) == 0 {
		return nil, nil
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

func (u *OrderModelHelper) GetOrderAll(user_id string, limit, page int) (*[]ResponseOrderHasProduct, int64, error) {

	var data []ResponseOrderHasProduct
	db := u.DB

	var count int64

	offset := (page - 1) * limit

	rows, err := db.Table("orders").
		Select("orders.id AS order_id, orders.user_id AS user_id, orders.create_at, orders.updated_at, orders.created_by, JSON_ARRAYAGG(JSON_OBJECT('id', products.id, 'name', products.name, 'description', products.description, 'quantity', order_has_products.order_product_total, 'price', products.price, 'image', products.image, 'total_products_price', order_has_products.order_product_total * products.price, 'category_id', products.category_id)) AS products, SUM(order_has_products.order_product_total) AS total_products, orders.total_price, orders.status").
		Joins("JOIN order_has_products ON orders.id = order_has_products.order_id").
		Joins("JOIN products ON order_has_products.product_id = products.id").
		Group("orders.id, orders.user_id").Where("orders.deleted_at IS NULL AND orders.user_id =?", user_id).
		Limit(limit).
		Offset(offset).
		Count(&count).
		Rows()
	if err != nil {
		return nil, 0, err
	}
	orders, err := rowUnmarshal(rows, data)
	if err != nil {
		return nil, 0, err
	}
	if orders == nil || len(*orders) == 0 {
		return nil, 0, nil
	}
	return orders, 0, nil
}
