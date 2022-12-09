package usecases

import (
	"github.com/google/uuid"
	domain "lean-oauth/internal/core/domains"
	"lean-oauth/internal/core/ports"
	"time"
)

type membersUseCase struct {
	membersRepo        ports.MembersRepository
	RegisterCategories ports.RegisterCategories
}

func NewMembersUseCase(members ports.MembersRepository, categories ports.RegisterCategories) ports.MembersUseCase {
	return &membersUseCase{membersRepo: members, RegisterCategories: categories}
}

func (m membersUseCase) NewMember(user string, pass string, fistName string, lastName string, dob time.Time) (*domain.Members, error) {
	member := domain.Members{uuid.New(), user, pass, fistName, lastName, dob, 1, time.Now()}
	result, err := m.membersRepo.Create(&member)
	return result, err
}

func (m membersUseCase) FindMemberById(id uuid.UUID) (*domain.Members, error) {
	//TODO implement me
	mem, err := m.membersRepo.Get(id)
	return mem, err
}
