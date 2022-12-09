package domains

import (
	"github.com/google/uuid"
	"time"
)

type Members struct {
	Mid          uuid.UUID
	Username     string
	Password     string
	FirstName    string
	LastName     string
	DateOfBird   time.Time
	RegisterType int
	CreatedAt    time.Time
}

func NewMember(id uuid.UUID, username string, password string, firstName string, lastName string, dateOfBird time.Time, registerType int, createdAt time.Time) *Members {
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
