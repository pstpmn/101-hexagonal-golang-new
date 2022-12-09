package ports

import (
	"github.com/google/uuid"
	domain "lean-oauth/internal/core/domains"
	"time"
)

type MembersUseCase interface {
	NewMember(user string, pass string, fistName string, lastName string, dob time.Time) (*domain.Members, error)
	FindMemberById(id uuid.UUID) (*domain.Members, error)
}
