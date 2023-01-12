package usecases

import (
	"errors"
	domain "learn-oauth2/internal/core/domains"
	ports2 "learn-oauth2/internal/core/ports"
	"time"
)

type membersUseCase struct {
	membersRepo        ports2.MembersRepository
	RegisterCategories ports2.RegisterCategories
	UidService         ports2.IUuidService
	CryptoService      ports2.ICryptoService
	JwtService         ports2.IJwtService
	requestService     ports2.IRequest
	LoggerService      ports2.ILogger
}

func NewMembersUseCase(members ports2.MembersRepository, categories ports2.RegisterCategories, uidService ports2.IUuidService, crypto ports2.ICryptoService, jwt ports2.IJwtService, request ports2.IRequest) ports2.MembersUseCase {
	return &membersUseCase{membersRepo: members, RegisterCategories: categories, UidService: uidService, CryptoService: crypto, JwtService: jwt, requestService: request}
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

func (m membersUseCase) Authorization(token string, key string) (data map[string]interface{}, err error) {
	data, err = m.JwtService.Extract(token, key)
	if err != nil {
		err = errors.New("invalid authorize token")
	}
	return
}

func (m membersUseCase) Authentication(user string, pass string, tokenKey string) (string, *domain.Members, error) {
	// find username
	// then not found stop flow and return error message to user
	mem := m.membersRepo.GetByUser(user)
	if mem.Mid == "" {
		return "", &domain.Members{}, errors.New("not found username")
	}

	// then found member
	// validate raw password and encript password
	if isValid := m.CryptoService.ValidateBcrypt(pass, mem.Password); isValid == false {
		return "", &domain.Members{}, errors.New("invalid username or password")
	}

	encript := map[string]interface{}{
		"memberId":  mem.Mid,
		"username":  mem.Username,
		"firstName": mem.FirstName,
		"lastName":  mem.LastName,
		"createdAt": mem.CreatedAt,
		"dob":       mem.DateOfBird,
	}
	expired := time.Now().Add(time.Hour + 3)
	token, err := m.JwtService.Generate(encript, tokenKey, expired)
	if err != nil {
		return "", &domain.Members{}, errors.New("error encript token")
	}

	return token, mem, nil
}
