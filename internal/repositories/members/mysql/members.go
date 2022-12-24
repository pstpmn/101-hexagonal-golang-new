package mysql

import (
	"gorm.io/gorm"
	domain "lean-oauth/internal/core/domains"
	"lean-oauth/internal/core/ports"
)

type membersMysqlRepo struct {
	db *gorm.DB
}

func (m membersMysqlRepo) GetByUser(user string) *domain.Members {
	model := &MembersModel{}
	m.db.Model(&MembersModel{}).Where(&MembersModel{Username: user}).Scan(&model)
	return &domain.Members{
		Mid:          model.Mid,
		Username:     model.Username,
		Password:     model.Password,
		FirstName:    model.FirstName,
		LastName:     model.LastName,
		DateOfBird:   model.DateOfBird,
		CreatedAt:    model.CreatedAt,
		RegisterType: model.RegisterId,
	}
}

func NewMembersMysqlRepo(db *gorm.DB) ports.MembersRepository {
	return &membersMysqlRepo{
		db: db,
	}
}

func (m membersMysqlRepo) Get(id string) (*domain.Members, error) {
	model := &MembersModel{}
	err := m.db.Model(&MembersModel{}).Where(&MembersModel{Mid: id}).Scan(&model).Error
	return &domain.Members{
		Mid:          model.Mid,
		Username:     model.Username,
		Password:     model.Password,
		FirstName:    model.FirstName,
		LastName:     model.LastName,
		DateOfBird:   model.DateOfBird,
		CreatedAt:    model.CreatedAt,
		RegisterType: model.RegisterId,
	}, err
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
