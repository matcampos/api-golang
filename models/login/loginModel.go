package loginModel

import (
	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l Login) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Email, validation.Required, is.Email),
		// Password cannot be empty.
		validation.Field(&l.Password, validation.Required),
	)
}

type JwtToken struct {
	Token string `json:"token"`
}

type CustomClain struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}
