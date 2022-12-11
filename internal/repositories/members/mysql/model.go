package mysql

import (
	"lean-oauth/internal/repositories/register_categories/mysql"
	"time"
)

type MembersModel struct {
	Mid                string                         `gorm:"column:mid;primary_key;"`
	Username           string                         `gorm:"column:username;unique"`
	Password           string                         `gorm:"column:password;"`
	FirstName          string                         `gorm:"column:firstName;"`
	LastName           string                         `gorm:"column:lastName;"`
	DateOfBird         time.Time                      `gorm:"column:dateOfBird;type:timestamp;"`
	RegisterId         int                            `gorm:"index"`
	RegisterCategories *mysql.RegisterCategoriesModel `gorm:"foreignKey:RegisterId;references:Rid;association_foreignkey:Rid;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt          time.Time                      `gorm:"column:createdAt;type:timestamp;default:current_timestamp"`
}

func (MembersModel) TableName() string {
	return "members"
}
