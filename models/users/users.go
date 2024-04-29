package users

import (
	"time"
)

type Users struct {
	ID           string     `gorm:"primaryKey" json:"id"`
	Firstname    string     `gorm:"not null" json:"firstname"`
	Lastname     string     `gorm:"not null" json:"lastname"`
	Address      string     `json:"address"`
	Email        string     `gorm:"unique" json:"email"`
	Password     string     `gorm:"not null" json:"password"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	PermissionID int        `json:"permission_id"`
	UpdatedBy    string     `json:"updated_by"`
}

type CreateUser struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Address   string `json:"address"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `gorm:"not null" json:"password"`
}

type UserUpdate struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Address   string `json:"address"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"password"`
}

type AdminUserMultiUpdate struct {
	ID           string `json:"id"`
	Firstname    string `json:"firstname,omitempty"`
	Lastname     string `json:"lastname,omitempty"`
	Address      string `json:"address,omitempty"`
	Email        string `gorm:"unique" json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	PermissionID int    `json:"permission_id,omitempty"`
}

type UserAuth struct {
	Email        string `gorm:"unique" json:"email"`
	Password     string `json:"password"`
	PermissionID int    `json:"permission_id"`
}

type UserDelete struct {
	ID string `json:"id"`
}

type GetUsersResponse struct {
	ID           string     `gorm:"primaryKey" json:"id"`
	Firstname    string     `gorm:"not null" json:"firstname"`
	Lastname     string     `gorm:"not null" json:"lastname"`
	Address      string     `json:"address"`
	Email        string     `gorm:"unique" json:"email"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	PermissionID int        `json:"permission_id" default:"0"`
}

func (Users) TableName() string {
	return "users"
}

func (GetUsersResponse) TableName() string {
	return "users"
}
