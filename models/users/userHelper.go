package users

import (
	"Intern_shopping/helper"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type DatabaseRequest struct {
	DB *gorm.DB
}

var now = time.Now()

// REVIEW - Function เช็คว่าเป็น Error Duplicate รึป่าว
func isDuplicateError(err error) bool {
	if err == nil {
		return false
	}
	// Check for specific MySQL error code 1062 or SQL state 23000
	return strings.Contains(err.Error(), "1062") || strings.Contains(err.Error(), "23000")
}

// SECTION - Create
// NOTE Insert
func (d DatabaseRequest) Insert(user *Users) error {
	tx := d.DB.Begin()
	result := tx.Debug().Create(&user)
	if result.Error != nil {
		if isDuplicateError(result.Error) {
			return fmt.Errorf("duplicate email")
		}
		tx.Rollback()
		return fmt.Errorf("error creating user: %v", result.Error)
	}
	tx.Commit()
	return nil
}

// NOTE - Insert array
func (d DatabaseRequest) InsertArray(users []*Users) error {
	tx := d.DB.Begin()
	if result := tx.Create(&users); result.Error != nil {
		if isDuplicateError(result.Error) {
			return fmt.Errorf("duplicate email")
		}
		tx.Rollback()
		return fmt.Errorf("error creating user: %v", result.Error)
	}

	tx.Commit()
	return nil
}

// !SECTION - Create

// SECTION - Read
// NOTE Select
func (d DatabaseRequest) SelectById(id string) (*Users, error) {
	user := &Users{}
	result := d.DB.First(&user, "id =?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// NOTE - Select Deleted user
func (d DatabaseRequest) SelectDeleted(p *helper.Pagination, f *helper.UserFilter) ([]*Users, error) {
	var users []*Users
	result := d.DB.Model(&users).Where("deleted_at IS NOT NULL").Where("firstname like ? and lastname like ? and email like ? and address like ?", "%"+f.Firstname+"%", "%"+f.Lastname+"%", "%"+f.Email+"%", "%"+f.Address+"%").Order(p.Sort).Count(&p.TotalRows)
	result.Debug().Limit(p.Row).Offset((p.Page - 1) * p.Row).Find(&users)
	p.TotalPages = math.Ceil(float64(p.TotalRows) / float64(p.Row))
	if p.Page >= int(p.TotalPages) {
		p.Page = int(p.TotalPages)
	} else if p.Page <= 0 {
		p.Page = 1
	}
	if result.Error != nil || p.TotalPages == 0 && p.TotalRows == 0 {
		return nil, fmt.Errorf("no users found")
	}
	return users, nil
}

// NOTE - Select All users
func (d DatabaseRequest) SelectAll(p *helper.Pagination, f *helper.UserFilter) ([]*GetUsersResponse, error) {
	var users []*GetUsersResponse

	result := d.DB.Model(&users).Where("deleted_at is null").Where("firstname like ? and lastname like ? and email like ? and address like ?", "%"+f.Firstname+"%", "%"+f.Lastname+"%", "%"+f.Email+"%", "%"+f.Address+"%").Order(p.Sort).Count(&p.TotalRows)

	if f.PermissionId != "" {
		permissionId, _ := strconv.Atoi(f.PermissionId)
		result.Where("permission_id = ?", permissionId)
	}

	result.Debug().Limit(p.Row).Offset((p.Page - 1) * p.Row).Find(&users)
	p.TotalPages = math.Ceil(float64(p.TotalRows) / float64(p.Row))
	if p.Page >= int(p.TotalPages) {
		p.Page = int(p.TotalPages)
	} else if p.Page <= 0 {
		p.Page = 1
	}
	if result.Error != nil || p.TotalPages == 0 && p.TotalRows == 0 {
		return nil, fmt.Errorf("no users found")
	}
	return users, nil
}

//!SECTION - Read

// SECTION - Update

// NOTE - แก้ไขข้อมูล User/ Update User
func (d DatabaseRequest) UpdateUser(user_id string, userReq Users) error {
	tx := d.DB.Begin()
	result := tx.Debug().Model(&Users{}).Where("id =?", user_id).Updates(userReq)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

// NOTE - แก้ไขข้อมูล Users/ Update Users หลายตัวพร้อมกัน
func (d *DatabaseRequest) UpdateUserArray(fields []*Users) error {
	// Start a new transaction
	tx := d.DB.Begin()

	defer func() {
		// If the function exits with an error, rollback the transaction
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()
	log.Print("fields: ", fields)
	for _, item := range fields {
		var user Users
		if result := tx.Debug().First(&user, "id = ?", item.ID).Error; result != nil {
			if errors.Is(result, gorm.ErrRecordNotFound) {
				tx.Rollback()
				return fmt.Errorf("%s Article with ID %s not found", "Error 404", fmt.Sprint(item.ID))
			}
			tx.Rollback()
			return fmt.Errorf("failed to update user %w", result)
		}
		if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(item.Password), bcrypt.DefaultCost); err != nil {
			return fmt.Errorf("%s %d Failed to hash password", "Error", 500)
		} else {
			user.Password = string(hashedPassword)
		}
		user.Firstname = item.Firstname
		user.Lastname = item.Lastname
		user.Email = item.Email
		user.Address = item.Address
		user.PermissionID = item.PermissionID
		if result := tx.Debug().Updates(&user).Error; result != nil {
			tx.Rollback()
			return fmt.Errorf("%s Failed to update article with ID %s", "Error 500", fmt.Sprint(item.ID))
		}
	}

	tx.Commit()
	return nil
}

//!SECTION  - Update

// SECTION - Delete

// NOTE - Soft Delete
// TODO - Check in controller before sending to func
func (d DatabaseRequest) SoftDelete(id string) (string, time.Time, error) {
	now := time.Now()
	tx := d.DB.Begin()
	user := &Users{}
	result := tx.Debug().Model(user).Where("id =?", id).Find(user)
	if user.DeletedAt != nil {
		return "This user already deleted", *user.DeletedAt, nil
	} else if result.Error != nil {
		tx.Rollback()
		return "", *user.DeletedAt, result.Error
	}
	result.Debug().Model(user).Where("id =? and deleted_at is null", id).Update("deleted_at", now)
	tx.Commit()
	return "", *user.DeletedAt, nil
}

// NOTE - Soft Delete แบบหลายคนพร้อมกัน
func (d DatabaseRequest) SoftArrayDelete(ids []UserDelete) error {
	tx := d.DB.Begin()
	listId := []string{}
	for _, v := range ids {
		listId = append(listId, v.ID)
	}
	if result := tx.Debug().Model(&Users{}).Where("id IN (?) and deleted_at IS NULL", listId).Update("deleted_at", now); result.RowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("%v", "Already deleted users")
	} else if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

// NOTE - Delete แบบ Remove from database
func (d DatabaseRequest) Remove(id string) error {
	tx := d.DB.Begin()
	if result := tx.Debug().Model(&Users{}).Where("id = ?", id).Delete(&Users{}); result.RowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("no user id %v in database", id)
	} else if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

// NOTE - Delete หลาย Users แบบ Remove from database
func (d DatabaseRequest) RemoveUsers(ids []UserDelete) error {
	tx := d.DB.Begin()
	listId := []string{}
	for _, v := range ids {
		listId = append(listId, v.ID)
	}
	if result := tx.Debug().Model(&Users{}).Where("id IN (?)", listId).Delete(&Users{}); result.RowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("%v", "No this user in the database")
	} else if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

// !SECTION - Delete
