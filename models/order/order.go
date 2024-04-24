package order

import "time"

type Order struct {
	Id         string     `json:"id" gorm:"id"`
	Create_at  *time.Time `json:"create_at" gorm:"create_at"`
	Updated_at *time.Time `json:"updated_at" gorm:"updated_at"`
	Deleted_at *time.Time `json:"deleted_at" gorm:"deleted_at"`
	User_id    string     `json:"user_id" gorm:"user_id"`
}

func (Order) TableName() string {
	return "order"
}

type OrderInsert struct {
	Create_at  *time.Time `json:"create_at" gorm:"create_at"`
	Updated_at *time.Time `json:"updated_at" gorm:"updated_at"`
	Deleted_at *time.Time `json:"deleted_at" gorm:"deleted_at"`
	User_id    string     `json:"user_id" gorm:"user_id"`
}
