package order

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type OrderModelHelper struct {
	DB *gorm.DB
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
func (u *OrderModelHelper) InsertOrderHasProduct(orderId int, products []RequestProducts) (*[]OrderHasProduct, error) {
	now := time.Now()

	tx := u.DB.Begin()
	hasorder := []OrderHasProduct{}
	for _, p := range products {
		orderhas := OrderHasProduct{
			ProductId:         p.Id,
			OrderId:           orderId,
			OrderProductTotal: p.Quantity,
			OrderProductPrice: p.Price * float64(p.Quantity),
		}
		if err := tx.Debug().Create(&orderhas).Error; err != nil {
			tx.Rollback()
			log.Println("Error creating order has product:", err)
			return nil, err
		}

		hasorder = append(hasorder, orderhas)

		order := Order{
			UpdatedAt:  &now,
			TotalPrice: p.Price * float64(p.Quantity),
		}

		if err := tx.Debug().Model(&Order{}).Where("id = ?", orderId).Updates(order).Error; err != nil {
			tx.Rollback()
			log.Println("Error updating order TotalPrice:", err)
			return nil, err
		}
	}

	tx.Commit()
	return &hasorder, nil
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
