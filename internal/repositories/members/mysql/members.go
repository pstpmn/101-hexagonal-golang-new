package mysql

import (
	domain "lean-oauth/internal/core/domains"
	"lean-oauth/internal/core/ports"

	"gorm.io/gorm"
)

type membersMysqlRepo struct {
	db *gorm.DB
}

func (m membersMysqlRepo) GetByUser(user string) *domain.Members {
	var mem *domain.Members
	m.db.Table("members").Where("username = ? ", user).Take(&mem)
	return mem
}

func NewMembersMysqlRepo(db *gorm.DB) ports.MembersRepository {
	return &membersMysqlRepo{
		db: db,
	}
}

func (m membersMysqlRepo) Get(id string) (*domain.Members, error) {
	var mem *domain.Members
	err := m.db.First(&MembersModel{Mid: id}).Scan(&mem).Error
	return mem, err
}

func (m membersMysqlRepo) List() ([]domain.Members, error) {
	var mem []domain.Members
	err := m.db.Model(&MembersModel{}).Scan(&mem).Error
	return mem, err
}

func (m membersMysqlRepo) Create(todo *domain.Members) (*domain.Members, error) {
	mem := MembersModel{todo.Mid, todo.Username, todo.Password, todo.FirstName, todo.LastName, todo.DateOfBird, 1, nil, todo.CreatedAt}
	err := m.db.Create(mem).Error
	return todo, err
}
