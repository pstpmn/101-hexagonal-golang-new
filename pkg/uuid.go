package pkg

import "github.com/google/uuid"

type IUuid interface {
	Random() string
}

type u struct {
}

func NewUuId() IUuid {
	return &u{}
}

func (u u) Random() string {
	return uuid.New().String()
}
