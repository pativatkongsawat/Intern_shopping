package order

import (
	"errors"

	"gorm.io/gorm"
)

type OrderModelHelper struct {
	DB *gorm.DB
}

func (u *OrderModelHelper) InsertOrder(orders *Order) (*Order, error) {

	tx := u.DB.Begin()

	if err := tx.Debug().Table("order").Create(&orders).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return orders, nil
}
func (u *OrderModelHelper) InsertOrderHasProduct(orderId int, products []RequestProducts) error {

	tx := u.DB.Begin()

	for _, p := range products {

		orderhas := OrderHasProduct{
			ProductId:         p.Id,
			OrderId:           orderId,
			OrderProductTotal: p.Quantity,
			OrderProductPrice: p.Price * float64(p.Quantity),
		}

		if err := tx.Debug().Create(&orderhas); err.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("error create order has product")

		}

	}

	tx.Commit()

	return nil
}
