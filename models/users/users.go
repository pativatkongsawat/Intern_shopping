package users

import (
	"Intern_shopping/models/permission"
	"time"
)

type Users struct {
	ID           string                `gorm:"primaryKey" json:"id"`
	Firstname    string                `gorm:"not null" json:"firstname"`
	Lastname     string                `gorm:"not null" json:"lastname"`
	Address      string                `json:"address"`
	Email        string                `gorm:"unique" json:"email"`
	Password     string                `gorm:"not null" json:"password"`
	CreatedAt    *time.Time            `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
	DeletedAt    *time.Time            `json:"deleted_at"`
	PermissionID int                   `json:"permission_id" default:"0"`
	Permission   permission.Permission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
	Firstname    string                `json:"firstname"`
	Lastname     string                `json:"lastname"`
	Address      string                `json:"address"`
	Email        string                `gorm:"unique" json:"email"`
	Password     string                `json:"password"`
	PermissionID int                   `json:"permission_id"`
	Permission   permission.Permission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
