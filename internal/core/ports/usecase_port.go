package ports

import (
	domain "lean-oauth/internal/core/domains"
	"time"
)

type MembersUseCase interface {
	NewMember(user string, pass string, fistName string, lastName string, dob time.Time) (*domain.Members, error)
	FindMemberById(id string) (*domain.Members, error)
}