package usecases

import (
	"errors"
	domain "lean-oauth/internal/core/domains"
	"lean-oauth/internal/core/ports"
	"time"
)

type membersUseCase struct {
	membersRepo        ports.MembersRepository
	RegisterCategories ports.RegisterCategories
	UidService         ports.IUuidService
}

func NewMembersUseCase(members ports.MembersRepository, categories ports.RegisterCategories) ports.MembersUseCase {
	return &membersUseCase{membersRepo: members, RegisterCategories: categories}
}

func (m membersUseCase) NewMember(user string, pass string, fistName string, lastName string, dob time.Time) (*domain.Members, error) {
	member := domain.Members{Mid: m.UidService.Random(), Username: user, Password: pass, FirstName: fistName, LastName: lastName, DateOfBird: dob, RegisterType: 1, CreatedAt: time.Now()}
	result, err := m.membersRepo.Create(&member)
	return result, err
}

func (m membersUseCase) FindMemberById(id string) (*domain.Members, error) {
	mem, err := m.membersRepo.Get(id)
	if err != nil {
		return mem, errors.New("error find member !!")
	} else if mem.Mid == "" {
		return mem, errors.New("not found member")
	}
	return mem, err
}
