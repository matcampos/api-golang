package userModel

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Birthdate string `json:"birthdate"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		// Name cannot be empty.
		validation.Field(&u.Name, validation.Required),
		// Email cannot be empty, and must be a valid email.
		validation.Field(&u.Email, validation.Required, is.Email),
		// Password cannot be empty.
		validation.Field(&u.Password, validation.Required),
		// Birthdate cannot be empty, and must be in format: yyyy-mm-dd.
		validation.Field(&u.Birthdate, validation.Required, validation.Match(regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])"))),
	)
}
