package mysql

import (
	"gorm.io/gorm"
	domain "lean-oauth/internal/core/domains"
	"lean-oauth/internal/core/ports"
)

type registerCategoriesMysqlRepo struct {
	db *gorm.DB
}

func NewRegisterCategoriesMysqlRepo(db *gorm.DB) ports.RegisterCategories {
	return &registerCategoriesMysqlRepo{db: db}
}

func (r registerCategoriesMysqlRepo) Get(id int) (*domain.RegisterCategories, error) {
	var result domain.RegisterCategories
	err := r.db.Model(&RegisterCategoriesModel{}).
		Scan(&result).Where(domain.RegisterCategories{Rid: id}).
		Error
	return &result, err
}

func (r registerCategoriesMysqlRepo) List() ([]domain.RegisterCategories, error) {
	var result []domain.RegisterCategories
	err := r.db.Model(&RegisterCategoriesModel{}).Scan(&result).Error
	return result, err
}
