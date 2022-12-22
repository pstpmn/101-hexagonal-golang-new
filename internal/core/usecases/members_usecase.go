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
	CryptoService      ports.ICryptoService
}

func NewMembersUseCase(members ports.MembersRepository, categories ports.RegisterCategories, uidService ports.IUuidService, crypto ports.ICryptoService) ports.MembersUseCase {
	return &membersUseCase{membersRepo: members, RegisterCategories: categories, UidService: uidService, CryptoService: crypto}
}

func (m membersUseCase) NewMember(user string, pass string, fistName string, lastName string, dob time.Time) (*domain.Members, error) {
	var uuid string = m.UidService.Random()
	encryptPass, err := m.CryptoService.Bcrypt(pass)
	member := domain.Members{Mid: uuid, Username: user, Password: encryptPass, FirstName: fistName, LastName: lastName, DateOfBird: dob, RegisterType: 1, CreatedAt: time.Now()}

	if err != nil {
		return &domain.Members{}, errors.New("error encrypt pass")
	}

	// validate username is used
	if findMember := m.membersRepo.GetByUser(user); findMember.Mid != "" {
		return &domain.Members{}, errors.New("username is already in used")
	}

	result, err := m.membersRepo.Create(&member)
	if err != nil {
		return &domain.Members{}, errors.New("error create member")
	}
	return result, nil
}

func (m membersUseCase) FindMemberById(id string) (*domain.Members, error) {
	mem, err := m.membersRepo.Get(id)
	if err != nil {
		return mem, errors.New("error find member ")
	} else if mem.Mid == "" {
		return mem, errors.New("not found member")
	}
	return mem, err
}

func (m membersUseCase) Authentication(user string, pass string) (*domain.Members, error) {
	// find username
	// then not found stop flow and return error message to user
	mem := m.membersRepo.GetByUser(user)
	if mem.Mid == "" {
		return &domain.Members{}, errors.New("not found username")
	}

	// then found member
	// validate raw password and encript password
	if isValid := m.CryptoService.ValidateBcrypt(pass, mem.Password); isValid == false {
		return &domain.Members{}, errors.New("invalid username or password")
	}
	return mem, nil
}
