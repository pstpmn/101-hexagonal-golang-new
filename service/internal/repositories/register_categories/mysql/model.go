package mysql

import (
	"time"
)

type RegisterCategoriesModel struct {
	Rid       int       `gorm:"column:rid;primary_key;"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:createdAt;type:timestamp;default:current_timestamp"`
}

func (RegisterCategoriesModel) TableName() string {
	return "register_categories"
}
