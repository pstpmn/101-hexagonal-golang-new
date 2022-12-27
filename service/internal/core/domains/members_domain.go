package domains

import (
	"time"
)

type Members struct {
	Mid          string
	Username     string
	Password     string
	FirstName    string
	LastName     string
	DateOfBird   time.Time
	RegisterType int
	CreatedAt    time.Time
}

func NewMember(id string, username string, password string, firstName string, lastName string, dateOfBird time.Time, registerType int, createdAt time.Time) *Members {
	return &Members{
		Mid:          id,
		Username:     username,
		Password:     password,
		FirstName:    firstName,
		LastName:     lastName,
		DateOfBird:   dateOfBird,
		RegisterType: registerType,
		CreatedAt:    createdAt,
	}
}
