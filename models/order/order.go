package order

import "time"

type Order struct {
	Id         int        `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	CreateAt   *time.Time `json:"create_at" gorm:"create_at"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at" gorm:"deleted_at"`
	UserId     string     `json:"user_id" gorm:"user_id"`
	TotalPrice float64    `json:"total_price" gorm:"total_price"`
}

func (Order) TableName() string {
	return "order"
}

type Requestorder struct {
	Id         string            `json:"id" gorm:"id"`
	UserId     string            `json:"user_id"`
	TotalPrice float64           `json:"total_price" gorm:"total_price"`
	Products   []RequestProducts `json:"products"`
}

type RequestProducts struct {
	Id       int     `json:"id"`
	Price    float64 `json:"price" gorm:"price"`
	Quantity int     `json:"quantity" gorm:"quantity"`
}
