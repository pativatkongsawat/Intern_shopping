package category

import (
	"log"

	"gorm.io/gorm"
)

type CategoryModelHelper struct {
	DB *gorm.DB
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

	if err := tx.Debug().Where("id = ?", id).Delete(&category).Error; err != nil {
		tx.Rollback()
		return nil, nil
	}
	tx.Commit()
	return category, nil
}

func (u *CategoryModelHelper) UpdateCategory(categoryData []Category) ([]Category, error) {
	tx := u.DB.Begin()

	updateCategory := []Category{}

	for _, category := range categoryData {

		updateValues := map[string]interface{}{
			"Name": category.Name,
		}

		if err := tx.Debug().Model(&Category{}).Where("id = ?", category.Id).Updates(updateValues).Error; err != nil {
			log.Println("Error updating category:", err)
			tx.Rollback()
			return nil, err
		}

		updateCategory = append(updateCategory, category)
	}

	tx.Commit()

	return updateCategory, nil
}
