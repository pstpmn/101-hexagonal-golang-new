package ports

import (
	domain "lean-oauth/internal/core/domains"
)

type MembersRepository interface {
	Get(id string) (*domain.Members, error)
	GetByUser(user string) *domain.Members
	List() ([]domain.Members, error)
	Create(todo *domain.Members) (*domain.Members, error)
}

type RegisterCategories interface {
	Get(id int) (*domain.RegisterCategories, error)
	List() ([]domain.RegisterCategories, error)
}
