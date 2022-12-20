package handlers

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type RegistrationRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	FistName string `json:"firstName" form:"firstName"`
	LastName string `json:"lastName" form:"LastName"`
	Dob      string `json:"dob" form:"dob"`
}

func (c RegistrationRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Username, validation.Required),
		validation.Field(&c.Password, validation.Required),
		validation.Field(&c.FistName, validation.Required),
		validation.Field(&c.LastName, validation.Required),
		validation.Field(&c.Dob, validation.Required, validation.Date("2006-01-02")),
	)
}
