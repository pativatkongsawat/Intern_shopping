package helper

import "time"

type UserFilter struct {
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	PermissionId string `json:"permission_id"`
}

type ProductFilter struct {
	Id          int     `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name        string  `json:"name" gorm:"name"`
	Price       float64 `json:"price" gorm:"price"`
	Quantity    int     `json:"quantity" gorm:"quantity"`
	Category_id int     `json:"category_id" gorm:"category_id"`
}

type CategoryFilter struct {
	Id   int    `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"`
}

type OrderFilter struct {
	Id         int        `json:"id"`
	UserId     string     `json:"user_id" gorm:"user_id"`
	CreateAt   *time.Time `json:"create_at" gorm:"create_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" gorm:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" gorm:"deleted_at"`
	TotalPrice float64    `json:"total_price" gorm:"total_price"`
	Operator   string     `json:"operator" gorm:"operator"`
	Status     string     `json:"status" gorm:"status"`
	MinPrice   int        `json:"min_price"`
	MaxPrice   int        `json:"max_price"`
}
