package userRequest

import (
	"Intern_shopping/helper"
	"fmt"
	"math"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string     `gorm:"primaryKey" json:"id"`
	Firstname    string     `gorm:"not null" json:"firstname"`
	Lastname     string     `gorm:"not null" json:"lastname"`
	Address      string     `json:"address"`
	Email        string     `gorm:"unique" json:"email"`
	Password     string     `gorm:"not null" json:"password"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	PermissionID int        `json:"permission_id" default:"0"`
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

type UserMultiUpdate struct {
	ID           string `json:"id"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Address      string `json:"address"`
	Email        string `gorm:"unique" json:"email"`
	Password     string `json:"password"`
	PermissionID int    `json:"permission_id" default:"0"`
}

type UserAuth struct {
	Email        string `gorm:"unique" json:"email"`
	Password     string `json:"password"`
	PermissionID int    `json:"permission_id"`
}

func (User) TableName() string {
	return "users"
}

type DatabaseRequest struct {
	DB *gorm.DB
}

// SECTION - Create
// NOTE Insert
func (d DatabaseRequest) Insert(user *User) error {
	tx := d.DB.Begin()
	result := tx.Debug().Create(&user)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

// NOTE - Insert array
func (d DatabaseRequest) InsertArray(users []*User) error {
	tx := d.DB.Begin()
	result := tx.Create(&users)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

// !SECTION - Create

// SECTION - Read
// NOTE Select
func (d DatabaseRequest) SelectById(id string) (*User, error) {
	tx := d.DB.Debug().Begin()
	user := &User{}
	result := tx.First(&user, "id =?", id)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return user, nil
}
func (d DatabaseRequest) SelectAll(p *helper.Pagination, f *helper.UserFilter) ([]*User, error) {
	tx := d.DB.Begin()
	var users []*User

	if p.Page > int(p.TotalPages) {
		p.Page = int(p.TotalPages)
	} else if p.Page <= 0 {
		p.Page = 1
	}
	tx.Debug().Model(&users).Where("deleted_at", nil).Where("firstname like ? and lastname like ?", "%"+f.Firstname+"%", "%"+f.Lastname+"%").Order(p.Sort).Count(&p.TotalRows)
	result := tx.Limit(p.Row).Offset((p.Page - 1) * p.Row).Find(&users)
	p.TotalPages = math.Ceil(float64(p.TotalRows) / float64(p.Row))
	if result.Error != nil || p.TotalRows == 0 || p.TotalPages == 0 {
		tx.Rollback()
		return nil, fmt.Errorf("no users found")
	}
	tx.Commit()
	return users, nil
}

//!SECTION - Read

// SECTION - Update

// NOTE - แก้ไขข้อมูล User/ Update User
func (d DatabaseRequest) UpdateUser(user_id string, fields User) error {
	tx := d.DB.Begin()
	result := tx.Debug().Model(&User{}).Where("id =?", user_id).Updates(fields)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

//!SECTION  - Update
