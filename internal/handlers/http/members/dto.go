package handlers

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type RegistrationRequest struct {
	Username string    `json:"username" form:"username"`
	Password string    `json:"password" form:"password"`
	FistName string    `json:"firstName" form:"fistName"`
	LastName string    `json:"lastName" form:"LastName"`
	Dob      time.Time `json:"dob" form:"dob"`
}

func (c RegistrationRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Username, validation.Required),
		validation.Field(&c.Password, validation.In("Female", "Male")),
		validation.Field(&c.FistName, validation.Required, is.Email),
		validation.Field(&c.LastName),
	)
}
