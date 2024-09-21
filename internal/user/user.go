package user

import (
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (e User) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Email, validation.Required, is.Email),
		validation.Field(&e.Password, validation.Required, validation.Length(5, 0)),
	)
}
