package request

import (
	"mini-project/businesses/auth"

	"github.com/go-playground/validator/v10"
)

type Auth struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (req *Auth) ToDomain() *auth.Domain {
	return &auth.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *Auth) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}