package mysql

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	domain "lean-oauth/internal/core/domains"
	"lean-oauth/internal/core/ports"
)

type membersMysqlRepo struct {
	db *gorm.DB
}

func NewTodoMysqlRepo(db *gorm.DB) ports.MembersRepository {
	return &membersMysqlRepo{
		db: db,
	}
}

func (m membersMysqlRepo) Get(id uuid.UUID) (*domain.Members, error) {
	var mem *domain.Members
	err := m.db.First(&MembersModel{Mid: id}).Scan(&mem).Error
	return mem, err
}

func (m membersMysqlRepo) List() ([]domain.Members, error) {
	//TODO implement me
	panic("implement me")
}

func (m membersMysqlRepo) Create(todo *domain.Members) (*domain.Members, error) {
	mem := MembersModel{todo.Mid, todo.Username, todo.Password, todo.FirstName, todo.LastName, todo.DateOfBird, 1, nil, todo.CreatedAt}
	err := m.db.Create(mem).Error
	return todo, err
}
