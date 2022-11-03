package request

import (
	"mini-project/businesses/students"

	"github.com/go-playground/validator/v10"
)

type Student struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (req *Student) ToDomain() *students.Domain {
	return &students.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *Student) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}