package order

import "time"

type Order struct {
	Id          string     `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Create_at   *time.Time `json:"create_at" gorm:"create_at"`
	Updated_at  *time.Time `json:"updated_at" gorm:"updated_at"`
	Deleted_at  *time.Time `json:"deleted_at" gorm:"deleted_at"`
	User_id     string     `json:"user_id" gorm:"user_id"`
	Total_price float64    `json:"total_price" gorm:"total_price"`
}

func (Order) TableName() string {
	return "order"
}

type Requestorder struct {
	Id         string            `json:"id" gorm:"id"`
	Created_at *time.Time        `json:"created_at"`
	User_id    string            `json:"user_id"`
	Product    []RequestProducts `json:"product"`
}

type RequestProducts struct {
	Id       int `json:"id"`
	Quantity int `json:"quantity"`
}
