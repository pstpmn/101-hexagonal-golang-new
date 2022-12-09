package ports

import (
	"github.com/google/uuid"
	domain "lean-oauth/internal/core/domains"
)

type MembersRepository interface {
	Get(id uuid.UUID) (*domain.Members, error)
	List() ([]domain.Members, error)
	Create(todo *domain.Members) (*domain.Members, error)
}

type RegisterCategories interface {
	Get(id int) (*domain.RegisterCategories, error)
	List() ([]domain.RegisterCategories, error)
}
