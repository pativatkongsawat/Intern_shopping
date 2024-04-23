package category

type Category struct {
	Id   int    `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"`
}
type CategoryUpdate struct {
	Id   int    `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"`
}

func (c Category) TableName() string {
	return "category"
}
