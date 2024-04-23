package order

import "gorm.io/gorm"

type OrderModelHelper struct {
	DB *gorm.DB
}

func (u *OrderModelHelper) GetAllorder() ([]Order, error) {

	order := []Order{}

	if err := u.DB.Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil

}

func (u *OrderModelHelper) Insertorder([]Order) ([]Order, error) {

	return nil, nil
}
