package users

import (
	"Intern_shopping/helper"
	"fmt"
	"math"

	"gorm.io/gorm"
)

type DatabaseRequest struct {
	DB *gorm.DB
}

// SECTION - Create
// NOTE Insert
func (d DatabaseRequest) Insert(user *Users) error {
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
func (d DatabaseRequest) InsertArray(users []*Users) error {
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
func (d DatabaseRequest) SelectById(id string) (*Users, error) {
	tx := d.DB.Debug().Begin()
	user := &Users{}
	result := tx.First(&user, "id =?", id)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return user, nil
}
func (d DatabaseRequest) SelectAll(p *helper.Pagination, f *helper.UserFilter) ([]*Users, error) {
	tx := d.DB.Begin()
	var users []*Users

	result := tx.Model(&users).Where("deleted_at", nil).Where("firstname like ? and lastname like ? and email like ? and address like ?", "%"+f.Firstname+"%", "%"+f.Lastname+"%", "%"+f.Email+"%", "%"+f.Address+"%").Order(p.Sort).Count(&p.TotalRows)
	result.Debug().Limit(p.Row).Offset((p.Page - 1) * p.Row).Find(&users)
	p.TotalPages = math.Ceil(float64(p.TotalRows) / float64(p.Row))
	if p.Page >= int(p.TotalPages) {
		p.Page = int(p.TotalPages)
	} else if p.Page <= 0 {
		p.Page = 1
	}
	if result.Error != nil || p.TotalRows == 0 {
		tx.Rollback()
		return nil, fmt.Errorf("no users found")
	}
	tx.Commit()
	return users, nil
}

//!SECTION - Read

// SECTION - Update

// NOTE - แก้ไขข้อมูล User/ Update User
func (d DatabaseRequest) UpdateUser(user_id string, fields Users) error {
	tx := d.DB.Begin()
	result := tx.Debug().Model(&Users{}).Where("id =?", user_id).Updates(fields)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

//!SECTION  - Update
