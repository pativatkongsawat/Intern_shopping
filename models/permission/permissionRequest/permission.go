package permissionrequest

type Permission struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique;not null" json:"name"`
}

func (Permission) TableName() string {
	return "permissions"
}
