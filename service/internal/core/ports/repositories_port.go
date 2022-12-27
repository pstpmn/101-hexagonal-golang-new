package ports

import (
	"learn-oauth2/internal/core/domains"
)

type MembersRepository interface {
	Get(id string) (*domains.Members, error)
	GetByUser(user string) *domains.Members
	List() ([]domains.Members, error)
	Create(todo *domains.Members) (*domains.Members, error)
}

type RegisterCategories interface {
	Get(id int) (*domains.RegisterCategories, error)
	List() ([]domains.RegisterCategories, error)
}
