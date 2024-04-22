package categoryRequest

import "gorm.io/gorm"

type CategoryModelHelper struct {
	DB *gorm.DB
}

type Category struct {
	Id   int    `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"`
}
type CategoryUpdate struct {
	Name string `json:"name" gorm:"name"`
}

func (c Category) TableName() string {
	return "category"
}

func (u *CategoryModelHelper) GetAllCategory() ([]Category, error) {

	category := []Category{}

	if err := u.DB.Debug().Find(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (u *CategoryModelHelper) InsertCategory(category []Category) error {
	tx := u.DB.Begin()

	if err := tx.Debug().Create(&category).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (u *CategoryModelHelper) DeleleteCategory(id int) ([]*Category, error) {
	category := []*Category{}
	tx := u.DB.Begin()

	if err := tx.Debug().Where("id = ?", id).Delete(category).Error; err != nil {
		tx.Rollback()
		return nil, nil
	}
	tx.Commit()
	return category, nil
}

func (u *CategoryModelHelper) UpdateCategory(id int, categoryupdate []CategoryUpdate) ([]Category, error) {

	//tx.Debug().Model(&User{}).Where("id = ?", User_id).Updates(&user);
	return nil, nil
}
