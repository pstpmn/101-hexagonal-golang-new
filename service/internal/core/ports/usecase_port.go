package ports

import (
	domain "learn-oauth2/internal/core/domains"
	"time"
)

type MembersUseCase interface {
	NewMember(user string, pass string, fistName string, lastName string, dob time.Time) (*domain.Members, error)
	FindMemberById(id string) (*domain.Members, error)
	Authentication(user string, pass string, tokenKey string) (string, *domain.Members, error)
	Authorization(token string, key string) (map[string]interface{}, error)
}

type Oauth2UseCase interface {
	AuthzFacebook(accessTokenClient string, accessToken string) (string, error)
	//AuthzGoogle(accessTokenClient string, accessToken string) (bool, error)
}
