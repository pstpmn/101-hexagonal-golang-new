package domains

import (
	"time"
)

type RegisterCategories struct {
	Rid       int
	Name      string
	CreatedAt time.Time
}

func NewRegisterCategories(id int, name string, createdAt time.Time) *RegisterCategories {
	return &RegisterCategories{Rid: id, Name: name, CreatedAt: createdAt}
}
